package controllers

import (
	"TechnikalTes/src/models"
	"TechnikalTes/src/services"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ListController struct {
	listService services.ListService
}

func (controller ListController) Create(c echo.Context) error{
	// Handle File
	file, errFile := c.FormFile("file")
	if errFile != nil {
			log.Println("Error Get File: ", errFile)
	}


	// Save the file
	var fileName string
	if file != nil {
		fileName = file.Filename
		src, err := file.Open()
		if err != nil {
			log.Println("Error Opening File: ", err)
			return err
		}
		defer src.Close()

		dst, err := os.Create(fmt.Sprintf("./src/public/uploads/%s", file.Filename))
		if err != nil {
			log.Println("Error Creating Destination File: ", err)
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			log.Println("Error Copying File: ", err)
			return err
		}
	} else {
		log.Println("File is nil")
	}
	
	

	// Retrieve form values explicitly
	title := c.FormValue("title")
	description := c.FormValue("description")

	// Validate required fields
	if title == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "title is required")
	}
	if description == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "description is required")
	}

	// Create the List model
	result := controller.listService.Create(models.List{
			Title:       title,
			Description: description,
			File:        fileName,
	})

	return c.JSON(http.StatusOK, result)

}

func (controller ListController) Update(c echo.Context) error{
	// Handle File
	file, errFile := c.FormFile("file")
	if errFile != nil {
			log.Println("Error Get File: ", errFile)
	}


	// Save the file
	var fileName string
	if file != nil {
		fileName = file.Filename
		src, err := file.Open()
		if err != nil {
			log.Println("Error Opening File: ", err)
			return err
		}
		defer src.Close()

		dst, err := os.Create(fmt.Sprintf("./src/public/uploads/%s", file.Filename))
		if err != nil {
			log.Println("Error Creating Destination File: ", err)
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			log.Println("Error Copying File: ", err)
			return err
		}
	} else {
		fileName = ""
		log.Println("File is nil")
	}
	
	

	// Retrieve form values explicitly
	title := c.FormValue("title")
	description := c.FormValue("description")

	// Validate required fields
	if title == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "title is required")
	}
	if description == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "description is required")
	}

	id,_ := strconv.Atoi(c.Param("id"))

	result := controller.listService.Update(id, models.List{
		Title:       title,
		Description: description,
		File:        fileName,
	})

	return c.JSON(http.StatusOK, result)

}

func (controller ListController) Delete(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	result := controller.listService.Delete(id)
	return c.JSON(http.StatusOK, result)
}

func (controller ListController) GetAll(c echo.Context) error{
	result := controller.listService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller ListController) GetById(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	result := controller.listService.GetById(id)
	return c.JSON(http.StatusOK, result)
}

func NewListController(db *gorm.DB) ListController {
	return ListController{
		listService: services.NewListService(db),
	} 
}