package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/home", home)
	e.GET("/project", project)
	e.GET("/project-detail", projectDetail)
	e.GET("/contact", contact)
	e.GET("/contact-alt", contactAlt)
	e.GET("/testimonial", testimonials)

	e.POST("/add-project", addProject)

	e.Logger.Fatal(e.Start("localhost:6611"))
}
// func helloWorld(c echo.Context) error {
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"name":    "123",
// 		"address": "3",
// 	})
// }

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html");

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contactAlt(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact-alt.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	name := c.FormValue("name")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")
	// nodejs := c.FormValue("nodejs")
	// reactjs := c.FormValue("reactjs")
	// vuejs := c.FormValue("vuejs")
	// javascript := c.FormValue("javascript")

	var nodejs bool
	if c.FormValue("nodejs") == "nodejs" {
		nodejs = true
	}
	var reactjs bool
	if c.FormValue("reactjs") == "reactjs" {
		reactjs = true
	}
	var vuejs bool
	if c.FormValue("vuejs") == "vuejs" {
		vuejs = true
	}
	var javascript bool
	if c.FormValue("javascript") == "javascript" {
		javascript = true
	}

	fmt.Println("name: ", name)
	fmt.Println("startDate: ", startDate)
	fmt.Println("endDate: ", endDate)
	fmt.Println("description: ", description)
	fmt.Println("nodejs: ", nodejs)
	fmt.Println("reactjs: ", reactjs)
	fmt.Println("vuejs: ", vuejs)
	fmt.Println("javascript: ", javascript)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}