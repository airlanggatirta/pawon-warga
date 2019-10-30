package common

import "fmt"

func GenerateShareableLinks(sharedURL string, platform string) string {
	switch platform {
	case "facebook":
		return fmt.Sprintf("https://www.facebook.com/sharer/sharer.php?u=%s", sharedURL)
	}

	return ""
}
