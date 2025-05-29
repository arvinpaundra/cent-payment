package cmd

import "github.com/spf13/cobra"

var grpcPort string

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		select {}
	},
}

func init() {
	grpcCmd.Flags().StringVarP(&grpcPort, "port", "p", "8093", "bind grpc to port. default: 8093")
	rootCmd.AddCommand(grpcCmd)
}
