package domain

type ServiceEnv int

const (
	ServiceEnvStaging ServiceEnv = iota
	ServiceEnvProduction
)

func (se ServiceEnv) String() string {
	switch se {
	case ServiceEnvStaging:
		return "staging"
	case ServiceEnvProduction:
		return "production"
	}
	return "unknown"
}
