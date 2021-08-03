package facades

import "net/http"

// FacadeInterface defines the shape of struct that can be used by the ParamConverter middleware
type FacadeInterface interface {
	// Deserialize methods that should be implemented to populate concrete type instance with Request params
	Deserialize(r *http.Request) (FacadeInterface, error)
}

// FacadeContextKey key name in http.Request.Context in which facades are stored and retrieved
const FacadeContextKey = "facade"
