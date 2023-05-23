package models

type (
	Family struct {
		FamId   int    // will generate using util generate ID
		CustId  int    `json:"CustId"`
		FamRel  string `json:"Relationship"`
		FamName string `json:"Name"`
		FamDob  string `json:"DoB"` // format : YYYY-MM-DD
	}
)
