// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package main

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatAWSUpload is an auto-generated flat version of AWSUpload.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAWSUpload struct {
	AccessKeyID     *string `mapstructure:"access_key_id" cty:"access_key_id" hcl:"access_key_id"`
	SecretAccessKey *string `mapstructure:"secret_access_key" cty:"secret_access_key" hcl:"secret_access_key"`
	AmiName         *string `mapstructure:"ami_name" cty:"ami_name" hcl:"ami_name"`
	S3Bucket        *string `mapstructure:"s3_bucket" cty:"s3_bucket" hcl:"s3_bucket"`
	Region          *string `mapstructure:"region" cty:"region" hcl:"region"`
}

// FlatMapstructure returns a new FlatAWSUpload.
// FlatAWSUpload is an auto-generated flat version of AWSUpload.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AWSUpload) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatAWSUpload)
}

// HCL2Spec returns the hcl spec of a AWSUpload.
// This spec is used by HCL to read the fields of AWSUpload.
// The decoded values from this spec will then be applied to a FlatAWSUpload.
func (*FlatAWSUpload) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"access_key_id":     &hcldec.AttrSpec{Name: "access_key_id", Type: cty.String, Required: false},
		"secret_access_key": &hcldec.AttrSpec{Name: "secret_access_key", Type: cty.String, Required: false},
		"ami_name":          &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"s3_bucket":         &hcldec.AttrSpec{Name: "s3_bucket", Type: cty.String, Required: false},
		"region":            &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
	}
	return s
}

// FlatBuildHost is an auto-generated flat version of BuildHost.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatBuildHost struct {
	Hostname *string `mapstructure:"hostname,required" cty:"hostname" hcl:"hostname"`
	Username *string `mapstructure:"username,required" cty:"username" hcl:"username"`
}

// FlatMapstructure returns a new FlatBuildHost.
// FlatBuildHost is an auto-generated flat version of BuildHost.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*BuildHost) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatBuildHost)
}

// HCL2Spec returns the hcl spec of a BuildHost.
// This spec is used by HCL to read the fields of BuildHost.
// The decoded values from this spec will then be applied to a FlatBuildHost.
func (*FlatBuildHost) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"hostname": &hcldec.AttrSpec{Name: "hostname", Type: cty.String, Required: false},
		"username": &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
	}
	return s
}

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	BuildHost           *FlatBuildHost    `mapstructure:"build_host,required" cty:"build_host" hcl:"build_host"`
	ImageType           *string           `mapstructure:"image_type,required" cty:"image_type" hcl:"image_type"`
	Architecture        *string           `mapstructure:"architecture" cty:"architecture" hcl:"architecture"`
	Blueprint           *string           `mapstructure:"blueprint" cty:"blueprint" hcl:"blueprint"`
	ContainerRepository *string           `mapstructure:"container_repository" cty:"container_repository" hcl:"container_repository"`
	AWSUpload           *FlatAWSUpload    `mapstructure:"aws_upload" cty:"aws_upload" hcl:"aws_upload"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"build_host":                 &hcldec.BlockSpec{TypeName: "build_host", Nested: hcldec.ObjectSpec((*FlatBuildHost)(nil).HCL2Spec())},
		"image_type":                 &hcldec.AttrSpec{Name: "image_type", Type: cty.String, Required: false},
		"architecture":               &hcldec.AttrSpec{Name: "architecture", Type: cty.String, Required: false},
		"blueprint":                  &hcldec.AttrSpec{Name: "blueprint", Type: cty.String, Required: false},
		"container_repository":       &hcldec.AttrSpec{Name: "container_repository", Type: cty.String, Required: false},
		"aws_upload":                 &hcldec.BlockSpec{TypeName: "aws_upload", Nested: hcldec.ObjectSpec((*FlatAWSUpload)(nil).HCL2Spec())},
	}
	return s
}
