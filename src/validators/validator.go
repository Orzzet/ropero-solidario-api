package validators

import (
	"fmt"
	"github.com/orzzet/ropero-solidario-api/src/services"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type Validator struct {
	Service *services.Service
}

func New(service *services.Service) *Validator {
	validator := &Validator{
		Service: service,
	}
	govalidator.AddCustomRule("categories", validator.validateCategories)
	govalidator.AddCustomRule("string", validator.validateString)
	govalidator.AddCustomRule("categoryExists", validator.validateCategoryExists)
	govalidator.AddCustomRule("orderLines", validator.validateOrderLines)
	govalidator.AddCustomRule("uniqueUserEmail", validator.validateUniqueUserEmail)
	return validator
}

func validate(rules govalidator.MapData, r *http.Request) (data map[string]interface{}, validation url.Values) {
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

func (v *Validator) validateString(field string, rule string, message string, valueData interface{}) error {
	_, ok := valueData.(string)
	if !ok {
		return fmt.Errorf("string")
	}
	return nil
}
