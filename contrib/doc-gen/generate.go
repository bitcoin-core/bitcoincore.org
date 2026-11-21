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
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
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
	Structured  bool
	Raw         bool
	Openrpc     bool
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

	methods, rawMethods := loadSpec(version)

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
			m, ok := methods[name]
			if !ok {
				// classic page because there is no OpenRPC for this method.
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
				continue
			}
			err = tmpl.Execute(open(address), CommandData{
				Version:     version,
				Name:        name,
				Description: renderBody(m, command.Description),
				Group:       groupname,
				Permalink:   permalink,
				Structured:  true,
			})
			if err != nil {
				log.Fatalf("Cannot make command file %s: %s", name, err.Error())
			}
			err = tmpl.Execute(open(fmt.Sprintf("%s%s-raw.html", dirname, name)), CommandData{
				Version:     version,
				Name:        name,
				Description: command.Description,
				Group:       groupname,
				Permalink:   permalink + "raw/",
				Raw:         true,
			})
			if err != nil {
				log.Fatalf("Cannot make command file %s: %s", name, err.Error())
			}
			err = tmpl.Execute(open(fmt.Sprintf("%s%s-openrpc.html", dirname, name)), CommandData{
				Version:     version,
				Name:        name,
				Description: renderOpenrpcBody(rawMethods[name]),
				Group:       groupname,
				Permalink:   permalink + "openrpc/",
				Openrpc:     true,
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

// runErr is run for RPCs that may not exist on this node.
func runErr(args ...string) (string, error) {
	args = append([]string{BITCOIN_CHAINOPTION}, args...)
	out, err := exec.Command(BITCOIN_COMMAND, args...).CombinedOutput()
	return string(out), err
}

// OpenRPC document parsing
type Spec struct {
	OpenRPC string            `json:"openrpc"`
	Info    Info              `json:"info"`
	Methods []json.RawMessage `json:"methods"`
}

type Info struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type Method struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Params      []Param `json:"params"`
	Result      *Result `json:"result"`
	Category    string  `json:"x-bitcoin-category"`
	Deprecated  bool    `json:"deprecated"`
}

type Param struct {
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Description string   `json:"description"`
	Schema      Schema   `json:"schema"`
	Aliases     []string `json:"x-bitcoin-aliases"`
	Placeholder bool     `json:"x-bitcoin-placeholder"`
}

type Result struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type                 json.RawMessage `json:"type"`
	Description          string          `json:"description"`
	Pattern              string          `json:"pattern"`
	Properties           json.RawMessage `json:"properties"`
	Required             []string        `json:"required"`
	Items                *Schema         `json:"items"`
	PrefixItems          []Schema        `json:"prefixItems"`
	OneOf                []Schema        `json:"oneOf"`
	AnyOf                []Schema        `json:"anyOf"`
	AllOf                []Schema        `json:"allOf"`
	AdditionalProperties json.RawMessage `json:"additionalProperties"`
	MinItems             *int            `json:"minItems"`
	MaxItems             *int            `json:"maxItems"`
	Default              json.RawMessage `json:"default"`
	Unit                 string          `json:"x-bitcoin-unit"`
	DefaultHint          string          `json:"x-bitcoin-default-hint"`
	TypeOverride         string          `json:"x-bitcoin-type-override"`
}

// fetches the OpenRPC document from the node (getopenrpcinfo) or spec file.
// keeps the classic pages if none available.
// specs located at openrpc-specs/openrpc-v<version>.json
func loadSpec(version string) (map[string]Method, map[string]json.RawMessage) {
	data, err := runErr("getopenrpcinfo")
	if err != nil {
		raw, ferr := os.ReadFile("openrpc-specs/openrpc-v" + version + ".json")
		if ferr != nil {
			log.Printf("node has no getopenrpcinfo and no committed spec for %s; generating plain pages only", version)
			return nil, nil
		}
		data = string(raw)
	}
	var spec Spec
	if err := json.Unmarshal([]byte(data), &spec); err != nil {
		log.Fatalf("Cannot parse OpenRPC document: %s", err.Error())
	}
	methods := map[string]Method{}
	rawMethods := map[string]json.RawMessage{}
	for _, raw := range spec.Methods {
		var m Method
		if err := json.Unmarshal(raw, &m); err != nil {
			log.Fatalf("Cannot parse OpenRPC method: %s", err.Error())
		}
		methods[m.Name] = m
		rawMethods[m.Name] = raw
	}
	return methods, rawMethods
}

func (s Schema) typeNames() []string {
	if len(s.Type) == 0 {
		return nil
	}
	var single string
	if err := json.Unmarshal(s.Type, &single); err == nil {
		return []string{single}
	}
	var multi []string
	if err := json.Unmarshal(s.Type, &multi); err == nil {
		return multi
	}
	return nil
}

func (s Schema) hasType(t string) bool {
	for _, n := range s.typeNames() {
		if n == t {
			return true
		}
	}
	return false
}

func (s Schema) additionalSchema() *Schema {
	if len(s.AdditionalProperties) == 0 {
		return nil
	}
	var b bool
	if err := json.Unmarshal(s.AdditionalProperties, &b); err == nil {
		return nil
	}
	var sch Schema
	if err := json.Unmarshal(s.AdditionalProperties, &sch); err != nil {
		return nil
	}
	return &sch
}

func (s Schema) variants() []Schema {
	switch {
	case len(s.OneOf) > 0:
		return s.OneOf
	case len(s.AnyOf) > 0:
		return s.AnyOf
	case len(s.AllOf) > 0:
		return s.AllOf
	}
	return nil
}

type propEntry struct {
	Name   string
	Schema Schema
}

func orderedProps(raw json.RawMessage) []propEntry {
	if len(raw) == 0 {
		return nil
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	if _, err := dec.Token(); err != nil { // opening '{'
		return nil
	}
	var out []propEntry
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			break
		}
		key, _ := keyTok.(string)
		var sch Schema
		if err := dec.Decode(&sch); err != nil {
			break
		}
		out = append(out, propEntry{Name: key, Schema: sch})
	}
	return out
}

func toSet(xs []string) map[string]bool {
	m := make(map[string]bool, len(xs))
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func esc(s string) string { return html.EscapeString(s) }

var urlRe = regexp.MustCompile(`https?://[^\s<>"'` + "`" + `)]+`)

func linkify(s string) string {
	var b strings.Builder
	last := 0
	for _, loc := range urlRe.FindAllStringIndex(s, -1) {
		start, end := loc[0], loc[1]
		url := strings.TrimRight(s[start:end], ".,:;")
		end = start + len(url)
		b.WriteString(esc(s[last:start]))
		b.WriteString(`<a href="` + esc(url) + `">` + esc(url) + `</a>`)
		last = end
	}
	b.WriteString(esc(s[last:]))
	return b.String()
}

func descInline(s string) string {
	return strings.ReplaceAll(linkify(strings.TrimRight(s, "\n")), "\n", "<br>")
}

var blankLineRe = regexp.MustCompile(`\n[ \t]*\n[ \t\n]*`)

// renders a method description like Core's
// wrapped lines.
// keeps blank lines and indented blocks.
func descBlock(s string) string {
	s = strings.TrimRight(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
	if s == "" {
		return ""
	}
	var b strings.Builder
	for _, para := range blankLineRe.Split(s, -1) {
		if strings.TrimSpace(para) == "" {
			continue
		}
		b.WriteString("<p>")
		prevIndent := 0
		for i, ln := range strings.Split(para, "\n") {
			indent := len(ln) - len(strings.TrimLeft(ln, " \t"))
			trimmed := strings.TrimSpace(ln)
			switch {
			case i == 0:
				b.WriteString(linkify(trimmed))
			case indent > 0 || prevIndent > 0:
				b.WriteString("<br>" + strings.Repeat("&nbsp;", indent) + linkify(trimmed))
			default:
				b.WriteString(" " + linkify(trimmed))
			}
			if trimmed == "" {
				indent = 0
			}
			prevIndent = indent
		}
		b.WriteString("</p>")
	}
	return b.String()
}

func typeLabel(s Schema) string {
	if s.TypeOverride != "" {
		return s.TypeOverride
	}
	if len(s.variants()) > 0 {
		switch {
		case len(s.OneOf) > 0:
			return "one of"
		case len(s.AnyOf) > 0:
			return "any of"
		default:
			return "all of"
		}
	}
	names := s.typeNames()
	if len(names) == 0 {
		switch {
		case len(s.Properties) > 0 || s.additionalSchema() != nil:
			names = []string{"object"}
		case s.Items != nil || len(s.PrefixItems) > 0:
			names = []string{"array"}
		default:
			return "any"
		}
	}
	label := strings.Join(names, " or ")
	if s.Unit != "" {
		label += ", " + s.Unit
	}
	return label
}

func annotation(s Schema, optional bool) string {
	a := typeLabel(s)
	if optional {
		a += ", optional"
	}
	if len(s.Default) > 0 {
		a += ", default=" + compactJSON(s.Default)
	} else if s.DefaultHint != "" {
		a += ", default: " + s.DefaultHint
	}
	return "(" + a + ")"
}

func compactJSON(raw json.RawMessage) string {
	var buf bytes.Buffer
	if err := json.Compact(&buf, raw); err != nil {
		return string(raw)
	}
	return buf.String()
}

func synopsis(m Method) string {
	parts := []string{m.Name}
	for _, p := range m.Params {
		tok := p.Name
		if p.Schema.hasType("string") {
			tok = `"` + tok + `"`
		}
		if !p.Required {
			tok = "( " + tok + " )"
		}
		parts = append(parts, tok)
	}
	return strings.Join(parts, " ")
}

// writes a newline and indentation.
func nl(b *strings.Builder, depth int) {
	b.WriteByte('\n')
	for i := 0; i < depth; i++ {
		b.WriteString("  ")
	}
}

func writeType(b *strings.Builder, s Schema, optional bool) {
	b.WriteString(`<span class="rpc-type">` + esc(annotation(s, optional)) + `</span>`)
}

func writeDesc(b *strings.Builder, s Schema) {
	if s.Description != "" {
		b.WriteString(` <span class="rpc-desc">` + descInline(s.Description) + `</span>`)
	}
}

// writes a schema's nested structure as an indented <ul> tree.
func renderChildren(b *strings.Builder, s Schema, depth int) {
	switch {
	case len(s.variants()) > 0:
		nl(b, depth)
		b.WriteString(`<ul class="rpc-schema rpc-variants">`)
		for i, v := range s.variants() {
			nl(b, depth+1)
			b.WriteString(`<li><span class="rpc-variant">option ` + fmt.Sprint(i+1) + `</span> `)
			writeType(b, v, false)
			writeDesc(b, v)
			renderChildren(b, v, depth+2)
			b.WriteString(`</li>`)
		}
		nl(b, depth)
		b.WriteString(`</ul>`)

	case len(s.Properties) > 0:
		req := toSet(s.Required)
		nl(b, depth)
		b.WriteString(`<ul class="rpc-schema">`)
		for _, p := range orderedProps(s.Properties) {
			renderProp(b, p.Name, p.Schema, req[p.Name], depth+1)
		}
		if add := s.additionalSchema(); add != nil {
			renderProp(b, "…", *add, false, depth+1)
		}
		nl(b, depth)
		b.WriteString(`</ul>`)

	case len(s.PrefixItems) > 0:
		nl(b, depth)
		b.WriteString(`<ul class="rpc-schema">`)
		for i, v := range s.PrefixItems {
			renderProp(b, fmt.Sprintf("[%d]", i), v, true, depth+1)
		}
		nl(b, depth)
		b.WriteString(`</ul>`)

	case s.Items != nil:
		nl(b, depth)
		b.WriteString(`<ul class="rpc-schema rpc-array">`)
		nl(b, depth+1)
		b.WriteString(`<li><span class="rpc-variant">each item</span> `)
		writeType(b, *s.Items, false)
		writeDesc(b, *s.Items)
		renderChildren(b, *s.Items, depth+2)
		b.WriteString(`</li>`)
		nl(b, depth)
		b.WriteString(`</ul>`)

	case s.additionalSchema() != nil:
		nl(b, depth)
		b.WriteString(`<ul class="rpc-schema">`)
		renderProp(b, "…", *s.additionalSchema(), false, depth+1)
		nl(b, depth)
		b.WriteString(`</ul>`)
	}
}

// writes a single named field as an <li> with recursion.
func renderProp(b *strings.Builder, name string, s Schema, required bool, depth int) {
	nl(b, depth)
	b.WriteString(`<li><code class="rpc-key">` + esc(name) + `</code> `)
	writeType(b, s, !required)
	writeDesc(b, s)
	renderChildren(b, s, depth+1)
	b.WriteString(`</li>`)
}

func renderResult(r *Result) string {
	if r == nil {
		return `<p class="rpc-empty">No return value.</p>`
	}
	var b strings.Builder
	b.WriteString(`<p class="rpc-result-head">`)
	writeType(&b, r.Schema, false)
	writeDesc(&b, r.Schema)
	b.WriteString(`</p>`)
	renderChildren(&b, r.Schema, 0)
	return b.String()
}

// writes the arguments table.
func renderArgs(params []Param, named map[string]bool) string {
	if len(params) == 0 {
		return `<p class="rpc-empty">None.</p>`
	}
	var b strings.Builder
	b.WriteString(`<table class="rpc-args"><thead><tr>` +
		`<th>#</th><th>Name</th><th>Type</th><th>Required</th><th>Description</th>` +
		`</tr></thead><tbody>`)
	for i, p := range params {
		nl(&b, 1)
		b.WriteString(`<tr><td>` + fmt.Sprint(i+1) + `</td>`)
		b.WriteString(`<td><code>` + esc(p.Name) + `</code>`)
		if len(p.Aliases) > 0 {
			b.WriteString(`<br><small class="rpc-note">alias: ` + esc(strings.Join(p.Aliases, ", ")) + `</small>`)
		}
		if p.Placeholder {
			b.WriteString(`<br><small class="rpc-note">positional placeholder</small>`)
		}
		if named[p.Name] {
			b.WriteString(`<br><small class="rpc-note">fields may also be passed as named arguments</small>`)
		}
		b.WriteString(`</td>`)
		b.WriteString(`<td><span class="rpc-type">` + esc(typeLabel(p.Schema)) + `</span></td>`)
		if p.Required {
			b.WriteString(`<td class="rpc-req-yes">yes</td>`)
		} else {
			b.WriteString(`<td class="rpc-req-no">no`)
			if len(p.Schema.Default) > 0 {
				b.WriteString(`<br><small class="rpc-note">default: ` + esc(compactJSON(p.Schema.Default)) + `</small>`)
			} else if p.Schema.DefaultHint != "" {
				b.WriteString(`<br><small class="rpc-note">default: ` + esc(p.Schema.DefaultHint) + `</small>`)
			}
			b.WriteString(`</td>`)
		}
		b.WriteString(`<td>`)
		if p.Description != "" {
			b.WriteString(descInline(p.Description))
		}
		renderChildren(&b, p.Schema, 1)
		b.WriteString(`</td></tr>`)
	}
	b.WriteString("\n</tbody></table>")
	return b.String()
}

// extracts the Examples section from help text.
func helpExamples(help string) string {
	_, after, found := strings.Cut(help, "\nExamples:\n")
	if !found {
		return ""
	}
	return strings.TrimRight(after, "\n")
}

var namedArgRe = regexp.MustCompile(`(?m)^([a-z0-9_]+)\s+\(`)

// extracts field names under the help text's "Named Arguments" section.
func helpNamedArgs(help string) map[string]bool {
	_, after, found := strings.Cut(help, "\nNamed Arguments:\n")
	if !found {
		return nil
	}
	for _, end := range []string{"\nResult", "\nExamples:"} {
		if i := strings.Index(after, end); i >= 0 {
			after = after[:i]
		}
	}
	named := map[string]bool{}
	for _, m := range namedArgRe.FindAllStringSubmatch(after, -1) {
		named[m[1]] = true
	}
	return named
}

// marks params whose fields appear under "Named Arguments".
func namedParams(params []Param, named map[string]bool) map[string]bool {
	if len(named) == 0 {
		return nil
	}
	out := map[string]bool{}
	for _, p := range params {
		for _, prop := range orderedProps(p.Schema.Properties) {
			if named[prop.Name] {
				out[p.Name] = true
				break
			}
		}
	}
	return out
}

// assembles HTML body of structured page
func renderBody(m Method, help string) string {
	var b strings.Builder
	b.WriteString(`<div class="rpc-doc">`)

	nl(&b, 1)
	b.WriteString(`<p class="rpc-meta"><span class="rpc-badge">` + esc(m.Category) + `</span>`)
	if m.Deprecated {
		b.WriteString(` <span class="rpc-badge rpc-badge-deprecated">deprecated</span>`)
	}
	b.WriteString(`</p>`)

	nl(&b, 1)
	b.WriteString(`<pre class="rpc-synopsis">` + esc(synopsis(m)) + `</pre>`)

	if desc := descBlock(m.Description); desc != "" {
		nl(&b, 1)
		b.WriteString(`<div class="rpc-description">` + desc + `</div>`)
	}

	nl(&b, 1)
	b.WriteString(`<h2 class="rpc-section">Arguments</h2>`)
	nl(&b, 1)
	b.WriteString(renderArgs(m.Params, namedParams(m.Params, helpNamedArgs(help))))

	nl(&b, 1)
	b.WriteString(`<h2 class="rpc-section">Result</h2>`)
	nl(&b, 1)
	b.WriteString(renderResult(m.Result))

	if ex := helpExamples(help); ex != "" {
		nl(&b, 1)
		b.WriteString(`<h2 class="rpc-section">Examples</h2>`)
		nl(&b, 1)
		b.WriteString(`<pre class="rpc-examples">` + esc(ex) + `</pre>`)
	}

	b.WriteString("\n</div>")
	return b.String()
}

func renderOpenrpcBody(raw json.RawMessage) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		buf.Reset()
		buf.Write(raw)
	}
	return `<pre class="rpc-json">` + esc(buf.String()) + `</pre>`
}
