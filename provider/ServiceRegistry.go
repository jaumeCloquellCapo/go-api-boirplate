package provider

import (
	"ApiRest/config"
	"fmt"
	"reflect"
)

type Service interface {
	// Start spawns the main process done by the service.
	Start(config.Config)
	// Stop terminates all processes belonging to the service,
	// blocking until they are all terminated.
	Stop() error
	// Returns error if the service is not considered healthy.
	Status() error
}

// ServiceRegistry provides a useful pattern for managing services.
// It allows for ease of dependency management and ensures services
// dependent on others use the same references in memory.
type ServiceRegistry struct {
	services     map[reflect.Type]Service // map of types to services.
	serviceTypes []reflect.Type           // keep an ordered slice of registered service types.
}

// NewServiceRegistry starts a registry instance for convenience
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[reflect.Type]Service),
	}
}


// RegisterService appends a service constructor function to the service
// registry.
func (s *ServiceRegistry) RegisterService(service Service) error {
	kind := reflect.TypeOf(service)
	if _, exists := s.services[kind]; exists {
		return fmt.Errorf("service already exists: %v", kind)
	}
	s.services[kind] = service
	s.serviceTypes = append(s.serviceTypes, kind)
	return nil
}

func (s *ServiceRegistry) StartAll() {
	fmt.Printf("Starting %d services: %v", len(s.serviceTypes), s.serviceTypes)
	for _, kind := range s.serviceTypes {
		fmt.Printf("Starting service type %v", kind)
		go s.services[kind].Start()
	}
}

func (s *ServiceRegistry) StopAll() {
	for i := len(s.serviceTypes) - 1; i >= 0; i-- {
		kind := s.serviceTypes[i]
		service := s.services[kind]
		if err := service.Stop(); err != nil {
			fmt.Errorf("Could not stop the following service: %v, %v", kind, err)
		}
	}
}

func (s *ServiceRegistry) FetchService(service interface{}) error {
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", service)
	}
	element := reflect.ValueOf(service).Elem()
	if running, ok := s.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return fmt.Errorf("unknown service: %T", service)
}