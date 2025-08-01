// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"slices"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_ebs_volume", name="EBS Volume")
// @Tags
// @Testing(tagsTest=false)
func dataSourceEBSVolume() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceEBSVolumeRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrAvailabilityZone: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrCreateTime: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrEncrypted: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			names.AttrFilter: customFiltersSchema(),
			names.AttrIOPS: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			names.AttrKMSKeyID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrMostRecent: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"multi_attach_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"outpost_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrSize: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			names.AttrSnapshotID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
			names.AttrThroughput: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			names.AttrVolumeType: {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_initialization_rate": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceEBSVolumeRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*conns.AWSClient)
	conn := c.EC2Client(ctx)

	input := ec2.DescribeVolumesInput{}

	input.Filters = append(input.Filters, newCustomFilterList(
		d.Get(names.AttrFilter).(*schema.Set),
	)...)

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	output, err := findEBSVolumes(ctx, conn, &input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EBS Volumes: %s", err)
	}

	if len(output) < 1 {
		return sdkdiag.AppendErrorf(diags, "Your query returned no results. Please change your search criteria and try again.")
	}

	var volume awstypes.Volume

	if len(output) > 1 {
		recent := d.Get(names.AttrMostRecent).(bool)

		if !recent {
			return sdkdiag.AppendErrorf(diags, "Your query returned more than one result. Please try a more "+
				"specific search criteria, or set `most_recent` attribute to true.")
		}

		volume = mostRecentVolume(output)
	} else {
		// Query returned single result.
		volume = output[0]
	}

	d.SetId(aws.ToString(volume.VolumeId))
	d.Set(names.AttrARN, ebsVolumeARN(ctx, c, d.Id()))
	d.Set(names.AttrAvailabilityZone, volume.AvailabilityZone)
	d.Set(names.AttrCreateTime, volume.CreateTime.Format(time.RFC3339))
	d.Set(names.AttrEncrypted, volume.Encrypted)
	d.Set(names.AttrIOPS, volume.Iops)
	d.Set(names.AttrKMSKeyID, volume.KmsKeyId)
	d.Set("multi_attach_enabled", volume.MultiAttachEnabled)
	d.Set("outpost_arn", volume.OutpostArn)
	d.Set(names.AttrSize, volume.Size)
	d.Set(names.AttrSnapshotID, volume.SnapshotId)
	d.Set(names.AttrThroughput, volume.Throughput)
	d.Set("volume_id", volume.VolumeId)
	d.Set("volume_initialization_rate", volume.VolumeInitializationRate)
	d.Set(names.AttrVolumeType, volume.VolumeType)

	setTagsOut(ctx, volume.Tags)

	return diags
}

func mostRecentVolume(volumes []awstypes.Volume) awstypes.Volume {
	return slices.MaxFunc(volumes, func(a, b awstypes.Volume) int {
		return a.CreateTime.Compare(aws.ToTime(b.CreateTime))
	})
}
