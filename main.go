package main

import (
	"encoding/json"
	"flag"
	"github.com/NubeDev/go-flow/server"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port = flag.String("port", "7071", "stream tools port")
)

func main() {

	flag.Parse()

	// Unpack settings file, or create a new one if necessary
	var settings server.Settings

	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)

	}
	fName := dir + "/stt.json"
	d, err := ioutil.ReadFile(fName)
	if err != nil {
		if os.IsNotExist(err) {
			// make a fresh settings file
			log.Println("creating new settings file at", fName)
			newSettings := server.NewSettings()
			d, err = json.Marshal(newSettings)
			err = ioutil.WriteFile(fName, d, 0644)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	err = json.Unmarshal(d, &settings)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(settings)
	r := s.NewRouter()

	http.Handle("/", r)

	log.Println("serving on", *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Panicf(err.Error())
	}
}
