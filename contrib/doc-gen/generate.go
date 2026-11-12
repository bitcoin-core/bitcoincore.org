// Generates Bitcoin Core's RPC and IPC documentation.
//
// What is necessary to run this:
//   (1) install golang
//   (2) install Cap'n Proto (`capnp` in PATH)
//   (3) install bitcoin core, set it up to use regtest
//   (4) run bitcoind
//   (5) checkout Bitcoin Core source matching the running node
//   (6) from contrib/doc-gen, with bitcoin-cli in PATH:
//         go run . -bitcoin /path/to/bitcoin
//   (7) add the generated files to git
//
// IPC docs are only generated for Bitcoin Core v31.0 and newer (the
// libmultiprocess capnp schemas were not previously stable).
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CommandData struct {
	Version     string
	Name        string
	Description string
	Group       string
	DocType     string
	HTML        bool
	IfaceIndex  bool     // per-IPC-interface landing page (sidebar hides it)
	Structs     []string // populated for the IPC structs index page
	Permalink   string
}

func main() {
	bitcoin := flag.String("bitcoin", "", "Path to a Bitcoin Core source checkout (required for IPC docs)")
	flag.Parse()

	version := generateRPC()

	if _, err := exec.LookPath("capnp"); err != nil {
		log.Fatalf("`capnp` not found in PATH: install Cap'n Proto to generate IPC docs")
	}
	if *bitcoin == "" {
		log.Fatalf("-bitcoin <path> is required (points at a Bitcoin Core source checkout matching the running node)")
	}
	if !ipcSupported(version) {
		log.Printf("Skipping IPC docs: not supported before v31.0 (running %s)", version)
		return
	}
	generateIPC(*bitcoin, version)
}

// ipcSupported reports whether the given X.Y.Z version string is >= 31.0.
func ipcSupported(version string) bool {
	parts := strings.SplitN(version, ".", 3)
	if len(parts) < 1 {
		return false
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}
	return major >= 31
}

func open(path string) io.Writer {
	f, err := os.Create(path)
	// not closing, program will close sooner
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", path, err.Error())
	}
	return f
}
