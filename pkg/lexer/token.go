package lexer

// TokenType represents the different types of tokens in the DSL.
type TokenType int

const (
	TokenEOF TokenType = iota
	TokenDeployApp
	TokenNamespace
	TokenReplicas
	TokenImage
	TokenEnv
	TokenPorts
	TokenResources
	TokenStorage
	TokenLBrace // {
	TokenRBrace // }
	TokenColon  // :
	TokenString // "value" or value
	TokenNumber // 123
	TokenIdentifier
	TokenMemory
	TokenCPU
	TokenVolume
	TokenSize
	TokenLimits
	TokenRequests
	TokenArgs
	TokenService // service
	TokenPort
	TokenTargetPort
	TokenTypeString
	TokenSeparator // ---
)
