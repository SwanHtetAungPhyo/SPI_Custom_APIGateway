package proxy

import (
	"bytes"
	logger2 "github.com/SwanHtetAungPhyo/api/log"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"

	"net/http"
)

var logger = logger2.GetLogger()

func DoWithClient(ctx *fiber.Ctx, targetUrl string, client *http.Client) error {
	req, err := http.NewRequest(ctx.Method(), targetUrl, bytes.NewReader(ctx.Body()))
	logger.Info("Proxying request to: ", ctx)
	logger.Info("Target URL: ", targetUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}

	for key, values := range ctx.GetReqHeaders() {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if ctx.Method() == "POST" || ctx.Method() == "PUT" {
		if contentType := ctx.Get("Content-Type"); contentType != "" {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error during proxying: %v", err)
		return ctx.Status(fiber.StatusGatewayTimeout).SendString("Request timed out")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error reading response body")
	}

	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))
	ctx.Status(resp.StatusCode)
	return ctx.Send(body)
}
