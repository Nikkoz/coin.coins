package coin

type (
	Response struct {
		ID string `json:"id" binding:"required,uint"`

		Short
	}

	Short struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	List struct {
		Total  uint64 `json:"total" default:"0"`
		Limit  uint64 `json:"limit" default:"10"`
		Offset uint64 `json:"offset" default:"0"`

		Data []*Response `json:"data"`
	}
)
