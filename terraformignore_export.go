package slug

// ParseIgnoreFile exports unexported parseIgnoreFile
func ParseIgnoreFile(rootPath string) []rule {
	return parseIgnoreFile(rootPath)
}

// MatchIgnoreRule exports unexported matchIgnoreRule
func MatchIgnoreRule(path string, rules []rule) bool {
	return matchIgnoreRule(path, rules)
}
