package azure_keyvault

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func Auth() keyvault.BaseClient {

	keyvault_client := keyvault.New()
	authorizer, err := auth.NewAuthorizerFromCLIWithResource("https://vault.azure.net")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	keyvault_client.Authorizer = authorizer

	return keyvault_client

}
