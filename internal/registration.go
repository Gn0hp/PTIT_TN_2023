package internal

import (
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

// AppManagement Use remote config like consul to manage service, it'd be better to access registered app
var AppManagement map[string]interface{}

func init() {
	if AppManagement == nil {
		AppManagement = make(map[string]interface{})
		logrus.Info("Init new app management")
	}
}
func RegisterApp(name string, app interface{}) {
	mux := sync.RWMutex{}
	mux.Lock()
	defer mux.Unlock()
	AppManagement[strings.ToLower(name)] = app
}
func GetApp(name string) interface{} {
	return AppManagement[strings.ToLower(name)]
}
