package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	ID   string
	Name string
}

type LineItem struct {
	Product  Product
	Quantity int
}

type Cart struct {
	LineItems []LineItem
	ID        string
}

func getCart(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, Cart{
		ID: id,
		LineItems: []LineItem{
			LineItem{
				Product: Product{
					ID:   "1",
					Name: "Product 1",
				},
				Quantity: 1,
			},
		},
	})
}

func addToCart(c *gin.Context) {
	id := c.Param("id")
	var lineItem LineItem
	var cart Cart = Cart{
		ID:        id,
		LineItems: []LineItem{},
	}

	if err := c.BindJSON(&lineItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart.LineItems = append(cart.LineItems, lineItem)

	c.IndentedJSON(http.StatusCreated, cart)
}

func main() {
	router := gin.Default()
	router.GET("/cart/:id", getCart)
	router.PUT("/cart/:id", addToCart)

	router.Run(":8080")
}
