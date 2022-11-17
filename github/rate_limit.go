package github

import (
	"encoding/json"
)

type RateLimit struct {
	Limit     int
	Remaining int
	Reset     int
	Used      int
}

type Resource struct {
	// Other object (e.g. search) is not needed
	// https://docs.github.com/en/rest/rate-limit
	Core RateLimit
}

type RateLimits struct {
	Resources Resource
	// Deprecated
	Rate RateLimit
}

func FetchRateLimit(client Client) (*RateLimits, error) {
	body, err := client.Get("https://api.github.com/rate_limit")
	if err != nil {
		return nil, err
	}

	r := RateLimits{}
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}