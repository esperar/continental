package persist

import "continential/app/models"

type DBLayer interface {
	GetAllConnections() ([]models.Content, error)
}
