package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rootxrishabh/chocoGram/controllers"
)

var SocialNetwork = func(r *gin.Engine) {
	r.POST("/add/:userA/:userB", controllers.SendFriendRequest)
	r.POST("/create", controllers.CreateUser)
	r.GET("/friends/:userA", controllers.GetFriends)
}
