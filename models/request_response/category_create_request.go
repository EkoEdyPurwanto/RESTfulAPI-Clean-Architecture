package request_response

type CategoryCreateRequest struct {
	Name string `validate:"required, min=1, max=255"`
}
