package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	dmmapi "github.com/syohex/dmm_actress_search/go"
)

var (
	bustParam    string
	waistParam   string
	hipParam     string
	heightParam  string
	ageParam     string
	sortParam    string
	keywordParam string
	debug        bool
)

func init() {
	flag.StringVar(&bustParam, "bust", "", "bust")
	flag.StringVar(&waistParam, "waist", "", "waist")
	flag.StringVar(&hipParam, "hip", "", "hip")
	flag.StringVar(&heightParam, "height", "", "height")
	flag.StringVar(&ageParam, "age", "", "age")
	flag.StringVar(&sortParam, "sort", "name", "how to sort result")
	flag.StringVar(&keywordParam, "keyword", "", "keyword")
	flag.BoolVar(&debug, "debug", false, "debug")
}

func readConfigFile() (*dmmapi.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(home, ".config", "dmm", "config.json")
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", configPath, err)
	}

	var config dmmapi.Config
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil, fmt.Errorf("could not parse config file as JSON: %w", err)
	}

	if config.ApiID == "" || config.AffiliateID == "" {
		return nil, fmt.Errorf("both API ID and affiliate ID must not be empty string")
	}

	config.Debug = debug
	return &config, nil
}

func readConfig() (*dmmapi.Config, error) {
	config, err := readConfigFile()
	if err == nil {
		return config, nil
	}

	apiID := os.Getenv("DMM_API_ID")
	affiliateID := os.Getenv("DMM_AFFILIATE_ID")
	if apiID == "" || affiliateID == "" {
		return nil, fmt.Errorf("neither config file nor environment variables are found")
	}

	return &dmmapi.Config{ApiID: apiID, AffiliateID: affiliateID}, nil
}

func main() {
	os.Exit(_main())
}

func constructSearchParam() (*dmmapi.SearchParam, error) {
	minBust, maxBust, err := dmmapi.ParseNumericParam(bustParam)
	if err != nil {
		return nil, fmt.Errorf("invalid bust param: %w", err)
	}

	minWaist, maxWaist, err := dmmapi.ParseNumericParam(waistParam)
	if err != nil {
		return nil, fmt.Errorf("invalid waist param: %w", err)
	}

	minHip, maxHip, err := dmmapi.ParseNumericParam(hipParam)
	if err != nil {
		return nil, fmt.Errorf("invalid hip param: %w", err)
	}

	minHeight, maxHeight, err := dmmapi.ParseNumericParam(heightParam)
	if err != nil {
		return nil, fmt.Errorf("invalid height param: %w", err)
	}

	minAge, maxAge, err := dmmapi.ParseNumericParam(ageParam)
	if err != nil {
		return nil, fmt.Errorf("invalid age param: %w", err)
	}

	if !dmmapi.IsValidSortType(sortParam) {
		return nil, fmt.Errorf("invalid sort param: %s", sortParam)
	}

	return dmmapi.NewSearchParam(
		dmmapi.WithBust(minBust, maxBust),
		dmmapi.WithWaist(minWaist, maxWaist),
		dmmapi.WithHip(minHip, maxHip),
		dmmapi.WithHeight(minHeight, maxHeight),
		dmmapi.WithAge(minAge, maxAge),
		dmmapi.WithSortType(sortParam),
		dmmapi.WithKeyword(keywordParam),
	), nil
}

func _main() int {
	flag.Parse()

	conf, err := readConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	searchParam, err := constructSearchParam()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	client := dmmapi.NewClient(conf)
	res, err := client.Get(searchParam)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	for i, info := range res {
		fmt.Printf("%2d: %s(%s)\n", i+1, info.Name, info.LargeImageURL)
	}

	return 0
}
