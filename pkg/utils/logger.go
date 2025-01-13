package utils

import "k8s.io/klog"

var (
	DebugMode = true
)

func Debug(format string, v ...interface{}) {
	if DebugMode {
		klog.V(2).Infof("[DEBUG] "+format, v...)
	}
}
