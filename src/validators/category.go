package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateCategories(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{}, r)
}
