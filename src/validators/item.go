package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateItem(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"name":     []string{"required", "string"},
		"category": []string{"required"},
		"amount":   []string{"required"},
	}, r)
}
