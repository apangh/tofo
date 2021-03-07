package kmskeystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

type KmsKeyStore struct {
	encCtx map[string]string
	client *kms.Client
}

// NewKeyStore creates a new kms keystore
func NewKeyStore(client *kms.Client, encCtx map[string]string) *KmsKeyStore {
	return &KmsKeyStore{
		encCtx: encCtx,
		client: client,
	}
}

// New generates a new pair of ciphertext and plaintext key
func (k *KmsKeyStore) New(ctx context.Context, keyId string) (
	[]byte, []byte, string, error) {
	i := &kms.GenerateDataKeyInput{
		KeyId:             aws.String(keyId),
		EncryptionContext: k.encCtx,
		KeySpec:           types.DataKeySpecAes256,
	}
	o, e := k.client.GenerateDataKey(ctx, i)
	if e != nil {
		return nil, nil, "", e
	}
	return o.Plaintext, o.CiphertextBlob, *o.KeyId, nil
}

// Decrypt ciphertext key and generates plaintext key
func (k *KmsKeyStore) Decrypt(ctx context.Context, ciphertextKey []byte,
	keyId string) ([]byte, string, error) {
	i := &kms.DecryptInput{
		KeyId:             aws.String(keyId),
		EncryptionContext: k.encCtx,
		CiphertextBlob:    ciphertextKey,
	}
	o, e := k.client.Decrypt(ctx, i)
	if e != nil {
		return nil, "", e
	}
	return o.Plaintext, *o.KeyId, nil
}
