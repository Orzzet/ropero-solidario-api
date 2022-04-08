package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func Init() {
	govalidator.AddCustomRule("categories", validateCategories)
}

func Validate(rules govalidator.MapData, r *http.Request) (data map[string]interface{}, validation url.Values) {
	validator := govalidator.New(govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &data,
	})
	validation = validator.ValidateJSON()
	if len(validation) > 0 {
		return data, validation
	}
	return data, nil
}
