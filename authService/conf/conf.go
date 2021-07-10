package conf

type baseConf struct {
	ServiceName string
	Port        string
}

var _BaseConf = baseConf{
	ServiceName: "",
	Port:        "",
}

func BaseConf() baseConf {
	return _BaseConf
}
