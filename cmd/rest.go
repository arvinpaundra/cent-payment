package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arvinpaundra/cent/payment/api/route"
	"github.com/arvinpaundra/cent/payment/config"
	"github.com/arvinpaundra/cent/payment/core"
	"github.com/arvinpaundra/cent/payment/core/grpc"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"github.com/arvinpaundra/cent/payment/database/sqlpkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var restPort string

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		pgsql := sqlpkg.NewPostgres()

		sqlpkg.NewConnection(pgsql)

		g := gin.New()

		_ = route.NewRoutes(g, sqlpkg.GetConnection(), validator.NewValidator()).
			WithPublic().
			WithPrivate().
			WithInternal()

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", restPort),
			Handler: g,
		}

		userClientAddr := viper.GetString("USER_SERVICE_ADDR")

		grpcClient, err := grpc.NewClientFactory(grpc.ClientConfig{UserClientAddr: userClientAddr})
		if err != nil {
			log.Fatalf("failed to dial grpc client: %s", err.Error())
		}

		go func() {
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("failed to start server: %s", err.Error())
			}
		}()

		wait := core.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"rest-server": func(_ context.Context) error {
				return srv.Close()
			},
			"postgres": func(_ context.Context) error {
				return pgsql.Close()
			},
			"grpc-client": func(ctx context.Context) error {
				return grpcClient.Close()
			},
		})

		_ = <-wait
	},
}

func init() {
	restCmd.Flags().StringVarP(&restPort, "port", "p", "8090", "bind rest server to port")
	rootCmd.AddCommand(restCmd)
}
