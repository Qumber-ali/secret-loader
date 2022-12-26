package aws_secrets_manager

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func LoadSecrets(aws_profile string, vault_name string, keys []string, secrets []string) {

	session := Auth(aws_profile)

	client := secretsmanager.New(session)

	var secret_inputs *secretsmanager.CreateSecretInput = &secretsmanager.CreateSecretInput{}

	fmt.Println("*******Adding Secret Manager Secrets*******")

	for index, key := range keys {

		secret_inputs.SecretString = &secrets[index]
		secret_inputs.Name = &key
		fmt.Println(fmt.Sprintf("%s ---------------> %s", key, *secret_inputs.SecretString))
		output, err := client.CreateSecret(secret_inputs)
		_ = output

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}

	}
}
