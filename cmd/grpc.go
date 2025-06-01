package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arvinpaundra/cent/payment/config"
	"github.com/arvinpaundra/cent/payment/core"
	"github.com/arvinpaundra/cent/payment/database/sqlpkg"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcPort string

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		pgsql := sqlpkg.NewPostgres()

		sqlpkg.NewConnection(pgsql)

		srv := grpc.NewServer()

		go func() {
			addr := fmt.Sprintf(":%s", grpcPort)

			listener, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("failed to listen: %s", err.Error())
			}

			err = srv.Serve(listener)
			if err != nil {
				log.Fatalf("failed to start grpc server: %s", err.Error())
			}
		}()

		wait := core.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"grpc-server": func(_ context.Context) error {
				srv.GracefulStop()

				return nil
			},
			"postgres": func(_ context.Context) error {
				return pgsql.Close()
			},
		})

		_ = <-wait
	},
}

func init() {
	grpcCmd.Flags().StringVarP(&grpcPort, "port", "p", "8093", "bind grpc to port")
	rootCmd.AddCommand(grpcCmd)
}
