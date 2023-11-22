// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure KubernetesProvider satisfies various provider interfaces.
var _ provider.Provider = &KubernetesProvider{}

// KubernetesProvider defines the provider implementation.
type KubernetesProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KubernetesProviderModel describes the provider data model.
type KubernetesProviderModel struct {
}

func (p *KubernetesProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "kubernetes"
	resp.Version = p.version
}

func (p *KubernetesProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The hostname (in form of URI) of Kubernetes master.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Optional:    true,
			},
			"insecure": schema.BoolAttribute{
				Description: "Whether server should be accessed without verifying the TLS certificate.",
				Optional:    true,
			},
			"tls_server_name": schema.StringAttribute{
				Description: "Server name passed to the server for SNI and is used in the client to check server certificates against.",
				Optional:    true,
			},
			"client_certificate": schema.StringAttribute{
				Description: "PEM-encoded client certificate for TLS authentication.",
				Optional:    true,
			},
			"client_key": schema.StringAttribute{
				Description: "PEM-encoded client certificate key for TLS authentication.",
				Optional:    true,
			},
			"cluster_ca_certificate": schema.StringAttribute{
				Description: "PEM-encoded root certificates bundle for TLS authentication.",
				Optional:    true,
			},
			"config_paths": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.",
				Optional:    true,
			},
			"config_path": schema.StringAttribute{
				Description: "Path to the kube config file. Can be set with KUBE_CONFIG_PATH.",
				Optional:    true,
			},
			"config_context": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"config_context_auth_info": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"config_context_cluster": schema.StringAttribute{
				Description: "",
				Optional:    true,
			},
			"token": schema.StringAttribute{
				Description: "Token to authenticate an service account",
				Optional:    true,
			},
			"proxy_url": schema.StringAttribute{
				Description: "URL to the proxy to be used for all API requests",
				Optional:    true,
			},
			"ignore_annotations": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of Kubernetes metadata annotations to ignore across all resources handled by this provider for situations where external systems are managing certain resource annotations. Each item is a regular expression.",
				Optional:    true,
			},
			"ignore_labels": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of Kubernetes metadata labels to ignore across all resources handled by this provider for situations where external systems are managing certain resource labels. Each item is a regular expression.",
				Optional:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"exec": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"api_version": schema.StringAttribute{
							Required: true,
						},
						"command": schema.StringAttribute{
							Required: true,
						},
						"env": schema.MapAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"args": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			"experiments": schema.ListNestedBlock{
				Description: "Enable and disable experimental features.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"manifest_resource": schema.BoolAttribute{
							Description: "Enable the `kubernetes_manifest` resource.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (p *KubernetesProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *KubernetesProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *KubernetesProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) provider.Provider {
	return &KubernetesProvider{
		version: version,
	}
}