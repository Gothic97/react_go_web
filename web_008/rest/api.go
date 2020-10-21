package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()
	h, _ := MakeHandler()
	r.GET("/", h.GetMainPage)
	r.GET("/test", h.GetTest)
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	//r.GET("/promos", h.GetPromos)
	//
	//userGroup := r.Group("/user")
	//{
	//	userGroup.POST("/:id/signout", h.SignOut)
	//	userGroup.GET("/:id/orders", h.GetOrders)
	//}
	//
	//usersGroup := r.Group("/users")
	//{
	//	//usersGroup.POST("/charge", h.Charge)
	//	usersGroup.POST("/create", h.Create)
	//	usersGroup.POST("/read", h.Read)
	//	usersGroup.POST("/update", h.Update)
	//	usersGroup.POST("/delete", h.Delete)
	//	//usersGroup.POST("", h.AddUser)
	//}

	r.POST("/crud", h.CRUD)

	return r.Run(address)
}
