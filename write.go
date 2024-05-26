package gofile

import (
	"encoding/json"
	"os"
)

func WriteAsJson(path string, d any) (err error) {
	dBytes, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return os.WriteFile(path, dBytes, 0644)
}

func WriteAsIndentedJson(path string, d any) (err error) {
	dBytes, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, dBytes, 0644)
}

func Write(path string, b []byte) (err error) {
	return os.WriteFile(path, b, 0644)
}

func WriteString(path string, s string) (err error) {
	return os.WriteFile(path, []byte(s), 0644)
}

func Append(path string, b []byte) (err error) {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

func AppendString(path string, s string) (err error) {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(s); err != nil {
		return err
	}
	return nil
}
