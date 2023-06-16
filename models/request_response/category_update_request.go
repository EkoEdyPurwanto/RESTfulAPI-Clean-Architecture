package request_response

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required, min=1, max=255" json:"name"`
}
