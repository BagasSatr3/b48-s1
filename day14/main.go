package main

import (
	"context"
	// "fmt"
	"html/template"
	"net/http"
	"personal/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id			int
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

func main() {
	e := echo.New()

	connection.DatabaseConnect()
	e.Static("/assets", "assets")
	// e.Static("/uploads", "uploads")

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
	dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT id, name, description, start_date, end_date, image FROM tb_project")

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Id ,&each.Name, &each.Description, &each.StartDate, &each.EndDate, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Image = ""
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

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT id, name, description, start_date, end_date, image FROM tb_project")

	if errProjects != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	for dataProjects.Next() {
		var each = Project{}

		err := dataProjects.Scan(&each.Id ,&each.Name, &each.Description, &each.StartDate, &each.EndDate, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		each.Image = ""
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
	idToInt, _ := strconv.Atoi(Id)

	projectDetail := Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT name, start_date, end_date, description, image FROM tb_project WHERE id=$1", idToInt).Scan(&projectDetail.Name, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Description, &projectDetail.Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectDetail.Start = projectDetail.StartDate.Format("2006-01-02")
	projectDetail.End = projectDetail.EndDate.Format("2006-01-02")
	projectDetail = Project{
		Name : projectDetail.Name,
		Description : projectDetail.Description,
		Image: projectDetail.Image,
		Start: projectDetail.Start,
		End: projectDetail.End,
	}

	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"Project": projectDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func formEdit(c echo.Context) error {
	Id := c.Param("id")
	
	idToInt, _ := strconv.Atoi(Id)

	projectEdit := Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT name, start_date, end_date, description, image FROM tb_project WHERE id=$1", idToInt).Scan(&projectEdit.Name, &projectEdit.StartDate, &projectEdit.EndDate, &projectEdit.Description, &projectEdit.Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	projectEdit.Start = projectEdit.StartDate.Format("2006-01-02")
	projectEdit.End = projectEdit.EndDate.Format("2006-01-02")
	projectEdit = Project{
		Name : projectEdit.Name,
		Description : projectEdit.Description,
		Image: projectEdit.Image,
		Start: projectEdit.Start,
		End: projectEdit.End,
	}
	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"Project": projectEdit,
	}

	tmpl, err := template.ParseFiles("views/form-edit.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
	// nodejs := c.FormValue("nodejs")
	// vuejs := c.FormValue("vuejs")
	// reactjs := c.FormValue("reactjs")
	// javascript := c.FormValue("javascript")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")

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

	// image, ok := c.Get("dataFile").(string)
	// if !ok  {
	// 	fmt.Println("error")
	// }

	image := "defauld.jpg"
	_ , err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (name, start_date, end_date, description, image) VALUES ($1, $2, $3, $4, $5)", name, startDate, endDate, description, image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func editProject(c echo.Context) error {
	Id := c.Param("id")
	idToInt, _ := strconv.Atoi(Id)
	name := c.FormValue("name")
	description := c.FormValue("description")
	// nodejs := c.FormValue("nodejs")
	// reactjs := c.FormValue("reactjs")
	// vuejs := c.FormValue("vuejs")
	// javascript := c.FormValue("javascript")

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
	
	
	image := "default.jpg"
	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name=$1, start_date=$2, end_date=$3, image=$4, description=$5 WHERE id=$6", name, startDate, endDate, image, description, idToInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

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