// This file was auto-generated by Fern from our API Definition.

package group

import (
	uuid "github.com/gofrs/uuid/v5"
	rivetgo "github.com/rivet-gg/rivet-go"
	upload "github.com/rivet-gg/rivet-go/upload"
)

type CreateRequest struct {
	// Represent a resource's readable display name.
	DisplayName string `json:"display_name"`
}

type CreateResponse struct {
	GroupId uuid.UUID `json:"group_id"`
}

type GetBansResponse struct {
	// A list of banned group members.
	BannedIdentities []*BannedIdentity `json:"banned_identities,omitempty"`
	// The pagination anchor.
	Anchor *string                `json:"anchor,omitempty"`
	Watch  *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type GetJoinRequestsResponse struct {
	// A list of group join requests.
	JoinRequests []*JoinRequest `json:"join_requests,omitempty"`
	// The pagination anchor.
	Anchor *string                `json:"anchor,omitempty"`
	Watch  *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type GetMembersResponse struct {
	// A list of group members.
	Members []*Member `json:"members,omitempty"`
	// The pagination anchor.
	Anchor *string                `json:"anchor,omitempty"`
	Watch  *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type GetProfileResponse struct {
	Group *Profile               `json:"group,omitempty"`
	Watch *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type GetSummaryResponse struct {
	Group *Summary `json:"group,omitempty"`
}

type ListSuggestedResponse struct {
	// A list of group summaries.
	Groups []*Summary             `json:"groups,omitempty"`
	Watch  *rivetgo.WatchResponse `json:"watch,omitempty"`
}

type PrepareAvatarUploadRequest struct {
	// The path/filename of the group avatar.
	Path string `json:"path"`
	// The MIME type of the group avatar.
	Mime *string `json:"mime,omitempty"`
	// Unsigned 64 bit integer.
	ContentLength int64 `json:"content_length"`
}

type PrepareAvatarUploadResponse struct {
	UploadId         uuid.UUID                `json:"upload_id"`
	PresignedRequest *upload.PresignedRequest `json:"presigned_request,omitempty"`
}

type SearchResponse struct {
	// A list of group handles.
	Groups []*Handle `json:"groups,omitempty"`
	Anchor *string   `json:"anchor,omitempty"`
}

type TransferOwnershipRequest struct {
	// Identity to transfer the group to.
	// Must be a member of the group.
	NewOwnerIdentityId string `json:"new_owner_identity_id"`
}

type UpdateProfileRequest struct {
	// Represent a resource's readable display name.
	DisplayName *string `json:"display_name,omitempty"`
	// Detailed information about a profile.
	Bio       *string    `json:"bio,omitempty"`
	Publicity *Publicity `json:"publicity,omitempty"`
}

type ValidateProfileRequest struct {
	// Represent a resource's readable display name.
	DisplayName *string `json:"display_name,omitempty"`
	// Detailed information about a profile.
	Bio       *string    `json:"bio,omitempty"`
	Publicity *Publicity `json:"publicity,omitempty"`
}

type ValidateProfileResponse struct {
	// A list of validation errors.
	Errors []*rivetgo.ValidationError `json:"errors,omitempty"`
}
