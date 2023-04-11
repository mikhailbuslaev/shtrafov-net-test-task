package cmd

import (
	"mihailbuslaev/pb-wrapper/internal/server"
	"net"

	service "mihailbuslaev/pb-wrapper/pkg/api"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var ServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "start server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		// prepare flags
		log.Info().Msg("run server command...")
		lis, err := net.Listen("tcp", "228")
		if err != nil {
			log.Error().Msgf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)
		service.RegisterRouteGuideServer(grpcServer, server.NewGrpcServerImplement())
		grpcServer.Serve(lis)
		return nil
	},
}
