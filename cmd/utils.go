package cmd

// SafeDeref is used to safely derefence a pointer where it might point to null
// In the case where it points to null it will return the default value for the type
// rather than a runtime error / segfault
func SafeDeref[T any](pointer *T) T {
	if pointer == nil {
		var defaultValue T
		return defaultValue
	}

	return *pointer
}