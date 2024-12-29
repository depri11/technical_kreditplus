package config

type Config struct {
	DB               DB  `mapstructure:"db"`
	APP              App `mapstructure:"app"`
	CUSTOMER_SERVICE App `mapstructure:"customer_service"`
}

type App struct {
	ServiceName string `mapstructure:"service_name"`
	Port        int    `mapstructure:"port"`
}

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}
