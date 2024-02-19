package types

import (
	"encoding/xml"
	wxModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
)

type MsgEvent struct {
	wxModels.CallbackMessageHeader
	SubscribeMsgPopupEvent SubscribeMsgPopupEvent
}

type SubscribeMsgPopupEvent struct {
	XMLName xml.Name               `xml:"SubscribeMsgPopupEvent"`
	List    []UserSubscribeSetting `xml:"List"`
}

type UserSubscribeSetting struct {
	TemplateId            string `xml:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString"`
	PopupScene            string `xml:"PopupScene"`
}
