package handler

import (
	"io/ioutil"
	"os"
	"strings"
)

func ignoresRoot(rootDir string, path string) bool {
	rootList := strings.Split(rootDir, "/")
	pathList := strings.Split(path, "/")

	rootLength := len(rootList)
	pathLength := 0

	for _, value := range pathList {
		if value == "." || value == "" {
			continue
		} else if value == ".." {
			pathLength--
		} else {
			pathLength++
		}
	}

	return (rootLength < pathLength)
}

func isFile(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getFile(path string) ([]byte, bool) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte{}, false
	}

	return content, true
}

func fileFromPath(path string) string {
	pathList := strings.Split(path, "/")
	return pathList[len(pathList)-1]
}

func getContentType(file string) string {
	extensionList := strings.Split(file, ".")
	switch extensionList[len(extensionList)-1] {

	//* Web Files
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	case "js":
		return "application/javascript"
	case "json":
		return "application/json"
	case "wasm":
		return "application/wasm"
	case "zip":
		return "application/zip"
	case "csv":
		return "text/csv"
	//* Fonts
	case "otf":
		return "font/otf"
	case "sfnt":
		return "font/sfnt"
	case "ttf":
		return "font/ttf"
	case "woff":
		return "font/woff"
	case "woff2":
		return "font/woff2"
	//* Audio
	case "mpeg":
		return "audio/mpeg"
	case "wav":
		return "audio/wave"
	case "ogg":
		return "application/ogg"
	//* Images
	case "ico":
		return "image/x-icon"
	case "cur":
		return "image/x-icon"
	case "apng":
		return "image/apng"
	case "bmp":
		return "image/bmp"
	case "gif":
		return "image/gif"
	case "jpeg":
		return "image/jpeg"
	case "jpg":
		return "image/jpeg"
	case "jfif":
		return "image/jpeg"
	case "pjpeg":
		return "image/jpeg"
	case "pjp":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "svg":
		return "image/svg+xml"
	case "tiff":
		return "image/tiff"
	case "tif":
		return "image/tiff"
	case "webp":
		return "image/webp"
	//* Video
	case "webm":
		return "video/webm"
	case "mp4":
		return "video/mp4"
	case "mov":
		return "video/quicktime"
	case "qt":
		return "video/quicktime"
	//* Other
	case "xml":
		return "application/xml"
	case "pdf":
		return "application/pdf"
	//* Default To Type
	default:
		return "text/plain"

	}
}
