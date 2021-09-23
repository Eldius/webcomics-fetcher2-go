package plugins

import (
	"encoding/json"
	"log"
	"os"
)

/*
GetAbsolutePath returns binary absolute path (to be used from plugins to return self location)
*/
func GetAbsolutePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return ex
}

/*
CreateOutputFile creates the output file and returns a function to clean up tem folder
*/
func CreateOutputFile() (string, func()) {
	tmpDir, err := os.MkdirTemp("", "*")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %s", err.Error())
	}

	o, err := os.CreateTemp(tmpDir, "*")
	if err != nil {
		log.Fatalf("Failed to create temp file at '%s': %s", tmpDir, err.Error())
	}

	return o.Name(), func() {
		os.RemoveAll(tmpDir)
	}
}

/*
ToOutputFile writes object to file in a JSON format
Created to be used from plugins to write its output to be read from main app
*/
func ToOutputFile(output string, obj interface{}) {
	f, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to write output to file (%s): %s", output, err.Error())
	}
	_ = json.NewEncoder(f).Encode(obj)
}
