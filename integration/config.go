package integration

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"testing"
)

var Cfg = Config{
	CommandService: struct {
		ServiceEndpoint
	}{
		ServiceEndpoint: ServiceEndpoint{
			protocol: "http",
			host:     "localhost",
			port:     "7080",
			actuator: "/health",
		},
	},
	ConsumerService: struct {
		ServiceEndpoint
	}{
		ServiceEndpoint: ServiceEndpoint{
			protocol: "http",
			host:     "localhost",
			port:     "7081",
			actuator: "/health",
		},
	},
	QueryService: struct {
		ServiceEndpoint
	}{
		ServiceEndpoint: ServiceEndpoint{
			protocol: "http",
			host:     "localhost",
			port:     "7082",
			actuator: "/health",
		},
	},
}

type ServiceEndpoint struct {
	protocol string
	host     string
	port     string
	actuator string
}

func (s ServiceEndpoint) endpointPath() string {
	return fmt.Sprintf("%s://%s:%s", s.protocol, s.host, s.port)
}

func (s ServiceEndpoint) healthPath() string {
	return fmt.Sprintf("%s%s", s.endpointPath(), s.actuator)
}

func (s ServiceEndpoint) health() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", s.healthPath(), nil)
	if err != nil {
		return nil, err
	}
	correlationId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	req.Header.Add("correlationId", correlationId.String())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status recieved: %d from endpoint %s", res.StatusCode, s.healthPath())
	}
	return res, nil
}

type Config struct {
	CommandService struct {
		ServiceEndpoint
	}
	ConsumerService struct {
		ServiceEndpoint
	}
	QueryService struct {
		ServiceEndpoint
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
