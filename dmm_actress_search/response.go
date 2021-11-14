package dmmapi

import (
	"strconv"
	"strings"
	"time"
)

type ActressInfo struct {
	ID            string
	Name          string
	Ruby          string
	Bust          int64
	Waist         int64
	Hip           int64
	Height        int64
	Age           int64
	SmallImageURL string
	LargeImageURL string
}

type apiResponse struct {
	Result apiResult `json:"result"`
}

type apiResult struct {
	Status    string       `json:"status"`
	Actresses []apiActress `json:"actress"`
}

type apiActress struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Ruby     string          `json:"ruby"`
	Bust     string          `json:"bust"`
	Waist    string          `json:"waist"`
	Hip      string          `json:"hip"`
	Height   string          `json:"height"`
	Birthday string          `json:"birthday"`
	ImageURL apiActressImage `json:"imageURL"`
}

type apiActressImage struct {
	Small string `json:"small"`
	Large string `json:"large"`
}

func extractFirstName(s string) string {
	if strings.Contains(s, "（") {
		parts := strings.SplitN(s, "（", 2)
		return parts[0]
	}

	return s
}

func birthDayToAge(s string, now time.Time) int64 {
	if s == "" {
		return -1
	}

	parts := strings.Split(s, "-") // s is like 1970-04-09
	yearStr := parts[0]
	monthStr := parts[1]
	dayStr := parts[2]
	if strings.HasPrefix(monthStr, "0") {
		monthStr = monthStr[1:]
	}
	if strings.HasPrefix(dayStr, "0") {
		dayStr = dayStr[1:]
	}

	num, _ := strconv.ParseInt(yearStr, 10, 32)
	year := int(num)
	num, _ = strconv.ParseInt(monthStr, 10, 32)
	month := int(num)
	num, _ = strconv.ParseInt(dayStr, 10, 32)
	day := int(num)

	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.Year() {
		years--
	}

	return int64(years)
}

func (a *apiResponse) toActressInfo() []*ActressInfo {
	ret := make([]*ActressInfo, 0)
	now := time.Now()

	for _, actress := range a.Result.Actresses {
		var info ActressInfo
		info.ID = actress.ID
		info.Name = extractFirstName(actress.Name)
		info.Ruby = extractFirstName(actress.Ruby)
		info.SmallImageURL = actress.ImageURL.Small
		info.LargeImageURL = actress.ImageURL.Large

		num, err := strconv.ParseInt(actress.Bust, 10, 64)
		if err != nil {
			info.Bust = -1
		} else {
			info.Bust = num
		}

		num, err = strconv.ParseInt(actress.Waist, 10, 64)
		if err != nil {
			info.Waist = -1
		} else {
			info.Waist = num
		}

		num, err = strconv.ParseInt(actress.Hip, 10, 64)
		if err != nil {
			info.Hip = -1
		} else {
			info.Hip = num
		}

		num, err = strconv.ParseInt(actress.Height, 10, 64)
		if err != nil {
			info.Height = -1
		} else {
			info.Height = num
		}

		info.Age = birthDayToAge(actress.Birthday, now)

		ret = append(ret, &info)
	}

	return ret
}
