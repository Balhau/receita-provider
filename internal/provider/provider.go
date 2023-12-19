package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Struct for the provider implementation
type ReceitaProvider struct {
	//Version identfier for our provider
	version string
}

// Datamodel for our provider
type ReceitaProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

// Now we need to define the contract defined by provider.Provider from terraform

// Set the metadata callback for the provider
// This callback is responsible to build metadata associated to the provider
func (p *ReceitaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	//set the metadata on the response object
	//get the version from the provider object
	resp.Version = p.version
	//set the name of the provider
	resp.TypeName = "receita"
}

// Set the schema of the provider
// This callback will be responsible to define the provider schema
func (p *ReceitaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Endpoint under which the receitas will be called upon",
				Required:            true,
			},
		},
	}
}

// Set the initializer for the provider
// Callback to build the provider

func (p *ReceitaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ReceitaProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Set the resources enabled by this provider

func (p *ReceitaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewReceitaResource,
	}
}

func (p *ReceitaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// This is the builder method for our ReceitaProvider instances
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ReceitaProvider{
			version: version,
		}
	}
}
