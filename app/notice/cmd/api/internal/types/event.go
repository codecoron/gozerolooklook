package types

import (
	wxModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

type MsgEvent struct {
	wxModels.CallbackMessageHeader
	List []UserSubscribeSetting `xml:"List"`
}

type UserSubscribeSetting struct {
	TemplateId            string `xml:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString"`
	PopupScene            string `xml:"PopupScene"`
}
