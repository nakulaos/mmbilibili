package conf

type Redis struct {
	Addrs          []string `yaml:"addrs"`
	ClientName     string   `yaml:"client_name"`
	DialTimeout    int      `yaml:"dial_timeout"`
	ReadTimeout    int      `yaml:"read_timeout"`
	WriteTimeout   int      `yaml:"write_timeout"`
	MaxActiveCoons int      `yaml:"max_active_coons"`
	MinIdleCoons   int      `yaml:"min_idle_coons"`
}
