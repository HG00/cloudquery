// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Records() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_domain_records",
		Resolver: fetchDomainsRecords,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Data"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "ttl",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TTL"),
			},
			{
				Name:     "weight",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Weight"),
			},
			{
				Name:     "flags",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Flags"),
			},
			{
				Name:     "tag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tag"),
			},
		},
	}
}
