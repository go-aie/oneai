package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/RussellLuo/appx"
	"github.com/RussellLuo/structool"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/yaml.v3"
)

func init() {
	appx.MustRegister(appx.New("oneai", new(OneAI)))
}

type OneAI struct {
	Addr string `structool:"addr"`

	server *http.Server
	router chi.Router
}

func (o *OneAI) Router() chi.Router {
	return o.router
}

func (o *OneAI) Init(ctx appx.Context) error {
	if err := structool.New().Decode(ctx.Config(), o); err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	o.server = &http.Server{
		Addr:    o.Addr,
		Handler: r,
	}
	o.router = r

	return nil
}

func (o *OneAI) Start(ctx context.Context) error {
	walkFunc := func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		log.Printf("Mounted route: %s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(o.router, walkFunc); err != nil {
		log.Printf("walk routes err: %v\n", err)
	}

	log.Printf("Starting HTTP server listening on %s\n", o.server.Addr)

	go o.server.ListenAndServe() // nolint:errcheck
	return nil

}

func (o *OneAI) Stop(ctx context.Context) error {
	log.Println("Stopping HTTP server")
	return o.server.Shutdown(ctx)
}

func Main() {
	var path string
	flag.StringVar(&path, "config", "./config.yaml", "path to the YAML config")
	flag.Parse()

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %s\n", path)
	}

	var config map[string]interface{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("failed to unmarshal config: %v\n", err)
	}

	appx.SetOptions(&appx.Options{
		ErrorHandler: func(err error) {
			log.Fatalf("error during stopping or uninstalling: %v\n", err)
		},
		AppConfigs: config,
	})

	ctx := context.Background()
	if err := appx.Install(ctx); err != nil {
		log.Fatalf("failed to install apps: %v\n", err)
	}
	defer appx.Uninstall()

	// Run the HTTP server.
	sig, err := appx.Run()
	if err != nil {
		log.Fatalf("failed to run apps: %v\n", err)
	}
	log.Printf("terminated since %s\n", sig.String())
}
