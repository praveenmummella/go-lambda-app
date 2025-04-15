package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Initialize connections (runs once per Lambda instance)
func init() {
	// Add DB/API clients here later
	log.Println("Initializing Lambda")
}

// Local development handler
func localHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go! (Local)\nPath: %s\n", r.URL.Path)
}

// Lambda handler with full request/response support
func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body: fmt.Sprintf(
			"Hello from Lambda!\nPath: %s\nQuery: %v\n",
			request.Path,
			request.QueryStringParameters,
		),
	}
	return response, nil
}

func main() {
	// Check if running in Lambda (with fallback for SAM local)
	isLambda := os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" || os.Getenv("AWS_EXECUTION_ENV") != ""

	if !isLambda {
		// Local mode
		http.HandleFunc("/", localHandler)
		log.Println("Starting local server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		// Lambda mode
		lambda.Start(lambdaHandler)
	}
}
