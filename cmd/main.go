package main

import (
	"fmt"
	"log"
	"meditrack/controller"
	"meditrack/database"
	_ "meditrack/docs"
	"meditrack/handlers"
	"meditrack/repository"
	"meditrack/usecase"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Meditrack API
// @version		1.0
// @description	API para consulta de itens por codigo de barras.
// @contact.name	Meditrack Team
//
//	@contact.url	localhost:3333
func main() {
	godotenv.Load()

	db, err := database.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Swagger Docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	//Handlers
	outerAPIHandler := handlers.OuterAPIHandler{
		OuterAPIURL:        os.Getenv("OUTER_API_URL"),
		OuterAPIToken:      os.Getenv("OUTER_API_TOKEN"),
		OuterAPIAuthHeader: os.Getenv("OUTER_API_AUTH_HEADER"),
	}

	// Controllers
	ProductController := controller.NewProductController(
		usecase.NewProductUsecase(
			repository.NewProductRepository(
				db,
				&outerAPIHandler,
			),
		),
	)
	NCMController := controller.NewNCMController(usecase.NewNCMUsecase(repository.NewNCMRepository(db)))

	// Routes
	apiGroup := e.Group("/v1")

	apiGroup.GET("/products", ProductController.GetAllProducts)
	apiGroup.GET("/product/:id", ProductController.GetProductById)
	apiGroup.GET("/product/gtin/:gtin", ProductController.GetProductByGtin)
	apiGroup.POST("/product", ProductController.CreateProduct)

	apiGroup.GET("/ncms", NCMController.GetAllNCM)
	apiGroup.GET("/ncm/:code", NCMController.GetNCMByCode)
	apiGroup.POST("/ncm", NCMController.CreateNCM)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
