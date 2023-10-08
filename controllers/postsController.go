package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//Get Data
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)
	//Create a post 
	posts := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&posts)
	
	if result.Error != nil{
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": posts,
	})
}

//Getting all the posts
func PostIndex (c *gin.Context){

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

//Getting single post by ID
func PostShow(c *gin.Context){

	//Get ID of Url
	id := c.Param("id")

	var posts models.Post
	initializers.DB.Find(&posts, id)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

//Update
func PostUpdate(c *gin.Context){

	//Get ID of Url
	id := c.Param("id")

	//Get Data
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//Find the post to update

	var posts models.Post
	initializers.DB.Find(&posts, id)

	//Update it
	initializers.DB.Model(&posts).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

//Delete

func PostDelete(c *gin.Context){

	//Get ID of Url
	id := c.Param("id")


	//Delete Data by id
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
