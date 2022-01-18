package v1

import (
	"encoding/json"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/sha3"

	"github.com/gin-gonic/gin"
)

func GetAllStates(c *gin.Context) {
	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var targetObject model.StateRequest
	decoder := json.NewDecoder(c.Request.Body)
	// r.PostForm is a map of our POST form values
	err := decoder.Decode(&targetObject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	t, s, err := sha3.Gethistory(true, targetObject.ID, targetObject.Type, model.AuditAddr, model.ZKPKEY)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, model.AllStatesResponse{
		Type:   t,
		States: s,
	})
}
