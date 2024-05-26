package gofile

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func ReadToBytes(path string) ([]byte, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return fileData, nil
}

func ReadToString(path string) (string, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}

func ReadCsv(path string) ([][]string, error) {
	var res [][]string
	f, err := os.Open(path)
	if err != nil {
		return res, fmt.Errorf("error opening the file: %s", err.Error())
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	res, err = csvReader.ReadAll()
	if err != nil {
		return res, fmt.Errorf("error parsing the csv: %s", err.Error())
	}

	return res, nil
}

func ReadCsvSimple(path string) ([]string, error) {
	var res []string
	f, err := os.Open(path)
	if err != nil {
		return res, fmt.Errorf("error opening the file: %s", err.Error())
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	unparsedRes, err := csvReader.ReadAll()
	if err != nil {
		return res, fmt.Errorf("error parsing the csv: %s", err.Error())
	}
	for _, line := range unparsedRes {
		res = append(res, line[0])
	}

	return res, nil
}

func ReadToml(path string, d any) error {
	tomlData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error laoding the toml file: %s", err.Error())
	}
	_, err = toml.Decode(string(tomlData), d)
	if err != nil {
		return fmt.Errorf("error decoding the toml file: %s", err.Error())
	}
	return nil
}

func ReadJson(path string, d any) error {
	jsonData, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("error loading the json file: %s", err.Error())
	}

	err = json.Unmarshal(jsonData, d)
	if err != nil {
		return fmt.Errorf("error unmarshalling the json data: %s", err.Error())
	}

	return nil
}
