package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Host : Host Object
type Host struct {
	ShortPath string `yaml:"shortpath"`
	FullPath  string `yaml:"fullpath"`
}

// Config : Config File
type Config struct {
	Hosts []Host `yaml:"hosts"`
}

func main() {
	configFileName := "config.yaml"
	checkExistFileConfig(configFileName)

	flag.String("a", "", "Short URL")
	flag.String("u", "", "Full URL")
	flag.Int("p", 0, "Port Number")
	deleteRedirect := flag.String("d", "", "Delete redirection")
	usageInfo := flag.Bool("h", false, "Prints usage information")
	listDirection := flag.Bool("l", false, "List redirections")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [<dir>]\nOptions are:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	//Implement append to the list
	if flag.Arg(0) == "configure" {
		shortLink := flag.Arg(2)
		fullLink := flag.Arg(4)
		if shortLink != "" && fullLink != "" {
			newHost := Host{shortLink, fullLink}
			addHostToConfig(newHost)
			os.Exit(0)
		} else {
			fmt.Println("Oops!\nAdd new redirect follow example below:\nurlshorten configure -a dogs -u www.dogs.com")
		}
	}

	//Implement remove from the list
	if *deleteRedirect != "" {
		removeHostFromConfig(*deleteRedirect)
	}

	//List redirections
	if *listDirection {
		hosts := getHosts()
		fmt.Println("List redirections:")
		for _, host := range hosts {
			fmt.Println(host.ShortPath + " -> " + host.FullPath)
		}
	}

	//Run HTTP server on a given port
	if flag.Arg(0) == "run" {
		port := flag.Arg(2)

		if port != "" {
			err := http.ListenAndServe(":"+port, nil)
			checkError(err)
		}
	}

	// Prints usage info
	if *usageInfo {
		flag.Usage()
		os.Exit(0)
	}
}

func checkExistFileConfig(configFileName string) {
	configPath := getConfigPath()
	_, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		config := Config{}
		createConfigFile(configPath, config)
	}
}

func createConfigFile(path string, config Config) {
	f, createConfigErr := os.Create(path)
	checkError(createConfigErr)

	data, mashlErr := yaml.Marshal(&config)
	checkError(mashlErr)

	_, writeErr := f.Write(data)
	checkError(writeErr)
	f.Sync()

	defer f.Close()
}

func getHosts() []Host {
	data, readFileErr := ioutil.ReadFile(getConfigPath())
	checkError(readFileErr)
	var config Config

	err := yaml.Unmarshal(data, &config)
	checkError(err)
	return config.Hosts
}

func updateFileConfig(hosts []Host) {
	err := os.Remove(getConfigPath())
	checkError(err)
	var config = Config{hosts}
	createConfigFile(getConfigPath(), config)
}

func addHostToConfig(host Host) {
	hosts := getHosts()
	hostIndex := getHostIndex(host, hosts)

	// Replace host if exist
	if hostIndex > -1 {
		hosts[hostIndex] = host
	} else {
		hosts = append(hosts, host)
	}
	updateFileConfig(hosts)
}

func removeHostFromConfig(shortPath string) {
	hosts := getHosts()
	tempHost := Host{shortPath, ""}
	hostIndex := getHostIndex(tempHost, hosts)
	hosts = append(hosts[:hostIndex], hosts[hostIndex+1:]...)
	updateFileConfig(hosts)
}

func getHostIndex(host Host, hosts []Host) int {
	hostIndex := -1
	for index, item := range hosts {
		if item.ShortPath == host.ShortPath {
			hostIndex = index
		}
	}
	return hostIndex
}

func getConfigPath() string {
	currentDir, err := os.Getwd()
	checkError(err)

	return path.Join(currentDir, "config.yaml")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
