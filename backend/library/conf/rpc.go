package conf

type RPCConfig struct {
	Name        string `json:"name" yaml:"name"`
	ConnTimeout int    `json:"conn_timeout" yaml:"conn_timeout"` // seconds
	RPCTimeout  int    `json:"rpc_timeout" yaml:"rpc_timeout"`   // seconds
}
