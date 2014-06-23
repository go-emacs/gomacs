package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Emacs string   `json:"emacs"`
	Args  []string `json:"args"`
}

func configFile() string {
	config := "settings.json"
	file := ".gomacs.json"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		if runtime.GOOS == "windows" {
			file = filepath.Join(os.Getenv("APPDATA"), "gomacs", config)
		} else {
			file = filepath.Join(os.Getenv("HOME"), ".config", "gomacs", config)
		}
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		file = filepath.Join(GomacsDir(), config)
	}

	return file
}

func GetConfig() Config {
	file := configFile()
	log.Printf("Load from %s\n", file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %s\n", err)
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatalf("Failed to unmarshal file: %s\n", err)
	}
	return c
}
