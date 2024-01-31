package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func Validate(dataStruct interface{}) error {
	zh_ch := zh.New()
	validate := validator.New()
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 在这里注册自定义结构体/字段校验方法

	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")

	// 在这里注册自定义tag翻译
	// 注意！因为这里会使用到trans实例
	// 所以这一步注册要放到trans初始化的后面

	// 验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	// 注册特定验证标签的翻译
	registerEnableVoteTranslation(validate, trans)

	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 获取字段的值并替换错误消息中的占位符
			fieldValue := err.Value()
			errorMsg := ""

			// 根据字段名自定义错误消息
			switch err.Field() {
			case "LotteryId":
				errorMsg = fmt.Sprintf("奖品id范围只能是0和1，当前值：%v", fieldValue)
			case "EnableVote":
				errorMsg = fmt.Sprintf("开启按钮：1 开启 0关闭，当前值：%v", fieldValue)
			default:
				errorMsg = err.Translate(trans)
			}

			//return errors.New(err.Translate(trans))
			return errors.New(errorMsg)
		}
	}
	return nil
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func registerEnableVoteTranslation(validate *validator.Validate, trans ut.Translator) error {
	return validate.RegisterTranslation(
		"eq=0|eq=1", trans, func(ut ut.Translator) error {
			return ut.Add("eq=0|eq=1", "{0}只能是0或1", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("eq=0|eq=1", fe.Field())
			return t
		},
	)
}
