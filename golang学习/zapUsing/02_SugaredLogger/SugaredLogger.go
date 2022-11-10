package main

import (
	"go.uber.org/zap"
	"net/http"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

// 我们通过调用主 logger 的.Sugar()方法来获取一个SugaredLogger
func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url) //sugarLogger的debug日志
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err) //sugarLogger的error日志
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url) //sugarLogger的Info日志
		resp.Body.Close()
	}
}
