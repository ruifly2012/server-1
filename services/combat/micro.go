package combat

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	micro_logger "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/transport"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/rs/zerolog/log"
	ucli "github.com/urfave/cli/v2"
	"github.com/east-eden/server/logger"
)

type MicroService struct {
	srv micro.Service
	c   *Combat
}

func NewMicroService(c *Combat, ctx *ucli.Context) *MicroService {
	// set metadata
	servID, err := strconv.Atoi(ctx.String("combat_id"))
	if err != nil {
		log.Fatal().Str("combat_id", ctx.String("combat_id")).Msg("wrong combat_id")
		return nil
	}

	section := servID / 10
	metadata := make(map[string]string)
	metadata["combatId"] = fmt.Sprintf("%d", servID)
	metadata["section"] = fmt.Sprintf("%d", section)

	// cert
	certPath := ctx.String("cert_path_release")
	keyPath := ctx.String("key_path_release")

	if ctx.Bool("debug") {
		certPath = ctx.String("cert_path_debug")
		keyPath = ctx.String("key_path_debug")
	}

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("load certificates failed")
	}
	tlsConf.Certificates = []tls.Certificate{cert}

	s := &MicroService{c: c}

	micro_logger.Init(micro_logger.WithOutput(logger.Logger))
	s.srv = micro.NewService(
		micro.Name("combat"),

		micro.Transport(grpc.NewTransport(
			transport.TLSConfig(tlsConf),
		)),

		micro.Metadata(metadata),

		micro.Flags(&cli.StringFlag{
			Name:  "config_file",
			Usage: "config file path",
		}),
	)

	// set environment
	os.Setenv("MICRO_SERVER_ID", ctx.String("combat_id"))

	if ctx.Bool("debug") {
		os.Setenv("MICRO_REGISTRY", ctx.String("registry_debug"))
		os.Setenv("MICRO_BROKER", ctx.String("broker_debug"))
	} else {
		os.Setenv("MICRO_REGISTRY", ctx.String("registry_release"))
		os.Setenv("MICRO_BROKER", ctx.String("broker_release"))
	}

	s.srv.Init()

	return s
}

func (s *MicroService) Run() error {

	// Run service
	if err := s.srv.Run(); err != nil {
		return err
	}

	return nil
}