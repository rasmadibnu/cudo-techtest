package controller

import (
	"cudo-techtest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Name   string
	Plural string
}

func NewTransactionController(r *gin.RouterGroup, s service.TransactionService) {
	this := &TransactionController{
		Name:   "Transaction",
		Plural: "Transactions",
	}

	route := r.Group("/fraud-detection")
	route.GET("", this.index(s))
}

// @Summary REST API Get data
// @Description Get data Pagination
// @Author rasmadibnu
// @Success 200 {object} array entity
// @Failure 404 {object} nil
// @method [GET]
// @Router /fraud-detection
func (c *TransactionController) index(service service.TransactionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := service.GetDataTransaction()

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err)
		}

		ctx.JSON(http.StatusOK, data)
	}
}
