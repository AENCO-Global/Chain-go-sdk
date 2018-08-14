package namespace

import "github.com/slackve/nem2-sdk-go/sdk/model"

// The namespace id structure describes namespace id
// since 1.0
type NamespaceId struct {
	// Namespace id
	id model.Id
	// Namespace full name
	FullName string
}
