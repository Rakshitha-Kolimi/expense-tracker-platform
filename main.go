package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IExpense struct {
	ID       int    `json:"id"`
	Amount   int    `json:"amount"`
	Category string `json:"category"`
	Date     string `json:"date"`
}

type IExpenseInput struct {
	Amount   int    `json:"amount"`
	Category string `json:"category"`
	Date     string `json:"date"`
}

var dummyData = []IExpense{
	{ID: 1, Amount: 200, Category: "Food", Date: "2025-09-04"},
	{ID: 2, Amount: 1500, Category: "Rent", Date: "2025-09-01"},
}

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/expenses", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dummyData)
	})

	r.POST("/expense", func(ctx *gin.Context) {
		// Bind the input
		var in IExpenseInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// add the expense
		expense := IExpense{
			ID:       len(dummyData) + 1,
			Amount:   in.Amount,
			Category: in.Category,
			Date:     in.Date,
		}
		dummyData = append(dummyData, expense)
		ctx.JSON(http.StatusOK, expense)
	})
	http.ListenAndServe(":8080", r)
}
