package v1

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/sha3"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func CalcFileHashHandler(c *gin.Context) {
	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
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

	err = c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	h, err := sha3.CalcFileByteshash(buff.Bytes(), c.Request.PostForm.Get("hashtype"), model.ZKPKEY)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, model.CalcTextHashResponse{
		Hash: h,
	})
}
