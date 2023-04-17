package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/traefik/hub-agent-kubernetes/pkg/version"
)

// GetUserGroups get the groups of a user given its email address.
func (c *Client) GetUserGroups(ctx context.Context, userEmail string) ([]string, error) {
	baseURL, err := c.baseURL.Parse(path.Join(c.baseURL.Path, "users", userEmail, "groups"))
	if err != nil {
		return nil, fmt.Errorf("parse endpoint: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	version.SetUserAgent(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		all, _ := io.ReadAll(resp.Body)

		apiErr := APIError{StatusCode: resp.StatusCode}
		if err = json.Unmarshal(all, &apiErr); err != nil {
			apiErr.Message = string(all)
		}

		return nil, apiErr
	}

	var groups []string
	if err = json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, fmt.Errorf("decode groups: %w", err)
	}

	return groups, nil
}
