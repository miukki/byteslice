package server

import (
	"net/http"
	"test-zkp/cmd/poc/api/model"
	v1 "test-zkp/cmd/poc/api/v1"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	apiversion = "v1"
)

var (
	GIT_COMMIT string
)

func validateENVVars() {
	if model.PORT == "" {
		log.Warn("ENV VAR PORT is nil")
		model.PORT = "3003"
	}
	if model.AuditAddr == "" {
		panic("ENV VAR TEST_AUDIT_ADDR is nil")
	}
	if model.HeaderZKPKEY == "" {
		panic("ENV VAR TEST_HEADER_ZKP_KEY is nil")
	}
	if model.ZKPKEY == "" {
		panic("ENV VAR TEST_ZKP_KEY is nil")
	}
	if model.MNEMONIC == "" {
		panic("ENV VAR MNEMONIC is nil")
	}
	if model.HDPATH == "" {
		panic("ENV VAR HDPATH is nil")
	}
}

// Server to start a server to handle web requests
func Server(addr string) (*gin.Engine, error) {

	validateENVVars()

	// Override model.AuditAddr when address (addr) was provided as an argument
	if addr != "" {
		model.AuditAddr = addr
	}

	// set the router
	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "test-ZKP api is running",
			"version":    apiversion,
			"GIT_COMMIT": GIT_COMMIT,
		})
	})

	// set up routes to api functions
	api := router.Group("/api/" + apiversion)

	api.GET("/address", v1.DeployHandler)
	api.POST("/hash", v1.GetHashHandler)
	api.POST("/lastlog", v1.GetLogHandler)
	api.POST("/allstates", v1.GetAllStates)
	api.POST("/logstate", v1.LogStateHandler)
	api.POST("/verifystate", v1.VerifyState)
	api.POST("/signtext", v1.SignTextHandler)
	api.POST("/signfile", v1.SignFileHandler)
	api.POST("/texthash", v1.CalcTextHashHandler)
	api.POST("/filehash", v1.CalcFileHashHandler)
	api.POST("/checktextsignature", v1.CheckTextSignatureHandler)
	api.POST("/checkfilesignature", v1.CheckFileSignatureHandler)

	log.Info("Server running on port: " + model.PORT)

	return router, nil
}

func Start(router *gin.Engine) error {
	// Start the server
	err := router.Run(":" + model.PORT)
	if err != nil {
		return err
	}

	return nil
}
