package incident

import (
	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

type Store interface {
	// CreateIncident Creates a new incident.
	CreateIncident(incident *Incident) (*Incident, error)

	// GetIncident Gets an incident by ID.
	GetIncident(ID string) (*Incident, error)

	// GetAllIncidents Gets all incidents
	GetAllIncidents() ([]Incident, error)
}

var _ Store = &incidentStore{}

type incidentStore struct {
	pluginAPI *pluginapi.Client
	helpers   plugin.Helpers
}

// NewStore creates a new store for incident service.
func NewStore(pluginAPI *pluginapi.Client, helpers plugin.Helpers) *incidentStore {
	newStore := &incidentStore{
		pluginAPI: pluginAPI,
		helpers:   helpers,
	}
	return newStore
}

// CreateIncident Creates a new incident.
func (s *incidentStore) CreateIncident(incident *Incident) (*Incident, error) {
	return nil, errors.New("not implemented")
}

// GetIncident Gets an incident by ID.
func (s *incidentStore) GetIncident(ID string) (*Incident, error) {
	return nil, errors.New("not implemented")
}

// GetAllIncidents Gets all incidents
func (s *incidentStore) GetAllIncidents() ([]Incident, error) {
	return nil, errors.New("not implemented")
}