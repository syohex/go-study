package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app01/models"
)

func getPersons(c *gin.Context) {
	persons, err := models.GetPersons(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")

	person, err := models.GetPersonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if person == nil {
		msg := fmt.Sprintf("Person ID=%s is not found", id)
		c.JSON(http.StatusNotFound, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func addPerson(c *gin.Context) {
	var p models.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.AddPerson(&p); err == nil {
		c.JSON(http.StatusCreated, gin.H{"person": p})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func updatePerson(c *gin.Context) {
	var p models.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdatePerson(&p, int(id)); err == nil {
		c.JSON(http.StatusCreated, gin.H{"person": p})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func deletePerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DeletePerson(int(id)); err == nil {
		c.JSON(http.StatusNoContent, gin.H{"success": "ok"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func options(c *gin.Context) {
	c.Header("Allow", "GET,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.String(200, "")
}

func main() {
	if err := models.ConnectDatabase(); err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		v1.GET("person/:id", getPersonById)
		v1.POST("person", addPerson)
		v1.PUT("person/:id", updatePerson)
		v1.DELETE("person/:id", deletePerson)
		v1.OPTIONS("person", options)
	}

	r.Run()
}
