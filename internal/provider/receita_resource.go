package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	providerData *ReceitaProviderData
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
				Required:            true,
			},
			"author": schema.StringAttribute{
				MarkdownDescription: "Name of the author",
				Required:            true,
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
	providerData, ok := req.ProviderData.(*ReceitaProviderData)

	// if we can't fetch the client from the provider just notifies the terraform diagnostics and return silently
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected provider.ReceitaProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.providerData = providerData
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

	// Do specific stuff
	endpoint := r.providerData.Model.Endpoint
	hResp, _ := http.Get(endpoint.ValueString() + "/create")

	fmt.Println(hResp.Body)

	//End specific stuff

	//log the creation of the resource in the terraform logging system
	tflog.Trace(ctx, "created a receita resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReceitaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ReceitaResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Do read specific operation

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *ReceitaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ReceitaResourceModel

	//Read terraform state into model
	//Append is a varadic method and Get returns an array we need to use ... to transform an array into varadic representation
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	// Return in case of error
	if resp.Diagnostics.HasError() {
		return
	}

	// Do specific stuff
	endpoint := r.providerData.Model.Endpoint
	hResp, _ := http.Get(endpoint.ValueString() + "/update")

	fmt.Println(hResp.Body)

	//End specific stuff

	//Finally update the terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReceitaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ReceitaResourceModel

	// Load tf state into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	// Do specific stuff
	endpoint := r.providerData.Model.Endpoint
	hResp, _ := http.Get(endpoint.ValueString() + "/delete")

	fmt.Println(hResp.Body)

	//End specific stuff

	if resp.Diagnostics.HasError() {
		return
	}

	// Do specific delete logic

}

func (r *ReceitaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
