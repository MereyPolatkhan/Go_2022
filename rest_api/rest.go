package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
}

var students []Student = []Student{
	{ID: "001", Name: "John Patrik", Faculty: "FIT"},
	{ID: "002", Name: "Michael Jackson", Faculty: "BS"},
	{ID: "003", Name: "Stive Jobs", Faculty: "MCM"},
}

func get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func post(c *gin.Context) {
	var newStudent Student

	err := c.BindJSON(&newStudent)
	if err != nil {
		fmt.Print("ERROR ---->  :: ", err, "\n")
		// log.Fatal(err)
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("id %s", id)

	for _, a := range students {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "student is not found to GET")
}

func deleteByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("id %s", id)

	for i, a := range students {
		if a.ID == id {
			students = append(students[0:i], students[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "student is not found to DELETE")
}

func main() {
	r := gin.Default()
	r.GET("/students", get)
	r.POST("/post_students", post)
	r.GET("/students/:id", getByID)
	r.DELETE("/delete_students/:id", deleteByID)

	err := r.Run("localhost:8080")
	if err != nil {
		fmt.Print("ERROR ---->  :: ", err, "\n")
		// log.Fatal(err)~
	}

}
