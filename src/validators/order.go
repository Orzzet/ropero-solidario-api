package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateOrder(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"status":         []string{"required", "string"},
		"requesterName":  []string{"required", "string"},
		"requesterPhone": []string{"required", "string"},
		"items":          []string{"required"},
	}, r)
}

func PatchOrder(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"status":         []string{"string"},
		"requesterName":  []string{"string"},
		"requesterPhone": []string{"string"},
	}, r)
}
