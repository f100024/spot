package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnsibleVaultProvider_Get(t *testing.T) {
	vault_path := "testdata/test_ansible-vault"
	vault_secret := "password"
	p, err := NewAnsibleVaultProvider(vault_path, vault_secret)
	require.NoError(t, err, "failed to create AnsibleVaultProvider")

	t.Run("secret found", func(t *testing.T) {
		encrypted_secret, err := p.Get("secret")
		require.NoError(t, nil, err, "Get method should not return an error")
		assert.Equal(t, "test-secret-data", encrypted_secret, "Get method should return the correct secret value")
	})

	t.Run("secret not found", func(t *testing.T) {
		_, err := p.Get("secret-2")
		require.EqualError(t, err, "not found key: secret-2")
	})
}

func TestAnsibleVaultProvider_Create(t *testing.T) {
	vault_path := "testdata/test_ansible-vault"
	vault_path_invalid_yaml := "testdata/test_ansible-vault-invalid-yaml"
	wrong_vault_file_path := "testdata/wrong-test_ansible-vault"
	vault_file_path_is_not_regular_file := "testdata/"
	vault_secret := "password"
	wrong_vault_secret := "password0"

	t.Run("ansible vault not found", func(t *testing.T) {
		_, err := NewAnsibleVaultProvider(wrong_vault_file_path, vault_secret)
		require.EqualError(t, err, "error get fileinfo of: testdata/wrong-test_ansible-vault")
	})

	t.Run("ansible vault is not a file", func(t *testing.T) {
		_, err := NewAnsibleVaultProvider(vault_file_path_is_not_regular_file, vault_secret)
		require.EqualError(t, err, "testdata/ is not a regular file")
	})

	t.Run("ansible vault wrong password", func(t *testing.T) {
		_, err := NewAnsibleVaultProvider(vault_path, wrong_vault_secret)
		require.EqualError(t, err, "error decrypting file: testdata/test_ansible-vault")
	})

	t.Run("ansible vault error unmarshaling yaml", func(t *testing.T) {
		_, err := NewAnsibleVaultProvider(vault_path_invalid_yaml, vault_secret)
		require.EqualError(t, err, "error during unmarshaling yaml file")
	})

}
