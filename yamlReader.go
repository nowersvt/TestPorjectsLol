package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
    "strings"
)

type EnvironmentDockerCompose struct {
    User     string `yaml:"POSTGRES_USER"`
    Password string `yaml:"POSTGRES_PASSWORD"`
    DbName   string `yaml:"POSTGRES_DB"`
}

type ServiceInDockerCompose struct {
    ContainerName string                   `yaml:"container_name"`
    Image         string                   `yaml:"image"`
    Ports         []string                 `yaml:"ports"`
    Environment   EnvironmentDockerCompose `yaml:"environment"`
}

type ComposeFile struct {
    Postgres ServiceInDockerCompose `yaml:"services"`
}

func LoadConfig(filename string) (ComposeFile, error) {
    _, err := os.Stat(filename)
	if err != nil {
		return ComposeFile{}, fmt.Errorf("Error finding file %s: %w", filename, err)
	}

    data, err := os.ReadFile(filename)
    if err != nil {
        return ComposeFile{}, fmt.Errorf("Error parsing YAML: %w", err)
    }

    var config ComposeFile
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return ComposeFile{}, fmt.Errorf("Error parsing YAML: %w", err)
    }

    fmt.Printf("Parsed YAML: %+v\n", config.Postgres.Ports)

    return config, nil
}

func extractPort(ports []string) (int, error) {
    if len(ports) == 0 {
        return 0, fmt.Errorf("no ports defined")
    }

    fmt.Println("Ports: ", ports)

    parts := strings.Split(ports[0], ":")
    if len(parts) != 2 {
        return 0, fmt.Errorf("invalid port format: %s", ports[0])
    }

    var port int
    _, err := fmt.Sscanf(parts[1], "%d", &port)
    if err != nil {
        return 0, fmt.Errorf("failed to parse port: %w", err)
    }

    return port, nil
}