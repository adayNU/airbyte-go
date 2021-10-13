package types

type OAuth2Specification struct {
	RootObject                []string
	OAuthFlowInitParameters   [][]string
	OAuthFlowOutputParameters [][]string
}

type ConnectorSpecification struct {
	DocumentationURL              string
	ChangelogURL                  string
	ConnectionSpecification       JSONData
	SupportsIncremental           bool
	SupportsNormalization         bool
	SupportsDBT                   bool
	SupportedDestinationSyncModes DestinationSyncMode
}
