package azure_keyvault

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
)

func LoadSecrets(vault_name string, keys []string, secrets []string) {

	ctx := context.Background()

	keyvault_client := Auth()
	var secret_parameters keyvault.SecretSetParameters
	var vault_uri string

	vault_uri = fmt.Sprintf("https://%s.vault.azure.net", vault_name)

	fmt.Println("*******Adding Keyvault Secrets*******")

	for index, key := range keys {

		secret_parameters.Value = &secrets[index]
		fmt.Println(fmt.Sprintf("%s ---------------> %s", key, *secret_parameters.Value))
		secret_bundle, err := keyvault_client.SetSecret(ctx, vault_uri, key, secret_parameters)
		_ = secret_bundle

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}

	}

}
