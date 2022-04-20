package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateToken(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}, r)
}
