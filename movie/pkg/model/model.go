package model

import "github.com/vaidik-bajpai/green-screen/metadata/pkg/model"

// MovieDetails includes movie metadata its aggregated
// rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating"`
	Metadata model.Metadata `json:"metadata"`
}
