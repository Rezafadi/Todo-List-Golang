package main

import (
	"TechnikalTes/config"
	"TechnikalTes/src/controllers"
	"TechnikalTes/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	db := config.InitDB()

	route := echo.New()

	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	apiV1 := route.Group("/api/v1")

	// list
	listController := controllers.NewListController(db)
	apiV1.GET("/lists", listController.GetAll)
	apiV1.GET("/list/:id", listController.GetById)
	apiV1.POST("/list", listController.Create)
	// apiV1.PUT("/list/:id", listController.Update)
	apiV1.DELETE("/list/:id", listController.Delete)

	// subList
	subListController := controllers.NewSubListController(db)
	apiV1.GET("/sublists", subListController.GetAll)
	apiV1.GET("/sublist/:id", subListController.GetById)
	apiV1.POST("/sublist", subListController.Create)
	apiV1.PUT("/sublist/:id", subListController.Update)
	apiV1.DELETE("/sublist/:id", subListController.Delete)

	Automigrate(db)

	route.Start(":8080")

}

func Automigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&models.List{},
		&models.SubList{},
	)
}
