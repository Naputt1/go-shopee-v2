package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
	"time"
)

type Config struct {
	Modules         map[string]ModuleConfig   `json:"modules"`
	StaticModules   map[string]ModuleConfig   `json:"static_modules"`
	TypeOverrides   map[string]string         `json:"type_overrides"`
	TypeDefinitions map[string]TypeDefinition `json:"type_definitions"`
}

type ModuleConfig struct {
	FileName   string   `json:"file_name"`
	IgnoreAPIs []string `json:"ignore_apis"`
}

type TypeDefinition struct {
	BaseType string            `json:"base_type"`
	Values   map[string]string `json:"values"`
}

type ModuleListResponse struct {
	Modules []Module `json:"modules"`
}

type Module struct {
	ModuleID   int    `json:"module_id"`
	ModuleName string `json:"module_name"`
	Type       int    `json:"type"`
	Items      []Item `json:"items"`
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type APIInfo struct {
	APIName        string `json:"api_name"`
	APIType        string `json:"api_type"`
	Define         string `json:"define"`
	Params         string `json:"params"` // JSON string
	Path           string `json:"path"`
	RequestSample  string `json:"request_sample"`  // JSON string
	ResponseSample string `json:"response_sample"` // JSON string
	ModuleID       int    `json:"module_id"`
}

type Sample struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Params struct {
	RequestParams  []Param `json:"request_params"`
	ResponseParams []Param `json:"response_params"`
}

type Param struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Required    string  `json:"required"`
	Children    []Param `json:"children"`
}

type MethodData struct {
	Name         string
	Path         string
	Method       string // GET or POST
	APIType      string // Shop or Merchant
	Description  string
	DocURL       string
	RequestType  string
	ResponseType string
	IsGet        bool
	FullAPIName  string
	FixtureName  string
	ModuleID     int
	Type         int
}

type StructData struct {
	Name   string
	Fields []FieldData
}

type FieldData struct {
	Name        string
	Type        string
	JSONTag     string
	URLTag      string
	Description string
}

type TemplateData struct {
	PackageName string
	ModuleName  string
	Methods     []MethodData
	Structs     []StructData
	Constants   []ConstantData
}

type ConstantData struct {
	TypeName string
	BaseType string
	Values   []ConstantValue
}

type ConstantValue struct {
	Name  string
	Value string
}

type ServiceInfo struct {
	Name      string
	Interface string
	Impl      string
}

type APIContext struct {
	Module    Module
	ModConfig ModuleConfig
	Info      *APIInfo
	Item      Item
	Method    *MethodData
	Params    Params
	ModName   string
}

var (
	structSignatures        = make(map[string]string)     // signature -> chosen struct name
	allGeneratedStructs     = make(map[string]StructData) // name -> struct data
	structNamesToSignatures = make(map[string]string)     // name -> signature
)

func main() {
	// Load config
	configData, err := ioutil.ReadFile("generator_config.json")
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Printf("Error parsing config: %v\n", err)
		os.Exit(1)
	}

	// Fetch modules
	modules, err := fetchModules()
	if err != nil {
		fmt.Printf("Error fetching modules: %v\n", err)
		os.Exit(1)
	}

	var services []ServiceInfo
	var apiContexts []APIContext

	// 1. Discover all APIs
	for _, mod := range modules {
		modConfig, ok := config.Modules[mod.ModuleName]
		if ok {
			fmt.Printf("Discovering module %s\n", mod.ModuleName)
			moduleNameClean := toCamelCase(strings.ReplaceAll(strings.ReplaceAll(mod.ModuleName, "(CB seller only)", ""), " ", ""))
			if moduleNameClean == "Media" {
				moduleNameClean = "MediaSpace"
			}

			for _, item := range mod.Items {
				item.Name = strings.TrimSpace(item.Name)
				if !strings.HasPrefix(item.Name, "v2.") {
					continue
				}

				ignored := false
				for _, ignore := range modConfig.IgnoreAPIs {
					if item.Name == ignore {
						ignored = true
						break
					}
				}
				if ignored {
					continue
				}

				time.Sleep(10 * time.Millisecond)
				fmt.Printf("  Fetching API %s\n", item.Name)
				info, err := fetchAPIInfo(item.Name)
				if err != nil {
					fmt.Printf("    Warning: skipping API %s due to error: %v\n", item.Name, err)
					continue
				}

				var params Params
				if info.Params != "" {
					json.Unmarshal([]byte(info.Params), &params)
				}

				parts := strings.Split(item.Name, ".")
				baseName := ""
				if len(parts) >= 3 {
					baseName = strings.Join(parts[2:], "_")
				} else {
					baseName = parts[len(parts)-1]
				}
				methodName := toCamelCase(baseName)

				var requestSamples []Sample
				if info.RequestSample != "" {
					json.Unmarshal([]byte(info.RequestSample), &requestSamples)
				}
				var responseSamples []Sample
				if info.ResponseSample != "" {
					json.Unmarshal([]byte(info.ResponseSample), &responseSamples)
				}

				isGet := false
				for _, s := range requestSamples {
					if s.Type == "cURL" && strings.Contains(s.Value, "GET") {
						isGet = true
						break
					}
				}

				fixtureName := item.Name
				fixtureName = strings.ReplaceAll(fixtureName, " ", "_")
				fixtureName = strings.ReplaceAll(fixtureName, "(", "_")
				fixtureName = strings.ReplaceAll(fixtureName, ")", "_")
				fixtureName = strings.Map(func(r rune) rune {
					if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '.' || r == '_' {
						return r
					}
					return '_'
				}, fixtureName)

				method := &MethodData{
					Name:         methodName,
					Path:         strings.TrimPrefix(info.Path, "/api/v2/"),
					Method:       "POST",
					APIType:      info.APIType,
					Description:  info.Define,
					DocURL:       fmt.Sprintf("https://open.shopee.com/documents/v2/%s?module=%d&type=%d", item.Name, mod.ModuleID, mod.Type),
					RequestType:  moduleNameClean + methodName + "Request",
					ResponseType: moduleNameClean + methodName + "Response",
					IsGet:        isGet,
					FullAPIName:  item.Name,
					FixtureName:  fixtureName,
					ModuleID:     mod.ModuleID,
					Type:         mod.Type,
				}
				if isGet {
					method.Method = "GET"
				}

				apiContexts = append(apiContexts, APIContext{
					Module:    mod,
					ModConfig: modConfig,
					Info:      info,
					Item:      item,
					Method:    method,
					Params:    params,
					ModName:   moduleNameClean,
				})

				// Save Fixture
				for _, s := range responseSamples {
					if s.Type == "JSON" {
						fixturePath := fmt.Sprintf("fixtures/%s_resp.json", fixtureName)
						ioutil.WriteFile(fixturePath, []byte(s.Value), 0644)
					}
				}
			}

			services = append(services, ServiceInfo{
				Name:      moduleNameClean,
				Interface: moduleNameClean + "Service",
				Impl:      moduleNameClean + "ServiceOp",
			})
		} else if _, isStatic := config.StaticModules[mod.ModuleName]; isStatic {
			fmt.Printf("Module %s is static, skipping generation...\n", mod.ModuleName)
			auditStaticModule(mod, config)
			serviceName := toCamelCase(strings.ReplaceAll(strings.ReplaceAll(mod.ModuleName, "(CB seller only)", ""), " ", ""))
			if serviceName == "Media" {
				serviceName = "MediaSpace"
			}
			found := false
			for _, s := range services {
				if s.Name == serviceName {
					found = true
					break
				}
			}
			if !found {
				services = append(services, ServiceInfo{
					Name:      serviceName,
					Interface: serviceName + "Service",
					Impl:      serviceName + "ServiceOp",
				})
			}
		}
	}

	// 2. Process all Requests first (to prioritize shorter names)
	fmt.Println("Processing Request types...")
	for i := range apiContexts {
		ctx := &apiContexts[i]
		chain := []string{ctx.ModName, ctx.Method.Name, "Request"}
		reqStruct := generateStruct(chain, ctx.Params.RequestParams, ctx.Method.IsGet, config, true)
		if reqStruct.Name != "" {
			allGeneratedStructs[reqStruct.Name] = reqStruct
		}
		ctx.Method.RequestType = reqStruct.Name // Update the method data with the chosen (possibly deduplicated) name
	}

	// 3. Process all Responses
	fmt.Println("Processing Response types...")
	for i := range apiContexts {
		ctx := &apiContexts[i]

		mainResp := StructData{
			Name: ctx.Method.ResponseType,
			Fields: []FieldData{
				{Name: "", Type: "BaseResponse", JSONTag: ",inline", Description: "Common response fields"},
			},
		}

		commonFields := map[string]bool{
			"request_id": true,
			"error":      true,
			"message":    true,
			"warning":    true,
		}

		for _, p := range ctx.Params.ResponseParams {
			if commonFields[p.Name] {
				continue
			}

			if p.Name == "response" {
				respDataChain := []string{ctx.ModName, ctx.Method.Name, "ResponseData"}
				dataResp := generateStruct(respDataChain, p.Children, false, config, false)
				if dataResp.Name != "" {
					allGeneratedStructs[dataResp.Name] = dataResp
					mainResp.Fields = append(mainResp.Fields, FieldData{
						Name:        "Response",
						Type:        dataResp.Name,
						JSONTag:     "response",
						Description: "Actual response data",
					})
				} else {
					mainResp.Fields = append(mainResp.Fields, FieldData{
						Name:        "Response",
						Type:        "interface{}",
						JSONTag:     "response",
						Description: "Actual response data",
					})
				}
			} else {
				// Field outside of "response"
				fieldName := toCamelCase(p.Name)
				fieldType := mapType(p.Type, p.Name, p.Children, []string{ctx.ModName, ctx.Method.Name, "Response"}, config, false, p.Required)
				mainResp.Fields = append(mainResp.Fields, FieldData{
					Name:        fieldName,
					Type:        fieldType,
					JSONTag:     p.Name + ",omitempty",
					Description: p.Description,
				})
			}
		}

		// Re-run pickName for the main response AFTER its structure is fully determined
		mainResp.Name = pickName(getStructSignature(mainResp), []string{ctx.ModName, ctx.Method.Name, "Response"})
		ctx.Method.ResponseType = mainResp.Name

		allGeneratedStructs[mainResp.Name] = mainResp
	}

	// 4. Render all modules
	moduleMethods := make(map[string][]MethodData)
	for _, ctx := range apiContexts {
		moduleMethods[ctx.ModConfig.FileName] = append(moduleMethods[ctx.ModConfig.FileName], *ctx.Method)
	}

	for fileName, methods := range moduleMethods {
		serviceName := ""
		for _, ctx := range apiContexts {
			if ctx.ModConfig.FileName == fileName {
				serviceName = ctx.ModName
				break
			}
		}

		data := TemplateData{
			PackageName: "goshopee",
			ModuleName:  serviceName,
			Methods:     methods,
		}
		renderModule(fileName, data)
		renderTest(strings.TrimSuffix(fileName, ".go")+"_test.go", data)
	}

	if err := generateSharedTypesFile(config); err != nil {
		fmt.Printf("Error generating types.go: %v\n", err)
	}

	// Add static services first
	staticOrder := []string{"Util", "Auth", "MediaSpace"}
	var finalServices []ServiceInfo
	for _, name := range staticOrder {
		finalServices = append(finalServices, ServiceInfo{
			Name:      name,
			Interface: name + "Service",
			Impl:      name + "ServiceOp",
		})
	}
	for _, s := range services {
		isStatic := false
		for _, sn := range staticOrder {
			if s.Name == sn {
				isStatic = true
				break
			}
		}
		if !isStatic {
			finalServices = append(finalServices, s)
		}
	}

	updateGoshopee(finalServices)
}

func fetchModules() ([]Module, error) {
	resp, err := http.Get("https://open.shopee.com/opservice/api/v1/doc/module/?version=2")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res ModuleListResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res.Modules, nil
}

func fetchAPIInfo(apiName string) (*APIInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://open.shopee.com/opservice/api/v1/doc/api/?version=2&api_name=%s", apiName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API %s returned status %d: %s", apiName, resp.StatusCode, string(body))
	}

	var info APIInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("API %s: error unmarshaling JSON: %v", apiName, err)
	}
	return &info, nil
}

func auditStaticModule(mod Module, config Config) {
	// Placeholder
}

func isParamRequired(req string) bool {
	req = strings.ToLower(req)
	return req == "yes" || req == "true"
}

func generateStruct(chain []string, params []Param, isGet bool, config Config, isRequest bool) StructData {
	if len(chain) == 0 {
		return StructData{}
	}

	s := StructData{}
	for _, p := range params {
		if p.Name == "partner_id" || p.Name == "shop_id" || p.Name == "merchant_id" || p.Name == "access_token" || p.Name == "timestamp" || p.Name == "sign" {
			continue
		}

		fieldName := toCamelCase(p.Name)
		fieldType := mapType(p.Type, p.Name, p.Children, chain, config, isRequest, p.Required)

		required := isParamRequired(p.Required)
		desc := p.Description
		jsonTag := p.Name

		if required {
			desc = "[Required] " + desc
		} else {
			desc = "[Optional] " + desc
			if !strings.HasPrefix(fieldType, "*") && !strings.HasPrefix(fieldType, "[]") && !strings.HasPrefix(fieldType, "map[") {
				fieldType = "*" + fieldType
			}
			jsonTag += ",omitempty"
		}

		field := FieldData{
			Name:        fieldName,
			Type:        fieldType,
			JSONTag:     jsonTag,
			Description: desc,
		}
		if isGet && isRequest {
			field.URLTag = p.Name
		}
		s.Fields = append(s.Fields, field)
	}

	if len(s.Fields) == 0 {
		return StructData{}
	}

	sig := getStructSignature(s)
	s.Name = pickName(sig, chain)
	allGeneratedStructs[s.Name] = s
	return s
}

func pickName(signature string, chain []string) string {
	isTopLevel := false
	if len(chain) > 0 {
		suffix := chain[len(chain)-1]
		if suffix == "Request" || suffix == "Response" || suffix == "ResponseData" {
			isTopLevel = true
		}
	}

	if !isTopLevel {
		if name, ok := structSignatures[signature]; ok {
			return name
		}
	}

	// Prepare candidate parts
	var parts []string
	for _, p := range chain {
		parts = append(parts, toCamelCase(p))
	}

	// Try increasing length from the end
	// For top-level, we start with 2 parts (e.g. MethodName + Request)
	start := 1
	if isTopLevel && len(parts) >= 2 {
		start = 2
	}

	for i := start; i <= len(parts); i++ {
		candidateParts := parts[len(parts)-i:]
		candidate := strings.Join(candidateParts, "")
		if candidate == "" {
			continue
		}
		candidate = applyManualPrefix(candidate)

		existingSig, taken := structNamesToSignatures[candidate]
		if !taken {
			structNamesToSignatures[candidate] = signature
			if !isTopLevel {
				structSignatures[signature] = candidate
			}
			return candidate
		}

		// If taken, and it's NOT top-level, we can deduplicate if signature matches
		if !isTopLevel && existingSig == signature {
			return candidate
		}

		// If taken and it IS top-level, we MUST try a longer name even if signature matches
		// because we want unique types for each API.
	}

	// Fallback to full name
	fullName := strings.Join(parts, "")
	fullName = applyManualPrefix(fullName)
	return fullName
}

func applyManualPrefix(name string) string {
	collides := map[string]bool{
		"ImageInfo":               true,
		"Option":                  true,
		"Item":                    true,
		"Order":                   true,
		"UploadImageResponseData": true,
		"UploadImageResponse":     true,
	}
	if collides[name] {
		return "Shared" + name
	}
	return name
}

func getStructSignature(s StructData) string {
	var fieldSigs []string
	for _, f := range s.Fields {
		fieldSigs = append(fieldSigs, fmt.Sprintf("%s:%s:%s", f.JSONTag, f.Type, f.URLTag))
	}
	sort.Strings(fieldSigs)
	return strings.Join(fieldSigs, "|")
}

func mapType(shopeeType, fieldName string, children []Param, chain []string, config Config, isRequest bool, requiredStr string) string {
	if override, ok := config.TypeOverrides[fieldName]; ok {
		return override
	}

	switch shopeeType {
	case "int", "int32", "int64", "timestamp":
		return "int64"
	case "float", "double":
		return "float64"
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "object":
		if len(children) == 0 {
			return "interface{}"
		}
		newChain := append(chain, fieldName)
		subStruct := generateStruct(newChain, children, false, config, isRequest)
		if subStruct.Name == "" {
			return "interface{}"
		}
		return "*" + subStruct.Name
	case "object[]":
		if len(children) == 0 {
			return "[]interface{}"
		}
		shortName := toCamelCase(fieldName)
		if strings.HasSuffix(shortName, "List") && len(shortName) > 4 {
			shortName = strings.TrimSuffix(shortName, "List")
		}
		newChain := append(chain, shortName)
		subStruct := generateStruct(newChain, children, false, config, isRequest)
		if subStruct.Name == "" {
			return "[]interface{}"
		}
		return "[]" + subStruct.Name
	case "int[]", "int64[]":
		return "[]int64"
	case "string[]":
		return "[]string"
	default:
		return "interface{}"
	}
}

func toCamelCase(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, ".", "_")
	s = strings.ReplaceAll(s, "(", "_")
	s = strings.ReplaceAll(s, ")", "_")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, ":", "_")
	s = strings.ReplaceAll(s, "：", "_")
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		p = strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
				return r
			}
			return -1
		}, p)
		if p == "" {
			continue
		}
		if p[0] >= '0' && p[0] <= '9' {
			p = "_" + p
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}

func formatFile(fileName string) {
	fmt.Printf("Formatting %s\n", fileName)
	exec.Command("go", "fmt", fileName).Run()
}

func renderModule(fileName string, data TemplateData) error {
	tmpl := `// Code generated. DO NOT EDIT.
package goshopee

type {{.ModuleName}}Service interface {
{{- range .Methods}}
	// {{.Name}} {{replace .Description "\n" "\n\t// "}}
	// {{.DocURL}}
	{{.Name}}(sid uint64, {{if .RequestType}}{{if .IsGet}}opt{{else}}req{{end}} {{.RequestType}}, {{end}}tok string) (*{{.ResponseType}}, error)
{{- end}}
}

type {{.ModuleName}}ServiceOp struct {
	client *Client
}

{{range .Methods}}
func (s *{{$.ModuleName}}ServiceOp) {{.Name}}(sid uint64, {{if .RequestType}}{{if .IsGet}}opt{{else}}req{{end}} {{.RequestType}}, {{end}}tok string) (*{{.ResponseType}}, error) {
	path := "/{{.Path}}"
	resp := new({{.ResponseType}})
	{{if eq .APIType "Shop" -}}
	err := s.client.WithShop(sid, tok).{{if .IsGet}}Get{{else}}Post{{end}}(path, {{if .IsGet}}resp, {{if .RequestType}}opt{{else}}nil{{end}}{{else}}{{if .RequestType}}req{{else}}nil{{end}}, resp{{end}})
	{{- else -}}
	err := s.client.WithMerchant(sid, tok).{{if .IsGet}}Get{{else}}Post{{end}}(path, {{if .IsGet}}resp, {{if .RequestType}}opt{{else}}nil{{end}}{{else}}{{if .RequestType}}req{{else}}nil{{end}}, resp{{end}})
	{{- end}}
	return resp, err
}
{{end}}
`
	t, err := template.New("module").Funcs(template.FuncMap{
		"replace": func(s, old, new string) string {
			return strings.ReplaceAll(s, old, new)
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		return err
	}
	formatFile(fileName)
	return nil
}

func renderTest(fileName string, data TemplateData) error {
	tmpl := `// Code generated. DO NOT EDIT.
package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

{{range .Methods}}
func Test_{{$.ModuleName}}_{{.Name}}(t *testing.T) {
	setup()
	defer teardown()

	fixture := "{{.FixtureName}}_resp.json"
	responder, err := httpmock.NewJsonResponder(200, loadFixtureInterface(fixture))
	if err != nil {
		t.Skipf("Skipping {{.Name}} due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("{{.Method}}", fmt.Sprintf("%s/api/v2/{{.Path}}", app.APIURL), responder)

	{{if .RequestType -}}
	var req {{.RequestType}}
	{{- end}}
	res, err := client.{{$.ModuleName}}.{{.Name}}({{if eq .APIType "Shop"}}shopID{{else}}merchantID{{end}}, {{if .RequestType}}req, {{end}}accessToken)
	if err != nil {
		t.Logf("{{$.ModuleName}}.{{.Name}} returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("{{$.ModuleName}}.{{.Name}} response: %#v", res)
}
{{end}}
`
	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		return err
	}
	formatFile(fileName)
	return nil
}

func updateGoshopee(services []ServiceInfo) {
	content, err := ioutil.ReadFile("goshopee.go")
	if err != nil {
		panic(err)
	}

	var serviceFields bytes.Buffer
	for _, s := range services {
		serviceFields.WriteString(fmt.Sprintf("\t%s %s\n", s.Name, s.Interface))
	}

	startMarker := "// BEGIN GENERATED SERVICES"
	endMarker := "// END GENERATED SERVICES"
	newContent := replaceBetween(string(content), startMarker, endMarker, serviceFields.String()+"\t")

	var serviceInit bytes.Buffer
	for _, s := range services {
		serviceInit.WriteString(fmt.Sprintf("\tc.%s = &%s{client: c}\n", s.Name, s.Impl))
	}

	startMarker = "// BEGIN GENERATED SERVICES INIT"
	endMarker = "// END GENERATED SERVICES INIT"
	newContent = replaceBetween(newContent, startMarker, endMarker, serviceInit.String()+"\t")

	ioutil.WriteFile("goshopee.go", []byte(newContent), 0644)
	formatFile("goshopee.go")
}

func replaceBetween(content, startMarker, endMarker, replacement string) string {
	startIdx := strings.Index(content, startMarker)
	if startIdx == -1 {
		return content
	}
	startIdx += len(startMarker) + 1

	endIdx := strings.Index(content, endMarker)
	if endIdx == -1 {
		return content
	}

	return content[:startIdx] + replacement + content[endIdx:]
}

func generateSharedTypesFile(config Config) error {
	data := TemplateData{
		PackageName: "goshopee",
	}

	var structNames []string
	for name := range allGeneratedStructs {
		structNames = append(structNames, name)
	}
	sort.Strings(structNames)

	for _, name := range structNames {
		data.Structs = append(data.Structs, allGeneratedStructs[name])
	}

	for typeName, def := range config.TypeDefinitions {
		constData := ConstantData{
			TypeName: typeName,
			BaseType: def.BaseType,
		}
		var valKeys []string
		for k := range def.Values {
			valKeys = append(valKeys, k)
		}
		sort.Strings(valKeys)
		for _, k := range valKeys {
			constData.Values = append(constData.Values, ConstantValue{Name: def.Values[k], Value: k})
		}
		data.Constants = append(data.Constants, constData)
	}

	tmpl := `// Code generated. DO NOT EDIT.
package {{.PackageName}}

{{range .Structs}}
type {{.Name}} struct {
{{- range .Fields}}
	{{if .Name}}{{.Name}} {{.Type}}{{else}}{{.Type}}{{end}} ` + "`" + `{{if .JSONTag}}json:"{{.JSONTag}}"{{end}}{{if .URLTag}} url:"{{.URLTag}}"{{end}}` + "`" + ` // {{replace .Description "\n" " "}}
{{- end}}
}
{{end}}

{{range $const := .Constants}}
type {{$const.TypeName}} {{$const.BaseType}}

const (
{{- range $val := .Values}}
	{{$val.Name}} {{$const.TypeName}} = {{if eq $const.BaseType "string"}}"{{$val.Value}}"{{else}}{{$val.Value}}{{end}}
{{- end}}
)
{{end}}
`
	t, err := template.New("types").Funcs(template.FuncMap{
		"replace": func(s, old, new string) string {
			return strings.ReplaceAll(s, old, new)
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}

	f, err := os.Create("types.go")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		return err
	}
	formatFile("types.go")
	return nil
}
