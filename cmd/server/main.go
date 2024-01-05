package main

import (
	"context"
	"errors"
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/modules/health"
	"go-web-server-2-practice/modules/test"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	app := &core.App{}
	initModules(app)
	err := app.Init()
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}
	
	srv := &http.Server{
		Addr: ":5000",
		Handler: app.Engine,
	}
	
	go func() {
		log.Printf("server is listening on %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}

func initModules(app *core.App) {
	tM := test.New(app)
	app.AddModule(tM)
	hM := health.New(app, tM.Service)
	app.AddModule(hM)
}