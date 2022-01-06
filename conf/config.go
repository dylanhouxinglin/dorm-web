package conf

import "github.com/spf13/viper"

type ServerConf struct {
	HttpAddress  string   `yaml:"httpaddress"`
	EtcdAddress	 string	  `yaml:"etcdaddress"`
}

type Config struct {
	ConfPath 	string
	ConfName 	string
}

func ParseYamlConfig(conf *Config) (*ServerConf, error) {
	v := viper.New()
	v.AddConfigPath(conf.ConfPath)
	v.SetConfigName(conf.ConfName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var sv ServerConf
	if err := v.Unmarshal(&sv); err != nil {
		return nil, err
	}
	return &sv, nil
}
