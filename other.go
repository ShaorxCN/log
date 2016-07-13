package log

import (
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
)

// ======== 兼容 qiniu/log   ===============
const (
	Ldebug int = int(LevelDebug)
	Linfo      = int(LevelInfo)
	Lwarn      = int(LevelWarn)
	Lerror     = int(LevelError)
	Lpanic     = int(LevelPanic)
	Lfatal     = int(LevelFatal)
)

// ======== 兼容 qiniu/log   ===============
func SetOutputLevel(l int) { v = Level(l) }

// ======== 兼容 wothing/log ===============

// TraceIn and TraceOut use in function in and out,reduce code line
// Example:
//	func test() {
//		user := User{Name: "zhangsan", Age: 21, School: "xayddx"}
//		service := "verification.GetVerifiCode"
//		defer log.TraceOut(log.TraceIn("12345", service, "user:%v", user))
//		....
//	}

// TraceIn 方法入口打印日志
func TraceIn(tag string, method string, format string, m ...interface{}) (string, string, time.Time) {
	startTime := time.Now()
	std.Tprintf(LevelInfo, tag, "calling "+method+", "+format, m...)
	return tag, method, startTime
}

// TraceCtx 方法入口打印日志
func TraceCtx(ctx context.Context, method string, format string, m ...interface{}) (string, string, time.Time) {
	tag := "-"
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			tag = md["tid"][0]
		}
	}
	startTime := time.Now()
	std.Tprintf(LevelInfo, tag, "calling "+method+", "+format, m...)
	return tag, method, startTime
}

// TraceOut 方法退出记录下消耗时间
func TraceOut(tag string, method string, startTime time.Time) {
	std.Tprintf(LevelInfo, tag, "finished "+method+", took %v", time.Since(startTime))
}

func Println(m ...interface{}) { std.Tprintf(LevelPrint, "", "", m...) }

func getTracerIDFromCtx(ctx context.Context) string {
	nid := "00000000-0000-0000-0000-000000000000"

	if ctx == nil {
		return nid
	}

	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			return md["tid"][0]
		}
	}
	return nid
}

func CtxDebugf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelDebug, getTracerIDFromCtx(ctx), format, m...)
}

func CtxDebug(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelDebug, getTracerIDFromCtx(ctx), "", m...)
}

func CtxInfof(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelInfo, getTracerIDFromCtx(ctx), format, m...)
}

func CtxInfo(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelInfo, getTracerIDFromCtx(ctx), "", m...)
}

func CtxWarnf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelWarn, getTracerIDFromCtx(ctx), format, m...)
}

func CtxWarn(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelWarn, getTracerIDFromCtx(ctx), "", m...)
}

func CtxErrorf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelError, getTracerIDFromCtx(ctx), format, m...)
}

func CtxError(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelError, getTracerIDFromCtx(ctx), "", m...)
}

func CtxFatal(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelFatal, getTracerIDFromCtx(ctx), "", m...)
}

func CtxFatalf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelFatal, getTracerIDFromCtx(ctx), format, m...)
}

func CtxFatalln(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelFatal, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPanic(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelPanic, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPanicf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelPanic, getTracerIDFromCtx(ctx), format, m...)
}

func CtxPanicln(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelPanic, getTracerIDFromCtx(ctx), "", m...)
}

func CtxStack(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelStack, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPrint(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelPrint, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPrintf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(LevelPrint, getTracerIDFromCtx(ctx), format, m...)
}

func CtxPrintln(ctx context.Context, m ...interface{}) {
	std.Tprintf(LevelPrint, getTracerIDFromCtx(ctx), "", m...)
}
