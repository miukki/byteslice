package v1

import (
	"encoding/json"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/sha3"

	"github.com/gin-gonic/gin"
)

func VerifyState(c *gin.Context) {
	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var targetObject model.VerifyStateRequest
	decoder := json.NewDecoder(c.Request.Body)
	// r.PostForm is a map of our POST form values
	err := decoder.Decode(&targetObject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	result, err := sha3.VerifyLatest(true, targetObject.ID, targetObject.Type, targetObject.TXN, model.AuditAddr, model.ZKPKEY)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Return status
	c.JSON(http.StatusOK, model.VerifyStateResponse{
		State:  "latest",
		Result: result,
	})
}
