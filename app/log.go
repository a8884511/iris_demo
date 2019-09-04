package app

import (
	"Goo/app/util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"os"
	"path"
	"strings"
	"time"
)

const deleteFileOnExit = true

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

var consoleLogger = logger.New(logger.Config{
	Status:             true,
	IP:                 true,
	Method:             true,
	Path:               true,
	Query:              true,
	MessageContextKeys: []string{"logger_message"},
	MessageHeaderKeys:  []string{"User-Agent"},
})

func todayFileName() string {
	today := time.Now().Format("2006_01_02")
	return today + ".log"
}

func newLogFile() *os.File {
	fileName := todayFileName()
	fileFolder := "log"
	exist, err := util.PathExist(fileFolder)
	if !exist {
		err = os.Mkdir(fileFolder, os.ModePerm)
	}
	f, err := os.OpenFile(path.Join(fileFolder, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return f
}

func newFileLogger() (handler iris.Handler, close func()) {
	c := logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	}

	logFile := newLogFile()
	close = func() {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		if err != nil {
			panic(err)
		}
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}
	c.AddSkipper(func(ctx iris.Context) bool {
		p := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(p, ext) {
				return true
			}
		}
		return false
	})
	handler = logger.New(c)
	return handler, close
}
