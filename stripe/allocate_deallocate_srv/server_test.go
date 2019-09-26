package main

import (
	"reflect"
	"testing"
)

func TestGetFreeResourceID(t *testing.T) {
	tests := []struct {
		input map[int]bool
		expID int
	}{
		{map[int]bool{}, 0},
		{map[int]bool{1: true}, 0},
		{map[int]bool{0: true, 2: true}, 1},
	}
	for _, test := range tests {
		id := getFreeResourceID(test.input)
		if id != test.expID {
			t.Errorf("getFreeResourceID() returned %d but expected %d", id, test.expID)
		}
	}

}

func TestCallOrder(t *testing.T) {
	tests := []struct {
		services       Services
		f              func(Services)
		expectedResult Services
	}{
		{
			services: make(Services),
			f: func(s Services) {
				s.Allocate("seal")
				s.Allocate("whale")
				s.Allocate("whale")
				s.Allocate("rhino")
			},
			expectedResult: Services{
				"seal": Resources{
					0: true,
				},
				"whale": Resources{
					0: true,
					1: true,
				},
				"rhino": Resources{
					0: true,
				},
			},
		},
		{
			services: make(Services),
			f: func(s Services) {
				s.Allocate("seal")
				s.Allocate("whale")
				s.Allocate("whale")
				s.Deallocate("whale0")
				s.Deallocate("whale1")
				s.Allocate("rhino")
			},
			expectedResult: Services{
				"seal": Resources{
					0: true,
				},
				"rhino": Resources{
					0: true,
				},
			},
		},
	}

	for _, test := range tests {
		test.f(test.services)
		if eq := reflect.DeepEqual(test.services, test.expectedResult); !eq {
			t.Errorf("returned %v but expected %v", test.services, test.expectedResult)
		}
	}
}
