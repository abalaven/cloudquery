// Code generated by codegen; DO NOT EDIT.

package {{.SubService}}

import (
  "context"

  "github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
  "github.com/cloudquery/plugin-sdk/schema"
)

func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
  cqClient := meta.(*client.Client)

  response, err := cqClient.PagerdutyClient.{{if ne .ListFunctionNameOverride ""}}{{.ListFunctionNameOverride}}{{else}}List{{.StructName}}sPaginated{{end}}(ctx)
  if err != nil {
    return err
  }

  res <- response

  return nil
}