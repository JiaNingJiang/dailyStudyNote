package cmd

import (
	"github.com/grpc-gateway-example/server"
	"github.com/spf13/cobra"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()
		//server.Serve()
		server.ServeWithoutTLS()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.ServerPemPath, "cert-pem", "", "./cert/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.ServerKeyPath, "cert-key", "", "./cert/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.ServerCertName, "cert-name", "", "www.github.com", "server's hostname")
	rootCmd.AddCommand(serverCmd)
}
