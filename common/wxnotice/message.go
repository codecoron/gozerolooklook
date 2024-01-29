package wxnotice

const (
	TmplateIdMessageLotteryDraw = "QJRFIFSdWwJOj2h-YsQLv9gU2HDxKaz-jGXTzwRSedU"

	_ = iota
	TypeLotteryDraw
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
)

func (m *MessageLotteryDraw) TemplateId() string {
	return TmplateIdMessageLotteryDraw
}
func (m *MessageLotteryDraw) Type() int {
	return TypeLotteryDraw
}
