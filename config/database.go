package config

type Database struct {
	Host     string `required:"true" split_words:"true"`
	Database string `required:"true" split_words:"true"`
	Port     string `required:"true" split_words:"true"`
	Username string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	SslMode  bool   `required:"true" split_words:"true"`
}
