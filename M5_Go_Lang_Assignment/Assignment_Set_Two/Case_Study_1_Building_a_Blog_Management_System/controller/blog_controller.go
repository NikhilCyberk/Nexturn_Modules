package controller

import (
	"Case_Study_1_Building_a_Blog_Management_System/model"
	"Case_Study_1_Building_a_Blog_Management_System/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BlogController struct {
    BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
    return &BlogController{BlogService: blogService}
}

func (c *BlogController) HandleBlogs(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        c.CreateBlog(w, r)
    case http.MethodGet:
        c.GetAllBlogs(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (c *BlogController) HandleBlogByID(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL path
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/blogs/"))
    if err != nil {
        http.Error(w, "Invalid blog ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        c.GetBlog(w, r, id)
    case http.MethodPut:
        c.UpdateBlog(w, r, id)
    case http.MethodDelete:
        c.DeleteBlog(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (c *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
    var blog model.Blog
    if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    createdBlog, err := c.BlogService.CreateBlog(&blog)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdBlog)
}

func (c *BlogController) GetBlog(w http.ResponseWriter, r *http.Request, id int) {
    blog, err := c.BlogService.GetBlog(id)
    if err != nil {
        if err.Error() == "blog not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(blog)
}

func (c *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
    blogs, err := c.BlogService.GetAllBlogs()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(blogs)
}

func (c *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request, id int) {
    var blog model.Blog
    if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updatedBlog, err := c.BlogService.UpdateBlog(id, &blog)
    if err != nil {
        if err.Error() == "blog not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedBlog)
}

func (c *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request, id int) {
    err := c.BlogService.DeleteBlog(id)
    if err != nil {
        if err.Error() == "blog not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}