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
	"github.com/arvinpaundra/cent/payment/core/validator"
	"github.com/arvinpaundra/cent/payment/database/sqlpkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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

		route.NewRoutes(
			g,
			sqlpkg.GetConnection(),
			validator.NewValidator(),
		).GatherRoutes()

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", restPort),
			Handler: g,
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
		})

		_ = <-wait
	},
}

func init() {
	restCmd.Flags().StringVarP(&restPort, "port", "p", "8090", "bind rest server to port. default: 8090")
	rootCmd.AddCommand(restCmd)
}
