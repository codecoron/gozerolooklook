package translator

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"looklook/app/lottery/cmd/api/internal/logic/lottery"
	"looklook/app/lottery/cmd/api/internal/types"
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
	// 注册自定义结构体校验方法
	validate.RegisterStructValidation(lottery.SignUpParamStructLevelValidation, types.TestReq{})

	// 注册自定义结构体字段校验方法
	if err := validate.RegisterValidation("checkDate", lottery.CheckDate); err != nil {
		return err
	}

	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")

	// 在这里注册自定义tag翻译
	// 注意！因为这里会使用到trans实例
	// 所以这一步注册要放到trans初始化的后面

	if err := validate.RegisterTranslation(
		"checkDate",
		trans,
		registerTranslator("checkDate", "{0}必须要晚于当前日期"),
		translate,
	); err != nil {
		return err
	}

	// 验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
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
