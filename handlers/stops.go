package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mawinkhumdee/flowly-project/backend/database"
	"github.com/mawinkhumdee/flowly-project/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Seed dummy data for testing
func SeedStops(c *gin.Context) {
	collection := database.GetCollection("stops")
	tripCollection := database.GetCollection("trips")

	// Check if already seeded
	count, _ := tripCollection.CountDocuments(context.Background(), bson.M{})
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Already seeded"})
		return
	}

	tripID := primitive.NewObjectID()
	trip := models.Trip{
		ID:     tripID,
		Title:  "Tokyo Trip 🇯🇵",
		Status: "In Progress",
	}
	tripCollection.InsertOne(context.Background(), trip)

	stops := []interface{}{
		models.Stop{
			ID:          primitive.NewObjectID(),
			TripID:      tripID,
			Day:         1,
			Time:        "9:00 AM",
			Name:        "Hotel Gracery",
			Status:      "Check-out complete",
			IsCompleted: true,
			OrderIndex:  1,
			Lat:         35.6961,
			Lng:         139.7024,
			Transit:     &models.Transit{Mode: "directions_walk", Time: "15 min walk"},
		},
		models.Stop{
			ID:          primitive.NewObjectID(),
			TripID:      tripID,
			Day:         1,
			Time:        "9:15 AM",
			Name:        "Blue Bottle Coffee",
			Status:      "Cafe • Arrived",
			IsCompleted: false,
			IsActive:    true,
			OrderIndex:  2,
			Lat:         35.6946,
			Lng:         139.7001,
			Actions:     []string{"Menu", "Notes"},
			Transit:     &models.Transit{Mode: "train", Time: "12 min train", Color: "bg-accent/10"},
		},
	}

	_, err := collection.InsertMany(context.Background(), stops)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Seeded dummy data successfully"})
}

func GetStops(c *gin.Context) {
	tripIDStr := c.Query("tripId")
	userIDStr := c.Query("userId")

	filter := bson.M{}
	if tripIDStr != "" {
		tripID, _ := primitive.ObjectIDFromHex(tripIDStr)
		filter["trip_id"] = tripID
	}
	if userIDStr != "" {
		userID, _ := primitive.ObjectIDFromHex(userIDStr)
		filter["user_id"] = userID
	}

	collection := database.GetCollection("stops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "orderIndex", Value: 1}})
	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	stops := []models.Stop{}
	if err = cursor.All(ctx, &stops); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stops)
}

func UpdateStop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Stop ID"})
		return
	}

	var updatedStop models.Stop
	if err := c.BindJSON(&updatedStop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("stops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        updatedStop.Name,
			"time":        updatedStop.Time,
			"status":      updatedStop.Status,
			"category":    updatedStop.Category,
			"day":         updatedStop.Day,
			"orderIndex":  updatedStop.OrderIndex,
			"isCompleted": updatedStop.IsCompleted,
			"transit":     updatedStop.Transit,
			"lat":         updatedStop.Lat,
			"lng":         updatedStop.Lng,
			"isActive":    updatedStop.IsActive,
			"image":       updatedStop.Image,
			"trip_id":     updatedStop.TripID,
			"user_id":     updatedStop.UserID,
			"city":        updatedStop.City,
			"country":     updatedStop.Country,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update successful"})
}

func CreateStop(c *gin.Context) {
	var stop models.Stop
	if err := c.BindJSON(&stop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if stop.ID.IsZero() {
		stop.ID = primitive.NewObjectID()
	}
	if stop.Day == 0 {
		stop.Day = 1
	}

	collection := database.GetCollection("stops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stop)
}

func DeleteStop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Stop ID"})
		return
	}

	collection := database.GetCollection("stops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete successful"})
}

func ReorderStops(c *gin.Context) {
	var payload []models.Stop
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	collection := database.GetCollection("stops")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, stop := range payload {
		id := stop.ID
		update := bson.M{
			"$set": bson.M{
				"orderIndex": stop.OrderIndex,
				"day":        stop.Day,
				"transit":    stop.Transit,
			},
		}
		_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stop sequence"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sequence updated successfully"})
}
