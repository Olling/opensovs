package main

import (
	"os"
	"sync"
	"io/ioutil"
	"encoding/json"
	"github.com/olling/slog"
)

var (
        outputMutex sync.Mutex
)

func ReadJsonFile (path string, output interface{}) (error) {
	slog.PrintDebug("Reading JSON file:", path)
        file,fileerr := os.Open(path)
	if fileerr != nil {
		slog.PrintDebug("Error opening JSON file for reading: ", fileerr)
		return fileerr
	}

        decoder := json.NewDecoder(file)
        decodererr := decoder.Decode(&output)
	if decodererr != nil {
		slog.PrintDebug("Error decoding JSON file", decodererr)
		return decodererr
	}

	slog.PrintDebug("Done reading JSON file:", path, output)
	return nil
}


func WriteJsonFile(s interface{}, path string) (err error){
	slog.PrintDebug("Writing to path: " + path, "content:", s)
        outputMutex.Lock()
        defer outputMutex.Unlock()

        bytes, marshalErr := json.MarshalIndent(s,"","\t")
        if marshalErr != nil {
                slog.PrintError("Could not convert struct to bytes", marshalErr)
                return marshalErr
        }

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		slog.PrintError("Could not write file: " + path + ". Error:", err)
		return err
	}
	return nil
}
