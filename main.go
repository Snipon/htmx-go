package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type Product struct {
	SKU   string `bson:"sku" json:"sku"`
	Name  string `bson:"name" json:"name"`
	Image string `bson:"image" json:"image"`
}

type LineItem struct {
	Product  Product `bson:"product" json:"product"`
	Quantity int     `bson:"quantity" json:"quantity"`
}

type Cart struct {
	LineItems []LineItem `bson:"lineItems" json:"lineItems"`
	ID        string     `bson:"_id" json:"id"`
	CreatedAt string     `bson:"createdAt" json:"createdAt"`
}

func Connection() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return client
}

func getCart(c *gin.Context) {
	id := c.Param("id")
	var cnx = Connection()
	collection := cnx.Database("store").Collection("carts")

	var cart Cart
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, cart)

}

func addToCart(c *gin.Context) {
	id := c.Param("id")

	var cnx = Connection()
	collection := cnx.Database("store").Collection("carts")

	var lineItem LineItem
	var cart Cart

	if err := c.BindJSON(&lineItem); err != nil {
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.D{{"lineItems", lineItem}}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&cart)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, cart)
}

func createCart(c *gin.Context) {
	var id = uuid.New()
	var cart Cart = Cart{
		ID:        id.String(),
		LineItems: []LineItem{},
		CreatedAt: time.Now().String(),
	}

	var cnx = Connection()
	collection := cnx.Database("store").Collection("carts")
	_, err := collection.InsertOne(context.TODO(), cart)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	c.IndentedJSON(http.StatusCreated, cart)
}

func deleteCart(c *gin.Context) {
	id := c.Param("id")
	var cnx = Connection()
	collection := cnx.Database("store").Collection("carts")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{})
}

func main() {
	router := gin.Default()
	router.GET("/cart", createCart)
	router.GET("/cart/:id", getCart)
	router.PUT("/cart/:id", addToCart)
	router.DELETE("/cart/:id", deleteCart)

	router.Run(":8080")
}
