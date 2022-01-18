package model

type AddressResponse struct {
	AuditAddress string `json:"address"` // SECURITY
}

type HashResponse struct {
	Hash string `json:"hash"`
}

type CalcTextHashResponse struct {
	Hash string `json:"hash"`
}

type CalcFileHashResponse struct {
	Hash string `json:"hash"`
}

type LogResponse struct {
	Hash  string `json:"hash"`
	Type  string `json:"type"`
	State string `json:"state"`
}

type AllStatesResponse struct {
	Type   string `json:"type"`
	States string `json:"states"`
}

type StateResponse struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

type VerifyStateResponse struct {
	State  string `json:"state"`
	Result bool   `json:"result"`
}

type SignTextResponse struct {
	Signature string `json:"signature"`
}

type SignFileResponse struct {
	Signature string `json:"signature"`
}

type CheckTextSigResponse struct {
	ValidSignature bool `json:"validsignature"`
}

type CheckFileSigResponse struct {
	ValidSignature bool `json:"validsignature"`
}

// LogStateRequest Will only be called by Test API (testapi)
type LogStateRequest struct {
	Txn string `json:"txn"` // Serialized json of order, otcorder or trade, all have an _id field where the following
	// must be true:
	// the serialized json must have a field "_id" with the format; <string>/<ID> where:
	// if _id == trades/<ID> then type must be "TRADE" and Id must be <ID>.
	// if _id == orders/<ID> then type must be "ORDER" and Id must be <ID>.
	// if _id == orders_otc/<ID> then type must be "OTCORDER" and Id must be <ID>.
	// json must be legal and should represent the full detail as stored in ADB.
	Type string `json:"type"` // record type ["ORDER"|"OTCORDER"|"TRADE"] which corresponding to ADB tables;
	//              			 orders,  orders_otc, trades
	Version string `json:"version"` // version of passed data-structure which defines the available attributes
	State   string `json:"state"`
}

type CalcTextHashHandler struct {
	Text  string `json:"text"`
	HType string `json:"hashtype"`
}

type SignatureRequest struct {
	Text      string `json:"text"`
	PubKey    string `json:"pubkey"`
	Signature string `json:"signature"`
}

type StateRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type VerifyStateRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	TXN  string `json:"txn"`
}
