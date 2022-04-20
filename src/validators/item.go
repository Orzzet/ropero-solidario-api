package validators

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateItem(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"name":     []string{"required", "string"},
		"category": []string{"required", "categoryExists"},
		"amount":   []string{"required"},
	}, r)
}

func (v *Validator) DeleteItem(ID uint) (validation url.Values) {
	isInUse := v.Service.IsItemInUse(ID)
	if isInUse {
		return map[string][]string{"itemId": []string{"in use"}}
	}
	return nil
}
