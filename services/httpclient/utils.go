package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

func DieLog(req *http.Request, res *http.Response) string {
	now := time.Now().Format(time.ANSIC)
	return fmt.Sprintf("[error] %v %s %v - %v\n", req.Method, req.URL, now, res.Status)
}
