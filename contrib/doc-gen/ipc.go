package main

import (
	"fmt"
	"html"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// Schema files to document, in the order they should appear.
var schemaFiles = []string{"init", "echo", "mining", "common"}

// Header files providing doxygen-style comments per interface, by Cap'n
// Proto interface name (the schema's `interface Foo` becomes a C++
// `class Foo`). Set the value to "" to opt an interface out of doxygen
// extraction. Adding a new interface to the schemas without a matching
// entry here is treated as an error.
var interfaceHeaders = map[string]string{
	"Init":          "interfaces/init.h",
	"Echo":          "interfaces/echo.h",
	"Mining":        "interfaces/mining.h",
	"BlockTemplate": "interfaces/mining.h",
}

// Cap'n Proto built-in types we render with the keyword-type style.
var builtinTypes = map[string]bool{
	"Void": true, "Bool": true,
	"Int8": true, "Int16": true, "Int32": true, "Int64": true,
	"UInt8": true, "UInt16": true, "UInt32": true, "UInt64": true,
	"Float32": true, "Float64": true,
	"Data": true, "Text": true,
	"List": true,
}

// Methods to hide from generated docs. `construct` is auto-injected by
// libmultiprocess for plumbing. The pattern catches deprecated stubs such as
// `makeMiningOld2`.
var (
	hiddenMethodNames = map[string]bool{"construct": true}
	deprecatedSuffix  = regexp.MustCompile(`Old\d*$`)
)

type Method struct {
	Name string
	Body string
	Doc  string // doxygen description extracted from header (HTML)
	// MakesIface is non-empty when the method is an Init "factory" returning an
	// interface (e.g. `makeMining` returning `Mining.Mining`). The value is the
	// lowercase interface group name we link to.
	MakesIface string
}

type Interface struct {
	Schema  string
	Name    string
	Methods []Method
}

type Struct struct {
	Schema string
	Name   string
	Body   string
}

func generateIPC(bitcoin, version string) {
	srcDir := filepath.Join(bitcoin, "src")
	importDir := filepath.Join("ipc", "libmultiprocess", "include")
	if _, err := os.Stat(filepath.Join(srcDir, importDir)); err != nil {
		log.Fatalf("Cannot find libmultiprocess include dir at %s: %s", importDir, err)
	}

	tmpl := template.Must(template.ParseFiles("command-template.html"))

	var (
		ifaces  []Interface
		structs []Struct
	)

	for _, schema := range schemaFiles {
		i, s := loadSchema(srcDir, importDir, schema)
		ifaces = append(ifaces, i...)
		structs = append(structs, s...)
	}

	// Build cross-reference sets used by the highlighter and link-resolver.
	structSet := map[string]bool{}
	for _, s := range structs {
		structSet[s.Name] = true
	}
	ifaceSet := map[string]bool{}
	for _, i := range ifaces {
		ifaceSet[i.Name] = true
	}

	// Refuse to silently skip a newly-added interface: every interface in
	// the schemas must be mapped to a header file (or to "" to opt out).
	for _, i := range ifaces {
		if _, known := interfaceHeaders[i.Name]; !known {
			log.Fatalf("Unknown interface %q has no entry in interfaceHeaders. "+
				"Add it to contrib/doc-gen/ipc.go (use \"\" to skip doxygen for it).", i.Name)
		}
	}

	// Filter hidden/deprecated methods, and detect `make<Iface>` factories
	// whose return type names a known interface — those get linked.
	for ii := range ifaces {
		var keep []Method
		for _, m := range ifaces[ii].Methods {
			if hiddenMethodNames[m.Name] || deprecatedSuffix.MatchString(m.Name) {
				continue
			}
			if strings.HasPrefix(m.Name, "make") {
				returned := returnedTypeName(clean(m.Body))
				if returned != "" && ifaceSet[returned] {
					m.MakesIface = strings.ToLower(returned)
				}
			}
			keep = append(keep, m)
		}
		ifaces[ii].Methods = keep
	}

	// Attach doxygen documentation harvested from the C++ headers.
	for ii := range ifaces {
		hdr := interfaceHeaders[ifaces[ii].Name]
		if hdr == "" {
			continue
		}
		comments := extractDoxygen(filepath.Join(srcDir, hdr), ifaces[ii].Name)
		for mi := range ifaces[ii].Methods {
			if doc, ok := comments[ifaces[ii].Methods[mi].Name]; ok {
				ifaces[ii].Methods[mi].Doc = linkifyDoc(doc, ifaceSet, structSet, "../../")
			}
		}
	}

	docRoot := filepath.Join("..", "..", "_doc", "en", version, "ipc")

	// Per-method pages and per-interface index pages.
	for _, iface := range ifaces {
		group := strings.ToLower(iface.Name)
		dirname := filepath.Join(docRoot, group)
		if err := os.MkdirAll(dirname, 0o777); err != nil {
			log.Fatalf("Cannot make directory %s: %s", dirname, err)
		}
		// Per-method page.
		for _, m := range iface.Methods {
			page := filepath.Join(dirname, m.Name+".html")
			permalink := fmt.Sprintf("en/doc/%s/ipc/%s/%s/", version, group, m.Name)
			body := renderMethodPage(iface, m, structSet, ifaceSet)
			if err := tmpl.Execute(open(page), CommandData{
				Version:     version,
				Name:        m.Name,
				Description: body,
				Group:       group,
				DocType:     "IPC",
				HTML:        true,
				Permalink:   permalink,
			}); err != nil {
				log.Fatalf("Cannot make page %s: %s", page, err)
			}
		}
		// Per-interface landing index that lists its methods (target for
		// `makeFoo` cross-links from the Init interface).
		ifaceIndex := filepath.Join(dirname, "index.html")
		permalink := fmt.Sprintf("en/doc/%s/ipc/%s/", version, group)
		if err := tmpl.Execute(open(ifaceIndex), CommandData{
			Version:     version,
			Name:        iface.Name,
			Description: renderInterfaceIndex(iface, structSet, ifaceSet),
			Group:       group,
			DocType:     "IPC",
			HTML:        true,
			IfaceIndex:  true,
			Permalink:   permalink,
		}); err != nil {
			log.Fatalf("Cannot make interface index %s: %s", ifaceIndex, err)
		}
	}

	// Single structs page (anchors per struct).
	structsDir := filepath.Join(docRoot, "structs")
	if err := os.MkdirAll(structsDir, 0o777); err != nil {
		log.Fatalf("Cannot make directory %s: %s", structsDir, err)
	}
	sortedStructs := append([]Struct(nil), structs...)
	sort.Slice(sortedStructs, func(i, j int) bool { return sortedStructs[i].Name < sortedStructs[j].Name })
	structNames := make([]string, len(sortedStructs))
	for i, s := range sortedStructs {
		structNames[i] = s.Name
	}
	if err := tmpl.Execute(open(filepath.Join(structsDir, "index.html")), CommandData{
		Version:     version,
		Name:        "structs",
		Description: renderStructsPage(sortedStructs, structSet, ifaceSet),
		Group:       "structs",
		DocType:     "IPC",
		HTML:        true,
		Structs:     structNames,
		Permalink:   fmt.Sprintf("en/doc/%s/ipc/structs/", version),
	}); err != nil {
		log.Fatalf("Cannot make structs page: %s", err)
	}

	// Per-version IPC index.
	if err := tmpl.Execute(open(filepath.Join(docRoot, "index.html")), CommandData{
		Version:   version,
		Name:      "ipcindex",
		Group:     "index",
		DocType:   "IPC",
		Permalink: fmt.Sprintf("en/doc/%s/ipc/", version),
	}); err != nil {
		log.Fatalf("Cannot make ipc index file: %s", err)
	}

	// Site-wide IPC landing page (lists all IPC versions). Created once.
	ipcLanding := filepath.Join("..", "..", "_doc", "en", "ipc-index.html")
	if _, err := os.Stat(ipcLanding); os.IsNotExist(err) {
		if err := tmpl.Execute(open(ipcLanding), CommandData{
			Version:   "ipc-index",
			Name:      "ipcdocsindex",
			Group:     "index",
			DocType:   "IPC",
			Permalink: "en/doc/ipc/",
		}); err != nil {
			log.Fatalf("Cannot make ipc landing page: %s", err)
		}
	}
}

func loadSchema(srcDir, importDir, schema string) ([]Interface, []Struct) {
	schemaRel := filepath.Join("ipc", "capnp", schema+".capnp")
	if _, err := os.Stat(filepath.Join(srcDir, schemaRel)); err != nil {
		log.Fatalf("Cannot find schema %s: %s", schemaRel, err)
	}
	cmd := exec.Command("capnp", "compile",
		"-I", importDir,
		"--src-prefix=ipc/capnp",
		"-ocapnp",
		schemaRel,
	)
	cmd.Dir = srcDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("capnp compile %s failed: %s\n%s", schemaRel, err, out)
	}
	return parseDecls(string(out), schema)
}

// -----------------------------------------------------------------------------
// Cap'n Proto parsing
// -----------------------------------------------------------------------------

func parseDecls(text, schema string) ([]Interface, []Struct) {
	var ifaces []Interface
	var structs []Struct

	lines := strings.Split(text, "\n")
	i := 0
	for i < len(lines) {
		trimmed := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmed, "interface ") {
			iface, end := parseInterface(lines, i, schema)
			ifaces = append(ifaces, iface)
			i = end + 1
			continue
		}
		if strings.HasPrefix(trimmed, "struct ") {
			st, end := parseStruct(lines, i, schema)
			structs = append(structs, st)
			i = end + 1
			continue
		}
		i++
	}
	return ifaces, structs
}

func parseInterface(lines []string, start int, schema string) (Interface, int) {
	header := strings.TrimSpace(lines[start])
	name := strings.Fields(header)[1]
	iface := Interface{Schema: schema, Name: name}
	depth := 0
	end := start
	for j := start; j < len(lines); j++ {
		depth += braceDelta(lines[j])
		end = j
		if depth == 0 {
			break
		}
		if j == start {
			continue
		}
		body := strings.TrimSpace(lines[j])
		if body == "" || body == "}" {
			continue
		}
		fields := strings.Fields(body)
		if len(fields) == 0 {
			continue
		}
		iface.Methods = append(iface.Methods, Method{Name: fields[0], Body: body})
	}
	return iface, end
}

func parseStruct(lines []string, start int, schema string) (Struct, int) {
	header := strings.TrimSpace(lines[start])
	name := strings.Fields(header)[1]
	depth := 0
	end := start
	var body []string
	for j := start; j < len(lines); j++ {
		depth += braceDelta(lines[j])
		body = append(body, lines[j])
		end = j
		if depth == 0 {
			break
		}
	}
	return Struct{Schema: schema, Name: name, Body: strings.Join(body, "\n")}, end
}

func braceDelta(line string) int {
	if i := strings.Index(line, "#"); i >= 0 {
		line = line[:i]
	}
	delta := 0
	for _, r := range line {
		switch r {
		case '{':
			delta++
		case '}':
			delta--
		}
	}
	return delta
}

// returnedTypeName extracts the type name from a method body's return tuple,
// e.g. `makeMining @3 (context :Proxy.Context) -> (result :Mining.Mining);`
// returns "Mining" (the trailing identifier of the return type).
var reReturnType = regexp.MustCompile(`->\s*\([^)]*:\s*([A-Za-z_][A-Za-z0-9_.]*)\s*\)`)

func returnedTypeName(body string) string {
	m := reReturnType.FindStringSubmatch(body)
	if m == nil {
		return ""
	}
	parts := strings.Split(m[1], ".")
	return parts[len(parts)-1]
}

// -----------------------------------------------------------------------------
// Doxygen extraction from C++ headers
// -----------------------------------------------------------------------------

var (
	reClassStart  = regexp.MustCompile(`^\s*class\s+(\w+)\b`)
	reMethodDecl  = regexp.MustCompile(`^\s*virtual\b.*?\b(\w+)\s*\(`)
	reSlashSlash3 = regexp.MustCompile(`^\s*///\s?`)
	reSlashBang   = regexp.MustCompile(`^\s*//!\s?`)
	reBlockStart  = regexp.MustCompile(`^\s*/\*\*+\s?`)
	reBlockMid    = regexp.MustCompile(`^\s*\*\s?`)
	reBlockEnd    = regexp.MustCompile(`\*+/\s*$`)
)

// extractDoxygen returns a map of methodName -> rendered HTML description for
// methods declared inside `class <className>` of the given header file.
func extractDoxygen(path, className string) map[string]string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("warn: cannot read %s: %s", path, err)
		return nil
	}
	lines := strings.Split(string(data), "\n")

	classStart := -1
	for i, line := range lines {
		m := reClassStart.FindStringSubmatch(line)
		if m != nil && m[1] == className {
			classStart = i
			break
		}
	}
	if classStart < 0 {
		return nil
	}
	classEnd := len(lines)
	depth := 0
	started := false
	for i := classStart; i < len(lines); i++ {
		line := stripStringsAndComments(lines[i])
		for _, r := range line {
			if r == '{' {
				depth++
				started = true
			} else if r == '}' {
				depth--
				if started && depth == 0 {
					classEnd = i
					break
				}
			}
		}
		if started && depth == 0 {
			break
		}
	}

	out := map[string]string{}
	var pendingComment []string
	flush := func() { pendingComment = nil }

	for i := classStart + 1; i < classEnd; i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			continue
		}
		if reBlockStart.MatchString(line) {
			pendingComment = nil
			if reBlockEnd.MatchString(line) {
				inner := line
				inner = reBlockStart.ReplaceAllString(inner, "")
				inner = reBlockEnd.ReplaceAllString(inner, "")
				pendingComment = append(pendingComment, strings.TrimSpace(inner))
				continue
			}
			inner := reBlockStart.ReplaceAllString(line, "")
			pendingComment = append(pendingComment, strings.TrimSpace(inner))
			for j := i + 1; j < classEnd; j++ {
				i = j
				ln := lines[j]
				if reBlockEnd.MatchString(ln) {
					ln = reBlockEnd.ReplaceAllString(ln, "")
					ln = reBlockMid.ReplaceAllString(ln, "")
					pendingComment = append(pendingComment, strings.TrimRight(ln, " \t"))
					break
				}
				ln = reBlockMid.ReplaceAllString(ln, "")
				pendingComment = append(pendingComment, strings.TrimRight(ln, " \t"))
			}
			continue
		}
		if reSlashBang.MatchString(line) {
			pendingComment = append(pendingComment, reSlashBang.ReplaceAllString(line, ""))
			continue
		}
		if reSlashSlash3.MatchString(line) {
			pendingComment = append(pendingComment, reSlashSlash3.ReplaceAllString(line, ""))
			continue
		}
		if m := reMethodDecl.FindStringSubmatch(line); m != nil {
			name := m[1]
			if len(pendingComment) > 0 {
				out[name] = renderDoxygen(pendingComment)
			}
			flush()
			continue
		}
		if !strings.HasPrefix(trimmed, "//") {
			flush()
		}
	}
	return out
}

func stripStringsAndComments(line string) string {
	if i := strings.Index(line, "//"); i >= 0 {
		line = line[:i]
	}
	for {
		s := strings.Index(line, "/*")
		if s < 0 {
			break
		}
		e := strings.Index(line[s:], "*/")
		if e < 0 {
			line = line[:s]
			break
		}
		line = line[:s] + line[s+e+2:]
	}
	return line
}

func renderDoxygen(lines []string) string {
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return ""
	}

	type tagged struct {
		Kind string // "prose", "param", "retval", "returns"
		Key  string
		Dir  string // "in", "out", "in,out" for @param
		Text []string
	}
	var sections []tagged
	cur := tagged{Kind: "prose"}
	push := func() {
		if len(cur.Text) > 0 || cur.Key != "" {
			sections = append(sections, cur)
		}
		cur = tagged{Kind: "prose"}
	}

	tagRe := regexp.MustCompile(`^\s*@(param(?:\[[^\]]+\])?|retval|returns?)(?:\s+|$)(\S*)\s*(.*)$`)
	for _, ln := range lines {
		if m := tagRe.FindStringSubmatch(ln); m != nil {
			push()
			tag := m[1]
			key := m[2]
			text := strings.TrimSpace(m[3])
			switch {
			case strings.HasPrefix(tag, "param"):
				dir := ""
				if strings.HasPrefix(tag, "param[") {
					dir = strings.TrimSuffix(strings.TrimPrefix(tag, "param["), "]")
				}
				cur = tagged{Kind: "param", Key: key, Dir: dir, Text: []string{text}}
			case tag == "retval":
				cur = tagged{Kind: "retval", Key: key, Text: []string{text}}
			default:
				cur = tagged{Kind: "returns", Key: key, Text: []string{text}}
			}
			continue
		}
		cur.Text = append(cur.Text, strings.TrimSpace(ln))
	}
	push()

	var b strings.Builder
	var inDL bool
	closeDL := func() {
		if inDL {
			b.WriteString("</dl>\n")
			inDL = false
		}
	}
	openDL := func() {
		if !inDL {
			b.WriteString(`<dl class="ipc-doc-tags">` + "\n")
			inDL = true
		}
	}
	for _, sec := range sections {
		text := strings.TrimSpace(strings.Join(sec.Text, " "))
		switch sec.Kind {
		case "prose":
			closeDL()
			for _, p := range splitParas(sec.Text) {
				if p == "" {
					continue
				}
				fmt.Fprintf(&b, "<p>%s</p>\n", html.EscapeString(p))
			}
		case "param":
			openDL()
			label := "param"
			if sec.Dir != "" {
				label = fmt.Sprintf("param [%s]", sec.Dir)
			}
			fmt.Fprintf(&b, `  <dt><strong>%s</strong> <code>%s</code></dt>`+"\n",
				html.EscapeString(label), html.EscapeString(sec.Key))
			fmt.Fprintf(&b, "  <dd>%s</dd>\n", html.EscapeString(text))
		case "retval":
			openDL()
			fmt.Fprintf(&b, `  <dt><strong>returns</strong> <code>%s</code></dt>`+"\n",
				html.EscapeString(sec.Key))
			fmt.Fprintf(&b, "  <dd>%s</dd>\n", html.EscapeString(text))
		case "returns":
			openDL()
			b.WriteString("  <dt><strong>returns</strong></dt>\n")
			full := strings.TrimSpace(sec.Key + " " + text)
			fmt.Fprintf(&b, "  <dd>%s</dd>\n", html.EscapeString(full))
		}
	}
	closeDL()
	return b.String()
}

func splitParas(lines []string) []string {
	var paras []string
	var cur []string
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			if len(cur) > 0 {
				paras = append(paras, strings.Join(cur, " "))
				cur = nil
			}
			continue
		}
		cur = append(cur, ln)
	}
	if len(cur) > 0 {
		paras = append(paras, strings.Join(cur, " "))
	}
	return paras
}

// linkifyDoc rewrites already-rendered doxygen HTML so that whole-word
// occurrences of known interface or struct names become links to their
// documentation pages. Linking happens inside text nodes only — existing
// tags and href attributes are left untouched.
func linkifyDoc(htmlIn string, ifaceSet, structSet map[string]bool, ifaceHrefBase string) string {
	if len(ifaceSet) == 0 && len(structSet) == 0 {
		return htmlIn
	}
	var b strings.Builder
	i := 0
	for i < len(htmlIn) {
		if htmlIn[i] == '<' {
			end := strings.IndexByte(htmlIn[i:], '>')
			if end < 0 {
				b.WriteString(htmlIn[i:])
				break
			}
			b.WriteString(htmlIn[i : i+end+1])
			i += end + 1
			continue
		}
		nextTag := strings.IndexByte(htmlIn[i:], '<')
		var seg string
		if nextTag < 0 {
			seg = htmlIn[i:]
			i = len(htmlIn)
		} else {
			seg = htmlIn[i : i+nextTag]
			i += nextTag
		}
		b.WriteString(linkifySegment(seg, ifaceSet, structSet, ifaceHrefBase))
	}
	return b.String()
}

var reIdentToken = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)

func linkifySegment(seg string, ifaceSet, structSet map[string]bool, ifaceHrefBase string) string {
	return reIdentToken.ReplaceAllStringFunc(seg, func(tok string) string {
		switch {
		case ifaceSet[tok]:
			return fmt.Sprintf(`<a href="%s%s/">%s</a>`,
				ifaceHrefBase, strings.ToLower(tok), tok)
		case structSet[tok]:
			return fmt.Sprintf(`<a href="%sstructs/#%s">%s</a>`,
				ifaceHrefBase, tok, tok)
		}
		return tok
	})
}

// -----------------------------------------------------------------------------
// Cleanup of canonical capnp output
// -----------------------------------------------------------------------------

var (
	reLongID       = regexp.MustCompile(`\s+@0x[0-9a-fA-F]+`)
	reMpAnno       = regexp.MustCompile(`\s+\$import\s+"/mp/proxy\.capnp"\.\w+\([^)]*\)`)
	reImportRef    = regexp.MustCompile(`import\s+"[^"]+"\.`)
	reTrailComment = regexp.MustCompile(`\s*#[^\n]*`)
	reMpBare       = regexp.MustCompile(`\s+\$import\s+"/mp/proxy\.capnp"`)
)

func clean(s string) string {
	s = reMpAnno.ReplaceAllString(s, "")
	s = reMpBare.ReplaceAllString(s, "")
	s = reLongID.ReplaceAllString(s, "")
	s = reImportRef.ReplaceAllString(s, "")
	s = reTrailComment.ReplaceAllString(s, "")
	return s
}

// wrapMethodSig pretty-prints a single-line method signature so each parameter
// goes on its own line.
func wrapMethodSig(s string) string {
	open := strings.IndexByte(s, '(')
	if open < 0 {
		return s
	}
	close := matchParen(s, open)
	if close < 0 {
		return s
	}
	prefix := s[:open]
	args := s[open+1 : close]
	rest := s[close+1:]

	var b strings.Builder
	b.WriteString(prefix)
	b.WriteString("(\n")
	for _, a := range splitTopLevel(args, ',') {
		fmt.Fprintf(&b, "  %s,\n", strings.TrimSpace(a))
	}
	out := strings.TrimRight(b.String(), ",\n") + "\n)"

	rest = strings.TrimSpace(rest)
	if strings.HasPrefix(rest, "->") {
		rt := strings.TrimSpace(rest[2:])
		if strings.HasPrefix(rt, "(") {
			rclose := matchParen(rt, 0)
			if rclose > 0 {
				inner := rt[1:rclose]
				tail := rt[rclose+1:]
				var rb strings.Builder
				rb.WriteString(out)
				rb.WriteString(" -> (\n")
				for _, a := range splitTopLevel(inner, ',') {
					fmt.Fprintf(&rb, "  %s,\n", strings.TrimSpace(a))
				}
				out = strings.TrimRight(rb.String(), ",\n") + "\n)" + tail
				return out
			}
		}
		out += " " + rest
		return out
	}
	if rest != "" {
		out += rest
	}
	return out
}

func matchParen(s string, open int) int {
	depth := 0
	for i := open; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

func splitTopLevel(s string, sep byte) []string {
	var out []string
	depth := 0
	start := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			depth++
		case ')', ']', '}':
			depth--
		case sep:
			if depth == 0 {
				out = append(out, s[start:i])
				start = i + 1
			}
		}
	}
	out = append(out, s[start:])
	return out
}

// -----------------------------------------------------------------------------
// IPC page rendering
// -----------------------------------------------------------------------------

func renderMethodPage(iface Interface, m Method, structSet, ifaceSet map[string]bool) string {
	cleaned := wrapMethodSig(clean(m.Body))
	highlighted := highlightCode(cleaned, structSet, ifaceSet, "../../structs/", "../../")

	var b strings.Builder
	fmt.Fprintf(&b, `<p class="ipc-context">Method on interface <a href="../"><code>%s</code></a>.</p>`+"\n",
		html.EscapeString(iface.Name))
	b.WriteString(`<div class="highlight"><pre><code>`)
	b.WriteString(highlighted)
	b.WriteString(`</code></pre></div>` + "\n")
	if m.Doc != "" {
		b.WriteString(`<div class="ipc-doc">` + "\n")
		b.WriteString(m.Doc)
		b.WriteString("</div>\n")
	}
	if m.MakesIface != "" {
		fmt.Fprintf(&b, `<p class="ipc-makes">Returns the <a href="../../%s/">%s</a> interface.</p>`+"\n",
			html.EscapeString(m.MakesIface), html.EscapeString(strings.Title(m.MakesIface)))
	}
	return b.String()
}

func renderInterfaceIndex(iface Interface, structSet, ifaceSet map[string]bool) string {
	var b strings.Builder
	fmt.Fprintf(&b, `<p>Methods on the <code>%s</code> IPC interface:</p>`+"\n",
		html.EscapeString(iface.Name))
	b.WriteString(`<ul>` + "\n")
	for _, m := range iface.Methods {
		fmt.Fprintf(&b, `  <li><a href="%s/">%s</a></li>`+"\n",
			html.EscapeString(m.Name), html.EscapeString(m.Name))
	}
	b.WriteString(`</ul>` + "\n")
	return b.String()
}

func renderStructsPage(structs []Struct, structSet, ifaceSet map[string]bool) string {
	var b strings.Builder
	b.WriteString(`<p>Shared Cap'n Proto structs used by the IPC interfaces.</p>` + "\n")
	for _, s := range structs {
		cleaned := clean(s.Body)
		highlighted := highlightCode(cleaned, structSet, ifaceSet, "", "")
		fmt.Fprintf(&b, `<h2 id="%s">%s</h2>`+"\n", html.EscapeString(s.Name), html.EscapeString(s.Name))
		b.WriteString(`<div class="highlight"><pre><code>`)
		b.WriteString(highlighted)
		b.WriteString(`</code></pre></div>` + "\n")
	}
	return b.String()
}

var (
	tokenRe  = regexp.MustCompile(`@\d+|0x[0-9a-fA-F]+|\d+\.\d+e[+-]?\d+|\d+\.\d+|\d+|"[^"]*"|->|[A-Za-z_][A-Za-z0-9_]*|[(){},:;=]|\S`)
	keywords = map[string]bool{
		"interface": true, "struct": true, "const": true,
		"using": true, "annotation": true, "enum": true, "union": true, "group": true,
	}
	literals = map[string]bool{"true": true, "false": true}
)

// highlightCode tokenises a single declaration and emits rouge-compatible
// span/anchor markup. structHrefBase is prepended to "#Name" anchors when a
// token names a known struct; ifaceHrefBase is prepended to "name/" links when
// a token names a known interface (use "" to disable that kind of link).
func highlightCode(src string, structSet, ifaceSet map[string]bool, structHrefBase, ifaceHrefBase string) string {
	var b strings.Builder
	for li, line := range strings.Split(src, "\n") {
		if li > 0 {
			b.WriteString("\n")
		}
		i := 0
		for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
			b.WriteByte(line[i])
			i++
		}
		rest := line[i:]
		matches := tokenRe.FindAllStringIndex(rest, -1)
		prevEnd := 0
		var prevToken string
		for _, m := range matches {
			if m[0] > prevEnd {
				b.WriteString(html.EscapeString(rest[prevEnd:m[0]]))
			}
			tok := rest[m[0]:m[1]]
			cls, link := classify(tok, prevToken, structSet, ifaceSet, structHrefBase, ifaceHrefBase)
			if link != "" {
				fmt.Fprintf(&b, `<a class="%s" href="%s">%s</a>`, cls, link, html.EscapeString(tok))
			} else if cls != "" {
				fmt.Fprintf(&b, `<span class="%s">%s</span>`, cls, html.EscapeString(tok))
			} else {
				b.WriteString(html.EscapeString(tok))
			}
			prevToken = tok
			prevEnd = m[1]
		}
		if prevEnd < len(rest) {
			b.WriteString(html.EscapeString(rest[prevEnd:]))
		}
	}
	return b.String()
}

func isIdent(tok string) bool {
	if tok == "" {
		return false
	}
	c := tok[0]
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_'
}

func classify(tok, prevToken string, structSet, ifaceSet map[string]bool, structHrefBase, ifaceHrefBase string) (cls, link string) {
	switch {
	case tok == "->" || tok == "=":
		return "o", ""
	case tok == "(" || tok == ")" || tok == "{" || tok == "}" || tok == "," || tok == ";" || tok == ":":
		return "p", ""
	case strings.HasPrefix(tok, "@"):
		return "nd", ""
	case strings.HasPrefix(tok, `"`):
		return "s", ""
	case len(tok) > 0 && tok[0] >= '0' && tok[0] <= '9':
		if strings.Contains(tok, ".") || strings.ContainsAny(tok, "eE") {
			return "mf", ""
		}
		return "mi", ""
	case keywords[tok]:
		return "kd", ""
	case literals[tok]:
		return "kc", ""
	case builtinTypes[tok]:
		return "kt", ""
	case isIdent(tok):
		if prevToken == "interface" || prevToken == "struct" || prevToken == "const" {
			return "nc", ""
		}
		first := tok[0]
		if first >= 'a' && first <= 'z' {
			return "nv", ""
		}
		if structSet[tok] {
			href := "#" + tok
			if structHrefBase != "" {
				href = structHrefBase + "#" + tok
			}
			return "nc", href
		}
		if ifaceSet[tok] && ifaceHrefBase != "" {
			return "nc", ifaceHrefBase + strings.ToLower(tok) + "/"
		}
		return "nc", ""
	}
	return "", ""
}
