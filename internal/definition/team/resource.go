// Copyright Splunk, Inc.
// SPDX-License-Identifier: MPL-2.0

package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/signalfx/signalfx-go/team"

	"github.com/splunk-terraform/terraform-provider-signalfx/internal/common"
	pmeta "github.com/splunk-terraform/terraform-provider-signalfx/internal/providermeta"
)

const (
	ResourceName = "signalfx_team"
	AppPath      = "/team"
)

func NewResource() *schema.Resource {
	return &schema.Resource{
		SchemaFunc:    newSchema,
		SchemaVersion: 1,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: newResourceCreate(),
		ReadContext:   newResourceRead(),
		UpdateContext: newResourceUpdate(),
		DeleteContext: newResourceDelete(),
	}
}

func newResourceCreate() schema.CreateContextFunc {
	return func(ctx context.Context, rd *schema.ResourceData, meta any) diag.Diagnostics {
		payload, err := decodeTerraform(rd)
		if err != nil {
			return diag.FromErr(err)
		}
		client, err := pmeta.LoadClient(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		tm, err := client.CreateTeam(ctx, &team.CreateUpdateTeamRequest{
			Name:              payload.Name,
			Description:       payload.Description,
			Members:           payload.Members,
			NotificationLists: payload.NotificationLists,
		})
		if common.HandleError(ctx, err, rd) != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, "Created new team", map[string]any{
			"id": tm.Id,
		})

		if err := rd.Set("url", pmeta.LoadApplicationURL(ctx, meta, AppPath, tm.Id)); err != nil {
			return diag.FromErr(err)
		}

		return diag.FromErr(encodeTerraform(tm, rd))
	}
}

func newResourceRead() schema.ReadContextFunc {
	return func(ctx context.Context, rd *schema.ResourceData, meta any) diag.Diagnostics {
		client, err := pmeta.LoadClient(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}

		tm, err := client.GetTeam(ctx, rd.Id())
		if common.HandleError(ctx, err, rd) != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, "Successfully fetched team data")

		if err := rd.Set("url", pmeta.LoadApplicationURL(ctx, meta, AppPath, tm.Id)); err != nil {
			return diag.FromErr(err)
		}

		return diag.FromErr(encodeTerraform(tm, rd))
	}
}

func newResourceUpdate() schema.UpdateContextFunc {
	return func(ctx context.Context, rd *schema.ResourceData, meta any) diag.Diagnostics {
		payload, err := decodeTerraform(rd)
		if err != nil {
			return diag.FromErr(err)
		}

		client, err := pmeta.LoadClient(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}

		tm, err := client.UpdateTeam(ctx, rd.Id(), &team.CreateUpdateTeamRequest{
			Name:              payload.Name,
			Description:       payload.Description,
			Members:           payload.Members,
			NotificationLists: payload.NotificationLists,
		})
		if common.HandleError(ctx, err, rd) != nil {
			return diag.FromErr(err)
		}

		if err := rd.Set("url", pmeta.LoadApplicationURL(ctx, meta, AppPath, tm.Id)); err != nil {
			return diag.FromErr(err)
		}

		tflog.Debug(ctx, "Successfully updated team data")

		return diag.FromErr(encodeTerraform(tm, rd))
	}
}

func newResourceDelete() schema.DeleteContextFunc {
	return func(ctx context.Context, rd *schema.ResourceData, meta any) diag.Diagnostics {
		client, err := pmeta.LoadClient(ctx, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Debug(ctx, "Deteting team resource", map[string]any{
			"team-id": rd.Id(),
		})

		err = common.HandleError(ctx, client.DeleteTeam(ctx, rd.Id()), rd)

		return diag.FromErr(err)
	}
}
