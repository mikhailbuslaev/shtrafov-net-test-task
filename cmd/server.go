package cmd

import (
	"context"
	"mihailbuslaev/sntt/internal/flags"
	"mihailbuslaev/sntt/internal/server"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "mihailbuslaev/sntt/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var cfg Config

var ServerCmd = &cobra.Command{
	Use:          "server",
	Short:        "start server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Info().Msg("run service")
		// flags
		flags.BindEnv(cmd)

		// grpc service prepare and run
		lis, err := net.Listen("tcp", cfg.TcpAddr)
		if err != nil {
			log.Err(err).Msgf("failed to listen: %v", err)
		}
		var grpcOpts []grpc.ServerOption

		grpcServer := grpc.NewServer(grpcOpts...)
		pb.RegisterRouteGuideServer(grpcServer, server.NewGrpcServerImplement())
		go grpcServer.Serve(lis)

		// http server over grpc prepare and run
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()

		dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err = pb.RegisterRouteGuideHandlerFromEndpoint(ctx, mux, cfg.TcpAddr, dialOpts)
		if err != nil {
			log.Err(err).Msg("http server registration failed")
		}

		log.Info().Msgf("http server listening at %s", cfg.HttpPort)
		if err := http.ListenAndServe(cfg.HttpPort, mux); err != nil {
			log.Err(err).Msg("http server launch failed")
		}

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
