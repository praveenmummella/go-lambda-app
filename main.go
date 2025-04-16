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

// Shared business logic
func generateResponse(message string) string {
	return fmt.Sprintf("Hello from Infovision %s!", message)
}

// Local handler (unchanged)
func localHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, generateResponse("Local Go Server"))
}

// Lambda handler
func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // Get the path from the request
    path := request.Path
    if path == "" {
        path = "/"  // Default to root path
    }

    // Prepare the response (matches your local output format)
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       fmt.Sprintf("Hello from Infovision AWS Lambda!\nPath: %s\n", path),
    }, nil
}

func main() {
	// Auto-detect environment
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		// Local
		http.HandleFunc("/", localHandler)
		log.Println("Running locally on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		// Lambda
		lambda.Start(lambdaHandler)
	}
}
