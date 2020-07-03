package kmsenv

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

type KmsEnv struct {
	ProjectID string `json:"project_id"`
	Location  string `json:"location"`
	Keyring   string `json:"keyring"`
	Key       string `json:"key"`
}

func New() (*KmsEnv, error) {
	raw, err := ioutil.ReadFile(".kmsenvrc")
	if err != nil {
		return nil, err
	}

	kmsenv := &KmsEnv{}
	err = json.Unmarshal(raw, kmsenv)
	if err != nil {
		return nil, err
	}

	return kmsenv, nil
}

func (ke *KmsEnv) Encrypt() (map[string]string, error) {
	envMap, err := environment()
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)
	for k, v := range envMap {
		if v == "" {
			continue
		}

		ciphertext, err := ke.encryption(v)
		if err != nil {
			return nil, err
		}
		encStringEncode := base64.StdEncoding.EncodeToString(ciphertext)
		res[k] = encStringEncode
	}

	return res, nil
}

func (ke *KmsEnv) encryption(str string) ([]byte, error) {
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}

	request := &kmspb.EncryptRequest{
		Name:      ke.resource(),
		Plaintext: []byte(str),
	}

	response, err := client.Encrypt(ctx, request)
	return response.GetCiphertext(), err
}

func (ke *KmsEnv) resource() string {
	return fmt.Sprintf(
		"projects/%v/locations/%v/keyRings/%v/cryptoKeys/%v",
		ke.ProjectID,
		ke.Location,
		ke.Keyring,
		ke.Key,
	)
}

// func decryption(ciphertext []byte) (string, error) {
// ctx := context.Background()
// client, err := kms.NewKeyManagementClient(ctx)
// if err != nil {
// return "", err
// }
// request := &kmspb.DecryptRequest{
// Name:       "projects/[PROJECT_NAME]/locations/global/keyRings/[KEY_RING_NAME]/cryptoKeys/[KEY_NAME]",
// Ciphertext: ciphertext,
// }
// response, err := client.Decrypt(ctx, request)
// return string(response.GetPlaintext()), err
// }

func environment() (map[string]string, error) {
	godotenv.Overload(".env")
	envMap, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	return envMap, nil
}
