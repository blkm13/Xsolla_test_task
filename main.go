package main

import (
	"Test_task/Methods"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main(){

	//gin part
	r := gin.Default()

	r.POST("/additem", Methods.AddItem)
	r.PUT("/edititem", Methods.EditItem)
	r.DELETE("/deleteitem", Methods.DeleteItem)
	r.GET("/getitem", Methods.GetItem)
	r.GET("/getallitems", Methods.GetAllItems)

	r.Run()
}
