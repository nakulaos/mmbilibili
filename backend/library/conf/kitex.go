package conf

type Kitex struct {
	Service string `yaml:"service"`
	Address string `yaml:"address"`
	Log     Log    `yaml:"log"`
}
