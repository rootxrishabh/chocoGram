package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rootxrishabh/chocoGram/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	Username string `gorm:"type:varchar(255);unique_index"`
}

type Friendships struct {
	User_a_username string `gorm:"type:varchar(255);unique_index"`
	User_b_username string `gorm:"type:varchar(255);unique_index"`
	Status          int    `gorm:"default:null"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() (interface{}, error) {
	d := db.Create(u)
	if d.Error != nil {
		return nil, d.Error
	}
	return d.Value, nil
}

func GetAllFriends(username string) ([]string, bool) {
	if username[0] == '$' {
		username = strings.TrimPrefix(username, "$")
	}

	var users []User

	var check User
	db.Where("username = ?", username).First(&check)
	if check.Username == "" {
		return nil, true
	}

	var users1, users2 []User

	db.Table("users").Select("users.*").
		Joins("JOIN friendships ON users.username = friendships.user_b_username").
		Where("friendships.User_a_username = ? AND friendships.status = ?", username, 2).
		Find(&users1)

	db.Table("users").Select("users.*").
		Joins("JOIN friendships ON users.username = friendships.user_a_username").
		Where("friendships.User_b_username = ? AND friendships.status = ?", username, 2).
		Find(&users2)

	users = append(users1, users2...)

	// users now contains the result of the query
	var ans []string
	for _, s := range users {
		s.Username = "$" + s.Username
		ans = append(ans, s.Username)
	}
	return ans, false
}

func CreateFriendRequest(usernameA string, usernameB string) (error, bool) {
	var friendship Friendships

	if usernameA[0] == '$' {
		usernameA = strings.TrimPrefix(usernameA, "$")
	}
	if usernameB[0] == '$' {
		usernameB = strings.TrimPrefix(usernameB, "$")
	}

	checkA := db.Where("username = ?", usernameA).First(&User{}).RecordNotFound()
	if checkA {
		return nil, true
	}

	// Check if usernameB exists
	checkB := db.Where("username = ?", usernameB).First(&User{}).RecordNotFound()
	if checkB {
		return nil, true
	}

	result := db.Where("User_a_username = ? AND User_b_username = ?", usernameB, usernameA).First(&friendship)
	if friendship.User_a_username != usernameA && friendship.Status == 1 {
		// Update the status to 2
		update := db.Model(&friendship).Where("user_a_username = ? AND user_b_username = ?", usernameB, usernameA).Update("status", 2)
		if update.Error != nil {
			fmt.Println("Error updating friendship status:", update.Error)
		} else {
			fmt.Println("Friendship status updated successfully")
		}
	} else if friendship.Status == 2 {
		// Return an error
		fmt.Println("A friend request or friendship already exists")
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newFriendship := Friendships{User_a_username: usernameA, User_b_username: usernameB, Status: 1}
		result = db.Create(&newFriendship)
		if result.Error != nil {
			fmt.Println("Error creating friend request:", result.Error)
		} else {
			fmt.Println("Friend request created successfully")
		}
	}
	return result.Error, false
}
