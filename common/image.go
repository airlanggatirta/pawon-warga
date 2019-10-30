package common

import (
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var (
	// ImageURL defines imagine url
	ImageURL string
	// ImgixURL defines Imgix URL
	ImgixURL string
)

func GenerateImageUrl(image string, imageType string) (fullPath string, err error) {
	_, err = url.ParseRequestURI(image)
	if err != nil {
		// image uploaded using imagine
		var width, height int64
		switch imageType {
		case "cover":
			width = 664
			height = 357
		case "campaign-card":
			width = 368
			height = 196
		case "avatar":
			width = 150
			height = 150
		}

		fullPath = GetImageURL(image, width, height)
	} else {
		// old image
		fullPath, err = GetOldImageURL(image, imageType)
		if err != nil {
			return "", err
		}
	}

	return fullPath, nil
}

// GenerateImgixURL will generate image link using imgix
func GenerateImgixURL(image string, imageType string) (fullPath string, err error) {
	_, err = url.ParseRequestURI(image)
	if err != nil {
		// image uploaded using Imgix
		fullPath = GetImgixURL(image)
	} else {
		// old image
		fullPath, err = GetOldImageURL(image, imageType)
		if err != nil {
			return "", err
		}
	}
	return fullPath, nil
}

// GetImageURL returns image url using imagine service
func GetImageURL(key string, width int64, height int64) (fullPath string) {
	endpoint := ImageURL
	prefix := "size"

	if height != 0 {
		fullPath = fmt.Sprintf("%s/%s/%dx%d/%s", endpoint, prefix, width, height, key)
	} else {
		fullPath = fmt.Sprintf("%s/%s/%d/%s", endpoint, prefix, width, key)
	}

	return
}

// GetOldImageURL returns image url if using old S3 url
func GetOldImageURL(image, imageType string) (fullPath string, err error) {
	componentURL, err := url.Parse(image)
	if err != nil {
		return "", err
	}

	scheme := componentURL.Scheme
	hostname := componentURL.Hostname()
	path := componentURL.Path
	s := strings.Split(path, ".")

	var suffix string
	switch imageType {
	case "cover":
		suffix = "f"
	case "campaign-card", "avatar":
		suffix = "s"
	}

	fullPath = fmt.Sprintf("%s://%s%s_%s.%s", scheme, hostname, s[0], suffix, s[1])
	return fullPath, nil
}

// GetImgixURL returns master image URL from Imgix
func GetImgixURL(key string) (fullPath string) {
	fullPath = fmt.Sprintf("%s/%s", ImgixURL, key)
	return
}

func GenerateImageFromBase64(b64 string) (uuidImage uuid.UUID, image string, err error) {
	coI := strings.Index(b64, ",")
	rawImage := b64[coI+1:]
	decodeBase64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(rawImage))

	uuidImage = uuid.Must(uuid.NewRandom())
	prefix := "temp_avatar_"

	var imageName string

	switch strings.TrimSuffix(b64[5:coI], ";base64") {
	case "image/png":

		pngImage, err := png.Decode(decodeBase64)
		if err != nil {
			return uuid.UUID{}, image, err
		}

		ext := "png"
		imageName = fmt.Sprintf("temp/%s%s.%s", prefix, uuidImage, ext)

		file, err := os.OpenFile(imageName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return uuid.UUID{}, image, err
		}

		png.Encode(file, pngImage)

	case "image/jpeg":

		jpgImage, err := jpeg.Decode(decodeBase64)
		if err != nil {
			return uuid.UUID{}, image, err
		}

		ext := "png"
		imageName = fmt.Sprintf("temp/%s%s.%s", prefix, uuidImage, ext)

		file, err := os.OpenFile(imageName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return uuid.UUID{}, image, err
		}

		jpeg.Encode(file, jpgImage, &jpeg.Options{Quality: 75})
	}

	image = imageName

	return uuidImage, image, err
}

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

func GetImageFromContentHTML(content string) []string {
	imgs := imgRE.FindAllStringSubmatch(content, -1)
	out := make([]string, len(imgs))

	for i := range out {
		out[i] = imgs[i][1]
	}

	return out
}
