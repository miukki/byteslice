package v1

import (
	"encoding/json"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/signing"

	"github.com/gin-gonic/gin"
)

func SignTextHandler(c *gin.Context) {

	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var targetObject model.SignatureRequest
	decoder := json.NewDecoder(c.Request.Body)
	// r.PostForm is a map of our POST form values
	err := decoder.Decode(&targetObject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	signature, err := signing.Signer(true, targetObject.Text, model.ZKPKEY)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, model.SignTextResponse{
		Signature: signature,
	})
}
