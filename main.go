package main

import (
	"github.com/Fabio4breu/opentelemetry-crud/config"
	"github.com/Fabio4breu/opentelemetry-crud/routes"
	"github.com/Fabio4breu/opentelemetry-crud/tracing" // <-- novo
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin" // <-- novo
)

func main() {
	// Conecta ao MongoDB
	config.Connect()

	// Inicializa o tracer do OpenTelemetry
	shutdown := tracing.InitTracer()
	defer shutdown()

	// Inicializa o servidor Gin com o middleware do OpenTelemetry
	r := gin.Default()
	r.Use(otelgin.Middleware("crud-api"))

	// Registra as rotas
	routes.SetupRoutes(r)

	// Inicia o servidor
	r.Run(":8080")
}
