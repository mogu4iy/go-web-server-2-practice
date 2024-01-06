package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/modules/health"
	"go-web-server-2-practice/modules/test"
	"go-web-server-2-practice/modules/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	PORT string
	VERSION string
	DSN string
)
var V1VERSION = "v1"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not loaded")
		err = nil
	}

	if Port, ok := os.LookupEnv("PORT"); ok {
		PORT = Port
	} else {
		log.Fatal("PORT env is absent")
	}
	if Version, ok := os.LookupEnv("VERSION"); ok {
		VERSION = Version
	}else {
		log.Fatal("VERSION env is absent")
	}
	if dsn, ok := os.LookupEnv("DATA_SOURCE_NAME"); ok {
		DSN = dsn
	} else {
		log.Fatal("DATA_SOURCE_NAME env is absent")
	}

	db := &core.DB{
		DSN: DSN,
	}
	err = db.Init()
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	app := &core.App{
		Version: VERSION,
		V1Version: V1VERSION,
	}
	initModules(app, db)
	err = app.Init()
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}
	
	srv := &http.Server{
		Addr: ":" + PORT,
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

func initModules(app *core.App, db *core.DB) {
	tM := test.New(app)
	app.AddModule(tM)
	hM := health.New(app, tM.Service)
	app.AddModule(hM)
	uM := user.New(app, db)
	app.AddModule(uM)
}