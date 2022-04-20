package validators

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateCategories(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"categories": []string{"required", "categories"},
	}, r)
}

func (v *Validator) CreateCategory(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"name": []string{"required", "string"},
	}, r)
}

func (v *Validator) validateCategories(field string, rule string, message string, valueData interface{}) error {
	values, ok := valueData.([]interface{})
	if !ok {
		return fmt.Errorf("array")
	}
	for i, value := range values {
		category, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%d object", i)
		}
		name, ok := category["name"]
		if !ok {
			return fmt.Errorf("%d.name required", i)
		}
		if _, ok := name.(string); !ok {
			return fmt.Errorf("%d.name string", i)
		}
	}
	return nil
}

func (v *Validator) validateCategoryExists(field string, rule string, message string, valueData interface{}) error {
	categoryID, ok := valueData.(float64)
	if !ok {
		return fmt.Errorf("categoryId should be a number and its: '%s'", valueData.(string))
	}

	if _, err := v.Service.GetCategory(uint(categoryID)); err != nil {
		return fmt.Errorf("category with id %d not found", uint(categoryID))
	}
	return nil
}
