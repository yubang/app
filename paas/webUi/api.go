package webUi

import "net/http"
import "../../ctsFrame/webTools"

func indexHtml(w http.ResponseWriter, r *http.Request){
	webTools.SendFile("index.html", 0, w)
}
