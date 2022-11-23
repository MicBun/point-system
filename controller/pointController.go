package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pointSystem/model"
	"pointSystem/util/token"
)

type PointInterface interface {
	GetPoint(ctx *gin.Context)
	UpdatePoint(ctx *gin.Context)
}

type GetPointOutput struct {
	Amount int `json:"amount"`
}

// GetPoint godoc
// @Summary Get Point
// @Description Get Point
// @Tags Point
// @Produce json
// @Success 200 {object} GetPointOutput
// @Router /point [get]
func GetPoint(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	decoded, _ := token.ExtractTokenID(ctx)
	var point model.Point
	if err := db.Where("user_id = ?", decoded).Find(&point).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Point not found!"})
		return
	}
	ctx.JSON(200, gin.H{"data": GetPointOutput{
		Amount: point.Amount,
	}})
}

type UpdatePointInput struct {
	Amount int `json:"amount"`
}

// UpdatePoint godoc
// @Summary Add Point
// @Description Add Point
// @Tags Point
// @Param Body body UpdatePointInput true "the body to add Point"
// @Produce json
// @Success 200 {object} GetPointOutput
// @Router /point [post]
func UpdatePoint(ctx *gin.Context) {
	var input UpdatePointInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	decoded, _ := token.ExtractTokenID(ctx)
	var point model.Point
	if err := db.Where("user_id = ?", decoded).Find(&point).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Point not found!"})
		return
	}
	if input.Amount < 0 && point.Amount < -input.Amount {
		ctx.JSON(400, gin.H{"error": "Point is not enough!"})
		return
	}
	point.Amount += input.Amount
	if err := db.Save(&point).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "Update Point failed!"})
		return
	}
	ctx.JSON(200, gin.H{"success": "Update Point success!"})
}
