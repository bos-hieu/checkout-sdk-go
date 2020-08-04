package common

const (
	// LiveSecretKeyRegex ...
	LiveSecretKeyRegex = `^sk_?(\w{8})-(\w{4})-(\w{4})-(\w{4})-(\w{12})$`
	// LivePublicKeyRegex ...
	LivePublicKeyRegex = `^pk_?(\w{8})-(\w{4})-(\w{4})-(\w{4})-(\w{12})$`
	// SandboxSecretKeyRegex ...
	SandboxSecretKeyRegex = `^sk_test_?(\w{8})-(\w{4})-(\w{4})-(\w{4})-(\w{12})$`
	// SandboxPublicKeyRegex ...
	SandboxPublicKeyRegex = `^pk_test_?(\w{8})-(\w{4})-(\w{4})-(\w{4})-(\w{12})$`
	// CardTokenRegex ...
	CardTokenRegex = `^(tok)_(\w{26})$`
	// SourceIDRegex ...
	SourceIDRegex = `^(src)_(\w{26})$`
)
