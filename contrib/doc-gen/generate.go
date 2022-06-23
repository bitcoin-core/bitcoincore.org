// This is a golang script, needed for generating the RPC bitcoin documentation
//
// What is necessary to run this:
// (1) install golang
// (2) install bitcoin core, set it up to use regtest
// (3) run bitcoind
// (4) run this script with `go run generate.go` while being in contrib/doc-gen, and with bitcoin-cli in PATH
// (5) add the generated files to git
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const BITCOIN_COMMAND = "bitcoin-cli"
const BITCOIN_CHAINOPTION = "-regtest"

type Command struct {
	Name        string
	Description string
}

type Group struct {
	Index    int
	Name     string
	Commands []Command
}

type CommandData struct {
	Version     string
	Name        string
	Description string
	Group       string
	Permalink   string
}

func getVersion() string {
	allInfo := run("getnetworkinfo")
	var f interface{}
	err := json.Unmarshal([]byte(allInfo), &f)
	if err != nil {
		panic("Cannot read network info as JSON")
	}
	m := f.(map[string]interface{})

	numv := int(m["version"].(float64))
	v := fmt.Sprintf("%d.%d.%d", (numv/10000)%100, (numv/100)%100, numv%100)
	return v
}

func main() {
	version := getVersion()

	first := run("help")
	split := strings.Split(first, "\n")

	groups := make([]Group, 0)
	commands := make([]Command, 0)
	lastGroupName := ""

	for _, line := range split {
		if len(line) > 0 {
			if strings.HasPrefix(line, "== ") {
				if len(commands) != 0 {
					g := Group{
						Name:     lastGroupName,
						Commands: commands,
						Index:    len(groups),
					}
					groups = append(groups, g)
					commands = make([]Command, 0)
				}
				lastGroupName = strings.ToLower(line[3 : len(line)-3])
			} else {
				name := strings.Split(line, " ")[0]
				desc := run("help", name)
				comm := Command{
					Name:        name,
					Description: desc,
				}
				commands = append(commands, comm)
			}
		}
	}

	g := Group{
		Name:     lastGroupName,
		Commands: commands,
		Index:    len(groups),
	}
	groups = append(groups, g)

	tmpl := template.Must(template.ParseFiles("command-template.html"))

	for _, group := range groups {
		groupname := group.Name
		dirname := fmt.Sprintf("../../_doc/en/%s/rpc/%s/", version, groupname)
		err := os.MkdirAll(dirname, 0777)
		if err != nil {
			log.Fatalf("Cannot make directory %s: %s", dirname, err.Error())
		}
		for _, command := range group.Commands {
			name := command.Name
			address := fmt.Sprintf("%s%s.html", dirname, name)
			permalink := fmt.Sprintf("en/doc/%s/rpc/%s/%s/", version, groupname, name)
			err = tmpl.Execute(open(address), CommandData{
				Version:     version,
				Name:        name,
				Description: command.Description,
				Group:       groupname,
				Permalink:   permalink,
			})
			if err != nil {
				log.Fatalf("Cannot make command file %s: %s", name, err.Error())
			}
		}
		address := fmt.Sprintf("../../_doc/en/%s/rpc/index.html", version)
		permalink := fmt.Sprintf("en/doc/%s/rpc/", version)
		err = tmpl.Execute(open(address), CommandData{
			Version:     version,
			Name:        "rpcindex",
			Description: "",
			Group:       "index",
			Permalink:   permalink,
		})
		if err != nil {
			log.Fatalf("Cannot make index file: %s", err.Error())
		}

		address = fmt.Sprintf("../../_doc/en/%s/index.html", version)
		permalink = fmt.Sprintf("en/doc/%s/", version)
		err = tmpl.Execute(open(address), CommandData{
			Version:     version,
			Name:        "index",
			Description: "",
			Group:       "index",
			Permalink:   permalink,
		})
		if err != nil {
			log.Fatalf("Cannot make index file: %s", err.Error())
		}
	}
}

func open(path string) io.Writer {
	f, err := os.Create(path)
	// not closing, program will close sooner
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", path, err.Error())
	}
	return f
}

func run(args ...string) string {
	args = append([]string{BITCOIN_CHAINOPTION}, args...)
	out, err := exec.Command(BITCOIN_COMMAND, args...).CombinedOutput()
	if err != nil {
		log.Fatalf("Cannot run bitcoin-cli: %s, is bitcoind (regtest) running?", err.Error())
	}

	return string(out)
}
