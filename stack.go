package main

type Stack struct {
	Version    int
	Deployment Deployment
}

type Deployment struct {
	Manifest         Manifest
	SecretsProviders SecretsProviders
	Resources        []Resource
}

type Manifest struct {
	Time    string
	Magic   string
	Version string
}

type SecretsProviders struct {
	ProviderType string
	State        map[string]string
}

type Resource struct {
	Urn                     string
	Outputs                 map[string]interface{}
	Custom                  bool
	Id                      string
	AdditionalSecretOutputs []string
	Type                    string
	PropertyDependencies    map[string]interface{}
	External                bool
	Parent                  string
	Inputs                  map[string]interface{}
	Provider                string
	Modified                string
	Dependencies            []string
	Created                 string
}
