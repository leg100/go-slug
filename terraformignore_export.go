package slug

// ParseIgnoreFile exports unexported parseIgnoreFile
func ParseIgnoreFile(rootPath string) []rule {
	return parseIgnoreFile(rootPath)
}
