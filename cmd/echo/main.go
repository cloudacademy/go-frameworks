package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Course struct {
	Title  string `json:"title"`
	Length int    `json:"length"`
}

func addCourse(c echo.Context) error {
	course := new(Course)
	if err := c.Bind(course); err != nil {
		return err
	}

	json := c.JSON(http.StatusCreated, course)

	fmt.Println("POST request...")
	fmt.Printf("payload JSON: %#v\n", course)

	c.Response().Header().Add("Content-Type", "application/json")
	return json
}

func getCourse(c echo.Context) error {
	id := c.Param("id")
	msg := fmt.Sprintf("GET request... id=%s\n", id)
	fmt.Println(msg)
	c.Response().Header().Add("Content-Type", "text/plain")
	return c.String(http.StatusOK, msg)
}

func updateCourse(c echo.Context) error {
	id := c.Param("id")
	msg := fmt.Sprintf("PUT request... id=%s\n", id)
	fmt.Println(msg)
	c.Response().Header().Add("Content-Type", "text/plain")
	return c.String(http.StatusOK, msg)
}

func deleteCourse(c echo.Context) error {
	id := c.Param("id")
	msg := fmt.Sprintf("DELETE request... id=%s\n", id)
	fmt.Println(msg)
	c.Response().Header().Add("Content-Type", "text/plain")
	return c.String(http.StatusOK, msg)
}

func main() {
	e := echo.New()

	e.GET("/cloudacademy", func(c echo.Context) error {
		return c.String(http.StatusOK, "CloudAcademy + Go + Echo!!\n")
	})

	e.POST("/cloudacademy/courses", addCourse)
	e.GET("/cloudacademy/courses/:id", getCourse)
	e.PUT("/cloudacademy/courses/:id", updateCourse)
	e.DELETE("/cloudacademy/courses/:id", deleteCourse)

	e.Logger.Fatal(e.Start(":8080"))
}
