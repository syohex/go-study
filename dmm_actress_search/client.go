package dmmapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const baseURL = "https://api.dmm.com/affiliate/v3/ActressSearch"

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{config: config}
}

func ageToDateString(base time.Time, age int64) string {
	birth := base.AddDate(int(-age), 0, 0)
	return fmt.Sprintf("%d-%02d-%02d", birth.Year(), birth.Month(), birth.Day())
}

func (c *Client) setQueryParameters(req *http.Request, param *SearchParam) {
	q := req.URL.Query()

	q.Add("api_id", c.config.ApiID)
	q.Add("affiliate_id", c.config.AffiliateID)

	if param.minBust > 0 {
		q.Add("gte_bust", strconv.FormatInt(param.minBust, 10))
	}
	if param.maxBust > 0 {
		q.Add("lte_bust", strconv.FormatInt(param.maxBust, 10))
	}
	if param.minWaist > 0 {
		q.Add("gte_waist", strconv.FormatInt(param.minWaist, 10))
	}
	if param.maxWaist > 0 {
		q.Add("lte_waist", strconv.FormatInt(param.maxWaist, 10))
	}
	if param.minHip > 0 {
		q.Add("gte_hip", strconv.FormatInt(param.minHip, 10))
	}
	if param.maxHip > 0 {
		q.Add("lte_hip", strconv.FormatInt(param.maxHip, 10))
	}
	if param.minHeight > 0 {
		q.Add("gte_height", strconv.FormatInt(param.minHeight, 10))
	}
	if param.maxHeight > 0 {
		q.Add("ete_height", strconv.FormatInt(param.maxHeight, 10))
	}

	base := time.Now()
	if param.minAge > 0 {
		ageStr := ageToDateString(base, param.minAge)
		q.Add("lte_birthday", ageStr)
	}
	if param.maxAge > 0 {
		ageStr := ageToDateString(base, param.maxAge)
		q.Add("gte_birthday", ageStr)
	}
	if param.sortType != "" {
		q.Add("sort", param.sortType)
	}
	if param.keyword > "" {
		q.Add("keyword", param.keyword)
	}

	q.Add("output", "json")

	req.URL.RawQuery = q.Encode()
}

func (c *Client) Get(param *SearchParam) ([]*ActressInfo, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	c.setQueryParameters(req, param)

	if c.config.Debug {
		fmt.Printf("[DEBUG] send request to %s\n", req.URL.RequestURI())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if c.config.Debug {
		fmt.Printf("[DEBUG] response=%s\n", string(body))
	}

	var apiResponse apiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("could not parse response: %w", err)
	}

	return apiResponse.toActressInfo(), nil
}
