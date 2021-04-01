package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	AppConf   *AppConf   `mapstructure:"app"`
	DbConf    *DbConf    `mapstructure:"db"`
	LogConf   *LogConf   `mapstructure:"log"`
	RedisConf *RedisConf `mapstructure:"redis"`
}

type AppConf struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"startTime"`
	MachineId int64  `mapstructure:"machineId"`
}

type DbConf struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"maxOpenConn"`
	MaxIdleConn int    `mapstructure:"maxIdleConn"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	//viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config update...")
		err = viper.Unmarshal(Conf)
		if err != nil {
			panic(err)
		}
	})
}
