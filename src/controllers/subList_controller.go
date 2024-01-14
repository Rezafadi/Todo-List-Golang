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

type SubListController struct {
	subListService services.SubListService
}

func (controller SubListController) Create(c echo.Context) error {
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
	list_id := c.FormValue("list_id")

	// Convert list_id to int
	list_id_int, err := strconv.Atoi(list_id)
	if err != nil {
		log.Println("Error Converting list_id to int: ", err)
		return err
	}

	// Validate required fields
	if title == "" || description == "" || list_id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title, description and list_id are required")
	}

	result := controller.subListService.Create(models.SubList{
		Title:       title,
		Description: description,
		File:        fileName,
		ListID:				list_id_int,
	})

	return c.JSON(http.StatusOK, result)
}

func (controller SubListController) Update(c echo.Context) error {
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
	if title == "" || description == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title and description are required")
	}

	id,_ := strconv.Atoi(c.Param("id"))

	result := controller.subListService.Update(id, models.SubList{
		Title:       title,
		Description: description,
		File:        fileName,
	})

	return c.JSON(http.StatusOK, result)
}

func (controller SubListController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := controller.subListService.Delete(id)
	return c.JSON(http.StatusOK, result)
}

func (controller SubListController) GetAll(c echo.Context) error {
	result := controller.subListService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller SubListController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := controller.subListService.GetById(id)
	return c.JSON(http.StatusOK, result)
}

func NewSubListController(db *gorm.DB) SubListController {
	return SubListController{
		subListService: services.NewSubListService(db),
	}
}