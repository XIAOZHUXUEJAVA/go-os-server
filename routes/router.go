package routes

import (
	"log"
	v1 "oschina/api/v1"
	"oschina/middleware"
	"oschina/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	routeV1 := r.Group("api")
	{
		// 这里我们一定要和前端的url对接好
		// user相关
		routeV1.POST("user/login", v1.LoginHandler)
		routeV1.POST("user/register", v1.RegisterHandler)
		routeV1.GET("rss/all", v1.GetAllRss)
		// routeV1.GET("users/:email", v1.GetUserByEmail)
		// routeV1.POST("users/", v1.InsertUser)

		// // booktable相关
		// routeV1.POST("booking", v1.InsertTable)

		// // food相关
		// routeV1.GET("foods", v1.GetAllFood)
		// routeV1.GET("foods/:id", v1.GetFoodById)
	}
	log.Println("Server is running on port:", utils.HttpPort)
	r.Run(utils.HttpPort)

}
