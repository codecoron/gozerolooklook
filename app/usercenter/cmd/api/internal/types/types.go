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

type SetAdminReq struct {
	Id int64 `json:"id"`
}

type SetAdminResp struct {
}

type ContactInfo struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Type    int64  `json:"type"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type ContactInfoReq struct {
	Id int64 `json:"id"`
}

type ContactInfoResp struct {
	ContactInfo ContactInfo `json:"contactInfo"`
}

type ConvertAddressReq struct {
	OriginalAddressInfo string `json:"originalAddressInfo"`
}

type ConvertAddressResp struct {
	AddressInfo
}

type CreateReq struct {
	Type    int64  `json:"type"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

type CreateResp struct {
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
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
