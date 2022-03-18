package secretsmanager_interface

import (
	"context"

	. "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsManagerAPI interface {
	CancelRotateSecretAPI
	CreateSecretAPI
	DescribeSecretAPI
	GetSecretValueAPI
	UpdateSecretAPI
}

type CancelRotateSecretAPI interface {
	CancelRotateSecret(ctx context.Context, params *CancelRotateSecretInput, optFns ...func(*Options)) (*CancelRotateSecretOutput, error)
}

type CreateSecretAPI interface {
	CreateSecret(ctx context.Context, params *CreateSecretInput, optFns ...func(*Options)) (*CreateSecretOutput, error)
}

type DescribeSecretAPI interface {
	DescribeSecret(ctx context.Context, params *DescribeSecretInput, optFns ...func(*Options)) (*DescribeSecretOutput, error)
}

type GetSecretValueAPI interface {
	GetSecretValue(ctx context.Context, params *GetSecretValueInput, optFns ...func(*Options)) (*GetSecretValueOutput, error)
}

type UpdateSecretAPI interface {
	UpdateSecret(ctx context.Context, params *UpdateSecretInput, optFns ...func(*Options)) (*UpdateSecretOutput, error)
}

var _ SecretsManagerAPI = (*Client)(nil)
