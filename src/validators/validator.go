package validators

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func Init() {
	govalidator.AddCustomRule("categories", validateCategories)
	govalidator.AddCustomRule("string", validateString)
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

func validateString(field string, rule string, message string, valueData interface{}) error {
	_, ok := valueData.(string)
	if !ok {
		return fmt.Errorf("string")
	}
	return nil
}
