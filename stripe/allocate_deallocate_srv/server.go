package main

import (
	"regexp"
	"strconv"
)

// Some letters followed by some digits.
var machineRegex = regexp.MustCompile("([a-zA-Z]+)([0-9]+)")

// Services represents a map of service name to service esources
type Services map[string]Resources

// Resources represents a map of resource ID to status
type Resources map[int]bool

/*
 * get the ID of the next zero-value boolean entry
 */
func getFreeResourceID(resources map[int]bool) int {
	for i := 0; i < len(resources); i++ {
		if _, allocated := resources[i]; !allocated {
			return i
		}
	}
	return len(resources)
}

/*
 * O(1) insertion
 */
func (s Services) Allocate(serviceName string) string {
	// if service is not yet in our object, add it
	if _, ok := s[serviceName]; !ok {
		s[serviceName] = make(map[int]bool)
	}
	resources := s[serviceName]
	id := getFreeResourceID(resources)
	resources[id] = true
	return serviceName + strconv.Itoa(id)
}

/*
 * O(1) deletion
 */
func (s Services) Deallocate(strToParse string) {
	match := machineRegex.FindAllStringSubmatch(strToParse, -1)
	svc := match[0][1]
	resourceID, _ := strconv.Atoi(match[0][2])
	delete(s[svc], resourceID)
	if len(s[svc]) == 0 {
		delete(s, svc)
	}
}
