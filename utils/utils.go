package utils

import (
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

// Parses query parameters and only splits at & symbol
func ParseQueryParamsTags(c echo.Context) ([]string, error) {
	// Use the original URL query, where '+' is intact and '%2B' is not yet decoded.
	rawQuery := c.Request().URL.RawQuery

	// Replace '+' with a placeholder 'PLUS_SIGN' to prevent it from being decoded to space.
	preparedQuery := strings.ReplaceAll(rawQuery, "+", "PLUS_SIGN")

	// Now, decode query parameters manually.
	decodedQuery, err := url.QueryUnescape(preparedQuery)
	if err != nil {
		return nil, err
	}

	// Convert 'PLUS_SIGN' placeholders back to the '+' character
	decodedQuery = strings.ReplaceAll(decodedQuery, "PLUS_SIGN", "+")

	// Extract `tag` values
	values := strings.Split(decodedQuery, "&")
	var tags []string
	for _, value := range values {
		if strings.HasPrefix(value, "tag=") {
			tag := strings.TrimPrefix(value, "tag=")
			tags = append(tags, tag)
		}
	}

	return tags, nil
}
