package controller

import (
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/model"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
    service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
    return &ProductController{service: service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    var product model.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.service.CreateProduct(&product); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    product, err := c.service.GetProduct(id)
    if err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        ctx.JSON(status, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
    var query model.PaginationQuery
    if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    products, total, err := c.service.GetProducts(query.Page, query.Limit)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": products,
        "total": total,
        "page": query.Page,
        "limit": query.Limit,
    })
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    var product model.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = id

    if err := c.service.UpdateProduct(&product); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        ctx.JSON(status, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) UpdateStock(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    var stockUpdate struct {
        Stock int `json:"stock" binding:"required,gte=0"`
    }
    if err := ctx.ShouldBindJSON(&stockUpdate); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.service.UpdateStock(id, stockUpdate.Stock); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        ctx.JSON(status, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "stock updated successfully"})
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    if err := c.service.DeleteProduct(id); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        ctx.JSON(status, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusNoContent)
}