package utils

import (
	"net/url"
	"strconv"
	"time"
)

const ISO8601 = "2006-01-02T15:04:05Z0700"

func QueryInt(query url.Values, queryName string, fallback int) (int, error) {
	numberStr := query.Get(queryName)
	if numberStr == "" {
		return fallback, nil
	}

	return strconv.Atoi(numberStr)
}

func QueryTime(query url.Values, queryName string, fallback time.Time) (time.Time, error) {
	timeStr := query.Get(queryName)
	if timeStr == "" {
		return fallback, nil
	}

	return time.Parse(ISO8601, timeStr)
}
