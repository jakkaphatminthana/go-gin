package exception

import (
	"fmt"
	"net/http"
)

type ProviderNotFound struct {
	ProviderID string
	Provider   string
}

func (e *ProviderNotFound) Error() string {
	return fmt.Sprintf("providerID: %s, provider: %s not found", e.ProviderID, e.Provider)
}

func (e *ProviderNotFound) StatusCode() int {
	return http.StatusNotFound
}
