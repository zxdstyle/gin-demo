package v1

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"reflect"
	"strings"
)

func FormVerify(input interface{}) (bool, error) {
	setVerifyMessage()

	//_ = validation.AddCustomFunc("Unique", Unique)

	valid := validation.Validation{}
	b, _ := valid.Valid(input)
	if !b {
		arr := strings.Split(valid.Errors[0].Key, ".")
		st := reflect.TypeOf(input).Elem()
		field, _ := st.FieldByName(arr[0])
		return false, errors.New(field.Tag.Get("attr")+valid.Errors[0].Message)
	}
	return true, nil
}

func setVerifyMessage()  {
	var MessageTemplates = map[string]string{
		"Required":     "不能为空",
		"Min":          "最小值 为 %d",
		"Max":          "最大值 为 %d",
		"Range":        "范围 为 %d 到 %d",
		"MinSize":      "最短长度 为 %d",
		"MaxSize":      "最大长度 为 %d",
		"Length":       "长度必须 为 %d",
		"Alpha":        "必须是有效的字母",
		"Numeric":      "必须是有效的数字",
		"AlphaNumeric": "必须是有效的字母或数字",
		"Match":        "必须匹配 %s",
		"NoMatch":      "必须不匹配 %s",
		"AlphaDash":    "必须是有效的字母、数字或连接符号(-_)",
		"Email":        "必须是有效的电子邮件地址",
		"IP":           "必须是有效的IP地址",
		"Base64":       "必须是有效的base64字符",
		"Mobile":       "必须是有效的手机号码",
		"Tel":          "必须是有效的电话号码",
		"Phone":        "必须是有效的电话或移动电话号码",
		"ZipCode":      "必须是有效的邮政编码",
	}

	validation.SetDefaultMessage(MessageTemplates)
}

//var Unique validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
//
//	logs.Debug("传入的字段名称:",key)
//	logs.Debug("传入的字段内容:",obj)
//	logs.Debug("传入的验证:",v)
//	logs.Debug("验证当前字段在表中不重复")
//}

