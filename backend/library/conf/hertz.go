package conf

type Hertz struct {
	Service string `yaml:"service"`
	Address string `yaml:"address"`
	Log     Log    `yaml:"log"`
}
