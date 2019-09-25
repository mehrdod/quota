package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"alif/quota/api/config"
	"alif/quota/api/route"

	"github.com/gin-gonic/gin"
)

func Start(srvConfig *config.Service) {
	gin.SetMode(config.Peek().Service.ReleaseMode)

	router := gin.Default()

	route.AddRoutes(router, srvConfig.ApiVersion)

	srv := &http.Server{
		Addr:    srvConfig.Port,
		Handler: router,
	}

	log.Println(fmt.Sprintf("✓ Service: %s is begin running(PID: %d)", srvConfig.Name, os.Getpid()))

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Println(fmt.Sprintf("✓ Service: %s is started", srvConfig.Name))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("Shutdown service  %s...", srvConfig.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
