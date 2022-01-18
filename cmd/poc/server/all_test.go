package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"test-zkp/cmd/poc/api/model"
	"test-zkp/cmd/poc/server"
	"test-zkp/cmd/poc/sha3"
	"testing"
	"time"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/require"
)

var (
	apiKey = model.HeaderZKPKEY
)

// global var with the address of the smart contract
var auditAddress string

func TestZKPAPI(t *testing.T) {

	// Check if API key matches the one in our environment
	envkey := model.ZKPKEY
	require.NotEmpty(t, envkey)

	mnemonic := model.MNEMONIC
	require.NotEmpty(t, mnemonic)

	// Deploy new contract and capture addr in global var
	auditAddress, err := sha3.Deploy(true, envkey)
	require.NoError(t, err)
	require.NotEmpty(t, auditAddress)

	// first wait for 25 secs to let the block be mined
	time.Sleep(25 * time.Second)

	router, err := server.Server(auditAddress)
	require.NoError(t, err)

	testflight.WithServer(router, func(r *testflight.Requester) {

		header := http.Header{}
		header.Add(model.HEADER_TEST_ZKP_KEY, apiKey)

		request, err := http.NewRequest("GET", "/api/v1/address", nil)
		require.NoError(t, err)
		request.Header = header

		response := r.Do(request)
		require.NotEmpty(t, response)
		require.Equal(t, 200, response.StatusCode)

		var responseObj model.AddressResponse

		err = json.Unmarshal(response.RawBody, &responseObj)
		require.NoError(t, err)

		// require.Equal(t, responseObj.AuditAddress, auditAddress)

	})
}

func TestLogState(t *testing.T) {
	router, err := server.Server(auditAddress)
	require.NoError(t, err)

	testflight.WithServer(router, func(r *testflight.Requester) {
		header := http.Header{}
		header.Add(model.HEADER_TEST_ZKP_KEY, apiKey)

		sample := model.LogStateRequest{
			Txn:     `{"_id": "trades/5","recipients": ["1000000","111111","999999"]}`,
			Type:    "TRADE",
			Version: "1.0",
			State:   "CREATED",
		}
		raw, err := json.Marshal(&sample)
		require.NoError(t, err)
		require.NotEmpty(t, raw)

		request, err := http.NewRequest("POST", "/api/v1/logstate", bytes.NewBuffer(raw))
		require.NoError(t, err)
		request.Header = header

		response := r.Do(request)
		require.NotEmpty(t, response)
		require.Equal(t, 200, response.StatusCode)

	})
}
