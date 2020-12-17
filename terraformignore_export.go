package slug

import "io"

// Rule exports unexported rule
type Rule struct {
	rule
}

func (r *Rule) Match(path string) (bool, error) {
	return r.rule.match(path)
}

// ParseIgnoreFile exports unexported parseIgnoreFile
func ParseIgnoreFile(rootPath string) []rule {
	return parseIgnoreFile(rootPath)
}

// MatchIgnoreRule exports unexported matchIgnoreRule
func MatchIgnoreRule(path string, rules []rule) bool {
	return matchIgnoreRule(path, rules)
}

// ReadRules exports unexported readRules
func ReadRules(input io.Reader) []rule {
	return readRules(input)
}

// DefaultExclusions exports unexported defaultExclusions
var DefaultExclusions = defaultExclusions
