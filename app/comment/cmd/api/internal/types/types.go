// Code generated by goctl. DO NOT EDIT.
package types

type Comment struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`      // 用户id
	LotteryId   int64  `json:"lotteryId"`   // 抽奖id
	PrizeName   string `json:"prizeName"`   // 奖品名称
	Content     string `json:"content"`     // 晒单评论内容
	Pics        string `json:"pics"`        // 晒单评论图片
	PraiseCount int64  `json:"praiseCount"` // 点赞数量
	CreateTime  int64  `json:"createTime"`  // 创建时间
	UpdateTime  int64  `json:"updateTime"`  // 更新时间
	DeleteTime  int64  `json:"deleteTime"`  // 删除时间
	DelState    int64  `json:"delstate"`
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

type Comments struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`      // 用户id
	LotteryId   int64  `json:"lotteryId"`   // 抽奖id
	PrizeName   string `json:"prizeName"`   // 奖品名称
	Content     string `json:"content"`     // 晒单评论内容
	Pics        string `json:"pics"`        // 晒单评论图片
	PraiseCount int64  `json:"praiseCount"` // 点赞数量
	CreateTime  int64  `json:"createTime"`  // 创建时间
	UpdateTime  int64  `json:"updateTime"`  // 更新时间
	DeleteTime  int64  `json:"deleteTime"`  // 删除时间
	DelState    int64  `json:"delstate"`
	User        User   `json:"user"`     // 用户信息
	IsPraise    int64  `json:"isPraise"` // 是否点赞
}

type TestReq struct {
	Age        int64  `json:"age" validate:"gte=1,lte=130"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"repassword" validate:"required"`
	Date       string `json:"date" validate:"required,datetime=2006-01-02"`
}

type TestResp struct {
}

type CommentAddReq struct {
	LotteryId int64  `json:"lotteryId" validate:"required"`
	PrizeName string `json:"prizeName" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Pics      string `json:"pics" validate:"required"`
}

type CommentAddResp struct {
}

type CommentDelReq struct {
	Id int64 `json:"id" validate:"required"`
}

type CommentDelResp struct {
}

type CommentUpdateReq struct {
	Id        int64  `json:"id" validate:"required"`
	UserId    int64  `json:"userId" validate:"required"`
	LotteryId int64  `json:"lotteryId" validate:"required"`
	PrizeName string `json:"prizeName" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Pics      string `json:"pics" validate:"required"`
}

type CommentUpdateResp struct {
}

type CommentListReq struct {
	LastId   int64 `json:"lastId"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Sort     int64 `json:"sort"`
}

type CommentListResp struct {
	List []Comments `json:"list"`
}

type CommentPraiseReq struct {
	Id int64 `json:"id" validate:"required"`
}

type CommentPraiseResp struct {
}

type CommentDetailReq struct {
	Id int64 `json:"id" validate:"required"`
}

type CommentDetailResp struct {
	Comment Comment `json:"comment"`
}

type GetCommentLastIdReq struct {
}

type GetCommentLastIdResp struct {
	LastId int64 `json:"lastId"`
}
