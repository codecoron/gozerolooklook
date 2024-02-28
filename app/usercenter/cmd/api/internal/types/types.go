// Code generated by goctl. DO NOT EDIT.
package types

type AddAddressReq struct {
	AddressInfo
}

type AddAddressResp struct {
	Id int64 `json:"id"`
}

type AddressInfo struct {
	ContactName   string       `json:"contactName"`
	ContactMobile string       `json:"contactMobile"`
	District      DistrictInfo `json:"district"`
	Detail        string       `json:"detail"`
	Postcode      string       `json:"postcode"`
	IsDefault     int64        `json:"isDefault"`
}

type AddressListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type AddressListResp struct {
	List []UserAddress `json:"list"`
}

type Contact struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type ContactDelReq struct {
	Id []int64 `json:"id"`
}

type ContactDelResp struct {
}

type ContactDetailReq struct {
	Id int64 `json:"id"`
}

type ContactDetailResp struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type ContactListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type ContactListResp struct {
	List []Contact `json:"list"`
}

type ConvertAddressReq struct {
	OriginalAddressInfo string `json:"originalAddressInfo"`
}

type ConvertAddressResp struct {
	AddressInfo
}

type CreateContactReq struct {
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type CreateContactResp struct {
	Id int64 `json:"id"`
}

type CreateSponsorReq struct {
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type CreateSponsorResp struct {
	Id int64 `json:"id"`
}

type DistrictInfo struct {
	Province DistrictItem `json:"province"`
	City     DistrictItem `json:"city"`
	County   DistrictItem `json:"county"`
	Town     DistrictItem `json:"town,omitempty"`
}

type DistrictItem struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type SetAdminReq struct {
	Id int64 `json:"id"`
}

type SetAdminResp struct {
}

type SponosorDetailReq struct {
	Id int64 `json:"id"`
}

type SponosorDetailResp struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type Sponsor struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type UpDateContactReq struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type UpDateContactResp struct {
	Id int64 `json:"id"`
}

type User struct {
	Id       int64  `json:"id"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Sex      int64  `json:"sex"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
	IsAdmin  int64  `json:"isAdmin"`
}

type UserAddress struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	AddressInfo
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type UserUpdateReq struct {
	Nickname  string  `json:"nickname"`
	Sex       int64   `json:"sex"`
	Avatar    string  `json:"avatar"`
	Info      string  `json:"info"`
	Signature string  `json:"signature"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type UserUpdateResp struct {
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
	Nickname      string `json:"nickname, optional"`
	Avatar        string `json:"avatar, optional"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
