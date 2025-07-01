package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type SessionResponse struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Session        string `json:"session"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	Version        string `json:"version"`
}

// Function to create or destroy a session against a flaresolverr instance.
//
// url is the the url of flaresolverr instance.
// cmd is the command to send to flaresolverr; `sessions.create` or `sessions.destroy`.
// session is the name of the session.
func flaresolverrSession(url string, cmd string, session string) error {
	// Input validation
	if cmd != "sessions.create" && cmd != "sessions.destroy" {
		return fmt.Errorf("cmd string must be either 'sessions.create' or 'sessions.destroy'")
	}

	if len(session) == 0 || len(session) > 64 {
		return fmt.Errorf("session string must be between 1 and 64 characters")
	}

	matched, err := regexp.MatchString(`^[a-z0-9_-]+$`, session)
	if err != nil {
		return fmt.Errorf("failed to validate session: %w", err)
	}
	if !matched {
		return fmt.Errorf("session contains invalid characters: must contain only lowercase letters, digits, hyphens, or underscores")
	}

	// Construct JSON payload
	reqBody, err := json.Marshal(
		map[string]string{
			"cmd":     cmd,
			"session": session,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Send HTTP POST request && parse results
	response, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error in %s request in %s: %w", cmd, url, err)
	}
	defer response.Body.Close()
	respBody, _ := io.ReadAll(response.Body)
	respBodyStr := string(respBody)
	var respBodyJson SessionResponse
	if err := json.Unmarshal([]byte(respBodyStr), &respBodyJson); err != nil {
		return fmt.Errorf("error unmarshaling json: %w", err)
	}

	// debug
	fmt.Printf("response from flaresolverr: %d, %s\n", response.StatusCode, respBodyStr)

	// return result
	if response.StatusCode != http.StatusOK &&
		((cmd == "sessions.create" &&
			(respBodyJson.Message != "Session created successfully." &&
				respBodyJson.Message != "Session already exists.")) ||
			(cmd == "sessions.destroy" &&
				(respBodyJson.Message != "The session has been removed." &&
					respBodyJson.Message != "The session doesn't exist."))) {
		return fmt.Errorf("Unexpected flareresolverr response: statusCode: %d, body: %s", response.StatusCode, respBodyStr)
	}

	return nil
}
