package handler

import (
	_"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model"
	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"
	"github.com/Kaua-Matheus/fitstore/backend/controller/middleware"
)

func User(router *gin.Engine, db *gorm.DB) {

	// POST - Público
	router.POST("/user", func(ctx *gin.Context) {

		userReq := model.UserReq{}
		if err := ctx.BindJSON(&userReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to bind json",
			})
		} else {

			if userDb, err := model.GetUserByLogin(db, userReq.UserLogin); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "couldn't get information about the user_login",
				})
				return
			} else {
				if userDb.UserLogin == userReq.UserLogin  {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": "user already exists, try login", 
						// Podemos aplicar uma função para realizar o login automaticamente se as credenciais estiverem corretas
					})
					return
				}

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

				if token, err := utils.CreateToken(user.UserLogin); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": "error trying to create a token",
					})
				} else {
					err = model.AddUser(db, user); if err != nil {
						ctx.JSON(http.StatusBadRequest, gin.H{
							"message": "error trying to add user",
						})
					} else {
						
						utils.SetAuthCookie(ctx, token);

						ctx.JSON(http.StatusOK, gin.H{
							"message": "user added successfully",
						})
					}
				}
			}
		}

	})

	router.POST("/userlogin", func(ctx *gin.Context){

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
				return
			} else {
				var response = utils.ComparePassword(bindUser.UserPassword, user.UserPasswordHash);

				if response {
					if token, err := utils.CreateToken(bindUser.UserLogin); err != nil {
						ctx.JSON(http.StatusBadRequest, gin.H{
							"message": "error trying to create a token",
						})
						return
					} else {

						utils.SetAuthCookie(ctx, token);

						ctx.JSON(http.StatusOK, gin.H{
							"auth": response,
						})
					}
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"auth": response,
					})
				}

			}
		}
	})

	router.POST("/logout", func(ctx *gin.Context) {
		utils.ClearAuthCookie(ctx);

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logout successful",
		})
	})
	
	// GET - Protegido (Implementar Middleware)
	// Pode ser interessante aplicar regras por conta dos usuários especiais 
	// (um cliente não pode ver o perfil de um admin)
	router.GET("/user/:login", middleware.Auth(), func(ctx *gin.Context) {
		
		login := ctx.Param("login");
		user, err := model.GetUserByLogin(db, login); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to get the user",
			})
		} else {
			
			ctx.JSON(http.StatusOK, gin.H{
				"user_name": user.UserName,
			})
		}

	})
}