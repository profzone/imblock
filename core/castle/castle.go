package castle

import (
	"reflect"
	"fmt"
	"github.com/profzone/imblock/core"
)

type Castle struct {
	Name     string
	serviceFunc []ServiceConstructor
	services map[reflect.Type]Service
}

var GeneralCastle *Castle

func NewStack(name string) *Castle {
	if GeneralCastle != nil {
		return GeneralCastle
	}
	s := &Castle{
		Name:     name,
		services: make(map[reflect.Type]Service),
	}

	GeneralCastle = s

	return s
}

//func NewStackWithTemplate(presetName string) *Castle {
//	if GeneralCastle != nil {
//		return GeneralCastle
//	}
//	preset := presetStack[presetName]
//	s := NewStack(presetName)
//
//	for _, f := range preset {
//		s.RegisterService(f)
//	}
//
//	return s
//}

func (s *Castle) RegisterService(c ServiceConstructor) {
	s.serviceFunc = append(s.serviceFunc, c)
}

func (s *Castle) GetService(serviceType reflect.Type) Service {
	return s.services[serviceType]
}

func (s *Castle) Start() error {
	for _, constructor := range s.serviceFunc {
		service := constructor()
		t := reflect.TypeOf(service)
		if t.Kind() == reflect.Ptr {
			v := reflect.ValueOf(service)
			if v.IsNil() {
				panic(fmt.Sprintf("service must not be nil, type: %v", t))
			}
		}
		focusMessage := service.Protocols()
		for _, m := range focusMessage {
			core.GetProtocolManager().RegisterProtocol(m)
		}
		s.services[t] = service
		go service.Start()
	}
	return nil
}

func (s *Castle) Stop() error {
	for _, service := range s.services {
		service.Stop()
	}
	return nil
}
