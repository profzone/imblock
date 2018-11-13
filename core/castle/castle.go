package castle

import (
	"reflect"
	"github.com/profzone/imblock/core"
	"github.com/sirupsen/logrus"
)

type Castle struct {
	Name        string
	serviceFunc []ServiceConstructor
	services    map[reflect.Type]Service
}

var GeneralCastle *Castle

func NewStack(name string) *Castle {
	if GeneralCastle != nil {
		return GeneralCastle
	}
	s := &Castle{
		Name:       name,
		services:   make(map[reflect.Type]Service),
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

func (s *Castle) Start() error {
	for _, constructor := range s.serviceFunc {

		service := constructor()
		if service == nil {
			logrus.Panic("service must not be nil")
		}

		t := reflect.TypeOf(service)

		// register protocols
		focusProtocols := service.Protocols()
		for _, p := range focusProtocols {
			core.GetProtocolManager().RegisterProtocol(p)
		}

		// register messages
		focusMessages := service.Messages()
		for _, m := range focusMessages {
			core.GetMessageManager().Subscribe(m.Topic, m.Runner)
		}

		s.services[t] = service
		go service.Start()
	}
	return nil
}

func (s *Castle) Stop() error {
	for _, service := range s.services {
		if err := service.Stop(); err != nil {
			return err
		}
	}
	return nil
}
