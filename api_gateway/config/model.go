package config

type Config struct {
	APP                 App `mapstructure:"app"`
	CUSTOMER_SERVICE    App `mapstructure:"customer_service"`
	TRANSACTION_SERVICE App `mapstructure:"transaction_service"`
}

type App struct {
	ServiceName string `mapstructure:"service_name"`
	Port        int    `mapstructure:"port"`
}
