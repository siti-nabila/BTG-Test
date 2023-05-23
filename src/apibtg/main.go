package main

import (
	controllers "BTG-Test/src/apibtg/controllers"
	repos "BTG-Test/src/apibtg/repositories"
	r "BTG-Test/src/apibtg/routes"
	services "BTG-Test/src/apibtg/services"

	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	Conn "BTG-Test/src/apibtg/database"

	"github.com/rs/cors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

func main() {
	// Init DB Connection handler
	if err := Conn.InitConnection(); err != nil {
		logger.Println(err)
	}
	defer Conn.CloseConnection()

	// Setting config
	port := ":" + viper.GetString("port")
	host := viper.GetString("host")

	// Create route
	router := r.CreateRouter(host, port)
	// r.InitRouter(router)

	// Init Repositories
	commonRepo := repos.CreateCommonRepositoryImpl()
	natRepo := repos.CreateNationalityRepositoryImpl(commonRepo)
	cstRepo := repos.CreateCustomerRepositoryImpl(commonRepo)

	// Init Services
	natService := services.CreateNationalityServiceImpl(natRepo)
	cstService := services.CreateCustomerServiceImpl(cstRepo)
	// Init Controller
	controllers.CreateNationalityController(router, natService)
	controllers.CreateCustomerController(router, cstService)
	c := cors.AllowAll()
	handler := c.Handler(router)
	subhost := host[7:]

	newHandler := http.TimeoutHandler(handler, 75*time.Second, "Timeout!")
	server := &http.Server{
		Addr:         port,
		Handler:      newHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 75 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at ", subhost, port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", port, err)
	}

	<-done
	logger.Println("Server stopped")
}
