package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	// Static Routes
	r.LoadHTMLGlob("/NORDIC_CODER/Gin_Go_WebApp/templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")

		timesOff, ok := TimesOff[id]
		if !ok {
			c.String(http.StatusNotFound, "404 - Page Note Found")
		}

		c.HTML(http.StatusOK, "vacation-overview.html", map[string]interface{}{
			"TimesOff": timesOff,
		})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", map[string]interface{}{
			"Employees": Employees,
		})
	})

	admin.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}

		employee, ok := Employees[id]
		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
			return
		}

		c.HTML(http.StatusOK, "admin-employee-edit.html", map[string]interface{}{
			"Employee": employee,
		})
	})

	admin.POST("/employee/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			pto, err := strconv.ParseFloat(c.PostForm("pto"), 32)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			startDate, err := time.Parse("2006-01-02", c.PostForm("startDate"))
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			var emp Employee
			emp.ID = 42
			emp.FirstName = c.PostForm("firstName")
			emp.LastName = c.PostForm("lastName")
			emp.Position = c.PostForm("position")
			emp.Status = "Active"
			emp.TotalPTO = float32(pto)
			emp.StartDate = startDate

			fmt.Println(emp)

			Employees["42"] = emp

			c.Redirect(http.StatusMovedPermanently, "admin/employees/42")
		}
	})

	// Static Files
	r.Static("/public", "/NORDIC_CODER/Gin_Go_WebApp/public")

	return r
}
