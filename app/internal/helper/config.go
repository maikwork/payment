package helper

import (
	"log"
	"os"

	"github.com/maikwork/restPayments/internal/model"
	"sigs.k8s.io/yaml"
)

func ReadConfig(path string) (*model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Print("Can't read the file")
		return nil, err
	}

	cnf := &model.Config{}

	err = yaml.Unmarshal(data, cnf)
	if err != nil {
		log.Print("Can't unmarshal")
		return nil, err
	}

	return cnf, nil
}
