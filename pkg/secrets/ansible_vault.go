package secrets

import (
	"fmt"
	"log"
	"os"

	vault "github.com/sosedoff/ansible-vault-go"
	yaml "gopkg.in/yaml.v3"
)

// AnsibleVaultProvider is a provider for ansible-vault files
type AnsibleVaultProvider struct {
	data map[string]interface{}
}

// NewAnsibleVaultProvider creates ad new instance of AnsibleVaultProvider
func NewAnsibleVaultProvider(vault_path, secret string) (*AnsibleVaultProvider, error) {
	info, err := os.Stat(vault_path)
	if err != nil {
		return nil, fmt.Errorf("error get stat of ansible-vault file: %s", vault_path)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("ansible-vault file is directory: %s", vault_path)
	}
	// Decrypt ansible-vault
	decryptedVault, err := vault.DecryptFile(vault_path, secret)
	if err != nil {
		return nil, fmt.Errorf("error decrypting file: %s", vault_path)
	}
	log.Printf("[INFO] ansible vault file decrypted")

	// Unmarshal decrypted data
	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(decryptedVault), &m)
	if err != nil {
		return nil, fmt.Errorf("error during unmarshaling yaml file")
	}
	return &AnsibleVaultProvider{m}, nil
}

// Get decrypted data from ansible-vault file
func (p *AnsibleVaultProvider) Get(key string) (string, error) {
	if key_value, ok := p.data[key]; ok {
		return fmt.Sprintf("%v", key_value), nil
	}
	return "", fmt.Errorf("not found key: %v", key)
}
