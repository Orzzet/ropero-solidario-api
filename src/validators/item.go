package validators

import (
	"fmt"
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

func (v *Validator) checkItemExists(valueData interface{}) error {
	itemID, ok := valueData.(float64)
	if !ok {
		return fmt.Errorf("itemId should be a number and its: '%s'", valueData.(string))
	}

	if _, err := v.Service.GetItem(uint(itemID)); err != nil {
		return fmt.Errorf("not found")
	}
	return nil
}
