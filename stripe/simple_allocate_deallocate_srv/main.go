package simple_allocate_deallocate_srv

import (
	"log"
	"regexp"
	"strconv"
)

// System represents our system to
// allocate and deallocate servers
type System struct {
	// we will map service name to
	// service map, where the service
	// map acts as a set of server ID
	services map[string]map[int]bool
}

var (
	sidFmt = regexp.MustCompile("([a-zA-Z]+)([0-9]+)")
)

// NewSystem is the constructor for the system struct
func NewSystem() *System {
	return &System{
		services: make(map[string]map[int]bool),
	}
}

// Deallocate deallocaes a server on the system - O(1)
func (s *System) Deallocate(sid string) {
	// FindAllStringSubmatch returns []string{sid, service, id}
	// on the first index
	details := sidFmt.FindAllStringSubmatch(sid, -1)
	// validate input
	if len(details) < 1 || len(details[0]) != 3 {
		log.Fatalf("invalid server id format, expected <name><id>, got %s", sid)
	}
	// parse service name and server id
	service := details[0][1]
	id, err := strconv.Atoi(details[0][2])
	if err != nil {
		log.Fatalf("non numerical id: %s", details[0][2])
	}
	// perform the deallocation
	delete(s.services[service], id)
	if len(s.services[service]) == 0 {
		delete(s.services, service)
	}
}

// Allocate allocates a server on the system - O(1)
func (s *System) Allocate(sname string) int {
	id := 0
	serviceSet, exists := s.services[sname]
	if exists {
		for {
			if _, ok := serviceSet[id]; !ok {
				serviceSet[id] = true
				break
			}
			id++
		}
	} else {
		s.services[sname] = make(map[int]bool)
		s.services[sname][id] = true
	}
	return id
}
