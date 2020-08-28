package conf

import (
	"encoding/json"
	"base/leaf/log"
	"io/ioutil"
	//"path/filepath"
	//"os"
	// "log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
