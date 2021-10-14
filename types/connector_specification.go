package types

type OAuth2Specification struct {
	RootObject                []string
	OAuthFlowInitParameters   [][]string
	OAuthFlowOutputParameters [][]string
}

type AuthType int

const (
	OAuth2Dot0 AuthType = iota
)

type AuthSpecification struct {
	Type                AuthType
	OAuth2Specification *OAuth2Specification
}

type ConnectorSpecification struct {
	DocumentationURL              string
	ChangelogURL                  string
	ConnectionSpecification       JSONData
	SupportsIncremental           bool
	SupportsNormalization         bool
	SupportsDBT                   bool
	SupportedDestinationSyncModes []DestinationSyncMode
	AuthSpecification             *AuthSpecification
}
