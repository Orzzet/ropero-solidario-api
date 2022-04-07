package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateUser(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"name":     []string{"required", "alpha_space"},
		"email":    []string{"required", "email"},
		"password": []string{"required"},
		"role":     []string{"in:admin,superadmin"},
	}, r)
}
