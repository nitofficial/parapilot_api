package handlers

import (
	"net/http"
	"parapilot/config"
	"parapilot/services"

	"github.com/gin-gonic/gin"
)

// MongoDB connection settings
const (
	dbName         = "parapilotdb"
	collectionName = "passages"
)

func PassagesHandler(c *gin.Context) {
	// Create MongoDB client
	client, err := config.GetMongoDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Access the MongoDB database and collection
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	// Aggregate passages
	passages, err := services.AggregatePassages(collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the JSON response
	c.JSON(http.StatusOK, passages)
}
