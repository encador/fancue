package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"embed"

	"github.com/encador/fancue/internal/handler"
	"github.com/encador/fancue/internal/middleware"
)

type config struct {
	address string
	port    int
}

//go:embed internal/static
var staticFiles embed.FS

func main() {
	cnf := config{}
	flag.StringVar(&cnf.address, "address", "localhost", "IP-address on which the application runs")
	flag.IntVar(&cnf.port, "port", 8080, "Port on which the application runs")
	flag.Parse()

	mux := http.NewServeMux()

	sub, _ := fs.Sub(staticFiles, "internal/static")
	mux.Handle("/static/", http.StripPrefix("/static/", middleware.Cache(http.FileServerFS(sub), 24)))

	h := handler.NewHandler()
	mux.Handle("/{$}", h.HomePage())
	mux.Handle("/test", h.TestPage())

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.address, cnf.port),
		Handler: middleware.Logger(mux),
	}

	go func() {
		fmt.Println("[LOG] Serving on " + srv.Addr)
		srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fmt.Println("\nShutting Down...")
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Shutdown Complete")

}
