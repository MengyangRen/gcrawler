package zlog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"runtime"
	"time"
)

var _ Logger = &jsonLog{}

type jsonLog struct {
	prefix string
	level  Lvl
	output io.Writer
	levels []string
	dw     *elogWriter
}

func NewJsonLog(w io.WriteCloser, prefix string) *jsonLog {
	l := &jsonLog{
		level:  INFO,
		prefix: prefix,
	}
	l.initLevels()
	l.dw = NewElogWriter(w)
	l.SetOutput(l.dw)
	l.SetLevel(INFO)
	return l
}
func (p *jsonLog) With(ctx context.Context, kv ...interface{}) context.Context {
	om := FromContext(ctx)
	if om == nil {
		om = NewOrderMap()
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	for i := 0; i < len(kv); i += 2 {
		om.Set(fmt.Sprintf("%v", kv[i]), kv[i+1])
	}
	return setContext(ctx, om)
}
func (p *jsonLog) Debug(ctx context.Context, msg interface{}, kv ...interface{}) {
	p.logJson(DEBUG, ctx, msg, kv...)
}
func (p *jsonLog) Info(ctx context.Context, msg interface{}, kv ...interface{}) {
	p.logJson(INFO, ctx, msg, kv...)
}
func (p *jsonLog) Warn(ctx context.Context, msg interface{}, kv ...interface{}) {
	p.logJson(WARN, ctx, msg, kv...)
}
func (p *jsonLog) Error(ctx context.Context, msg interface{}, kv ...interface{}) {
	p.logJson(ERROR, ctx, msg, kv...)
}
func (p *jsonLog) Fatal(ctx context.Context, msg interface{}, kv ...interface{}) {
	p.logJson(FATAL, ctx, msg, kv...)
	panic(msg)
}

//kv 应该是成对的 数据, 类似: name,张三,age,10,...
func (p *jsonLog) logJson(v Lvl, ctx context.Context, msg interface{}, kv ...interface{}) (err error) {
	if v < p.level {
		return nil
	}
	om := NewOrderMap()
	_, file, line, _ := runtime.Caller(3)
	file = p.getFilePath(file)
	om.Set("prefix", p.Prefix())
	om.Set("level", p.levels[v])
	om.Set("cur_time", time.Now().Format(time.RFC3339Nano))
	om.Set("file", file)
	om.Set("line", line)
	//traceid,pspanid,spanid 应该提前写到context中区
	om.AddVals(FromContext(ctx))
	om.Set("std_msg", msg)
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	for i := 0; i < len(kv); i += 2 {
		om.Set(fmt.Sprintf("%v", kv[i]), kv[i+1])
	}
	str, _ := json.Marshal(om)
	str = append(str, []byte("\n")...)
	_, err = p.Output().Write(str)
	return
}
func (p *jsonLog) getFilePath(file string) string {
	dir, base := path.Dir(file), path.Base(file)
	return path.Join(path.Base(dir), base)
}
func (p *jsonLog) Close() error {
	if p.dw != nil {
		p.dw.Close()
		p.dw = nil
	}
	return nil
}
func (p *jsonLog) EnableDebug(b bool) {
	if b && Config().Mode != ModeOnLine && Config().Mode != ModePre {
		p.SetLevel(DEBUG)
	} else {
		p.SetLevel(INFO)
	}
}

type Lvl uint8

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

func (l *jsonLog) initLevels() {
	l.levels = []string{
		"-",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
	}
}

func (l *jsonLog) Prefix() string {
	return l.prefix
}

func (l *jsonLog) SetPrefix(p string) {
	l.prefix = p
}

func (l *jsonLog) Level() Lvl {
	return l.level
}

func (l *jsonLog) SetLevel(v Lvl) {
	l.level = v
}

func (l *jsonLog) Output() io.Writer {
	return l.output
}

func (l *jsonLog) SetOutput(w io.Writer) {
	l.output = w
}
