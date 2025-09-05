package main

import (
	"fmt"
	"net/http"
	"strconv"

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

var expenses = make([]IExpense,0)

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/expenses", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, expenses)
	})

	r.GET("/expense/:id", func(ctx *gin.Context) {
		var expense IExpense
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		for _, exp := range expenses {
			if exp.ID == id {
				expense = exp
				ctx.JSON(http.StatusOK, expense)
				return
			}
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Errorf("expense with ID %s doesnot exsist", strconv.Itoa(id))})
	})

	r.POST("/expense", func(ctx *gin.Context) {
		// Bind the input
		var in IExpenseInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		prevExpense := expenses[len(expenses)-1]
		// add the expense
		expense := IExpense{
			ID:        prevExpense.ID + 1,
			Amount:   in.Amount,
			Category: in.Category,
			Date:     in.Date,
		}
		expenses = append(expenses, expense)
		ctx.JSON(http.StatusOK, expense)
	})

	r.PUT("/expense/:id", func(ctx *gin.Context) {
		// Bind the input
		var in IExpenseInput
		if err := ctx.ShouldBindJSON(&in); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var expense IExpense
		index := 0
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		for i, exp := range expenses {
			if exp.ID == id {
				expense = exp
				index = i + 1
				break
			}
		}

		if index == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Errorf("expense with ID %s doesnot exsist", strconv.Itoa(id))})
		}

		// update the expense
		expense.Amount = in.Amount
		expense.Category = in.Category
		expense.Date = in.Date

		expenses[index-1] = expense
		ctx.JSON(http.StatusOK, expense)
	})

	r.DELETE("/expense", func(ctx *gin.Context) {
		// Bind the input
		index := 0
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		for i, exp := range expenses {
			if exp.ID == id {
				index = i + 1
				break
			}
		}

		if index == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Errorf("expense with ID %s doesnot exsist", strconv.Itoa(id))})
		}

		expenses = append(expenses[:index-1], expenses[index+1:]...) 
		ctx.JSON(http.StatusOK, "Expense is deleted successfully!")
	})

	http.ListenAndServe(":8080", r)
}
