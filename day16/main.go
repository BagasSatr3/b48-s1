package main

import (
	"context"
	// "testing/quick"
	// "fmt"
	"os"
	"log"
	"html/template"
	"net/http"
	"personal/connection"
	"personal/middleware"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id			int
	Name		string
	Description	string
	StartDate	time.Time
	EndDate		time.Time
	Start		string
	End			string
	Image 		string
	Duration	string
	Author      string
	Username 	string
}

type User struct {
	Id			int
	Name 		string
	Email		string
	Password	string
	HashedPassword	string
}

type UserLoginSession struct {
	IsLogin bool
	Name 	string
}

var userLoginSession = UserLoginSession{}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	e.GET("/", home)
	e.GET("/home", home)
	e.GET("/project", project)
	e.GET("/form-edit/:id", formEdit)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/contact", contact)
	e.GET("/contact-alt", contactAlt)
	e.GET("/testimonial", testimonials)
	e.GET("/form-login", formLogin)
	e.GET("/form-register", formRegister)

	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/logout", logout)
	e.POST("/edit-project/:id", middleware.UploadFiles(editProject))
	e.POST("/delete-project/:id",deleteProject)
	e.POST("/add-project", middleware.UploadFiles(addProject))

	e.Logger.Fatal(e.Start("localhost:6611"))
}

func formLogin(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html");

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
		dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, image, author FROM tb_project")

		if errProjects != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		for dataProjects.Next() {
			var each = Project{}

			err := dataProjects.Scan(&each.Id ,&each.Name, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Author)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			each.Start = each.StartDate.Format("2006-01-02")
			each.End = each.EndDate.Format("2006-01-02")
			each.Duration = duration(each.StartDate,each.EndDate)
			resultProject = append(resultProject, each)
		}
	} else {
		userLoginSession.IsLogin = true
		id := sess.Values["id"].(int)
		dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_project.name, tb_project.start_date, tb_project.end_date, tb_project.description, tb_project.image, tb_user.name AS author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_user.id=$1 ORDER BY tb_project.id DESC", id)
		if errProjects != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		for dataProjects.Next() {
			var each = Project{}
	
			err := dataProjects.Scan(&each.Id ,&each.Name, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Author)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			// each.Image = ""
			each.Start = each.StartDate.Format("2006-01-02")
			each.End = each.EndDate.Format("2006-01-02")
			each.Duration = duration(each.StartDate,each.EndDate)
			resultProject = append(resultProject, each)
		}
	}

	flash := map[string]interface{}{
		"FlashMessage" : sess.Values["message"],
		"FlashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"UserLoginSession" : userLoginSession,
		"Flash" : flash,
		"Projects" : resultProject,
		}
	return tmpl.Execute(c.Response(), data)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
		dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, image, author FROM tb_project")

		if errProjects != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		for dataProjects.Next() {
			var each = Project{}

			err := dataProjects.Scan(&each.Id ,&each.Name, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Author)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			each.Start = each.StartDate.Format("2006-01-02")
			each.End = each.EndDate.Format("2006-01-02")
			each.Duration = duration(each.StartDate,each.EndDate)
			resultProject = append(resultProject, each)
		}
	} else {
		userLoginSession.IsLogin = true
		id := sess.Values["id"].(int)
		dataProjects, errProjects := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_project.name, tb_project.start_date, tb_project.end_date, tb_project.description, tb_project.image, tb_user.name AS author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_user.id=$1 ORDER BY tb_project.id DESC", id)
		if errProjects != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		for dataProjects.Next() {
			var each = Project{}
	
			err := dataProjects.Scan(&each.Id ,&each.Name, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Author)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			// each.Image = ""
			each.Start = each.StartDate.Format("2006-01-02")
			each.End = each.EndDate.Format("2006-01-02")
			each.Duration = duration(each.StartDate,each.EndDate)
			resultProject = append(resultProject, each)
		}
	}

	// fmt.Println(resultProject)
	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}
	// fmt.Println(resultProject)
	data := map[string]interface{}{
		"UserLoginSession" : userLoginSession,
		"Projects" : resultProject,
		}
	return tmpl.Execute(c.Response(), data)
}

func projectDetail(c echo.Context) error {
	Id := c.Param("id")
	idToInt, _ := strconv.Atoi(Id)

	projectDetail := Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT name, start_date, end_date, description, image, author FROM tb_project WHERE id=$1", idToInt).Scan(&projectDetail.Name, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Description, &projectDetail.Image, &projectDetail.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectDetail.Start = projectDetail.StartDate.Format("2006-01-02")
	projectDetail.End = projectDetail.EndDate.Format("2006-01-02")
	projectDetail.Duration = duration(projectDetail.StartDate,projectDetail.EndDate)
	projectDetail = Project{
		Name : projectDetail.Name,
		Description : projectDetail.Description,
		Image: projectDetail.Image,
		Start: projectDetail.Start,
		End: projectDetail.End,
		Author: projectDetail.Author,
		Duration: projectDetail.Duration,
	}

	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"UserLoginSession" : userLoginSession,
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

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   Id,
		"UserLoginSession" : userLoginSession,
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

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{ 
		"UserLoginSession" : userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func contactAlt(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact-alt.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{ 
		"UserLoginSession" : userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonials(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{ 
		"UserLoginSession" : userLoginSession,
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")

	image := c.Get("dataFile").(string)

	sess, _ := session.Get("session", c)
	author := sess.Values["id"].(int)
	authorName := sess.Values["name"].(string)

	_ , err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (name, start_date, end_date, description, image, author_id, author) VALUES ($1, $2, $3, $4, $5, $6, $7)", name, startDate, endDate, description, image, author, authorName)
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

	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	
	projectedit := Project{}

	erredit := connection.Conn.QueryRow(context.Background(), "SELECT image FROM tb_project WHERE id=$1", idToInt).Scan(&projectedit.Image)
	if erredit != nil {
		return c.JSON(http.StatusInternalServerError, erredit.Error())
	}

	e := os.Remove("G:/dumbways/b48-s1/day16/uploads/" + projectedit.Image)
	if e != nil {
        log.Fatal(e)
   	} 
	
	image := c.Get("dataFile").(string)
	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name=$1, start_date=$2, end_date=$3, image=$4, description=$5 WHERE id=$6", name, startDate, endDate, image, description, idToInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	projectDelete := Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT image FROM tb_project WHERE id=$1", idToInt).Scan(&projectDelete.Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e := os.Remove("G:/dumbways/b48-s1/day16/uploads/" + projectDelete.Image)
	if e != nil {
        log.Fatal(e)
   	} 

	_, errdel := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)
	if errdel != nil {
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

func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name ,email, passwordHash)

	if err != nil {
		return redirectWithMessage(c, "Register gagal!", false, "/form-register")
	}

	return redirectWithMessage(c, "Register berhasil!", true, "/form-login")
}

func login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := User{}
	
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if errPassword != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10000 // how long to be expired
	sess.Values["message"] = "Login success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Log In berhasil!", true, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Logout berhasil!", true, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}