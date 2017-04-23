package webUi

import "../../ctsFrame/webTools"

func indexHtml(r webTools.HttpRequest)webTools.HttpRequest{
	return webTools.SendFile("index.html", 0, r)
}
