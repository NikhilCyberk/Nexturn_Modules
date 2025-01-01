package controller

import (
	"blogmanager/model"
	"blogmanager/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{BlogService: blogService}
}
func (controller *BlogController) CreateBlog(c *gin.Context) {
	fmt.Println("CreateBlog: Received request to create a Blog")

	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		fmt.Println("CreateBlog: Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("CreateBlog: Parsed Blog details: %+v\n", blog)

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		fmt.Printf("CreateBlog: Error creating Blog in service layer: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Blog"})
		return
	}

	fmt.Printf("CreateBlog: Successfully created Blog: %+v\n", createdBlog)
	c.JSON(http.StatusOK, createdBlog)
}

func (controller *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")
	BlogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	Blog, err := controller.BlogService.GetBlog(BlogID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, Blog)
}

func (controller *BlogController) GetAllBlogs(c *gin.Context) {
	Blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Blogs)
}

func (controller *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	BlogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var Blog model.Blog
	if err := c.ShouldBindJSON(&Blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Blog.ID = BlogID
	updatedBlog, err := controller.BlogService.UpdateBlog(&Blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (controller *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	BlogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.BlogService.DeleteBlog(BlogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
