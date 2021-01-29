package controller

import "github.com/gin-gonic/gin"

type Belanjaan struct {
	name     string `json:"name"`
	quantity string `json:"quantity"`
	price    string `json:"price"`
}

type ItemBelanjaan struct {
	ListOfItem []Belanjaan
}

func (list *ItemBelanjaan) CreateItem(c *gin.Context) {
	var item Belanjaan
	itemsName := c.PostForm("name")
	itemsQuantity := c.PostForm("quantity")
	itemsPrice := c.PostForm("price")

	item.name = itemsName
	item.quantity = itemsQuantity
	item.price = itemsPrice

	list.ListOfItem = append(list.ListOfItem, item)

	c.JSON(200, gin.H{
		"message":  "successfully created",
		"name":     itemsName,
		"quantity": itemsQuantity,
		"price":    itemsPrice,
	})
}

func (list *ItemBelanjaan) GetList(c *gin.Context) {
	c.JSON(200, gin.H{
		"List Shooping": list.ListOfItem,
	})
}

func (list *ItemBelanjaan) UpdateListShopping(c *gin.Context) {
	name := c.Param("name")
	quantity := c.PostForm("quantity")
	price := c.PostForm("price")

	for i, v := range list.ListOfItem {
		if name == v.name {
			list.ListOfItem[i].quantity = quantity
			list.ListOfItem[i].price = price
		}
	}

}
