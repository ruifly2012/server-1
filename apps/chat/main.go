package main

import (
	"log"
	"os"

	"github.com/yokaiio/yokai_server/chat"
	"github.com/yokaiio/yokai_server/entries"

	// micro plugins
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/consul"
	_ "github.com/micro/go-plugins/store/consul"
	_ "github.com/micro/go-plugins/transport/tcp"
)

func init() {
	// set working directory as yokai_server
	os.Chdir("../../")
}

func main() {
	// entries init
	entries.InitEntries()

	c, err := chat.NewChat()
	if err != nil {
		log.Fatal("chat new error:", err)
		os.Exit(1)
	}

	if err = c.Run(os.Args); err != nil {
		log.Fatal("chat run error:", err)
		os.Exit(1)
	}

	c.Stop()
}
