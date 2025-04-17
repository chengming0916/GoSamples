package types

type ServiceConfig struct {
	TcpCfg TcpConfig `json:"tcp,omitempty" yaml:"tcp,omitempty"`
}

type TcpConfig struct {
	Host      string `json:"host,omitempty" yaml:"host,omitempty"`
	Port      int    `default:"8000" json:"port" yaml:"port"`
	KeepAlive int    `json:"keepalive" yaml:"keepalive" default:"5"`
}

func (t TcpConfig) Complete() {

}

func (c ServiceConfig) Complete() {
	//TODO: complte config
	c.TcpCfg.Complete()
}
