package response

type (
	Version struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	BaseResponse struct {
		Status       string `json:"status"`
		Code         string `json:"code"`
		ErrorMessage string `json:"error_message"`
		Data         any    `json:"data"`
	}

	Login struct {
		Token string `json:"token"`
	}
)
