package model

import "os"

const (
	HEADER_TEST_ZKP_KEY = "X-TEST-ZKP-KEY"
)

const (
	ENV_TEST_HEADER_ZKP_KEY = "TEST_HEADER_ZKP_KEY"
	ENV_TEST_ZKP_KEY        = "TEST_ZKP_KEY"
	ENV_TEST_AUDIT_ADDR     = "TEST_AUDIT_ADDR"
	ENV_MNEMONIC              = "MNEMONIC"
	ENV_HDPATH                = "HDPATH"
	ENV_PORT                  = "PORT"
	ENV_GANACHEKEY            = "GANACHEKEY"
)

var (
	AuditAddr    = os.Getenv(ENV_TEST_AUDIT_ADDR)
	HeaderZKPKEY = os.Getenv(ENV_TEST_HEADER_ZKP_KEY)
	ZKPKEY       = os.Getenv(ENV_TEST_ZKP_KEY)
	MNEMONIC     = os.Getenv(ENV_MNEMONIC)
	HDPATH       = os.Getenv(ENV_HDPATH)
	PORT         = os.Getenv(ENV_PORT)
	GanacheKey   = os.Getenv(ENV_GANACHEKEY)
)
