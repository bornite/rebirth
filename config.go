package rebirth

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"golang.org/x/xerrors"
)

type Config struct {
	Host  *Host  `yaml:"host,omitempty"`
	Build *Build `yaml:"build,omitempty"`
	Run   *Run   `yaml:"run,omitempty"`
}

type Host struct {
	Docker string `yaml:"docker,omitempty"`
}

type Build struct {
	Env map[string]string `yaml:"env,omitempty"`
}

type Run struct {
	Env map[string]string `yaml:"env,omitempty"`
}

func LoadConfig(confPath string) (*Config, error) {
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, xerrors.Errorf("failed to read config file from %s: %w", confPath, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, xerrors.New(yaml.FormatError(err, true, true))
	}
	return &cfg, nil
}