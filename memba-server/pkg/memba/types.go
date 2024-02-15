package memba

type Config struct {
	Web      HTTPConfiguration `mapstructure:"http"`
	Database DBConfiguration   `mapstructure:"database"`
}

type DBConfiguration struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"db_name"`
}

type HTTPConfiguration struct {
	URL  string `mapstructure:"url"`
	Port string `mapstructure:"port"`
}
