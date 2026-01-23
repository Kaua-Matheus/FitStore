package handler

import (
	_"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
	"github.com/Kaua-Matheus/fitstore/backend/model/handler"
	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"
	"github.com/Kaua-Matheus/fitstore/backend/controller/middleware"
)

func User(router *gin.Engine, db *gorm.DB) {

	// POST - Público
	router.POST("/user/register", func(ctx *gin.Context) {

		userReq := entitie.UserReq{}
		if err := ctx.BindJSON(&userReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to bind json",
			})
			return
		} else {

			if userDb, err := handler.GetUserByLogin(db, userReq.UserLogin); err != nil {
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
					return
				}

				var user = entitie.User{
					UserName: userReq.UserName, 
					UserLogin: userReq.UserLogin, 
					UserPasswordHash: password,
				}

				if token, err := utils.CreateToken(user.UserLogin); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": "error trying to create a token",
					})
					return
				} else {
					err = handler.AddUser(db, user); if err != nil {
						ctx.JSON(http.StatusBadRequest, gin.H{
							"message": "error trying to add user",
						})
						return
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

	router.POST("/user/login", func(ctx *gin.Context){

		var bindUser = entitie.UserReq{}
		if err := ctx.BindJSON(&bindUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to bind json",
			})
			return
		} else {
			var user, err = handler.GetUserByLogin(db, bindUser.UserLogin); if err != nil {
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
							"authenticated": response,
						})
						return
					}
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"authenticated": response,
					})
					return
				}

			}
		}
	})

	// GET - Público
	// Auth
	router.GET("/user/auth", func(ctx *gin.Context) {
		var token, err = utils.GetAuthCookie(ctx); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authenticated": false,
				"message": "error trying to get token",
			})
			return
		}

		if token == "" {
			utils.ClearAuthCookie(ctx);
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authenticated": false,
				"message": "token doesn't exist, try register or login",
			})
			return
		}

		err = utils.VerifyToken(token); if err != nil {
			utils.ClearAuthCookie(ctx);
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authenticated": false,
				"message": "token doesn't seems right",
			})
			return
		}

		user_login, err := utils.GetUserLoginFromToken(token); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authenticated": false,
				"message": "can't take the user_login",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"authenticated": true,
			"user": gin.H{
				"user_login": user_login,
			},
			"message": "user is authenticated",
		})
	})

	// Logout
	router.GET("/user/logout", func(ctx *gin.Context) {
		utils.ClearAuthCookie(ctx);

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logout successful",
		})
	})

	
	// GET - Protegido (Implementar Middleware)
	// Pode ser interessante aplicar regras por conta dos usuários especiais 
	// (um cliente não pode ver o perfil de um admin)
	router.GET("/user/get/:login", middleware.Auth(), func(ctx *gin.Context) {
		
		login := ctx.Param("login");
		user, err := handler.GetUserByLogin(db, login); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to get the user",
			})
			return
		} 
			
		ctx.JSON(http.StatusOK, gin.H{
			"user_name": user.UserName,
			"profile_image": user.IdImage,
		})
	})
}