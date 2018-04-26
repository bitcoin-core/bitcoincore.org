// This is a golang script, needed for generating the RPC bitcoin documentation
//
// What is necessary to run this:
// (1) install golang
// (2) install bitcoin core, set it up to use regtest
// (3) run bitcoind
// (4) change BITCOIN_PATH to actual bitcoin-cli directory path, change VERSION
// (5) run this script with `go run generate.go` while being in contrib/doc-gen
// (6) add the generated files to git
package main

import (
	"fmt"
	"text/template"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const BITCOIN_PATH = "/home/g/dev/bitcoind/bitcoin-0.16.0/bin/bitcoin-cli"
const VERSION = "0.16.0"

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

func main() {
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
		dirname := fmt.Sprintf("../../_doc/en/%s/rpc/%s/", VERSION, groupname)
		err := os.MkdirAll(dirname, 0777)
		if err != nil {
			log.Fatalf("Cannot make directory %s: %s", dirname, err.Error())
		}
		for _, command := range group.Commands {
			name := command.Name
			address := fmt.Sprintf("%s%s.html", dirname, name)
			permalink := fmt.Sprintf("en/doc/%s/rpc/%s/%s/", VERSION, groupname, name)
			err = tmpl.Execute(open(address), CommandData{
				Version:     VERSION,
				Name:        name,
				Description: command.Description,
				Group:       groupname,
				Permalink:   permalink,
			})
			if err != nil {
				log.Fatalf("Cannot make command file %s: %s", name, err.Error())
			}
		}
		address := fmt.Sprintf("../../_doc/en/%s/rpc/index.html", VERSION)
		permalink := fmt.Sprintf("en/doc/%s/rpc/", VERSION)
		err = tmpl.Execute(open(address), CommandData{
			Version:     VERSION,
			Name:        "rpcindex",
			Description: "",
			Group:       "index",
			Permalink:   permalink,
		})
		if err != nil {
			log.Fatalf("Cannot make index file: %s", err.Error())
		}

		address = fmt.Sprintf("../../_doc/en/%s/index.html", VERSION)
		permalink = fmt.Sprintf("en/doc/%s/", VERSION)
		err = tmpl.Execute(open(address), CommandData{
			Version:     VERSION,
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
	out, err := exec.Command(BITCOIN_PATH, args...).CombinedOutput()
	if err != nil {
		log.Fatalf("Cannot run bitcoin-cli: %s, is bitcoind running?", err.Error())
	}

	return string(out)
}
