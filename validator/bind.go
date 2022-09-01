package validator

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Init() {
	// zh := zh.New()
	// uni = ut.New(zh, zh)
	// trans, _ := uni.GetTranslator("zh")
	// zh_translations.RegisterDefaultTranslations(validate, trans)

	Valids := make(map[string]validator.Func)
	Valids["username"] = username
	Valids["password"] = password

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for vname, vfunc := range Valids {
			v.RegisterValidation(vname, vfunc)
		}
	} else {
		panic("validator error")
	}
}

var username validator.Func = func(fl validator.FieldLevel) bool {

	return true
}

var password validator.Func = func(fl validator.FieldLevel) bool {
	return false
}
