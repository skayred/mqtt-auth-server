package main

import (
	"context"
	"fmt"
	"mqtt-auth-server/auth"
	"mqtt-auth-server/mqtt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var logger = logrus.New()

// fmt.Sprintf("panimone/%s/relay/1", mac),

type Server struct {
	auth.UnimplementedMqttAuthServer

	Logger         *logrus.Logger
	BrokerIP       string
	BrokerPort     string
	MasterUsername string
	MasterPassword string
}

func (server Server) CreateUser(ctx context.Context, req *auth.UserRequest) (*auth.Response, error) {
	logger.Info("Creating user", req.Login, req.Token)
	err := mqtt.CreateUser(
		server.BrokerIP,
		server.BrokerPort,
		server.MasterUsername,
		server.MasterPassword,
		req.Login,
		req.Token)

	if err != nil {
		logger.Error(err, "Error when user creating!")
	}

	return &auth.Response{}, err
}

func (server Server) UpdateUser(ctx context.Context, req *auth.UserRequest) (*auth.Response, error) {
	logger.Info("Updating user", req.Login, req.Token)
	err := mqtt.UpdateUser(
		server.BrokerIP,
		server.BrokerPort,
		server.MasterUsername,
		server.MasterPassword,
		req.Login,
		req.Token)

	if err != nil {
		logger.Error(err, "Error when updating user!")
	}

	return &auth.Response{}, err
}

func (server Server) ActivateDevice(ctx context.Context, req *auth.DeviceRequest) (*auth.Response, error) {
	logger.Info("Adding topic permisions for the device", req.Login, req.Mac)
	err := mqtt.GrantTopicAccess(
		server.BrokerIP,
		server.BrokerPort,
		server.MasterUsername,
		server.MasterPassword,
		req.Login,
		// TODO: change the topic to those you would like to manage!
		[]string{
			fmt.Sprintf("panimone/%s/relay/1", req.Mac),
			fmt.Sprintf("panimone/%s/relay/2", req.Mac),
			fmt.Sprintf("panimone/%s/relay/3", req.Mac)})

	if err != nil {
		logger.Error(err, "Error when activation device!")
	}

	return &auth.Response{}, err
}

func (server Server) DeactivateDevice(ctx context.Context, req *auth.DeviceRequest) (*auth.Response, error) {
	logger.Println("Removing topic permissions for the device device", req.Login, req.Mac)
	err := mqtt.DenyTopicAccess(
		server.BrokerIP,
		server.BrokerPort,
		server.MasterUsername,
		server.MasterPassword,
		req.Login,
		// TODO: change the topic to those you would like to manage!
		[]string{
			fmt.Sprintf("panimone/%s/relay/1", req.Mac),
			fmt.Sprintf("panimone/%s/relay/2", req.Mac),
			fmt.Sprintf("panimone/%s/relay/3", req.Mac)})

	if err != nil {
		logger.Error(err, "Error when deactivation device!")
	}

	return &auth.Response{}, err
}

func main() {
	var serverPort, portFound = os.LookupEnv("PORT")
	var brokerIP, brokerFound = os.LookupEnv("MQTT_BROKER_IP")
	var brokerPort, brokerPortFound = os.LookupEnv("MQTT_BROKER_PORT")
	var username, usernameFound = os.LookupEnv("MQTT_MASTER_USERNAME")
	var password, passwordFound = os.LookupEnv("MQTT_MASTER_PASSWORD")

	if !portFound {
		logger.Fatal("PORT environment variable is not set!")
	}

	if !brokerFound {
		logger.Fatal("MQTT_BROKER_IP environment variable is not set!")
	}

	if !brokerPortFound {
		logger.Fatal("MQTT_BROKER_PORT environment variable is not set!")
	}

	if !usernameFound {
		logger.Fatal("MQTT_MASTER_USERNAME environment variable is not set!")
	}

	if !passwordFound {
		logger.Fatal("MQTT_MASTER_PASSWORD environment variable is not set!")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", serverPort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterMqttAuthServer(s, &Server{
		Logger:         logger,
		BrokerIP:       brokerIP,
		BrokerPort:     brokerPort,
		MasterUsername: username,
		MasterPassword: password,
	})
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
