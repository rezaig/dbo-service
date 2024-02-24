package console

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/delivery/httpsvc"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(_ *cobra.Command, _ []string) {
	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	authHTTPSvc := httpsvc.NewAuthHTTPService()
	customerHTTPSvc := httpsvc.NewCustomerHTTPService()
	orderHTTPSvc := httpsvc.NewOrderHTTPService()

	go func() {
		// Handle graceful shutdown
		for {
			select {
			case <-sigCh:
				log.Println("server is shutting down because of interrupt signal")
				quitCh <- true
			case e := <-errCh:
				log.Printf("server is shutting down because of error, error: %v\n", e)
				quitCh <- true
			}
		}
	}()

	go func() {
		// Init HTTP server
		r := gin.Default()

		r.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong!")
		})

		authHTTPSvc.Routes(r)
		customerHTTPSvc.Routes(r)
		orderHTTPSvc.Routes(r)

		errCh <- r.Run(":4141") // TODO: handle port on config
	}()

	<-quitCh
	log.Printf("server has been shutdown")
}
