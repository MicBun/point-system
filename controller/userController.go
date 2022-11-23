package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pointSystem/model"
	"pointSystem/util/token"
)

type UserInterface interface {
	LoginUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser godoc
// @Summary Login with credential.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginUserInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func LoginUser(ctx *gin.Context) {
	var input LoginUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var user model.User
	if err := db.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}
	generatedToken, err := token.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "token": generatedToken})
}

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// CreateUser godoc
// @Summary Create a new user.
// @Description Create a new user.
// @Tags User
// @Param Body body CreateUserInput true "the body to create a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	var input CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
	}
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong."})
		return
	}
	point := model.Point{
		UserID: user.ID,
	}
	if err := db.Create(&point).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user created successfully", "user": user})
}
