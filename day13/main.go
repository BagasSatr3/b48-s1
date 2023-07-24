package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personal/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	// Id			int
	Name		string
	Description	string
	StartDate	time.Time
	EndDate		time.Time
	Start		string
	End		string
	Nodejs		string
	Vuejs		string
	Reactjs		string
	Javascript	string
	Technologies []string
	Image 		string
	Duration	string
}

var dataProjects = []Project{
	{
		Name			:	"Dumbways Mobile App 1",
		Description		:	"Lorem ipsum dolor sit amet",
	},
	{
		Name			:	"App 2",
		Description		:	"Lorem ipsum dolor sit ametaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	},
	{
		Name			:	"Dumbways Mobile App 3",
		Description		:	"Lorem ipsum dolor sit ametbbbbbbbbbbbbbbbbbbb",
	},
}

func main() {
	e := echo.New()

	connection.DatabaseConnect()
	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/home", home)
	e.GET("/project", project)
	e.GET("/form-edit/:id", formEdit)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/contact", contact)
	e.GET("/contact-alt", contactAlt)
	e.GET("/testimonial", testimonials)

	e.POST("/edit-project/:id", editProject)
	e.POST("/delete-project/:id",deleteProject)
	e.POST("/add-project", addProject)

	e.Logger.Fatal(e.Start("localhost:6611"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html");

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT name, description, start_date, end_date, image FROM tb_project")

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Name, &each.Description, &each.StartDate, &each.EndDate, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resultProject = append(resultProject, each)
	}
	// fmt.Println(resultProject)
	data := map[string]interface{}{
		// "Projects" : dataProjects,
		"Projects" : resultProject,
		}
	return tmpl.Execute(c.Response(), data)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT name, description, start_date, end_date, image FROM tb_project")

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Name, &each.Description, &each.StartDate, &each.EndDate, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Start = each.StartDate.Format("2006-01-02")
		each.End = each.EndDate.Format("2006-01-02")
		each.Duration = duration(each.StartDate,each.EndDate)
		resultProject = append(resultProject, each)
	}
	// fmt.Println(resultProject)
	data := map[string]interface{}{
		// "Projects" : dataProjects,
		"Projects" : resultProject,
		}
	return tmpl.Execute(c.Response(), data)
}

func projectDetail(c echo.Context) error {
	Id := c.Param("id")
	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(Id)

	projectDetail := Project{}

	for index, data := range dataProjects {
		if index == idToInt {
			projectDetail = Project{
				Name:        data.Name,
				Description: data.Description,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
			}
		}
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"Project": projectDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func formEdit(c echo.Context) error {
	Id := c.Param("id")
	tmpl, err := template.ParseFiles("views/form-edit.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(Id)

	projectEdit := Project{}

	for index, data := range dataProjects {
		if index == idToInt {
			projectEdit = Project{
				Name		: data.Name,
				Description	: data.Description,
				StartDate	: data.StartDate,
				EndDate		: data.EndDate,
				Duration	: data.Duration,
			}
		}
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"Project": projectEdit,
	}

	return tmpl.Execute(c.Response(), data)
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
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	vuejs := c.FormValue("vuejs")
	reactjs := c.FormValue("reactjs")
	javascript := c.FormValue("javascript")

	// var nodejs bool
	// if c.FormValue("nodejs") == "nodejs" {
	// 	nodejs = true
	// }
	// var reactjs bool
	// if c.FormValue("reactjs") == "reactjs" {
	// 	reactjs = true
	// }
	// var vuejs bool
	// if c.FormValue("vuejs") == "vuejs" {
	// 	vuejs = true
	// }
	// var javascript bool
	// if c.FormValue("javascript") == "javascript" {
	// 	javascript = true
	// }

	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	
	layout := "2006-01-02"
	dateString1 := startDate
	dateString2 := endDate

	date1, err := time.Parse(layout, dateString1)
	if err != nil {
		fmt.Println("Error parsing date string 1: ", err)
		return nil
	}
	date2, err := time.Parse(layout, dateString2)
	if err != nil {
		fmt.Println("Error parsing date string 2: ", err)
		return nil
	}
	duration := date2.Sub(date1)
	days := int(duration.Hours() / 24)
	days1 := string(days)

	newProject := Project{
		Name: name,
		Description: description,
		StartDate: date1,
		EndDate: date2,
		Nodejs: nodejs,
		Reactjs: reactjs,
		Vuejs: vuejs,
		Javascript: javascript,
		Duration: days1,
	}

	dataProjects = append(dataProjects, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func editProject(c echo.Context) error {
	Id := c.Param("id")
	idToInt, _ := strconv.Atoi(Id)
	name := c.FormValue("name")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	reactjs := c.FormValue("reactjs")
	vuejs := c.FormValue("vuejs")
	javascript := c.FormValue("javascript")

	// var nodejs bool
	// if c.FormValue("nodejs") == "nodejs" {
	// 	nodejs = true
	// }
	// var reactjs bool
	// if c.FormValue("reactjs") == "reactjs" {
	// 	reactjs = true
	// }
	// var vuejs bool
	// if c.FormValue("vuejs") == "vuejs" {
	// 	vuejs = true
	// }
	// var javascript bool
	// if c.FormValue("javascript") == "javascript" {
	// 	javascript = true
	// }

	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	
	layout := "2006-01-02"
	dateString1 := startDate
	dateString2 := endDate

	date1, err := time.Parse(layout, dateString1)
	if err != nil {
		fmt.Println("Error parsing date string 1: ", err)
		return nil
	}
	date2, err := time.Parse(layout, dateString2)
	if err != nil {
		fmt.Println("Error parsing date string 2: ", err)
		return nil
	}
	duration := date2.Sub(date1)
	days := int(duration.Hours() / 24)
	days1 := string(days)

	editProject := Project{
		Name: name,
		Description: description,
		StartDate: date1,
		EndDate: date2,
		Nodejs: nodejs,
		Reactjs: reactjs,
		Vuejs: vuejs,
		Javascript: javascript,
		Duration: days1,
	}

	dataProjects[idToInt] = editProject

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func duration (startDate time.Time, endDate time.Time) string {
	duration := endDate.Sub(startDate)
	days := int(duration.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " Year"
	}
	if months > 0 {
		return strconv.Itoa(months) + " Months"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " Weeks"
	}
	return strconv.Itoa(days) + " days"
}