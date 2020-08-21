package gate

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"
	"sync"
	"time"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/yokaiio/yokai_server/utils"
)

var (
	httpReadTimeout           = time.Second * 5
	httpWriteTimeout          = time.Second * 5
	ginConcurrentRequestLimit = 1000
)

var (
	opsSelectGameCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "gate",
		Name:      "select_game_ops",
		Help:      "选择服务器操作总数",
	})

	timeCounterHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "gate",
		Name:      "select_game_addr_latency",
		Help:      "请求延迟",
	},
		[]string{"method"},
	)
)

type GinServer struct {
	g         *Gate
	router    *gin.Engine
	tlsRouter *gin.Engine
	wg        utils.WaitGroupWrapper
}

// wrap http.HandlerFunc to gin.HandlerFunc
func ginHandlerWrapper(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {

				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *GinServer) setupHttpRouter() {
	s.router.Use(limit.MaxAllowed(ginConcurrentRequestLimit))

	// pprof
	s.router.GET("/debug/pprof", ginHandlerWrapper(pprof.Index))
	s.router.GET("/debug/cmdline", ginHandlerWrapper(pprof.Cmdline))
	s.router.GET("/debug/symbol", ginHandlerWrapper(pprof.Symbol))
	s.router.GET("/debug/profile", ginHandlerWrapper(pprof.Profile))
	s.router.GET("/debug/trace", ginHandlerWrapper(pprof.Trace))
	s.router.GET("/debug/allocs", ginHandlerWrapper(pprof.Handler("allocs").ServeHTTP))
	s.router.GET("/debug/heap", ginHandlerWrapper(pprof.Handler("heap").ServeHTTP))
	s.router.GET("/debug/goroutine", ginHandlerWrapper(pprof.Handler("goroutine").ServeHTTP))
	s.router.GET("/debug/block", ginHandlerWrapper(pprof.Handler("block").ServeHTTP))
	s.router.GET("/debug/threadcreate", ginHandlerWrapper(pprof.Handler("threadcreate").ServeHTTP))

	// metrics
	metricsHandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{Registry: prometheus.DefaultRegisterer})
	s.router.GET("/metrics", ginHandlerWrapper(metricsHandler.ServeHTTP))
}

func (s *GinServer) setupHttpsRouter() {
	s.tlsRouter.Use(limit.MaxAllowed(ginConcurrentRequestLimit))

	// store_write
	s.tlsRouter.POST("/store_write", func(c *gin.Context) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		if c.Bind(&req) == nil {
			s.g.mi.StoreWrite(req.Key, req.Value)
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		c.String(http.StatusBadRequest, "bad request")
	})

	// select_game_addr
	s.tlsRouter.POST("/select_game_addr", func(c *gin.Context) {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			us := v * 1000000 // make microseconds
			timeCounterHistogram.WithLabelValues("/select_game_addr").Observe(us)
		}))
		defer timer.ObserveDuration()

		opsSelectGameCounter.Inc()

		var req struct {
			UserID   string `json:"userId"`
			UserName string `json:"userName"`
		}

		if err := c.Bind(&req); err != nil {
			logger.Warn("select_game_addr request bind error:", err)
			c.String(http.StatusBadRequest, "bad request:%s", err.Error())
			return
		}

		if user, metadata := s.g.gs.SelectGame(req.UserID, req.UserName); user != nil {
			h := gin.H{
				"userId":        user.UserID,
				"userName":      req.UserName,
				"accountId":     user.AccountID,
				"gameId":        metadata["gameId"],
				"publicTcpAddr": metadata["publicTcpAddr"],
				"publicWsAddr":  metadata["publicWsAddr"],
				"section":       metadata["section"],
			}
			c.JSON(http.StatusOK, h)

			logger.Info("select_game_addr calling with result:", h)
			return
		}

		c.String(http.StatusBadRequest, fmt.Sprintf("cannot find account by userid<%s>", req.UserID))
	})

	// pub_gate_result
	s.tlsRouter.POST("/pub_gate_result", func(c *gin.Context) {
		s.g.GateResult()
		c.String(http.StatusOK, "status ok")
	})

	// update_player_exp
	s.tlsRouter.POST("/update_player_exp", func(c *gin.Context) {
		var req struct {
			Id string `json:"id"`
		}

		if c.Bind(&req) == nil {
			id, err := strconv.ParseInt(req.Id, 10, 64)
			if err != nil {
				c.String(http.StatusBadRequest, "request error")
				return
			}
			r, err := s.g.rpcHandler.CallUpdatePlayerExp(id)
			c.String(http.StatusOK, "UpdatePlayerExp result", r, err)

			// test storage
			//user := NewUserInfo().(*UserInfo)
			//if err := store.GetStore().LoadObject(store.StoreType_User, "_id", id, user); err != nil {
			//logger.Warn(err)
			//}

			//user.UserID = id
			//user.PlayerName = "dudu"
			//if err := store.GetStore().SaveObject(store.StoreType_User, user); err != nil {
			//logger.Warn(err)
			//}

			//user.PlayerLevel++
			//user.PlayerName += "."
			//fields := map[string]interface{}{
			//"player_level": user.PlayerLevel,
			//"player_name":  user.PlayerName,
			//}
			//if err := store.GetStore().SaveFields(store.StoreType_User, user, fields); err != nil {
			//logger.Warn(err)
			//}
		}

	})

	// get_lite_player
	s.tlsRouter.POST("/get_lite_player", func(c *gin.Context) {
		var req struct {
			PlayerId string `json:"playerId"`
		}

		if c.Bind(&req) == nil {
			id, err := strconv.ParseInt(req.PlayerId, 10, 64)
			if err != nil {
				c.String(http.StatusBadRequest, "request error")
				return
			}

			rep, err := s.g.rpcHandler.CallGetRemoteLitePlayer(id)
			if err == nil {
				c.JSON(http.StatusOK, rep)
				return
			}

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusBadRequest, "request error")
	})

}

func NewGinServer(g *Gate, ctx *cli.Context) *GinServer {
	if ctx.Bool("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &GinServer{
		g:         g,
		router:    gin.Default(),
		tlsRouter: gin.Default(),
	}

	s.setupHttpRouter()
	s.setupHttpsRouter()
	return s
}

func (s *GinServer) Main(ctx *cli.Context) error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal("GinServer Run() error:", err)
			}
			exitCh <- err
		})
	}

	s.wg.Wrap(func() {
		exitFunc(s.Run(ctx))
	})

	// listen https
	go func() {
		certPath := ctx.String("cert_path_release")
		keyPath := ctx.String("key_path_release")
		if ctx.Bool("debug") {
			certPath = ctx.String("cert_path_debug")
			keyPath = ctx.String("key_path_debug")
		}

		server := &http.Server{
			Addr:         ctx.String("https_listen_addr"),
			Handler:      s.tlsRouter,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
		}

		if err := server.ListenAndServeTLS(certPath, keyPath); err != nil {
			logger.Error("GinServer RunTLS error:", err)
			exitCh <- err
		}
	}()

	// listen http
	go func() {
		server := &http.Server{
			Addr:         ctx.String("http_listen_addr"),
			Handler:      s.router,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
		}

		if err := server.ListenAndServe(); err != nil {
			logger.Error("GinServer Run error:", err)
			exitCh <- err
		}
	}()

	return <-exitCh
}

func (s *GinServer) Run(ctx *cli.Context) error {

	for {
		select {
		case <-ctx.Done():
			logger.Info("GinServer context done...")
			return nil
		}
	}

	return nil
}

func (s *GinServer) Exit(ctx context.Context) {
	s.wg.Wait()
	logger.Info("gin server exit...")
}
