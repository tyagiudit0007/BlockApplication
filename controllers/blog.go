package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	RelatedBlogs string    `json:"relatedBlogs"`
	BlogType     string    `json:"type"`
	Rating       int       `json:"rating"`
	Completed    string    `json:"completed"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Create User Table
// func CreateBlogTable(db *pg.DB) error {
// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	createError := db.CreateTable(&Blog{}, opts)
// 	if createError != nil {
// 		log.Printf("Error while creating Blog table, Reason: %v\n", createError)
// 		return createError
// 	}
// 	log.Printf("Blog table created")
// 	return nil
// }

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *gorm.DB

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

func GetAllBlogs(c *gin.Context) {
	var Blogs []Blog
	result := dbConnect.Find(&Blogs)

	if result.Error != nil {
		log.Printf("Error while getting all Blogs, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Blogs",
		"data":    Blogs,
	})
	return
}

func CreateBlog(c *gin.Context) {
	var Blg Blog
	c.BindJSON(&Blg)
	fmt.Println(Blg, "retrived data")
	title := Blg.Title
	body := Blg.Body
	completed := Blg.Completed

	blgType := Blg.BlogType
	id := guuid.New().String()

	result := dbConnect.Create(&Blog{
		ID:           id,
		Title:        title,
		Body:         body,
		BlogType:     blgType,
		Rating:       Blg.Rating,
		RelatedBlogs: Blg.RelatedBlogs,
		Completed:    completed,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	fmt.Println(guuid.New().String(), "newId")
	if result.Error != nil {
		log.Printf("Error while inserting new Blog into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Blog created Successfully",
	})
	return
}

func GetSingleBlog(c *gin.Context) {
	BlogId := c.Param("blogId")
	fmt.Println("blog Id", BlogId)
	var Blg Blog
	result := dbConnect.First(&Blg, "id = ?", BlogId)

	if result.Error != nil {
		log.Printf("Error while getting a single Blog, Reason: %v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Blog",
		"data":    Blg,
	})
	return
}

func EditBlog(c *gin.Context) {
	BlogId := c.Param("blogId")
	var payload Blog
	c.BindJSON(&payload)

	var Blg Blog
	result := dbConnect.First(&Blg, "id = ?", BlogId)
	if result.Error != nil {
		log.Printf("Error, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	UpdatedBlog := Blog{
		ID:           BlogId,
		Body:         payload.Body,
		Title:        payload.Title,
		BlogType:     payload.BlogType,
		Rating:       payload.Rating,
		RelatedBlogs: payload.RelatedBlogs,
		UpdatedAt:    time.Now(),
		CreatedAt:    Blg.CreatedAt,
	}
	dbConnect.Model(&Blg).Updates(UpdatedBlog)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Blog Edited Successfully",
	})
	return
}

func DeleteBlog(c *gin.Context) {
	BlogId := c.Param("blogId")
	fmt.Println(BlogId, "ksdfn")
	result := dbConnect.Delete(&Blog{}, "id = ?", BlogId)
	if result.Error != nil {
		log.Printf("Error while deleting a single Blog, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Blog deleted successfully",
	})
	return
}
