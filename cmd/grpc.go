package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arvinpaundra/cent/payment/core"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcPort string

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
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
			"grpc-server": func(ctx context.Context) error {
				srv.Stop()

				return nil
			},
		})

		_ = <-wait
	},
}

func init() {
	grpcCmd.Flags().StringVarP(&grpcPort, "port", "p", "8093", "bind grpc to port. default: 8093")
	rootCmd.AddCommand(grpcCmd)
}
