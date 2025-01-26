package main

import (
	"fmt"
	"log"
	"meditrack/entities"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/JezzDiego/meditrack/docs"
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

	e := echo.New()

	// Swagger Docs
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "templates/assets",
		Browse: true,
		HTML5:  true,
		Index:  "index.html",
	}))
	e.Static("/static", "templates/assets")

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/doc", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	outerAPIHandler := entities.NewOuterAPIHandler(os.Getenv("OUTER_API_URL"), os.Getenv("OUTER_API_TOKEN"), os.Getenv("X-Cosmos-Token"))
	apiGroup := e.Group("/v1")

	apiGroup.GET("/ncm", entities.GetAllNCMs)
	apiGroup.GET("/medicine", entities.GetAllMedicines)

	apiGroup.GET("/ncm/:id", entities.GetNCMById)
	apiGroup.GET("/medicine/:id", entities.GetMedicineById)

	apiGroup.GET("/medicine/gtin/:gtin", outerAPIHandler.GetMedicineByGtin)

	apiGroup.POST("/ncm", entities.PostNCM)
	apiGroup.POST("/medicine", entities.PostMedicine)

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
