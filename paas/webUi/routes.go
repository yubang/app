package webUi

import (
	"../../ctsFrame/webTools"
)

var routes = map[string]func(request webTools.HttpRequest)webTools.HttpRequest{
	"/index.html": indexHtml,
}
