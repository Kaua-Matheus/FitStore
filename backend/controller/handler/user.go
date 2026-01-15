package handler

import (
	_"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model"
	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"
)

func User(router *gin.Engine, db *gorm.DB) {

	// GET
	router.GET("/user/:login", func(ctx *gin.Context) {
		
		login := ctx.Param("login");
		user, err := model.GetUserByLogin(db, login); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to get the user",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": user,
			})
		}

	})

	// POST
	router.POST("/user", func(ctx *gin.Context) {

		userReq := model.UserReq{}
		if err := ctx.BindJSON(&userReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to bind json",
			})
		} else {

			var password, err = utils.HashPassword(userReq.UserPassword); if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "error trying to hash the password",
				})
			}

			var user = model.User{
				UserName: userReq.UserName, 
				UserLogin: userReq.UserLogin, 
				UserPasswordHash: password,
			}

			err = model.AddUser(db, user); if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "error trying to add user",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "user added successfully",
				})
			}
		}

	})

	router.POST("/loguser", func(ctx *gin.Context){

		var bindUser = model.UserReq{}
		if err := ctx.BindJSON(&bindUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to bind json",
			})
		} else {
			var user, err = model.GetUserByLogin(db, bindUser.UserLogin); if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "error trying to get user",
				})	
			} else {
				var response = utils.ComparePassword(bindUser.UserPassword, user.UserPasswordHash);

				ctx.JSON(http.StatusOK, gin.H{
					"auth": response,
				})
			}
		}
	})
}