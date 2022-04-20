package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateUser(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"name":     []string{"required", "string"},
		"email":    []string{"required", "email"},
		"password": []string{"required"},
		"role":     []string{"in:admin,superadmin"},
	}, r)
}

func (v *Validator) ResetPassword(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"password": []string{"required", "min:6"},
	}, r)
}
