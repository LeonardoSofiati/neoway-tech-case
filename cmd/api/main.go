package main

import (
	"context"
	"fmt"
	"log"
	"neoway_test/internal/domain/customer/service"
	"neoway_test/internal/infrastructure/api/handlers"
	databaseConfig "neoway_test/internal/infrastructure/database/config"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	usecaseCreate "neoway_test/internal/usecase/customer/create"
	usecaseDelete "neoway_test/internal/usecase/customer/delete"
	usecaseFind "neoway_test/internal/usecase/customer/find"
	usecaseList "neoway_test/internal/usecase/customer/list"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "neoway_test/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// @title           Neoway recruitment process tech test
// @version         1.0
// @description     Documentation for Neoway test API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	timeZone := os.Getenv("APP_TIMEZONE")
	if timeZone == "" {
		timeZone = "America/Sao_Paulo"
	}

	location, err := time.LoadLocation(timeZone)

	if err != nil {
		panic("Error loading time zone: " + err.Error())
	}

	time.Local = location

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := databaseConfig.NewDb()

	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Repositório
	customerRepo, err := databaseRepository.NewPostgresCustomerRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	createCustomersBulkService := service.NewParseTxtFileService()
	createCustomersService := service.NewParseService()

	createCustomerUsecase := usecaseCreate.NewCreateCustomerUseCase(customerRepo, createCustomersService)
	createCustomersBulkUsecase := usecaseCreate.NewCreateCustomersBulkUseCase(customerRepo, createCustomersBulkService)
	getCustomerByCpfUsecase := usecaseFind.NewGetCustomerByCpfUseCase(customerRepo)
	getCustomerByIdUsecase := usecaseFind.NewGetCustomerByIdUseCase(customerRepo)
	getCustomersListUsecase := usecaseList.NewGetCustomersListUseCase(customerRepo)
	deleteCustomersUsecase := usecaseDelete.NewDeleteCustomerUseCase(customerRepo)

	// Handlers HTTP
	customerHandler := handlers.NewCustomerHandler(
		getCustomersListUsecase,
		createCustomerUsecase,
		createCustomersBulkUsecase,
		getCustomerByCpfUsecase,
		getCustomerByIdUsecase,
		deleteCustomersUsecase,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	r.Route("/api/v1/customer", func(r chi.Router) {
		r.Post("/", handlers.HandlerError(customerHandler.CustomerPost))
		r.Post("/bulkCreation", handlers.HandlerError(customerHandler.CustomerPostBulk))
		r.Get("/", handlers.HandlerError(customerHandler.CustomerGet))
		r.Get("/getById/{id}", handlers.HandlerError(customerHandler.CustomerGetById))
		r.Get("/getByCpf/{cpf}", handlers.HandlerError(customerHandler.CustomerGetByCpf))
		r.Delete("/{id}", handlers.HandlerError(customerHandler.CustomerDelete))
	})

	server := &http.Server{Addr: ":8080", Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Shutting down gracefully...")
	if err := server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error during server shutdown: %w", err)
	}

	return nil
}
