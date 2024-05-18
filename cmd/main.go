package main

import (
	"github.com/devlongs/batch-wave/config"
	"github.com/devlongs/batch-wave/handlers"
	"github.com/devlongs/batch-wave/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	ethURL := config.GetConfig("ethereum.url")
    privateKey := config.GetConfig("ethereum.private_key")
	
    r := gin.Default()

    ethService, err := services.NewEthereumService(ethURL, privateKey)
    if err != nil {
        panic(err)
    }

    handlers.InitEthereumService(ethService)

    r.POST("/batch", handlers.CreateBatch)
    r.POST("/batch/send", handlers.SendBatch)

    r.Run(":8080")
}
