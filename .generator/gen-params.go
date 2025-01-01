package main

// input: openapi v3 file
// output: param go types

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"
)

var (
	openapiV3Path string
	outputPath    string

	//go:embed params.tmpl
	paramTemp embed.FS
)

func main() {
	flag.StringVar(&openapiV3Path, "i", "", "openapi v3 path")
	flag.StringVar(&outputPath, "o", "./", "output path")
	flag.Parse()

	if openapiV3Path == "" {
		panic("openapi v3 path is required")
	}

	paramTmpl, err := paramTemp.ReadFile("params.tmpl")
	if err != nil {
		panic(fmt.Errorf("error reading params template: %w", err))
	}

	spec, err := util.LoadSwagger(openapiV3Path)
	if err != nil {
		panic(fmt.Errorf("error loading swagger: %w", err))
	}
	codegen.Generate(spec, codegen.Configuration{
		Generate: codegen.GenerateOptions{
			Models: true,
		},
	})
	ops, err := codegen.OperationDefinitions(spec, true)
	if err != nil {
		panic(fmt.Errorf("error generating operation definitions: %w", err))
	}
	fmt.Println(ops)
	t := template.Must(template.New("params").Parse(string(paramTmpl)))

	var outputs []string
	for _, op := range ops {
		if op.QueryParams != nil {
			var buf bytes.Buffer
			t.ExecuteTemplate(&buf, "param", op)
			outputs = append(outputs, buf.String())
		}
	}
	err = os.WriteFile(outputPath, []byte(strings.Join(outputs, "\n")), 0644)
	if err != nil {
		panic(fmt.Errorf("error writing output file: %w", err))
	}
}
