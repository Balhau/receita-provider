package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Type checking validation
// Used to enforce/validate compatibility between contracts
var _ resource.Resource = &ReceitaResource{}
var _ resource.ResourceWithImportState = &ReceitaResource{}

type ReceitaResource struct {
	client *http.Client
}

func NewReceitaResource() resource.Resource {
	return &ReceitaResource{}
}

// Resource data model
// This defines basically the contract between our resource and the user
type ReceitaResourceModel struct {
	Name   types.String `tfsdk:"name"`
	Author types.String `tfsdk:"author"`
	Id     types.String `tfsdk:"id"`
}

// Resource terraform contract definitions

// Callback to set the resource metadata information
func (r *ReceitaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_receita"
}

// Terraform contract method to setup the schema associated with the terraform resource
func (r *ReceitaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Receita data model",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the receita",
				Optional:            false,
			},
			"author": schema.StringAttribute{
				MarkdownDescription: "Name of the author",
				Optional:            false,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Receita identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Callback contract defintion that serve as terraform engine configuration
func (r *ReceitaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	// If the provider is not properly initialized then just return
	if req.ProviderData == nil {
		return
	}

	//Extract the http client from the provider data definition
	//This also casts the object, it will return fals in the ok field if the types are not castable
	client, ok := req.ProviderData.(*http.Client)

	// if we can't fetch the client from the provider just notifies the terraform diagnostics and return silently
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *ReceitaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ReceitaResourceModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Return if loading data into model turns into an error
	if resp.Diagnostics.HasError() {
		return
	}

	// Here just generates a random uuid
	data.Id = basetypes.NewStringValue(uuid.Must(uuid.NewRandom()).String())

	//log the creation of the resource in the terraform logging system
	tflog.Trace(ctx, "created a receita resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
