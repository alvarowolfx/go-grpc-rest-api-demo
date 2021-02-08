package validations

import (
	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Value       string `json:"value"`
	Error       string `json:"error"`
}

var (
	validate = validator.New()
)

// ValidateStruct Check if struct have errors and return properly formatted errors
func ValidateStruct(requestStruct interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(requestStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Value = err.Param()
			//element.Error = err.ActualTag()
			element.Error = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}

func GroupValidationErrorAsGrpcError(msg string, errors []*ErrorResponse) error {
	st := status.New(codes.InvalidArgument, msg)
	br := &errdetails.BadRequest{}
	for _, err := range errors {
		fv := &errdetails.BadRequest_FieldViolation{
			Field:       err.FailedField,
			Description: err.Error,
		}
		br.FieldViolations = append(br.FieldViolations, fv)
	}
	st, err := st.WithDetails(br)
	if err != nil {
		return status.Error(codes.Internal, "Failed to build error msg")
	}
	return st.Err()
}

func RegisterValidation(tag string, fn validator.Func) error {
	return validate.RegisterValidation(tag, fn)
}
