package hh

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type AppInfo struct {
	AuthType      string `json:"auth_type"`
	IsAdmin       bool   `json:"is_admin"`
	IsApplicant   bool   `json:"is_applicant"`
	IsApplication bool   `json:"is_application"`
	IsEmployer    bool   `json:"is_employer"`
}

type Client struct {
	BaseUrl   *url.URL // https://api.hh.ru
	UserAgent string   // MyApp/1.0 (my-app-feedback@example.com)

	httpClient *http.Client
}

func (c *Client) Me() (*AppInfo, error) {
	rel := &url.URL{Path: "/me"}
	u := c.BaseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var appInfo *AppInfo
	err = json.NewDecoder(resp.Body).Decode(&appInfo)
	return appInfo, err
}