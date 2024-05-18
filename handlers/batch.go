package handlers

import (
	"math/big"
	"net/http"

	"github.com/devlongs/batch-wave/models"
	"github.com/devlongs/batch-wave/services"

	"github.com/gin-gonic/gin"
)

var ethService *services.EthereumService

func InitEthereumService(service *services.EthereumService) {
    ethService = service
}

func CreateBatch(c *gin.Context) {
    var batch models.Batch
    if err := c.ShouldBindJSON(&batch); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
   
    c.JSON(http.StatusOK, gin.H{"batch": batch})
}

func SendBatch(c *gin.Context) {
    var batch models.Batch
    if err := c.ShouldBindJSON(&batch); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for _, tx := range batch.Transactions {
        amount, _ := new(big.Int).SetString(tx.Amount, 10)
        txHash, err := ethService.SendTransaction(tx.ToAddress, amount)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
 
        c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
    }
}
