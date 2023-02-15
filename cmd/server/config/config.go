package config

type Config struct {
	HatechServer HatechServer `mapstructure:"hatech_server"`
}
type HatechServer struct {
	Name   string `mapstructure:"name"`
	Listen string `mapstructure:"listen"`
}
