/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PhuMinh08082001/go-jwt-authen/config"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/controller"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/dal"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/middleware"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/repository"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/routes"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/server"
	"github.com/PhuMinh08082001/go-jwt-authen/internal/service"
	"go.uber.org/fx"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run HTTP server",
	Run:   initHttpServer,
}

func initHttpServer(cmd *cobra.Command, args []string) {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		config.Module,
		dal.Module,
		controller.Module,
		middleware.Module,
		service.Module,
		repository.Module,
		server.Module,
		routes.Module,
	)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
