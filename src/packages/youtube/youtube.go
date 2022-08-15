package youtubepkg

import (
	"errors"
	"fmt"
	"net/url"
)

func GetIdFromLink(link string) (string, error) {
	result, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	query, err := url.ParseQuery(result.RawQuery)
	if err != nil {
		return "", err
	}

	val, ok := query["v"]
	if !ok {
		return "", errors.New("url not valid")
	}

	if len(val) == 0 {
		return "", errors.New("url not valid")
	}

	return val[0], nil

}

func GetLinkThumbnailFromId(id string) string {
	return fmt.Sprintf("https://img.youtube.com/vi/%v/0.jpg", id)
}
