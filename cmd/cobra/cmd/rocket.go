package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type rocket struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Mission  string `json:"mission"`
	Fuel     int    `json:"fuel"`
	Maxspeed int    `json:"maxspeed"`
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return usr.HomeDir + "/.rocketctl"
}

func createRocket(r rocket) {
	homeDir := getHomeDir()
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		err = os.Mkdir(homeDir, 0700)
		if err != nil {
			panic("unable to create .rocketctl dir")
		}
	}

	file, _ := os.Create(homeDir + "/" + r.Name)

	defer file.Close()

	b, _ := json.Marshal(r)
	_, err := file.WriteString(string(b))
	if err != nil {
		panic("unable to write rocket to file")
	}
	fmt.Printf("rocket %s has been created\n", r.Name)
}

func getRocketByName(name string) rocket {
	data, _ := ioutil.ReadFile(getHomeDir() + "/" + name)
	fmt.Printf("Data: %s\n", data)
	rocket := rocket{}
	err := json.Unmarshal([]byte(data), &rocket)
	if err != nil {
		panic("unable to read rocket data")
	}
	return rocket
}
