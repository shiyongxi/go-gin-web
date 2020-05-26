package request

type (
	DemoRequest struct {
		Name string `json:"name" form:"name"`
	}
)
