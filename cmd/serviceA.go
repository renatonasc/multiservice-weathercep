package main

import (
	"context"
	"log"
	"renatonasc/multiservice-weathercep/internal/infra/web"
	"renatonasc/multiservice-weathercep/internal/infra/web/webserver"
	"renatonasc/multiservice-weathercep/pkg"

	"go.opentelemetry.io/otel"
)

func main() {

	shutdown, err := pkg.InitProvider("serviceA", "otel-collector:4317")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatal("failed to shutdown provider", err)
		}
	}()
	tracer := otel.Tracer("tracer")

	webserver := webserver.NewWebServer(":8081")

	cepHandler := web.NewCepHandlerPost(tracer)

	webserver.AddHandler("POST", "/weather/", cepHandler.GetWeatherByCep)

	webserver.Start()
}
