package main

import (
	"go.uber.org/zap"
	"net/http"
	"zapUsing/zapConfig"
)

var sugarLogger *zap.SugaredLogger

func main() {
	zapLogger := zapConfig.NewLogger()
	sugarLogger = zapLogger.GetSugarLogger()
	defer zapLogger.Close() //程序退出之前需要调用此方法更新日志条目,不然会造成丢失日志

	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
