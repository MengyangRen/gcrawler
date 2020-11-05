# elog
一个日志库,带有trace，输出有序json格式等

特性：
 - 带有trace信息,通过context实现，可以关联所有日志
 - 使用的时候支持key-val形式，方便
 - 输出的是json格式，并且能保证有序



使用样例

```
zlog.ChangeConfig(&zlog.Conf{Prefix: "lzy", Dir: "/tmp/log", Mode: zlog.ModeDev})
zlog.EnableDebug(true)
ctx := zlog.SetTraceInfo(context.Background(), "1", "2", "3")
zlog.Info(ctx, "adsfjasdfja;d", "aa", 1)
ctx = zlog.With(ctx, "uid", 123, "order_no", "1231231231231231456788")
zlog.Debug(ctx, "debug msg")
zlog.Info(ctx, "info msg")
zlog.Warn(ctx, "warn msg")
zlog.Error(ctx, "error msg")
zlog.Close()
```
输出日志格式

```
{"prefix":"lzy","level":"INFO","cur_time":"2017-11-03T18:05:27.463962711+08:00","file":"app/main.go","line":87,"trace_id":"1","pspan_id":"2","span_id":"3","std_msg":"adsfjasdfja;d","aa":1}
{"prefix":"lzy","level":"DEBUG","cur_time":"2017-11-03T18:05:27.464048262+08:00","file":"app/main.go","line":89,"trace_id":"1","pspan_id":"2","span_id":"3","uid":123,"order_no":"1231231231231231456788","std_msg":"debug msg"}
{"prefix":"lzy","level":"INFO","cur_time":"2017-11-03T18:05:27.464097096+08:00","file":"app/main.go","line":90,"trace_id":"1","pspan_id":"2","span_id":"3","uid":123,"order_no":"1231231231231231456788","std_msg":"info msg"}
{"prefix":"lzy","level":"WARN","cur_time":"2017-11-03T18:05:27.464131054+08:00","file":"app/main.go","line":91,"trace_id":"1","pspan_id":"2","span_id":"3","uid":123,"order_no":"1231231231231231456788","std_msg":"warn msg"}
2017-11-03 18:05:27.464191339 +0800 CST,elogWriter_close(w.closeStartCh)
{"prefix":"lzy","level":"ERROR","cur_time":"2017-11-03T18:05:27.46416251+08:00","file":"app/main.go","line":92,"trace_id":"1","pspan_id":"2","span_id":"3","uid":123,"order_no":"1231231231231231456788","std_msg":"error msg"}

```