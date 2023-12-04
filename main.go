package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"os"
	"strings"
)

func main() {
	format := flag.String("pattern", "json", "Pattern for filename to retrieve")
	inputPath := flag.String("dir", ".", "The input directory")
	flag.Parse()

	dir, err := os.ReadDir(*inputPath)
	if err != nil {
		slog.Error("Error while reading directory", "error", err.Error())
		panic(err)
	}
	for _, file := range dir {
		if file.IsDir() || !strings.HasSuffix(file.Name(), *format) {
			slog.Error("File is not valid for formatting", "filename", file.Name())
			continue
		}
		input, err := os.ReadFile(*inputPath + "/" + file.Name())
		if err != nil {
			slog.Error("Error while reading the file", "error", err.Error())
			continue
		}
		var data interface{}
		err = json.Unmarshal(input, &data)
		if err != nil {
			slog.Error("Error during json decoding: %v\n", "error", err.Error())
			continue
		}
		result, err := json.MarshalIndent(data, "", "   ")
		if err != nil {
			slog.Error("Error during formating", "error", err.Error())
			continue
		}
		filename := *inputPath + "/" + file.Name()
		err = os.WriteFile(filename, result, 0444)
		if err != nil {
			slog.Error("Error during write", "error", err.Error())
			continue
		}
	}
}
