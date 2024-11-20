package proxy

import (
	"bytes"
	"fmt"
	"github.com/SwanHtetAungPhyo/api/jwt"
	logging "github.com/SwanHtetAungPhyo/api/log"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"net/http"
	"strings"
	"time"
)

var logger = logging.GetLogger()

func DoWithClient(ctx *fiber.Ctx, targetUrl string, client *http.Client) error {
	req, err := createRequest(ctx, targetUrl)
	if err != nil {
		return err
	}

	resp, err := retryRequest(client, req)
	if err != nil {
		return ctx.Status(fiber.StatusGatewayTimeout).SendString(err.Error())
	}
	defer closeResponseBody(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body: ", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error reading response body")
	}

	logger.Debug("Raw response body: ", string(body))

	if strings.Contains(targetUrl, "/login") {
		return handleLoginRequest(ctx, resp, body)
	}

	return handleProxyResponse(ctx, resp, body)
}

func createRequest(ctx *fiber.Ctx, targetUrl string) (*http.Request, error) {
	req, err := http.NewRequest(ctx.Method(), targetUrl, bytes.NewReader(ctx.Body()))
	if err != nil {
		logger.Error("Error creating request: ", err)
		return nil, ctx.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}

	for key, values := range ctx.GetReqHeaders() {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if (ctx.Method() == "POST" || ctx.Method() == "PUT") && ctx.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func retryRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	retries := 3
	delay := 2 * time.Second

	for i := 0; i < retries; i++ {
		resp, err := client.Do(req)
		if err == nil {
			return resp, nil
		}

		logger.Error("Error during proxying: ", err)
		if i < retries-1 {
			logger.Info("Retrying request...")
			time.Sleep(delay)
			delay *= 2
		}
	}

	return nil, fmt.Errorf("failed after %d retries", retries)
}

func closeResponseBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		logger.Error("Error closing response body: ", err)
	}
}

func handleLoginRequest(ctx *fiber.Ctx, resp *http.Response, body []byte) error {
	modifiedBody, err := SetJwtForLogin(body)
	if err != nil {
		logger.Error("Error modifying response body for login request: ", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error modifying response body")
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(resp.StatusCode).Send(modifiedBody)
}

func handleProxyResponse(ctx *fiber.Ctx, resp *http.Response, body []byte) error {
	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))
	ctx.Status(resp.StatusCode)
	return ctx.Send(body)
}

func SetJwtForLogin(body []byte) ([]byte, error) {
	UUID := utils.UUID()

	access, err := jwt.GenerateJWT(UUID, "ACCESS")
	if err != nil {
		logger.Error("Error generating access token: ", err)
		return nil, err
	}
	refresh, err := jwt.GenerateJWT(UUID, "REFRESH")
	if err != nil {
		logger.Error("Error generating refresh token: ", err)
		return nil, err
	}
	var responseMap struct {
		Body    string `json:"body"`
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}

	responseMap.Body = string(body)
	responseMap.Access = access
	responseMap.Refresh = refresh

	modifiedMap, err := json.Marshal(&responseMap)
	if err != nil {
		logger.Error("Error marshaling modified response map: ", err)
		return nil, err
	}

	logger.Debug("JWT tokens successfully added to login response")
	return modifiedMap, nil
}
