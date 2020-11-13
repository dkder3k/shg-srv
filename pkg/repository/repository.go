package repository

import (
	consulapi "github.com/hashicorp/consul/api"
)

type HostsList interface {
}

type Host interface {
}

type ConfigurationsList interface {
}

type Configuration interface {
}

type Repository struct {
	HostsList
	Host
	ConfigurationsList
	Configuration
}

func NewConsulRepository(consul *consulapi.Client) *Repository {
	return &Repository{
		HostsList:          nil,
		Host:               nil,
		ConfigurationsList: nil,
		Configuration:      nil,
	}
}
