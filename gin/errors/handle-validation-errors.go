package ginErrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Kamaalio/kamaalgo/strings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(context *gin.Context, err error, placeOfFailure string) bool {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return false
	}

	for _, err := range validationErrors {
		ErrorHandler(context, Error{
			Status:  http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("'%s' is %s in the %s", strings.PascalToSnakeCase(err.Field()), err.Tag(), placeOfFailure),
		})

		return true
	}

	return false
}
