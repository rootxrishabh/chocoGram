package controllers

import (
	"encoding/json"
	"log"

	"github.com/rootxrishabh/chocoGram/models"
	"github.com/rootxrishabh/chocoGram/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	createUser := models.User{}
	utils.ParseBody(c.Request, &createUser)
	res, err := createUser.CreateUser()
	if err != nil {
		c.Writer.WriteHeader(400)
		c.Header("Content-Type", "application/json")
		c.Writer.Write([]byte(`{	"status": "failure",
 	"reason": "User already exists or invalid request format"
}`))
		return
	}
	val, _ := json.Marshal(res)
	c.Writer.WriteHeader(201)
	c.Header("Content-Type", "application/json")
	c.Writer.Write(val)
}

func SendFriendRequest(c *gin.Context) {
	err, present := models.CreateFriendRequest(c.Param("userA"), c.Param("userB"))
	if present {
		c.Writer.WriteHeader(400)
		c.Header("Content-Type", "application/json")
		c.Writer.Write([]byte(`{	"status": "failure",
 	"reason": "Invalid usernames"
}`))
		return
	} else if err != nil {
		log.Println("request already sent")
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte(`{	"status": "failure",
		"reason": "Invalid request format or user already has a pending request or is already friends with the user"
}`))
		return
	}
	c.Writer.WriteHeader(202)
	c.Header("Content-Type", "application/json")
	c.Writer.Write([]byte(`{	"status": "success" }`))
}

func GetFriends(c *gin.Context) {
	res, present := models.GetAllFriends(c.Param("userA"))
	if present {
		c.Writer.WriteHeader(400)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write([]byte(`{
			"status": "failure",
			"reason": "Invalid username or incorrect request format"
		}`))
		return
	}
	if len(res) == 0 {
		c.Writer.WriteHeader(404)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write([]byte(`{
			"status": "failure",
			"reason": "No friends found for this user"
		}`))
		return
	}
	val, err := json.Marshal(res)
	if err != nil {
		log.Println("Error marshalling response:", err)
		return
	}
	c.Writer.WriteHeader(200)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write([]byte(`{"friends": ` + string(val) + `}`))
}
