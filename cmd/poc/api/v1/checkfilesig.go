package v1

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/signing"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func CheckFileSignatureHandler(c *gin.Context) {

	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	checkfilepublickey, valid := c.GetPostForm("pubkey")
	if !valid {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	checkfilesignature, valid := c.GetPostForm("signature")
	if !valid {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	// single file form
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	// ToDo: Inspect header and file length. If the file is too big, throw error

	// open file and read its contents in the fileBytes array
	fh, err := file.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	defer func(file multipart.File) {
		_ = file.Close()
	}(fh)

	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(fh)
	if err != nil {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	// Log information on the file that we read
	log.WithField("Bytes read", fileSize).Info("we just read the entire file: " + file.Filename)

	// Check if the signature is correct for this file and public key
	validSig, err := signing.VerifyFileBytesSignature(buff.Bytes(), model.ZKPKEY, checkfilepublickey, checkfilesignature)
	if err != nil {
		//validation err back
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	c.JSON(http.StatusOK, model.CheckFileSigResponse{
		ValidSignature: validSig,
	})
}
