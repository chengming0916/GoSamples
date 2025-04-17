package config

import (
	"GoSamples/ServiceSample/internal/pkg/config/types"
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var glbEnvs map[string]string

func init() {
	glbEnvs = make(map[string]string)
	envs := os.Environ()
	for _, env := range envs {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) != 2 {
			continue
		}
		glbEnvs[pair[0]] = pair[1]
	}
}

type Values struct {
	Envs map[string]string
}

func GetValues() *Values {
	return &Values{
		Envs: glbEnvs,
	}
}

func LoadConfig(path string) (*types.ServiceConfig, error) {
	svrCfg := &types.ServiceConfig{}
	if err := LoadConfigFromFile(path, svrCfg); err != nil {
		return nil, err
	}

	svrCfg.Complete()

	return svrCfg, nil
}

func LoadConfigFromFile(path string, c any) error {
	content, err := LoadFileWithTemplate(path, GetValues())
	if err != nil {
		return err
	}
	return LoadConfigure(content, c)
}

func LoadFileWithTemplate(path string, values *Values) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return RenderWithTemplate(b, values)
}

func RenderWithTemplate(in []byte, values *Values) ([]byte, error) {
	tmpl, err := template.New("sample").Funcs(template.FuncMap{
		"parseNumberRange":     parseNumberRange,
		"parseNumberRangePair": parseNumberRangePair,
	}).Parse(string(in))
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("")
	if err := tmpl.Execute(buffer, values); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func LoadConfigure(b []byte, c any) error {

	var cfg types.ServiceConfig
	if err := json.Unmarshal(b, &cfg); err != nil {

	}

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return err
	}

	return nil
}
