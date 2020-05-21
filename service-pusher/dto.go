package main

import (
	"time"
)

type CmdName string

const (
	CmdWhoAreYouRequest    CmdName = "WHO_ARE_YOU_REQ"
	CmdWhoAreYouResponse   CmdName = "WHO_ARE_YOU_RES"
	CmdAreYouThereRequest  CmdName = "ARE_YOU_THERE_REQ"
	CmdAreYouThereResponse CmdName = "ARE_YOU_THERE_RES"
	CmdLetsDoThisRequest   CmdName = "LETS_DO_THIS_REQ"
	CmdLetsDoThisResponse  CmdName = "LETS_DO_THIS_RES"
)

type (
	RequestDto struct {
		ID        string      `json:"id"`
		IDRelated string      `json:"id_related"`
		Cmd       CmdName     `json:"cmd"`
		CreatedAt time.Time   `json:"created_at"`
		Data      interface{} `json:"data"`
	}
)
