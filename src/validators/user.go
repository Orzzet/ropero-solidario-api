package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateUser(r *http.Request) (map[string]interface{}, url.Values) {
	rules := govalidator.MapData{
		"name":     []string{"required", "alpha_space"},
		"email":    []string{"required", "email"},
		"password": []string{"required"},
		"role":     []string{"in:admin,superadmin"},
	}
	data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &data,
	}
	validator := govalidator.New(opts)
	validation := validator.ValidateJSON()
	if len(validation) > 0 {
		return data, validation
	}
	return data, nil
}
