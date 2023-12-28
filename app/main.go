package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type temperature struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
	Scale string  `json:"scale"`
}

// Simple middleware to check for API key based on if the sqlite3 database exists or not
// Not secure, not well thought out, but works well enough for my application
func tokenAuthMIddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		token := c.GetHeader("X-Api-Key")

		// If no token provided, abort with 401 Unauthorized
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key required"})
			return
		}

		// Check if token is valid
		dbPath := fmt.Sprintf("./%s.db", token)
		if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key invalid"})
			return
		}

		// If all of the above checks pass, continue
		c.Next()
	}
}

func getTemperature(c *gin.Context) {
	token := c.GetHeader("X-Api-Key")
	dbPath := fmt.Sprintf("./%s.db", token)

	// Open the database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Query all rows in the temperatures table
	rows, err := db.Query("SELECT * FROM temperatures;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	// Parse rows into a slice of Temperature structs
	var temps []temperature
	for rows.Next() {
		var temp temperature
		err = rows.Scan(&temp.Id, &temp.Value, &temp.Scale)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, temps)
}

func putTemperature(c *gin.Context) {
	token := c.GetHeader("X-Api-Key")
	dbPath := fmt.Sprintf("./%s.db", token)

	// Open the database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Insert the new temperature into the database
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Bind the JSON body to a temperature struct
	var newTemperature temperature
	if err := c.BindJSON(&newTemperature); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Insert the new temperature into the database
	_, err = db.Exec("INSERT INTO temperatures(id, value, scale) VALUES(?, ?, ?);", newTemperature.Id, newTemperature.Value, newTemperature.Scale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the newly inserted temperature
	c.JSON(http.StatusCreated, newTemperature)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.Use(tokenAuthMIddleWare())
		v1.GET("/temperature", getTemperature)
		v1.PUT("/temperature", putTemperature)

		v1.GET("/LED", func(c *gin.Context) {c.JSON("LED"})})
	}

	router.Run(":8080")
}
