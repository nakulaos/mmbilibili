package conf

type Registry struct {
	RegistryAddress string `yaml:"registry_address"`
	Weight          int    `yaml:"weight"`
}

type Resolve struct {
	ResolveAddress string `yaml:"resolve_address"`
}
