package main

import (
	"github.com/CESARBR/knot-babeltower/internal/config"
	"github.com/CESARBR/knot-babeltower/pkg/network"
	"github.com/CESARBR/knot-babeltower/pkg/server"

	"github.com/CESARBR/knot-babeltower/pkg/logging"
)

func main() {
	config := config.Load()
	logrus := logging.NewLogrus(config.Logger.Level)

	logger := logrus.Get("Main")
	logger.Info("Starting KNoT Babeltower")

	amqp := network.NewAmqpHandler(logrus.Get("AmqpHandler"))
	amqp.Start()

	server := server.NewServer(config.Server.Port, logrus.Get("Server"))
	server.Start()
}
