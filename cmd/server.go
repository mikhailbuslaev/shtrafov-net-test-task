package cmd

import (
	server "mihailbuslaev/sntt/internal"
	"net"
	"os"
	"os/signal"
	"syscall"

	service "mihailbuslaev/sntt/pkg/api"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var cfg Config

var ServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "start server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		// prepare flags
		log.Info().Msg("run server command...")
		lis, err := net.Listen("tcp", cfg.TcpAddr)
		if err != nil {
			log.Error().Msgf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)
		service.RegisterRouteGuideServer(grpcServer, server.NewGrpcServerImplement())
		grpcServer.Serve(lis)

		sigCh := make(chan os.Signal, 1)
		defer close(sigCh)
		go func() {
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)

			<-sigCh
			log.Info().Msg("stop signal received... exit")
			grpcServer.Stop()
		}()
		return nil
	},
}

func init() {
	ServerCmd.Flags().AddFlagSet(cfg.Flags())
}
