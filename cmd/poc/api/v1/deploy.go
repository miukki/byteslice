package v1

import (
	"net/http"
	"test-zkp/cmd/poc/api/model"

	"github.com/gin-gonic/gin"
)

func DeployHandler(c *gin.Context) {
	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, model.AddressResponse{
		AuditAddress: model.AuditAddr,
	})
}
