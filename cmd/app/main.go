package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kam1dere/CSVDataService/pkg/di"
	"github.com/kam1dere/CSVDataService/pkg/rest"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

func main() {
	injector := do.New()

	do.Provide(injector, di.NewConfig)
	do.Provide(injector, di.NewFlags)
	do.Provide(injector, di.NewLogger)
	do.Provide(injector, di.NewServer)
	do.Provide(injector, di.NewService)
	do.Provide(injector, di.NewRestServer)

	server := do.MustInvoke[*di.GrpcServer](injector)
	interruptChan := make(chan struct{})
	go func() {
		defer close(interruptChan)

		err := server.Start()
		if err != nil {
			panic(err)
		}
	}()

	restServer := do.MustInvoke[*rest.Rest](injector)
	go func() {
		restServer.NewRestServer()
	}()

	logger := do.MustInvoke[*zerolog.Logger](injector)
	logger.Info().Msg("Server started")

	<-interruptChan
}
