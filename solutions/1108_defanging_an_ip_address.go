func defangIPaddr(address string) string {
	ip := strings.ReplaceAll(address, ".", "[.]")
	return ip
}