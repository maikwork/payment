package model

type DataBase struct {
	DSN string `yaml:"dsn"`
}

type Config struct {
	DB DataBase `yaml:"db"`
}
