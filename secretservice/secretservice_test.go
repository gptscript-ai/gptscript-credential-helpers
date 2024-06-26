//go:build linux && cgo

package secretservice

import (
	"strings"
	"testing"

	"github.com/gptscript-ai/gptscript-credential-helpers/credentials"
)

func TestSecretServiceHelper(t *testing.T) {
	t.Skip("test requires gnome-keyring but travis CI doesn't have it")

	creds := &credentials.Credentials{
		ServerURL: "https://foobar.example.com:2376/v1",
		Username:  "foobar",
		Secret:    "foobarbaz",
	}

	helper := Secretservice{}

	// Check how many docker credentials we have when starting the test
	oldAuths, err := helper.List()
	if err != nil {
		t.Fatal(err)
	}

	// If any docker credentials with the tests values we are providing, we
	// remove them as they probably come from a previous failed test
	for k, v := range oldAuths {
		if strings.Compare(k, creds.ServerURL) == 0 && strings.Compare(v, creds.Username) == 0 {
			if err := helper.Delete(creds.ServerURL); err != nil {
				t.Fatal(err)
			}
		}
	}

	// Check again how many docker credentials we have when starting the test
	oldAuths, err = helper.List()
	if err != nil {
		t.Fatal(err)
	}

	// Add new credentials
	if err := helper.Add(creds); err != nil {
		t.Fatal(err)
	}

	// Verify that it is inside the secret service store
	username, secret, err := helper.Get(creds.ServerURL)
	if err != nil {
		t.Fatal(err)
	}

	if username != "foobar" {
		t.Fatalf("expected %s, got %s\n", "foobar", username)
	}

	if secret != "foobarbaz" {
		t.Fatalf("expected %s, got %s\n", "foobarbaz", secret)
	}

	// We should have one more credential than before adding
	newAuths, err := helper.List()
	if err != nil || (len(newAuths)-len(oldAuths) != 1) {
		t.Fatal(err)
	}
	oldAuths = newAuths

	// Deleting the credentials associated to current server url should succeed
	if err := helper.Delete(creds.ServerURL); err != nil {
		t.Fatal(err)
	}

	// We should have one less credential than before deleting
	newAuths, err = helper.List()
	if err != nil || (len(oldAuths)-len(newAuths) != 1) {
		t.Fatal(err)
	}
}

func TestMissingCredentials(t *testing.T) {
	t.Skip("test requires gnome-keyring but travis CI doesn't have it")

	helper := Secretservice{}
	_, _, err := helper.Get("https://adsfasdf.wrewerwer.com/asdfsdddd")
	if !credentials.IsErrCredentialsNotFound(err) {
		t.Fatalf("expected ErrCredentialsNotFound, got %v", err)
	}
}
