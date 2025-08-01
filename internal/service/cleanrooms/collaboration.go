// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanrooms

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_cleanrooms_collaboration", name="Collaboration")
// @Tags(identifierAttribute="arn")
// @Testing(tagsTest=false)
func ResourceCollaboration() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceCollaborationCreate,
		ReadWithoutTimeout:   resourceCollaborationRead,
		UpdateWithoutTimeout: resourceCollaborationUpdate,
		DeleteWithoutTimeout: resourceCollaborationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"analytics_engine": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: enum.Validate[types.AnalyticsEngine](),
			},
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrCreateTime: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creator_display_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"creator_member_abilities": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrDescription: {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_encryption_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_clear_text": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"allow_duplicates": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"allow_joins_on_columns_with_different_names": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"preserve_nulls": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			names.AttrID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"member": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						names.AttrAccountID: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						names.AttrDisplayName: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						names.AttrStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						"member_abilities": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			names.AttrName: {
				Type:     schema.TypeString,
				Required: true,
			},
			"query_log_status": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

const (
	ResNameCollaboration = "Collaboration"
)

func resourceCollaborationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).CleanRoomsClient(ctx)

	creatorAbilities := d.Get("creator_member_abilities").([]any)

	input := &cleanrooms.CreateCollaborationInput{
		Name:                   aws.String(d.Get(names.AttrName).(string)),
		CreatorDisplayName:     aws.String(d.Get("creator_display_name").(string)),
		CreatorMemberAbilities: expandMemberAbilities(creatorAbilities),
		Members:                *expandMembers(d.Get("member").(*schema.Set).List()),
		Tags:                   getTagsIn(ctx),
	}

	if v, ok := d.GetOk("analytics_engine"); ok {
		input.AnalyticsEngine = types.AnalyticsEngine(v.(string))
	}

	queryLogStatus, err := expandQueryLogStatus(d.Get("query_log_status").(string))
	if err != nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionCreating, ResNameCollaboration, d.Get(names.AttrName).(string), err)
	}
	input.QueryLogStatus = queryLogStatus

	if v, ok := d.GetOk("data_encryption_metadata"); ok {
		input.DataEncryptionMetadata = expandDataEncryptionMetadata(v.([]any))
	}

	if v, ok := d.GetOk(names.AttrDescription); ok {
		input.Description = aws.String(v.(string))
	}

	out, err := conn.CreateCollaboration(ctx, input)
	if err != nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionCreating, ResNameCollaboration, d.Get(names.AttrName).(string), err)
	}

	if out == nil || out.Collaboration == nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionCreating, ResNameCollaboration, d.Get(names.AttrName).(string), errors.New("empty output"))
	}
	d.SetId(aws.ToString(out.Collaboration.Id))

	return append(diags, resourceCollaborationRead(ctx, d, meta)...)
}

func resourceCollaborationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).CleanRoomsClient(ctx)

	out, err := findCollaborationByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] CleanRooms Collaboration (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionReading, ResNameCollaboration, d.Id(), err)
	}

	collaboration := out.Collaboration
	d.Set(names.AttrARN, collaboration.Arn)
	d.Set(names.AttrName, collaboration.Name)
	d.Set(names.AttrDescription, collaboration.Description)
	d.Set("analytics_engine", collaboration.AnalyticsEngine)
	d.Set("creator_display_name", collaboration.CreatorDisplayName)
	d.Set(names.AttrCreateTime, collaboration.CreateTime.String())
	d.Set("update_time", collaboration.UpdateTime.String())
	d.Set("query_log_status", collaboration.QueryLogStatus)
	if err := d.Set("data_encryption_metadata", flattenDataEncryptionMetadata(collaboration.DataEncryptionMetadata)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting data_encryption_metadata: %s", err)
	}

	membersOut, err := findMembersByCollaborationId(ctx, conn, d.Id())
	if err != nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionSetting, ResNameCollaboration, d.Id(), err)
	}

	if err := d.Set("member", flattenMembers(membersOut.MemberSummaries, collaboration.CreatorAccountId)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting member: %s", err)
	}
	if err := d.Set("creator_member_abilities", flattenCreatorAbilities(membersOut.MemberSummaries, collaboration.CreatorAccountId)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting creator_member_abilities: %s", err)
	}

	return diags
}

func resourceCollaborationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).CleanRoomsClient(ctx)

	if d.HasChangesExcept(names.AttrTags, names.AttrTagsAll) {
		input := &cleanrooms.UpdateCollaborationInput{
			CollaborationIdentifier: aws.String(d.Id()),
		}

		if d.HasChanges(names.AttrDescription) {
			input.Description = aws.String(d.Get(names.AttrDescription).(string))
		}

		if d.HasChanges(names.AttrName) {
			input.Name = aws.String(d.Get(names.AttrName).(string))
		}

		if d.HasChanges("analytics_engine") {
			input.AnalyticsEngine = types.AnalyticsEngine(d.Get("analytics_engine").(string))
		}

		_, err := conn.UpdateCollaboration(ctx, input)
		if err != nil {
			return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionUpdating, ResNameCollaboration, d.Id(), err)
		}
	}

	return append(diags, resourceCollaborationRead(ctx, d, meta)...)
}

func resourceCollaborationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	conn := meta.(*conns.AWSClient).CleanRoomsClient(ctx)

	log.Printf("[INFO] Deleting CleanRooms Collaboration %s", d.Id())
	input := cleanrooms.DeleteCollaborationInput{
		CollaborationIdentifier: aws.String(d.Id()),
	}
	_, err := conn.DeleteCollaboration(ctx, &input)

	if errs.IsA[*types.AccessDeniedException](err) {
		return diags
	}

	if err != nil {
		return create.AppendDiagError(diags, names.CleanRooms, create.ErrActionDeleting, ResNameCollaboration, d.Id(), err)
	}

	return diags
}

func findCollaborationByID(ctx context.Context, conn *cleanrooms.Client, id string) (*cleanrooms.GetCollaborationOutput, error) {
	in := &cleanrooms.GetCollaborationInput{
		CollaborationIdentifier: aws.String(id),
	}
	out, err := conn.GetCollaboration(ctx, in)

	if errs.IsA[*types.AccessDeniedException](err) {
		//We throw Access Denied for NFE in Cleanrooms for collaborations since they are cross account
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: in,
		}
	}

	if err != nil {
		return nil, err
	}

	if out == nil || out.Collaboration == nil {
		return nil, tfresource.NewEmptyResultError(in)
	}

	return out, nil
}

func findMembersByCollaborationId(ctx context.Context, conn *cleanrooms.Client, id string) (*cleanrooms.ListMembersOutput, error) {
	in := &cleanrooms.ListMembersInput{
		CollaborationIdentifier: aws.String(id),
	}
	out, err := conn.ListMembers(ctx, in)

	if errs.IsA[*types.ResourceNotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: in,
		}
	}

	if err != nil {
		return nil, err
	}

	if out == nil || out.MemberSummaries == nil {
		return nil, tfresource.NewEmptyResultError(in)
	}

	return out, nil
}

func expandMemberAbilities(data []any) []types.MemberAbility {
	mappedAbilities := []types.MemberAbility{}
	for _, v := range data {
		switch v.(string) {
		case "CAN_QUERY":
			mappedAbilities = append(mappedAbilities, types.MemberAbilityCanQuery)
		case "CAN_RECEIVE_RESULTS":
			mappedAbilities = append(mappedAbilities, types.MemberAbilityCanReceiveResults)
		}
	}
	return mappedAbilities
}

func expandQueryLogStatus(status string) (types.CollaborationQueryLogStatus, error) {
	switch status {
	case "ENABLED":
		return types.CollaborationQueryLogStatusEnabled, nil
	case "DISABLED":
		return types.CollaborationQueryLogStatusDisabled, nil
	default:
		return types.CollaborationQueryLogStatusDisabled, fmt.Errorf("Invalid query log status type: %s", status)
	}
}

func expandDataEncryptionMetadata(data []any) *types.DataEncryptionMetadata {
	dataEncryptionMetadata := &types.DataEncryptionMetadata{}
	if len(data) > 0 {
		metadata := data[0].(map[string]any)
		dataEncryptionMetadata.PreserveNulls = aws.Bool(metadata["preserve_nulls"].(bool))
		dataEncryptionMetadata.AllowCleartext = aws.Bool(metadata["allow_clear_text"].(bool))
		dataEncryptionMetadata.AllowJoinsOnColumnsWithDifferentNames = aws.Bool(metadata["allow_joins_on_columns_with_different_names"].(bool))
		dataEncryptionMetadata.AllowDuplicates = aws.Bool(metadata["allow_duplicates"].(bool))
	}
	return dataEncryptionMetadata
}

func expandMembers(data []any) *[]types.MemberSpecification {
	members := []types.MemberSpecification{}
	for _, member := range data {
		memberMap := member.(map[string]any)
		member := &types.MemberSpecification{
			AccountId:       aws.String(memberMap[names.AttrAccountID].(string)),
			MemberAbilities: expandMemberAbilities(memberMap["member_abilities"].([]any)),
			DisplayName:     aws.String(memberMap[names.AttrDisplayName].(string)),
		}
		members = append(members, *member)
	}
	return &members
}

func flattenDataEncryptionMetadata(dataEncryptionMetadata *types.DataEncryptionMetadata) []any {
	if dataEncryptionMetadata == nil {
		return nil
	}
	m := map[string]any{}
	m["preserve_nulls"] = aws.Bool(*dataEncryptionMetadata.PreserveNulls)
	m["allow_clear_text"] = aws.Bool(*dataEncryptionMetadata.AllowCleartext)
	m["allow_joins_on_columns_with_different_names"] = aws.Bool(*dataEncryptionMetadata.AllowJoinsOnColumnsWithDifferentNames)
	m["allow_duplicates"] = aws.Bool(*dataEncryptionMetadata.AllowDuplicates)
	return []any{m}
}

func flattenMembers(members []types.MemberSummary, ownerAccount *string) []any {
	flattenedMembers := []any{}
	for _, member := range members {
		if aws.ToString(member.AccountId) != aws.ToString(ownerAccount) {
			memberMap := map[string]any{}
			memberMap[names.AttrStatus] = member.Status
			memberMap[names.AttrAccountID] = member.AccountId
			memberMap[names.AttrDisplayName] = member.DisplayName
			memberMap["member_abilities"] = flattenMemberAbilities(member.Abilities)
			flattenedMembers = append(flattenedMembers, memberMap)
		}
	}
	return flattenedMembers
}

func flattenCreatorAbilities(members []types.MemberSummary, ownerAccount *string) []string {
	flattenedAbilities := []string{}
	for _, member := range members {
		if aws.ToString(member.AccountId) == aws.ToString(ownerAccount) {
			return flattenMemberAbilities(member.Abilities)
		}
	}
	return flattenedAbilities
}

func flattenMemberAbilities(abilities []types.MemberAbility) []string {
	flattenedAbilities := []string{}
	for _, ability := range abilities {
		flattenedAbilities = append(flattenedAbilities, string(ability))
	}
	return flattenedAbilities
}
