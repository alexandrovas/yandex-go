package main

import (
	"encoding/json"
	"os"
	"testing"
)

// Settings содержит настройки приложения.
type Settings struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

// Save сохраняет настройки в файле fname.
func (settings Settings) Save(fname string) error {
	// сериализуем структуру в JSON формат
	data, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		return err
	}
	// сохраняем данные в файл
	return os.WriteFile(fname, data, 0666)
}

// Load читает настройки из файла fname.
func (settings *Settings) Load(fname string) error {
	// прочитайте файл с помощью os.ReadFile
	// десериализуйте данные, используя json.Unmarshal
	data, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, settings); err != nil {
		return err
	}
	return nil
}

func TestSettings(t *testing.T) {
	fname := `settings.json`
	settings := Settings{
		Port: 3000,
		Host: `localhost`,
	}
	if err := settings.Save(fname); err != nil {
		t.Error(err)
	}
	var result Settings
	if err := (&result).Load(fname); err != nil {
		t.Error(err)
	}
	if settings != result {
		t.Errorf(`%+v не равно %+v`, settings, result)
	}
	// удалим файл settings.json
	if err := os.Remove(fname); err != nil {
		t.Error(err)
	}
}
