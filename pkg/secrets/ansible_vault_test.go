package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnsibleVaultProvider_Get(t *testing.T) {
	vault_path := "testdata/test_ansible_vault"
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
