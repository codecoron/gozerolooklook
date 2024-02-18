package wxnotice

import "strconv"

const (
	TmplateIdMessageLotteryDraw  = "QJRFIFSdWwJOj2h-YsQLv9gU2HDxKaz-jGXTzwRSedU"
	TmplateIdMessageLotteryStart = "kpcOV5_zgXIPgKY-FmTAOIldBr-b4bXIpcersk_poTY"
	TmplateIdMessageWishSign     = "9rh1Ou_Y8BamXI7DeL649pCw6m_On3opt73VBgEiv3I"
	_                            = iota
	TypeLotteryDraw
	TypeLotteryStart
	TypeWishSign
)

type (
	Message interface {
		Type() int
		TemplateId() string
	}

	Item struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}

	// MessageLotteryDraw 开奖提醒的消息内容
	MessageLotteryDraw struct {
		PrizeName   Item `json:"thing8"`
		LotteryName Item `json:"thing5"`
		RemindText  Item `json:"thing7"`
	}
	// MessageLotteryStart 抽奖开始提醒的消息内容
	MessageLotteryStart struct {
		LotteryName Item `json:"thing2"`
		RemindText  Item `json:"thing7"`
	}
	// MessageWishCheckin 签到开始提醒的消息内容
	MessageWishCheckin struct {
		CheckinType Item `json:"phrase12"`
		Reward      Item `json:"phrase1"`
		Accumulate  Item `json:"phrase11"`
		RemindText  Item `json:"thing14"`
	}
)

func (m *MessageLotteryDraw) TemplateId() string {
	return TmplateIdMessageLotteryDraw
}
func (m *MessageLotteryDraw) Type() int {
	return TypeLotteryDraw
}
func (m *MessageLotteryStart) TemplateId() string {
	return TmplateIdMessageLotteryStart
}
func (m *MessageLotteryStart) Type() int {
	return TypeLotteryStart
}

func (m *MessageWishCheckin) TemplateId() string {
	return TmplateIdMessageWishSign
}
func (m *MessageWishCheckin) Type() int {
	return TypeWishSign
}

// ConvertToChineseNumber 数字转汉字的方法，不大于10亿
func ConvertToChineseNumber(number int64) string {
	chineseNumbers := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	chineseUnits := []string{"", "十", "百", "千", "万", "十", "百", "千", "亿"}
	numberStr := strconv.FormatInt(number, 10)
	var result string
	unitIndex := len(numberStr) - 1
	for i, digit := range numberStr {
		digitInt, _ := strconv.Atoi(string(digit))
		if digitInt != 0 {
			result += chineseNumbers[digitInt] + chineseUnits[unitIndex-i]
		} else {
			// 处理零的情况
			if i > 0 && numberStr[i-1] != '0' {
				result += chineseNumbers[digitInt]
			}
		}
	}
	return result
}
