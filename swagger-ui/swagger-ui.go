package swaggerui

import (
	"flag"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

var apiurl = flag.String("api", os.Getenv("HOST"), "The base path URI of the API service")

// SwaggerUIController swagger ui controller
type SwaggerUIController struct{}

// NewSwaggerUIController mount swagger ui controller
func NewSwaggerUIController(e *echo.Echo) {
	handler := &SwaggerUIController{}
	flag.Parse()

	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler.IndexHandler)))
	e.Static("/swagger-ui", "./swagger-ui/")

	for apiKey := range apiDescriptionsJson {
		e.GET("/swagger-ui/"+apiKey, echo.WrapHandler(http.HandlerFunc(handler.APIDescriptionHandler)))
	}
}

// IndexHandler index handler
func (c *SwaggerUIController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	isJSONRequest := false

	if acceptHeaders, ok := r.Header["Accept"]; ok {
		for _, acceptHeader := range acceptHeaders {
			if strings.Contains(acceptHeader, "json") {
				isJSONRequest = true
				break
			}
		}
	}

	if isJSONRequest {
		w.Write([]byte(resourceListingJson))
	} else {
		http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
	}
}

// APIDescriptionHandler api description handler
func (c *SwaggerUIController) APIDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.Split(strings.Trim(r.RequestURI, "/"), "/")

	if json, ok := apiDescriptionsJson[apiKey[1]]; ok {
		t, e := template.New("desc").Parse(json)
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, *apiurl)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
