package config

type Config struct {
	Database Database `required:"true" split_words:"true"`
}
