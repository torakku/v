package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/torakku/v"
)

const (
	ConfigFile = "cagliostro.json"
)

func main() {
	var (
		err error

		config *Config
	)

	config, err = LoadConfigFile(ConfigFile)
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	if config.Token == "" {
		logger.Fatal("missing token")
	}
	if config.Prefix == "" {
		logger.Fatal("missing prefix")
	}

	c := &cagliostro.Cagliostro{
		Token:    config.Token,
		Prefix:   config.Prefix,
		EmojiDir: config.EmojiDir,
		Logger:   logger,
	}

	err = c.Open()
	if err != nil {
		log.Fatal(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-interrupt

	err = c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
