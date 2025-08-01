// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssm_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/hashicorp/go-version"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfssm "github.com/hashicorp/terraform-provider-aws/internal/service/ssm"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccSSMParameter_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "ssm", fmt.Sprintf("parameter/%s", name)),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrVersion),
					resource.TestCheckResourceAttr(resourceName, "data_type", "text"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTags), knownvalue.Null()),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTagsAll), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			// Test import with version.
			// https://github.com/hashicorp/terraform-provider-aws/issues/37812.
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateId:                        name + ":1",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: names.AttrName,
				ImportStateVerifyIgnore:              []string{names.AttrID, "overwrite", "has_value_wo"},
			},
		},
	})
}

// TestAccSSMParameter_multiple is mostly a performance benchmark
func TestAccSSMParameter_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_multiple(rName, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "ssm", fmt.Sprintf("parameter/%s-1", rName)),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrVersion),
					resource.TestCheckResourceAttr(resourceName, "data_type", "text"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
			},
		},
	})
}

func TestAccSSMParameter_updateValue(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "String", "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test2"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_updateDescription(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_description(name, names.AttrDescription, "String", "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, names.AttrDescription),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_description(name, "updated description", "String", "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "updated description"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test"),
					resource.TestCheckNoResourceAttr(resourceName, "overwrite"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_writeOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.SSMServiceID),
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_writeOnly(rName, "test", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					testAccCheckParameterWriteOnlyValueEqual(t, &param, "test"),
				),
			},
			{
				Config: testAccParameterConfig_writeOnly(rName, "testUpdated", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					testAccCheckParameterWriteOnlyValueEqual(t, &param, "testUpdated"),
				),
			},
		},
	})
}

func TestAccSSMParameter_changeValueToWriteOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.SSMServiceID),
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(version.Must(version.NewVersion("1.11.0"))),
		},
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_changeValueToWriteOnly1(rName, "SecureString", "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
			},
			{
				Config: testAccParameterConfig_changeValueToWriteOnly2(rName, "SecureString", "testUpdated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					testAccCheckParameterWriteOnlyValueEqual(t, &param, "testUpdated"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("has_value_wo"), knownvalue.Bool(true)),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("value_wo_version"),
						knownvalue.NumberFunc(func(v *big.Float) error {
							if v.IsInt() {
								if v == nil {
									return fmt.Errorf("version is nil")
								}
								if v.Cmp(big.NewFloat(0)) <= 0 { // Si v <= 0
									return fmt.Errorf("expected version to be greater than 0, got %s", v.String())
								}
								return nil
							} else {
								return fmt.Errorf("expected version to be an int value")
							}
						})),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value_wo"), knownvalue.Null()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("has_value_wo")),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("value_wo_version")),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("value_wo"), knownvalue.Null()),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config: testAccParameterConfig_changeValueToWriteOnly1(rName, "SecureString", "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("has_value_wo"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("value_wo_version"),
						knownvalue.NumberExact(big.NewFloat(float64(0))),
					),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("value_wo"), knownvalue.Null()),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("has_value_wo")),
						plancheck.ExpectKnownValue(
							resourceName,
							tfjsonpath.New("value_wo_version"),
							knownvalue.Null(),
						),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("value_wo"), knownvalue.Null()),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("insecure_value"), knownvalue.Null()),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New(names.AttrValue)),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New(names.AttrVersion)),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func TestAccSSMParameter_tier(t *testing.T) {
	ctx := acctest.Context(t)
	var parameter1, parameter2, parameter3 awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierAdvanced)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter1),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierStandard)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter2),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierAdvanced)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter3),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
		},
	})
}

func TestAccSSMParameter_Tier_intelligentTieringToStandard(t *testing.T) {
	ctx := acctest.Context(t)
	var parameter awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierIntelligentTiering)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierStandard)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierIntelligentTiering)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Tier_intelligentTieringToAdvanced(t *testing.T) {
	ctx := acctest.Context(t)
	var parameter1, parameter2 awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierIntelligentTiering)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter1),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierAdvanced)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter1),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
			{
				// Intelligent-Tiering will not downgrade an existing parameter to Standard
				Config: testAccParameterConfig_tier(rName, string(awstypes.ParameterTierIntelligentTiering)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter2),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Tier_intelligentTieringOnCreation(t *testing.T) {
	ctx := acctest.Context(t)
	var parameter awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	value := sdkacctest.RandString(5000) // Maximum size for Standard tier is 4 KB

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_tierWithValue(rName, string(awstypes.ParameterTierIntelligentTiering), value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Tier_intelligentTieringOnUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var parameter awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	standardSizedValue := sdkacctest.RandString(10)
	advancedSizedValue := sdkacctest.RandString(5000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_tierWithValue(rName, string(awstypes.ParameterTierIntelligentTiering), standardSizedValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierStandard)),
				),
			},
			{
				Config: testAccParameterConfig_tierWithValue(rName, string(awstypes.ParameterTierIntelligentTiering), advancedSizedValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &parameter),
					resource.TestCheckResourceAttr(resourceName, "tier", string(awstypes.ParameterTierAdvanced)),
				),
			},
		},
	})
}

func TestAccSSMParameter_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfssm.ResourceParameter(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSSMParameter_Overwrite_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					conn := acctest.Provider.Meta().(*conns.AWSClient).SSMClient(ctx)

					input := &ssm.PutParameterInput{
						Name:  aws.String(fmt.Sprintf("%s-%s", "test_parameter", name)),
						Type:  awstypes.ParameterTypeString,
						Value: aws.String("This value is set using the SDK"),
					}

					_, err := conn.PutParameter(ctx, input)

					if err != nil {
						t.Fatal(err)
					}
				},
				Config: testAccParameterConfig_basicOverwrite(name, "String", "This value is set using Terraform"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "2"),
					resource.TestCheckResourceAttr(resourceName, "overwrite", acctest.CtTrue),
				),
			},
			{
				Config: testAccParameterConfig_basicOverwrite(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "3"),
					resource.TestCheckResourceAttr(resourceName, "overwrite", acctest.CtTrue),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"overwrite", "has_value_wo"},
			},
			{
				Config: testAccParameterConfig_basicOverwrite(name, "String", "test3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test3"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "4"),
					resource.TestCheckResourceAttr(resourceName, "overwrite", acctest.CtTrue),
				),
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/12213
func TestAccSSMParameter_Overwrite_cascade(t *testing.T) {
	ctx := acctest.Context(t)
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test_upstream"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_cascadeOverwrite(name, "test1"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
			{
				Config: testAccParameterConfig_cascadeOverwrite(name, "test2"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				Config: testAccParameterConfig_cascadeOverwrite(name, "test2"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/18550
func TestAccSSMParameter_Overwrite_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_overwriteTags1(rName, true, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"overwrite", "has_value_wo"},
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/18550
func TestAccSSMParameter_Overwrite_noOverwriteTags(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_overwriteTags1(rName, false, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"overwrite", "has_value_wo"},
			},
		},
	})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/18550
func TestAccSSMParameter_Overwrite_updateToTags(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basicTags1(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_overwriteTags1(rName, true, acctest.CtKey1, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue2),
					})),
				},
			},
		},
	})
}
func TestAccSSMParameter_Overwrite_removeAttribute(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.SSMServiceID),
		CheckDestroy: testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"aws": {
						Source:            "hashicorp/aws",
						VersionConstraint: "4.67.0",
					},
				},
				Config: testAccParameterConfig_overwriteRemove_Setup(rName, "String", acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "overwrite", acctest.CtTrue),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				Config:                   testAccParameterConfig_overwriteRemove_Remove(rName, "String", acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "overwrite", acctest.CtFalse),
				),
			},
		},
	})
}

func TestAccSSMParameter_updateType(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "SecureString", "test2"),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
			},
		},
	})
}

func TestAccSSMParameter_Overwrite_updateDescription(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basicOverwrite(name, "String", "test2"),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"overwrite", "has_value_wo"},
			},
			{
				Config: testAccParameterConfig_basicOverwriteNoDescription(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, ""),
				),
			},
		},
	})
}

func TestAccSSMParameter_changeNameForcesNew(t *testing.T) {
	ctx := acctest.Context(t)
	var beforeParam, afterParam awstypes.Parameter
	before := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	after := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(before, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &beforeParam),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_basic(after, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &afterParam),
					testAccCheckParameterRecreated(t, &beforeParam, &afterParam),
				),
			},
		},
	})
}

func TestAccSSMParameter_fullPath(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("/path/%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					acctest.CheckResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "ssm", fmt.Sprintf("parameter%s", name)),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "test2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Secure_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "SecureString", "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "secret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
					resource.TestCheckResourceAttr(resourceName, names.AttrKeyID, "alias/aws/ssm"), // Default SSM key id
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Secure_insecure(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_insecure(rName, "String", "notsecret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "insecure_value", "notsecret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
			{
				Config: testAccParameterConfig_insecure(rName, "String", "newvalue"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "insecure_value", "newvalue"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				Config: testAccParameterConfig_insecure(rName, "String", "diff"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
			{
				Config:      testAccParameterConfig_insecure(rName, "SecureString", "notsecret"),
				ExpectError: regexache.MustCompile("invalid configuration"),
			},
		},
	})
}

func TestAccSSMParameter_Secure_insecureChangeSecure(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_insecure(rName, "String", "notsecret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "insecure_value", "notsecret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
			},
			{
				Config: testAccParameterConfig_secure(rName, "newvalue"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "newvalue"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
				),
			},
			{
				Config: testAccParameterConfig_insecure(rName, "String", "atlantis"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "insecure_value", "atlantis"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "String"),
				),
			},
		},
	})
}

func TestAccSSMParameter_DataType_ec2Image(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_dataTypeEC2Image(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "data_type", "aws:ec2:image"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_DataType_ssmIntegration(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	webhookName := sdkacctest.RandString(16)
	rName := fmt.Sprintf("/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/%s", webhookName)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_dataTypeSSMIntegration(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "data_type", "aws:ssm:integration"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_DataType_update(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_dataTypeUpdate(rName, "text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "data_type", "text"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_dataTypeUpdate(rName, "aws:ec2:image"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, "data_type", "aws:ec2:image"),
				),
			},
		},
	})
}

func TestAccSSMParameter_Secure_key(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	randString := sdkacctest.RandString(10)
	name := fmt.Sprintf("%s_%s", t.Name(), randString)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_secureKey(name, "secret", randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "secret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
					resource.TestCheckResourceAttr(resourceName, names.AttrKeyID, "alias/"+randString),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
		},
	})
}

func TestAccSSMParameter_Secure_keyUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	randString := sdkacctest.RandString(10)
	name := fmt.Sprintf("%s_%s", t.Name(), randString)
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_secure(name, "secret"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "secret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
					resource.TestCheckResourceAttr(resourceName, names.AttrKeyID, "alias/aws/ssm"), // Default SSM key id
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"has_value_wo"},
			},
			{
				Config: testAccParameterConfig_secureKey(name, "secret", randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, "secret"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "SecureString"),
					resource.TestCheckResourceAttr(resourceName, names.AttrKeyID, "alias/"+randString),
				),
			},
		},
	})
}

// lintignore:AT002
func TestAccSSMParameter_importByARN(t *testing.T) {
	ctx := acctest.Context(t)
	var param awstypes.Parameter
	name := fmt.Sprintf("%s_%s", t.Name(), sdkacctest.RandString(10))
	resourceName := "aws_ssm_parameter.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.SSMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckParameterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccParameterConfig_basic(name, "String", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParameterExists(ctx, resourceName, &param),
				),
			},
			// Test import by ARN.
			// https://github.com/hashicorp/terraform-provider-aws/issues/39050.
			{
				ResourceName: resourceName,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return aws.ToString(param.ARN), nil
				},
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: names.AttrName,
				ImportStateVerifyIgnore:              []string{names.AttrID, "overwrite", "has_value_wo"},
			},
		},
	})
}

func testAccCheckParameterRecreated(t *testing.T, before, after *awstypes.Parameter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *before.Name == *after.Name {
			t.Fatalf("Expected change of SSM Param Names, but both were %v", *before.Name)
		}
		return nil
	}
}

func testAccCheckParameterWriteOnlyValueEqual(t *testing.T, param *awstypes.Parameter, writeOnlyValue string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if aws.ToString(param.Value) != writeOnlyValue {
			t.Fatalf("Expected SSM Param Value to be %v, but got %v", writeOnlyValue, aws.ToString(param.Value))
		}
		return nil
	}
}

func testAccCheckParameterExists(ctx context.Context, n string, v *awstypes.Parameter) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SSMClient(ctx)

		output, err := tfssm.FindParameterByName(ctx, conn, rs.Primary.ID, true)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckParameterDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SSMClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_ssm_parameter" {
				continue
			}

			_, err := tfssm.FindParameterByName(ctx, conn, rs.Primary.ID, false)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SSM Parameter %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccParameterConfig_basic(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name  = %[1]q
  type  = %[2]q
  value = %[3]q
}
`, rName, pType, value)
}

func testAccParameterConfig_multiple(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name  = "%[1]s-1"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test2" {
  name  = "%[1]s-2"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test3" {
  name  = "%[1]s-3"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test4" {
  name  = "%[1]s-4"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test5" {
  name  = "%[1]s-5"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test6" {
  name  = "%[1]s-6"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test7" {
  name  = "%[1]s-7"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test8" {
  name  = "%[1]s-8"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test9" {
  name  = "%[1]s-9"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test10" {
  name  = "%[1]s-10"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test11" {
  name  = "%[1]s-11"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test12" {
  name  = "%[1]s-12"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test13" {
  name  = "%[1]s-13"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test14" {
  name  = "%[1]s-14"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test15" {
  name  = "%[1]s-15"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test16" {
  name  = "%[1]s-16"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test17" {
  name  = "%[1]s-17"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test18" {
  name  = "%[1]s-18"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test19" {
  name  = "%[1]s-19"
  type  = %[2]q
  value = %[3]q
}

resource "aws_ssm_parameter" "test20" {
  name  = "%[1]s-20"
  type  = %[2]q
  value = %[3]q
}
`, rName, pType, value)
}

func testAccParameterConfig_description(rName, description, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = %[1]q
  description = %[2]q
  type        = %[3]q
  value       = %[4]q
}
`, rName, description, pType, value)
}

func testAccParameterConfig_insecure(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name           = %[1]q
  type           = %[2]q
  insecure_value = %[3]q
}
`, rName, pType, value)
}

func testAccParameterConfig_tier(rName, tier string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name  = %[1]q
  tier  = %[2]q
  type  = "String"
  value = "test2"
}
`, rName, tier)
}

func testAccParameterConfig_tierWithValue(rName, tier, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name  = %[1]q
  tier  = %[2]q
  type  = "String"
  value = %[3]q
}
`, rName, tier, value)
}

func testAccParameterConfig_dataTypeEC2Image(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigLatestAmazonLinux2HVMEBSX8664AMI(),
		fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name      = %[1]q
  data_type = "aws:ec2:image"
  type      = "String"
  value     = data.aws_ami.amzn2-ami-minimal-hvm-ebs-x86_64.id
}
`, rName))
}

func testAccParameterConfig_dataTypeUpdate(rName, datatype string) string {
	return acctest.ConfigCompose(
		acctest.ConfigLatestAmazonLinux2HVMEBSX8664AMI(),
		fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name      = %[1]q
  data_type = %[2]q
  type      = "String"
  value     = data.aws_ami.amzn2-ami-minimal-hvm-ebs-x86_64.id
}
`, rName, datatype))
}

func testAccParameterConfig_dataTypeSSMIntegration(rName string) string { // nosemgrep:ci.ssm-in-func-name
	return acctest.ConfigCompose(
		fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name      = %[1]q
  data_type = "aws:ssm:integration"
  type      = "SecureString"
  value     = "{\"description\": \"My first webhook integration for Automation.\", \"url\": \"https://example.com\"}"
}
`, rName))
}

func testAccParameterConfig_basicTags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name  = %[1]q
  type  = "String"
  value = %[1]q

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccParameterConfig_basicOverwrite(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = "test_parameter-%[1]s"
  description = "description for parameter %[1]s"
  type        = "%[2]s"
  value       = "%[3]s"
  overwrite   = true
}
`, rName, pType, value)
}

func testAccParameterConfig_basicOverwriteNoDescription(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name      = "test_parameter-%[1]s"
  type      = "%[2]s"
  value     = "%[3]s"
  overwrite = true
}
`, rName, pType, value)
}

func testAccParameterConfig_overwriteTags1(rName string, overwrite bool, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name      = %[1]q
  overwrite = %[2]t
  type      = "String"
  value     = %[1]q
  tags = {
    %[3]q = %[4]q
  }
}
`, rName, overwrite, tagKey1, tagValue1)
}

func testAccParameterConfig_cascadeOverwrite(rName, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test_upstream" {
  name      = "test_parameter_upstream-%[1]s"
  type      = "String"
  value     = "%[2]s"
  overwrite = true
}

resource "aws_ssm_parameter" "test_downstream" {
  name      = "test_parameter_downstream-%[1]s"
  type      = "String"
  value     = aws_ssm_parameter.test_upstream.version
  overwrite = true
}
`, rName, value)
}

func testAccParameterConfig_overwriteRemove_Setup(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = "test_parameter-%[1]s"
  description = "description for parameter %[1]s"
  type        = "%[2]s"
  value       = "%[3]s"
  overwrite   = true
}
`, rName, pType, value)
}

func testAccParameterConfig_overwriteRemove_Remove(rName, pType, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = "test_parameter-%[1]s"
  description = "description for parameter %[1]s"
  type        = "%[2]s"
  value       = "%[3]s"
}
`, rName, pType, value)
}

func testAccParameterConfig_secure(rName string, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = "test_secure_parameter-%[1]s"
  description = "description for parameter %[1]s"
  type        = "SecureString"
  value       = "%[2]s"
}
`, rName, value)
}

func testAccParameterConfig_secureKey(rName string, value string, keyAlias string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name        = "test_secure_parameter-%[1]s"
  description = "description for parameter %[1]s"
  type        = "SecureString"
  value       = "%[2]s"
  key_id      = "alias/%[3]s"
  depends_on  = [aws_kms_alias.test_alias]
}

resource "aws_kms_key" "test_key" {
  description             = "KMS key 1"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_kms_alias" "test_alias" {
  name          = "alias/%[3]s"
  target_key_id = aws_kms_key.test_key.id
}
`, rName, value, keyAlias)
}

func testAccParameterConfig_writeOnly(rName string, value string, valueVersion int) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "test" {
  name             = %[1]q
  type             = "String"
  value_wo         = %[2]q
  value_wo_version = %[3]d
}
`, rName, value, valueVersion)
}

func testAccParameterConfig_changeValueToWriteOnly1(rName, typ, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "prereq" {
  name  = "%[1]s-prereq"
  type  = %[2]q
  value = %[3]q
}

data "aws_ssm_parameter" "prereq" {
  name = aws_ssm_parameter.prereq.name
}

resource "aws_ssm_parameter" "test" {
  name  = %[1]q
  type  = %[2]q
  value = data.aws_ssm_parameter.prereq.value
}
`, rName, typ, value)
}

func testAccParameterConfig_changeValueToWriteOnly2(rName, typ, value string) string {
	return fmt.Sprintf(`
resource "aws_ssm_parameter" "prereq" {
  name  = "%[1]s-prereq"
  type  = %[2]q
  value = %[3]q
}

data "aws_ssm_parameter" "prereq" {
  name = aws_ssm_parameter.prereq.name
}

resource "aws_ssm_parameter" "test" {
  name             = %[1]q
  type             = %[2]q
  value_wo         = data.aws_ssm_parameter.prereq.value
  value_wo_version = data.aws_ssm_parameter.prereq.version
}
`, rName, typ, value)
}
