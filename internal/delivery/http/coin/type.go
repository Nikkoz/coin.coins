package coin

type (
	ID struct {
		Value uint `json:"id" uri:"id" binding:"required"`
	}

	Response struct {
		ID string `json:"id" binding:"required"`

		Short
	}

	Short struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	List struct {
		Total uint64 `json:"total" default:"0"`
		Limit uint64 `json:"limit" default:"10"`
		Page  uint64 `json:"page" default:"0"`

		Data []*Response `json:"data"`
	}
)
