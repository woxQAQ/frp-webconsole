package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"
	"gopkg.in/yaml.v3"
)

var (
	swaggerPath string
	outputPath  string
	configPath  string
)

func main() {
	flag.StringVar(&swaggerPath, "f", "", "openapi file path")
	flag.StringVar(&outputPath, "o", "", "output file path")
	flag.StringVar(&configPath, "c", "", "config file path")
	flag.Parse()
	buf, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	var config struct {
		codegen.Configuration `yaml:",inline"`
		CustomServer          []string            `yaml:"custom-server,omitempty"`
		CustomServerTemplate  map[string][]string `yaml:"custom-server-template,omitempty"`
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		panic(err)
	}
	swagger, err := util.LoadSwagger(swaggerPath)
	if err != nil {
		panic(err)
	}
	res, err := codegen.Generate(swagger, config.Configuration)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outputPath, []byte(res), 0644)
	if err != nil {
		panic(err)
	}
	codegen.TemplateFunctions["opts"] = func() codegen.Configuration {
		return config.Configuration
	}
	ops, err := codegen.OperationDefinitions(swagger, config.OutputOptions.InitialismOverrides)
	if err != nil {
		panic(fmt.Errorf("error creating operation definitions: %w", err))
	}
	t := template.New("oapi-codegen").Funcs(codegen.TemplateFunctions)
	var customServerOut []string
	if len(config.CustomServer) > 0 {
		for _, customServer := range config.CustomServer {
			if len(config.CustomServerTemplate[customServer]) == 0 {
				panic(fmt.Errorf("custom server template %q is not defined", customServer))
			}
			res, err := GenerateCustomServer(t, ops, config.CustomServerTemplate[customServer])
			if err != nil {
				panic(fmt.Errorf("error generating Go handlers for Paths: %w", err))
			}
			customServerOut = append(customServerOut, res)
		}
	}
}

func GenerateCustomServer(t *template.Template,
	operations []codegen.OperationDefinition,
	customServer []string,
) (string, error) {
	return codegen.GenerateTemplates(customServer, t, operations)
}
