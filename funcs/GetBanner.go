package funcs

import (
	"os"
	"strings"
)

func GetBanner(bannerName string) ([]string, error) {
	// Read banner from file
	bannerData, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		return nil, err
	}

	// Split banner data into lines
	banner := strings.Split(string(bannerData), "\n")

	return banner, nil
}
