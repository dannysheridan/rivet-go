// This file was auto-generated by Fern from our API Definition.

package games

import (
	uuid "github.com/gofrs/uuid/v5"
	rivetgo "github.com/rivet-gg/rivet-go"
	version "github.com/rivet-gg/rivet-go/cloud/version"
)

type CreateGameVersionRequest struct {
	// Represent a resource's readable display name.
	DisplayName string          `json:"display_name"`
	Config      *version.Config `json:"config,omitempty"`
}

type CreateGameVersionResponse struct {
	VersionId uuid.UUID `json:"version_id"`
}

type GetGameVersionByIdResponse struct {
	Version *version.Full `json:"version,omitempty"`
}

type ValidateGameVersionRequest struct {
	// Represent a resource's readable display name.
	DisplayName string          `json:"display_name"`
	Config      *version.Config `json:"config,omitempty"`
}

type ValidateGameVersionResponse struct {
	// A list of validation errors.
	Errors []*rivetgo.ValidationError `json:"errors,omitempty"`
}
