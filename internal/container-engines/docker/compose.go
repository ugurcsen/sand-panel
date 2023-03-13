package docker

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
)

// composeYamlBuilder builds a docker-compose.yml file from a collection
type service struct {
	Image    string            `yaml:"image,omitempty"`
	Port     string            `yaml:"port,omitempty"`
	Labels   map[string]string `yaml:"labels,omitempty"`
	Networks []string          `yaml:"networks,omitempty"`
	Runtime  string            `yaml:"runtime,omitempty"`
	Volumes  []string          `yaml:"volumes,omitempty"`
}

// network is a docker network for a collection
type network struct {
	Name     string `yaml:"name,omitempty"`
	External bool   `yaml:"external,omitempty"`
	Driver   string `yaml:"driver,omitempty"`
}

// compose is a docker-compose.yml file
type compose struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]service `yaml:"services,omitempty"`
	Networks map[string]network `yaml:"networks,omitempty"`
}

// composeYamlBuilder builds a docker-compose.yml file from a collection
func composeYamlBuilder(c *domain.Collection, f io.Writer) error {
	comp := compose{}

	if len(c.Services) > 0 {
		comp.Services = make(map[string]service)
	}

	comp.Networks = make(map[string]network)

	comp.Version = "3.7"

	for _, s := range c.Services {
		var hosts []string
		for _, h := range s.Hosts {
			hosts = append(hosts, "Host(`"+h+"`)")
		}
		serv := service{
			Image:   s.Image,
			Runtime: "runsc",
			Networks: []string{
				"default",
			},
			Labels: map[string]string{
				"sand.panel.collection": c.Id,
				"sand.panel.service":    s.Id,
				"sand.panel.user":       c.UserId,
			},
		}

		for _, v := range s.Volumes {
			serv.Volumes = append(serv.Volumes, v.From+":"+v.To)
		}

		if len(hosts) > 0 {
			serv.Networks = append(serv.Networks, "traefik")
			serv.Labels["traefik.enable"] = "true"
			serv.Labels["traefik.http.routers."+s.Id+".entrypoints"] = "web, websecure"
			serv.Labels["traefik.http.routers."+s.Id+".rule"] = strings.Join(hosts, " || ")
		}

		comp.Services[s.Id] = serv
	}

	comp.Networks["default"] = network{
		Driver: "bridge",
	}

	comp.Networks["traefik"] = network{
		Name:     "traefiknet",
		External: true,
	}

	if err := yaml.NewEncoder(f).Encode(comp); err != nil {
		return err
	}
	return nil
}
