package config

type Database struct {
	Host     string `required:"true" split_words:"true"`
	Database string `required:"true" split_words:"true"`
	Port     int    `required:"true" split_words:"true"`
	Username string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	SslMode  string `required:"true" split_words:"true"`
}
