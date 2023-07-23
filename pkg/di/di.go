package di

import (
	"fmt"
	"github.com/kam1dere/CSVDataService/app"
	"github.com/kam1dere/CSVDataService/config"
	"github.com/kam1dere/CSVDataService/grpc/genproto/csvDataService"
	"github.com/kam1dere/CSVDataService/pkg/flags"
	"github.com/kam1dere/CSVDataService/pkg/logger"
	"github.com/kam1dere/CSVDataService/pkg/rest"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func NewLogger(i *do.Injector) (*zerolog.Logger, error) {
	cfg := do.MustInvoke[*config.Config](i)
	logger := logger.NewLogger(cfg)

	return &logger, nil
}

func NewFlags(_ *do.Injector) (*flags.Flags, error) {
	flags, err := flags.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create flags: %w", err)
	}

	return flags, nil
}

func NewConfig(i *do.Injector) (*config.Config, error) {
	allFlags := do.MustInvoke[*flags.Flags](i)

	viperConfig, err := config.LoadConfig(*allFlags.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("load config error: %w", err)
	}

	cfg, err := config.ParseConfig(viperConfig)
	if err != nil {
		return nil, fmt.Errorf("parse config error: %w", err)
	}

	return cfg, nil
}

func NewService(i *do.Injector) (*app.DataService, error) {
	service, err := app.NewDataService(
		do.MustInvoke[*zerolog.Logger](i),
		do.MustInvoke[*config.Config](i),
	)
	if err != nil {
		return nil, fmt.Errorf("creating service error: %w", err)
	}

	return service, nil
}

func NewRestServer(i *do.Injector) (*rest.Rest, error) {
	service := rest.NewRest(
		do.MustInvoke[*config.Config](i),
	)
	return service, nil
}

func NewServer(i *do.Injector) (*GrpcServer, error) {
	config := do.MustInvoke[*config.Config](i)
	service := do.MustInvoke[*app.DataService](i)

	listener, err := net.Listen(config.Server.Network, config.Server.Address)
	if err != nil {
		return nil, fmt.Errorf("create listener error: %w", err)
	}

	server := grpc.NewServer()

	csvDataService.RegisterServiceServer(server, service)
	reflection.Register(server)

	return &GrpcServer{
		Ln:   listener,
		Serv: server,
	}, nil
}
