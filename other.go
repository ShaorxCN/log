package log

import (
	"io"
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
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
	std.Tprintf(Linfo, tag, "calling "+method+", "+format, m...)
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
	std.Tprintf(Linfo, tag, "calling "+method+", "+format, m...)
	return tag, method, startTime
}

// TraceOut 方法退出记录下消耗时间
func TraceOut(tag string, method string, startTime time.Time) {
	std.Tprintf(Linfo, tag, "finished "+method+", took %v", time.Since(startTime))
}

func Println(m ...interface{}) { std.Tprintf(Lprint, "", "", m...) }

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
	std.Tprintf(Ldebug, getTracerIDFromCtx(ctx), format, m...)
}

func CtxDebug(ctx context.Context, m ...interface{}) {
	std.Tprintf(Ldebug, getTracerIDFromCtx(ctx), "", m...)
}

func CtxInfof(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Linfo, getTracerIDFromCtx(ctx), format, m...)
}

func CtxInfo(ctx context.Context, m ...interface{}) {
	std.Tprintf(Linfo, getTracerIDFromCtx(ctx), "", m...)
}

func CtxWarnf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lwarn, getTracerIDFromCtx(ctx), format, m...)
}

func CtxWarn(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lwarn, getTracerIDFromCtx(ctx), "", m...)
}

func CtxErrorf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lerror, getTracerIDFromCtx(ctx), format, m...)
}

func CtxError(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lerror, getTracerIDFromCtx(ctx), "", m...)
}

func CtxFatal(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), "", m...)
}

func CtxFatalf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), format, m...)
}

func CtxFatalln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPanic(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPanicf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), format, m...)
}

func CtxPanicln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), "", m...)
}

func CtxStack(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lstack, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPrint(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), "", m...)
}

func CtxPrintf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), format, m...)
}

func CtxPrintln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), "", m...)
}

// -----------------

var Ldefault = Linfo

type XLogger struct{}

func New(w io.Writer, prefix string, l Level, depth int) *XLogger {
	return &XLogger{}
}

// 打印日志
func (*XLogger) Trace(m ...interface{}) { std.Tprintf(Ltrace, "", "", m...) }
func (*XLogger) Debug(m ...interface{}) { std.Tprintf(Ldebug, "", "", m...) }
func (*XLogger) Info(m ...interface{})  { std.Tprintf(Linfo, "", "", m...) }
func (*XLogger) Warn(m ...interface{})  { std.Tprintf(Lwarn, "", "", m...) }
func (*XLogger) Error(m ...interface{}) { std.Tprintf(Lerror, "", "", m...) }
func (*XLogger) Panic(m ...interface{}) { std.Tprintf(Lpanic, "", "", m...) }
func (*XLogger) Fatal(m ...interface{}) { std.Tprintf(Lfatal, "", "", m...) }
func (*XLogger) Print(m ...interface{}) { std.Tprintf(Lprint, "", "", m...) }
func (*XLogger) Stack(m ...interface{}) { std.Tprintf(Lstack, "", "", m...) }

// 按一定格式打印日志
func (*XLogger) Tracef(format string, m ...interface{}) { std.Tprintf(Ltrace, "", format, m...) }
func (*XLogger) Debugf(format string, m ...interface{}) { std.Tprintf(Ldebug, "", format, m...) }
func (*XLogger) Infof(format string, m ...interface{})  { std.Tprintf(Linfo, "", format, m...) }
func (*XLogger) Warnf(format string, m ...interface{})  { std.Tprintf(Lwarn, "", format, m...) }
func (*XLogger) Errorf(format string, m ...interface{}) { std.Tprintf(Lerror, "", format, m...) }
func (*XLogger) Panicf(format string, m ...interface{}) { std.Tprintf(Lpanic, "", format, m...) }
func (*XLogger) Fatalf(format string, m ...interface{}) { std.Tprintf(Lfatal, "", format, m...) }
func (*XLogger) Printf(format string, m ...interface{}) { std.Tprintf(Lprint, "", format, m...) }
func (*XLogger) Stackf(format string, m ...interface{}) { std.Tprintf(Lstack, "", format, m...) }

func (*XLogger) CtxDebugf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Ldebug, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxDebug(ctx context.Context, m ...interface{}) {
	std.Tprintf(Ldebug, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxInfof(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Linfo, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxInfo(ctx context.Context, m ...interface{}) {
	std.Tprintf(Linfo, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxWarnf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lwarn, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxWarn(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lwarn, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxErrorf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lerror, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxError(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lerror, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxFatal(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxFatalf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxFatalln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lfatal, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxPanic(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxPanicf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxPanicln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lpanic, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxStack(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lstack, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxPrint(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), "", m...)
}

func (*XLogger) CtxPrintf(ctx context.Context, format string, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), format, m...)
}

func (*XLogger) CtxPrintln(ctx context.Context, m ...interface{}) {
	std.Tprintf(Lprint, getTracerIDFromCtx(ctx), "", m...)
}

// 按一定格式打印日志，并在打印日志时带上 tag
func (*XLogger) Ttracef(tag string, format string, m ...interface{}) {
	std.Tprintf(Ltrace, tag, format, m...)
}
func (*XLogger) Tdebugf(tag string, format string, m ...interface{}) {
	std.Tprintf(Ldebug, tag, format, m...)
}
func (*XLogger) Tinfof(tag string, format string, m ...interface{}) {
	std.Tprintf(Linfo, tag, format, m...)
}
func (*XLogger) Twarnf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lwarn, tag, format, m...)
}
func (*XLogger) Terrorf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lerror, tag, format, m...)
}
func (*XLogger) Tpanicf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lpanic, tag, format, m...)
}
func (*XLogger) Tfatalf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lfatal, tag, format, m...)
}
func (*XLogger) Tprintf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lprint, tag, format, m...)
}
func (*XLogger) Tstackf(tag string, format string, m ...interface{}) {
	std.Tprintf(Lstack, tag, format, m...)
}
