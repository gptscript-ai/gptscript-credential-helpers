package credentials

// Helper is the interface a credentials store helper must implement.
type Helper interface {
	// Add appends credentials to the store.
	Add(*Credentials) error
	// Delete removes credentials from the store.
	Delete(serverURL string) error
	// Get retrieves credentials from the store.
	// It returns username and secret as strings.
	Get(serverURL string) (string, string, error)
	// List returns the serverURLs of keys and their
	// associated usernames from the OS store as a
	// list of strings
	List() ([]string, []string, error)
}
