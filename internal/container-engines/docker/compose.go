package docker

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"gopkg.in/yaml.v3"
	"os"
)

type service struct {
	Image    string            `yaml:"image,omitempty"`
	Port     string            `yaml:"port,omitempty"`
	Labels   map[string]string `yaml:"labels,omitempty"`
	Networks []string          `yaml:"networks,omitempty"`
	Runtime  string            `yaml:"runtime,omitempty"`
}

type network struct {
	Name     string `yaml:"name,omitempty"`
	External bool   `yaml:"external,omitempty"`
}

type compose struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]service `yaml:"services,omitempty"`
	Networks map[string]network `yaml:"networks,omitempty"`
}

func composeYamlBuilder(c *domain.Collection, f *os.File) error {
	comp := compose{}

	if len(c.Services) > 0 {
		comp.Services = make(map[string]service)
	}

	comp.Networks = make(map[string]network)

	comp.Version = "3.7"

	for _, s := range c.Services {
		comp.Services[s.Id] = service{
			Image:   s.Image,
			Runtime: "runsc",
			Networks: []string{
				"trafeik",
			},
			Labels: map[string]string{
				"traefik.enable":                         "true",
				"traefik.http.routers." + s.Id + ".rule": "Host(`" + s.Host + "`)",
			},
		}
	}

	comp.Networks["trafeik"] = network{
		Name:     "trafeik_default",
		External: true,
	}

	if err := yaml.NewEncoder(f).Encode(comp); err != nil {
		return err
	}
	return nil
}
