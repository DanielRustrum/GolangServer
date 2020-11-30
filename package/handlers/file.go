package handler

// TODO: Gzip Control

import (
	"net/http"

	serverHttp "github.com/DanielRustrum/Https-Go-Server/package/server/http"
)

//* Public

//FileServer is ...
func FileServer(publicDir string, routeMap map[string]string) *serverHttp.HTTPHandler {
	getContent := func(path string) (string, []byte, int) {
		content := []byte{}
		success := false
		statusCode := 500
		filename := ""

		if routeMap[path] != "" {
			content, success = getFile(publicDir + "/" + routeMap[path])
			if success {
				statusCode = 200
			}
		} else {
			fileCheck := routeMap[path] == "" &&
				ignoresRoot(publicDir, publicDir+path) &&
				isFile(publicDir+path)

			if fileCheck {
				content, success = getFile(publicDir + "/" + routeMap[path])
				if success {
					filename = fileFromPath(path)
					statusCode = 200
				}
			}

			statusCode = 404
		}

		return filename, content, statusCode
	}

	getErrorContent := func(statusCode int) (string, []byte) {
		content := []byte{}
		filename := ""

		switch statusCode {
		case 400:
			filename = fileFromPath(routeMap["400"])
			content, _ = getFile(publicDir + "/" + routeMap["400"])
		case 404:
			filename = fileFromPath(routeMap["404"])
			content, _ = getFile(publicDir + "/" + routeMap["404"])
		case 500:
			filename = fileFromPath(routeMap["500"])
			content, _ = getFile(publicDir + "/" + routeMap["500"])
		default:
			filename = fileFromPath(routeMap[""])
			content, _ = getFile(publicDir + "/" + routeMap[""])
		}

		return filename, content
	}

	id := serverHttp.GetHandlerID()
	return &serverHttp.HTTPHandler{
		ID: id,
		Handler: func(response http.ResponseWriter, request *http.Request) {
			filename, content, statusCode := getContent(request.URL.Path)

			if statusCode != 200 {
				filename, content = getErrorContent(statusCode)
			}

			contentType := getContentType(filename)

			response.Write(content)
			response.Header().Set("Content-Type", contentType)
		},
	}
}
