// Code generated by codegen; DO NOT EDIT.

package secretsmanager

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:                "aws_secretsmanager_secrets",
		Resolver:            fetchSecretsmanagerSecrets,
		PreResourceResolver: getSecret,
		Multiplex:           client.ServiceAccountRegionMultiplexer("secretsmanager"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
				Description: `The list of user-defined tags associated with the secret`,
			},
			{
				Name:        "policy",
				Type:        schema.TypeJSON,
				Resolver:    fetchSecretsmanagerSecretPolicy,
				Description: `A JSON-formatted string that describes the permissions that are associated with the attached secret.`,
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "deleted_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "last_accessed_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAccessedDate"),
			},
			{
				Name:     "last_changed_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastChangedDate"),
			},
			{
				Name:     "last_rotated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastRotatedDate"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owning_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwningService"),
			},
			{
				Name:     "primary_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrimaryRegion"),
			},
			{
				Name:     "replication_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReplicationStatus"),
			},
			{
				Name:     "rotation_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RotationEnabled"),
			},
			{
				Name:     "rotation_lambda_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RotationLambdaARN"),
			},
			{
				Name:     "rotation_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RotationRules"),
			},
			{
				Name:     "version_ids_to_stages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VersionIdsToStages"),
			},
		},
	}
}
