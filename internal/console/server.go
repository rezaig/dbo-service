package console

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/db"
	"github.com/rezaig/dbo-service/internal/delivery/httpsvc"
	"github.com/rezaig/dbo-service/internal/repository"
	"github.com/rezaig/dbo-service/internal/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	dbConn := db.InitMySQLConn()
	defer func() {
		_ = dbConn.Close()
	}()

	// Initialize repositories
	authRepo := repository.NewAuthRepository(dbConn)
	customerRepo := repository.NewCustomerRepository(dbConn)
	orderRepo := repository.NewOrderRepository(dbConn)

	// Initialize usecases
	authUsecase := usecase.NewAuthUsecase(authRepo)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)

	// Initialize delivery HTTP services
	authHTTPSvc := httpsvc.NewAuthHTTPService(authUsecase)
	customerHTTPSvc := httpsvc.NewCustomerHTTPService(customerUsecase)
	orderHTTPSvc := httpsvc.NewOrderHTTPService(orderUsecase)

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		// Handle graceful shutdown
		for {
			select {
			case <-sigCh:
				log.Info("server is shutting down because of interrupt signal")
				quitCh <- true
			case e := <-errCh:
				log.Infof("server is shutting down because of error, error: %v\n", e)
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
	log.Info("server has been shutdown")
}
