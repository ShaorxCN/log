package main

import "github.com/gotips/log"

func execLevelExamples() {
	// 默认日志级别 debug
	log.Printf("default log level: %s", log.GetLevel())
	log.Tracef("IsTraceEnabled? %t", log.IsTraceEnabled())
	log.Debugf("IsDebugEnabled? %t", log.IsDebugEnabled())
	log.Infof("IsInfoEnabled? %t", log.IsInfoEnabled())

	// trace 级别
	log.SetLevel(log.LevelTrace)
	log.Tracef(msgFmt, 1)

	// info 级别
	log.SetLevel(log.LevelInfo)
	log.Debugf(msgFmt, 2)
	log.Infof(msgFmt, 2)

	// warn 级别
	log.SetLevel(log.LevelWarn)
	log.Infof(msgFmt, 3)
	log.Warnf(msgFmt, 3)

	// error 级别
	log.SetLevel(log.LevelError)
	log.Warnf(msgFmt, 4)
	log.Errorf(msgFmt, 4)

	// 恢复默认级别，防止影响其他测试
	// debug 级别
	log.SetLevel(log.LevelDebug)
	log.Tracef(msgFmt, 5)
	log.Debugf(msgFmt, 5)
}
