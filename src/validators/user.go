package validators

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

func (v *Validator) CreateUser(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"name":     []string{"required", "string"},
		"email":    []string{"required", "email", "uniqueUserEmail"},
		"password": []string{"required"},
		"role":     []string{"in:admin,superadmin"},
	}, r)
}

func (v *Validator) ResetPassword(r *http.Request) (data map[string]interface{}, validation url.Values) {
	return validate(govalidator.MapData{
		"password": []string{"required", "min:6"},
	}, r)
}

func (v *Validator) validateUniqueUserEmail(field string, rule string, message string, valueData interface{}) error {
	email, ok := valueData.(string)
	if !ok {
		return fmt.Errorf("string")
	}

	if isUnique := v.Service.IsUserEmailUnique(email); !isUnique {
		return fmt.Errorf("unique")
	}
	return nil
}
