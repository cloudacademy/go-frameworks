package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type rocket struct {
	Name     string
	Type     string
	Mission  string
	Fuel     int
	Maxspeed int
	Launched bool
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

	_, err := file.WriteString(fmt.Sprintf("%#v", r))
	if err != nil {
		panic("unable to write file")
	}
	fmt.Printf("rocket %s has been created\n", r.Name)
}

func getRocketByName(name string) {
	data, _ := ioutil.ReadFile(getHomeDir() + "/" + name)
	fmt.Printf("Data: %s\n", data)
}
