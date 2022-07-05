package validate

import "github.com/go-playground/validator/v10"

func Exec(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return err
	}

	return nil
}
