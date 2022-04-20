package validators

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateOrder(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"status":         []string{"required", "string"},
		"requesterName":  []string{"required", "string"},
		"requesterPhone": []string{"required", "string"},
		"lines":          []string{"required", "orderLines"},
	}, r)
}

func (v *Validator) PatchOrder(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"status":         []string{"string"},
		"requesterName":  []string{"string"},
		"requesterPhone": []string{"string"},
	}, r)
}

func (v *Validator) validateOrderLines(field string, rule string, message string, valueData interface{}) error {
	values, ok := valueData.([]interface{})
	if !ok {
		return fmt.Errorf("array")
	}
	for i, value := range values {
		line, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%d object", i)
		}
		itemId, ok := line["itemId"]
		if !ok {
			return fmt.Errorf("%d.itemId required", i)
		}
		itemIdFloat, ok := itemId.(float64)
		if !ok {
			return fmt.Errorf("%d.itemId number", i)
		}
		if _, err := v.Service.GetItem(uint(itemIdFloat)); err != nil {
			return fmt.Errorf("%d.itemId not found", i)
		}

		amount, ok := line["amount"]
		if !ok {
			return fmt.Errorf("%d.amount required", i)
		}
		_, ok = amount.(float64)
		if !ok {
			return fmt.Errorf("%d.amount number", i)
		}
	}
	return nil
}
