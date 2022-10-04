package config

type Database struct {
	Host       string
	Port       string
	UserName   string
	Password   string
	SchemaName string
}

type DBConfig struct {
	Db Database `mapstructure:"db"`
}

type MainConfig struct {
	Main Main `mapstructure:"main"`
}

type Main struct {
	HostGrpc string
	PortGrpc string
	Network  string
	PortHttp string
}
