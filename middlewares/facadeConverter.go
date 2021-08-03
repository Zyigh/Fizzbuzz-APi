package middlewares

import (
	"context"
	"encoding/json"
	"fizzbuzz-api/facades"
	"fizzbuzz-api/models"
	"log"
	"net/http"
)

// ParamConverter struct that helps build a net/http compatible middleware that converts Request parameters into a
// FacadeInterface and joins it to the Context of the Request so they are accessible as an object
type ParamConverter struct {
	// facade the facade that will be bind in the Request
	facade facades.FacadeInterface
}

// setupFacadeToContext creates a new Context from Request context and returns it with the facade joined to it at key
// "facade"
func (p *ParamConverter) setupFacadeToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, facades.FacadeContextKey, p.facade)
}

// addFacadeToRequest net/http compatible middleware that tries to convert raw Request params into an in-app defined
// FacadeInterface
// If the deserialization didn't work, the middleware stops the middleware process and returns a 400 (Bad Request)
// Response
func (p *ParamConverter) addFacadeToRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		p.facade, err = p.facade.Deserialize(r)
		if err != nil {
			log.Println(err)
			jsonData, _ := json.Marshal(models.ErrorModel{Message: "Couldn't process invalid arguments"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonData)

			return
		}

		ctx := p.setupFacadeToContext(r.Context())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetParamConverter builds an instance of ParamConverter with the facade set to the instance and returns the net/http
// compatible middleware to be used by the router
func GetParamConverter(facade facades.FacadeInterface, next http.Handler) http.Handler {
	paramConverter := ParamConverter{facade: facade}

	return paramConverter.addFacadeToRequest(next)
}
