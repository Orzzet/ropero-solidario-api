package validators

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func CreateCategories(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return Validate(govalidator.MapData{
		"categories": []string{"required", "categories"},
	}, r)
}

func validateCategories(field string, rule string, message string, valueData interface{}) error {
	values, ok := valueData.([]interface{})
	if !ok {
		return fmt.Errorf("array")
	}
	for i, value := range values {
		category, ok := value.(map[string]interface{})
		if !ok {
			fmt.Println(category["name"])
			return fmt.Errorf("%d object", i)
		}
		name, ok := category["name"]
		if !ok {
			return fmt.Errorf("%d.name required", i)
		}
		if _, ok := name.(string); !ok {
			return fmt.Errorf("%d.name string", i)
		}
		parentCategoryId, ok := category["parentCategoryId"]
		if !ok {
			return fmt.Errorf("%d.parentCategoryId required", i)
		}
		if _, ok := parentCategoryId.(float64); !ok {
			return fmt.Errorf("%d.parentCategoryId number", i)
		}
	}
	return nil
}
