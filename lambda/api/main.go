package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
)
var chiLambda *chiadapter.ChiLambda

// handler is the function called by the lambda.
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

// main is called when a new lambda starts, so don't
// expect to have something done for every query here.
func main() {
	// init go-chi router
	r := chi.NewRouter()
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		_ = render.Render(w, r, &apiResponse{
			Status: http.StatusOK,
			URL: r.URL.String(),
			RequestBody: string(body),
		})
	})

	// start the lambda with a context
	lambda.StartWithContext(context.Background(), handler)
}

// apiResponse is the response to the API.
type apiResponse struct {
	Status int `json:"status_code,omitempty"`
	URL string `json:"url,omitempty"`
	RequestBody string `json:"request_body,omitempty"`
}

// Render is used by go-chi-render to render the JSON response.
func (a apiResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, a.Status)
	return nil
}
