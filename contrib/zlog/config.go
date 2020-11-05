package zlog

type modeLog string

const (
	ModeDev    modeLog = "dev"
	ModeQa     modeLog = "qa"
	ModePre    modeLog = "pre"
	ModeOnLine modeLog = "online"
)

type Conf struct {
	Dir    string
	Prefix string
	Mode   modeLog
}

var conf = &Conf{
	Dir:    "/tmp/log",
	Prefix: "default",
	Mode:   ModeDev,
}

func ChangeConfig(c *Conf) {
	if _dLogger != nil {
		_dLogger.Close()
	}
	if c != nil {
		conf = c
	}
	file, err := NewFileBackend(conf.Dir, conf.Prefix+".log_json")
	if err != nil {
		panic(err)
	}
	_jsonLog := NewJsonLog(file, conf.Prefix)
	SetLogger(_jsonLog)
}

func Config() *Conf {
	return conf
}
