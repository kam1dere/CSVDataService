package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kam1dere/CSVDataService/config"
	"github.com/kam1dere/CSVDataService/grpc/genproto/CsvDataService"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Rest struct {
	cfg *config.Config
}

func NewRest(cfg *config.Config) *Rest {
	return &Rest{
		cfg: cfg,
	}
}

func (r *Rest) NewRestServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := CsvDataService.RegisterCsvDataServiceHandlerFromEndpoint(ctx, mux, r.cfg.Server.Address, opts)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msgf("server listening at %s", r.cfg.Server.RestServer)
	if err := http.ListenAndServe(r.cfg.Server.RestServer, mux); err != nil {
		log.Fatal().Err(err)
	}
}
