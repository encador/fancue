package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/encador/fancue/internal/components"
)

type config struct {
	address string
	port    int
}

func main() {
	cnf := config{}
	flag.StringVar(&cnf.address, "address", "localhost", "IP-address on which the application runs")
	flag.IntVar(&cnf.port, "port", 8080, "Port on which the application runs")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		components.Base().Render(r.Context(), w)
	})

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.address, cnf.port),
		Handler: mux,
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
