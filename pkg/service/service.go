package service

import "github.com/dkder3k/shg-srv/pkg/repository"

type HostsList interface {
}

type Host interface {
}

type ConfigurationsList interface {
}

type Configuration interface {
}

type Service struct {
	HostsList
	Host
	ConfigurationsList
	Configuration
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		HostsList:          nil,
		Host:               nil,
		ConfigurationsList: nil,
		Configuration:      nil,
	}
}
