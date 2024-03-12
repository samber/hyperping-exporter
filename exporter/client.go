package exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpRequest[T any](client *http.Client, token string, entity string) (T, error) {
	var output T

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", HyperPingEndpoint, entity), nil)
	if err != nil {
		return output, err
	}

	req.Header.Set("User-Agent", "samber/hyperping_exporter")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return output, fmt.Errorf("status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return output, err
	}

	return output, nil
}

// https://hyperping.notion.site/Hyperping-API-documentation-a0dc48fb818e4542a8f7fb4163ede2c3#b68fd1b79eeb4bcda5d36c78a54e6a0a
type monitor struct {
	Name               string   `json:"name"`
	URL                string   `json:"url"`
	UUID               string   `json:"uuid"`
	Paused             bool     `json:"paused"`
	Protocol           string   `json:"protocol"`
	ProjectUUID        string   `json:"projectUuid"`
	Port               int      `json:"port"`
	HTTPMethod         string   `json:"http_method"`
	Regions            []string `json:"regions"`
	CheckFrequency     int      `json:"check_frequency"`
	FollowRedirects    bool     `json:"follow_redirects"`
	ExpectedStatusCode string   `json:"expected_status_code"`
	// RequestBody        string   `json:"request_body"`
	// RequestHeaders     []string `json:"request_headers"`
	Status        string `json:"status"`
	SSLExpiration int    `json:"ssl_expiration"`
	AlertsWait    int    `json:"alerts_wait"`
}

// get monitors from hyperping
func (e *Exporter) getMonitors() ([]*monitor, error) {
	return httpRequest[[]*monitor](e.client, e.token, "monitors")
}
