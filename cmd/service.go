package cmd

import (
	"context"
	"log"
	"mikrotik-script-generator/config"
	"mikrotik-script-generator/internal/controller"
	"mikrotik-script-generator/pkg/httpserver"
	"mikrotik-script-generator/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var createServiceCommand = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Start service",
		Long:  "Start service",
		Run: func(_ *cobra.Command, _ []string) {
			runSevice()
		},
	}

	return cmd
}

func runSevice() {
	appCtx, canFunc := context.WithCancel(context.Background())
	conf := config.NewConfig()
	logger := logger.New(conf.Log.Level)
	app := fx.New(
		fx.StartTimeout(conf.App.StartTimeout), fx.StopTimeout(conf.App.StopTimeout),
		fx.Provide(
			newHttpServer,
			controller.NewHomeController,
		),
		fx.Supply(conf, logger),
		fx.Invoke(startServer),
		fx.Decorate(),
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	stopChan := make(chan os.Signal)

	err := app.Start(appCtx)
	if err != nil {
		log.Fatalf("error occurred when start app %+v", err)
	}

	go func() {
		val := <-quit

		logger.Info("stopping app")

		err := app.Stop(appCtx)
		if err != nil {
			logger.Error("error occurred when stop app %v", err)
		}

		canFunc()
		stopChan <- val
	}()

	<-stopChan
}

func startServer(lc fx.Lifecycle,
	conf *config.Config,
	logger *logger.Logger,
	server httpserver.Interface,
	homeController controller.HomeController,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Http server is listening on %v", conf.Http.Port)

			controller.SetRoutes(server, homeController)

			server.Start(ctx)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Http server is shutting down")

			return server.Stop(ctx)
		},
	})
}

func newHttpServer(_ fx.Lifecycle, conf *config.Config) httpserver.Interface {
	return httpserver.New(conf.Http.Port, conf.Http.ReadTimeout, conf.Http.WriteTimeout)
}
