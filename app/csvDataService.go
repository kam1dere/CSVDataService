package app

import (
	"github.com/kam1dere/CSVDataService/config"
	"github.com/kam1dere/CSVDataService/grpc/genproto/csvDataService"
	"github.com/rs/zerolog"
)

type DataService struct {
	logger *zerolog.Logger
	config *config.Config

	csvDataService.UnimplementedCsvDataServiceServer
}

func NewDataService(
	logger *zerolog.Logger,
	cfg *config.Config,
) (*DataService, error) {
	return &DataService{
		logger: logger,
		config: cfg,
	}, nil
}
