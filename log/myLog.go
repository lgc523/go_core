package main

import (
	"io"
	"log"
	"os"
	"sync/atomic"
)

//log level
const (
	LDEBUG = iota + 1
	LWARN
	LINFO
	LERROR
	LFATAL
)

type myLogger struct {
	level       int32
	w           io.Writer
	debugLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
}

func New(w io.Writer, level int32, flag int) *myLogger {
	if w == nil {
		w = os.Stderr
	}
	if flag <= 0 {
		flag = log.Lmicroseconds
	}
	return &myLogger{
		w:           w,
		level:       level,
		debugLogger: log.New(w, "[DEBUG] ", flag|log.Lmsgprefix),
		warnLogger:  log.New(w, "[WARN] ", flag|log.Lmsgprefix),
		infoLogger:  log.New(w, "[INFO] ", flag|log.Lmsgprefix),
		errLogger:   log.New(w, "[ERROR] ", flag|log.Lmsgprefix),
		fatalLogger: log.New(w, "[FATAL] ", flag|log.Lmsgprefix),
	}
}

func (l *myLogger) SetLevel(level int) {
	if level < LDEBUG || level > LFATAL {
		return
	}
	atomic.StoreInt32(&l.level, int32(level))
}

func (l *myLogger) Debugln(v ...interface{}) {
	if atomic.LoadInt32(&l.level) > LDEBUG {
		return
	}
	l.debugLogger.Println(v...)
}
func (l *myLogger) Debugf(format string, v ...interface{}) {
	if atomic.LoadInt32(&l.level) > LDEBUG {
		return
	}
	l.debugLogger.Printf(format, v...)
}
func (l *myLogger) InfoLn(v ...interface{}) {
	if atomic.LoadInt32(&l.level) > LINFO {
		return
	}
	l.infoLogger.Println(v...)
}
func (l *myLogger) InfoF(format string, v ...interface{}) {
	if atomic.LoadInt32(&l.level) > LINFO {
		return
	}
	l.infoLogger.Printf(format, v...)
}
func main() {
	lg := New(nil, LINFO, 0)
	lg.InfoLn("info level...")
	lg.Debugln("debug level...")
	log.Println("std log...")
}
