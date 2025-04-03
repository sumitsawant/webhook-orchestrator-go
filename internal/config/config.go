package config

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port         int `mapstructure:"port"`
        ReadTimeout  int `mapstructure:"read_timeout"`
        WriteTimeout int `mapstructure:"write_timeout"`
    }
    Callback struct {
        MaxRetries    int `mapstructure:"max_retries"`
        BaseDelaySecs int `mapstructure:"base_delay_secs"`
    }
    ThirdParty struct {
        TimeoutSecs int `mapstructure:"timeout_secs"`
    }
    TLS struct {
        Enabled  bool   `mapstructure:"enabled"`
        CertFile string `mapstructure:"cert_file"`
        KeyFile  string `mapstructure:"key_file"`
    }
}

var AppConfig Config

func Load() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Config read error: %v", err)
    }
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Config parse error: %v", err)
    }
}
