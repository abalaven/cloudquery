// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func OAuthAuthorizations() *schema.Table {
	return &schema.Table{
		Name:        "heroku_oauth_authorizations",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#o-auth-authorization-attributes",
		Resolver:    fetchOAuthAuthorizations,
		Columns: []schema.Column{
			{
				Name:     "access_token",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessToken"),
			},
			{
				Name:     "client",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Client"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "grant",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Grant"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "refresh_token",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RefreshToken"),
			},
			{
				Name:     "scope",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
		},
	}
}

func fetchOAuthAuthorizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.OAuthAuthorizationList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
