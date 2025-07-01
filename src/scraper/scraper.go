package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/therealnoob/novelGo/config"
)

type contents struct {
	name    string
	chapter uint16
	title   string
	content string
}

// FlareSolverrRequest represents the JSON payload sent to FlareSolverr
type FlareSolverrRequest struct {
	Cmd string `json:"cmd"`
	Url string `json:"url"`
}

// FlareSolverrResponse represents the JSON response from FlareSolverr
type FlareSolverrResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Solution struct {
		Url       string              `json:"url"`
		Status    int                 `json:"status"`
		Headers   map[string]string   `json:"headers"`
		Response  string              `json:"response"`
		Cookies   []map[string]string `json:"cookies"`
		UserAgent string              `json:"userAgent"`
	} `json:"solution"`
}

func Scrape(cfg *config.ConfigStruct) error {
	// create flaresolverr session (if it doesn't already exist)
	if err := flaresolverrSession("http://localhost:8191/v1", "sessions.create", "miketest"); err != nil {
		return err
	}
	defer flaresolverrSession("http://localhost:8191/v1", "sessions.destroy", "miketest")

	// Construct request body
	reqBody, _ := json.Marshal(
		map[string]string{
			"cmd":     "request.get",
			"url":     cfg.URL,
			"session": "miketest",
		},
	)

	// Send HTTP POST request && parse results
	response, err := http.Post(cfg.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("request error against %s: %w", cfg.URL, err)
	}
	defer response.Body.Close()
	respBody, _ := io.ReadAll(response.Body)
	respBodyStr := string(respBody)
	//var respBodyJson SessionResponse
	//if err := json.Unmarshal([]byte(respBodyStr), &respBodyJson); err != nil {
	//	return fmt.Errorf("error unmarshaling json: %w", err)
	//}

	fmt.Printf("response from request: %s\n", respBodyStr)
	fmt.Println("end of response")
	return nil
}
