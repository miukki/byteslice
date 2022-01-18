package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/sha3"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ORDERTRADE struct {
	Id string                 `json:"_id"` // just need to be able to capture the _id from the parsed json Txn
	X  map[string]interface{} `json:"-"`   // te remainder of the ORDERTRADE struct we can ignore
}

func LogStateHandler(c *gin.Context) {
	// Check API key and validate
	headerZKPKEY := c.GetHeader(model.HEADER_TEST_ZKP_KEY)
	if headerZKPKEY != model.HeaderZKPKEY {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var targetObject model.LogStateRequest
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&targetObject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//TODO: validate the targetObject is valid against the (orders|orders_otc|trades)model use swager API

	var tx ORDERTRADE
	err = json.Unmarshal([]byte(targetObject.Txn), &tx)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Here we have a valid json and should have isolated "_id" ignoring the rest

	s := strings.Split(tx.Id, "/")
	if len(s) != 2 { // "_id must have a '/' and therefore the split should result in s having 2 elements
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(s[1]) // check if second element of '/' split tx.Id is a number, if so capture in id
	if err != nil {
		log.WithError(err).Error("sha3.Loghash no valid ID error")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	switch s[0] {
	case "trades":
		if targetObject.Type != "TRADE" {
			err = errors.New("sha3.Loghash table missmatch in _id")
		}
	case "orders":
		if targetObject.Type != "ORDER" {
			err = errors.New("sha3.Loghash table missmatch in _id")
		}
	case "orders_otc":
		if targetObject.Type != "TRADE" {
			err = errors.New("sha3.Loghash table missmatch in _id")
		}
	default:
		err = errors.New("sha3.Loghash unknown table in _id")
	}
	if err != nil {
		log.WithError(err).Error("sha3.Loghash table missmatch in _id")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h, id2, err := sha3.Loghash(true, targetObject.Txn, targetObject.Type, targetObject.State, model.AuditAddr, model.ZKPKEY)
	if err != nil {
		log.WithError(err).Error("sha3.Loghash error")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id2val, err := strconv.Atoi(id2) // check if second element of '/' split tx.Id is a number, if so capture in id
	if err != nil {
		log.WithError(err).Error("sha3.Loghash no valid ID error")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if id != id2val {
		log.WithError(err).Error("sha3.Loghash ID missmatch error")
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, model.StateResponse{
		ID:   id2,
		Hash: h,
	})
}
