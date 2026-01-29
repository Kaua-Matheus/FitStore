package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"
)

func SetupFileRoutes(router *gin.Engine) {
	
	ip, err := utils.GetLocalIP(); if err != nil {
		ip = "localhost"
	}

	// GET
	// Adquire todas as imagens de dentro da pasta
	router.GET("/:dir", func(ctx *gin.Context) {
		dir := ctx.Param("dir")

		path := filepath.Join("..", "uploads", dir)

		files, err := os.ReadDir(path)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Error trying to read the '%s' path", dir),
			})
			return
		}

		var fileList []map[string]string

		for _, file := range files {
			if !file.IsDir() {
				fileInfo := map[string]string{
					"filename": file.Name(),
					"url":      fmt.Sprintf("http://%s:8080/files/%s/%s", ip, dir, file.Name()),
				}
				fileList = append(fileList, fileInfo)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": fileList,
		})
	})

	// GET
	// Retorna um tipo file contendo uma imagem com base na categoria e nome da imagem
	router.GET("/files/:categoria/:filename", func(ctx *gin.Context) {
		categoria := ctx.Param("categoria")
		filename := ctx.Param("filename")

		filePath := filepath.Join("..", "uploads", categoria, filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("File not found: %s", filePath),
			})
			return
		}

		ctx.File(filePath)
	})
}