package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateToken(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}, r)
}
