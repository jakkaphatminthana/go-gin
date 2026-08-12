package utils

// response data
type DocsDataObject[Payload any] struct {
	Data Payload `json:"data,omitempty"`
}

// response success message
type DocsSuccessMessage struct {
	Success DocsSuccessDetail `json:"success"`
}

type DocsSuccessDetail struct {
	Status  int               `json:"status"`
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

// response error
type (
	DocsErrorBadRequest struct {
		Error struct {
			Status  int               `json:"status" example:"400"`
			Message string            `json:"message" example:"Bad request"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}

	DocsErrorNotFound struct {
		Error struct {
			Status  int               `json:"status" example:"404"`
			Message string            `json:"message" example:"Not Found"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}

	DocsErrorInternalServerError struct {
		Error struct {
			Status  int               `json:"status" example:"500"`
			Message string            `json:"message" example:"Internal Server Error"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}
)
