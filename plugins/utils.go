package plugins

import "os"

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
