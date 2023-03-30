// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"terraform/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationResource{}
var _ resource.ResourceWithImportState = &DestinationResource{}

func NewDestinationResource() resource.Resource {
	return &DestinationResource{}
}

// DestinationResource defines the resource implementation.
type DestinationResource struct {
	client *sdk.SDK
}

// DestinationResourceModel describes the resource data model.
type DestinationResourceModel struct {
	GroupID        types.String              `tfsdk:"group_id"`
	ID             types.String              `tfsdk:"id"`
	Region         types.String              `tfsdk:"region"`
	Service        types.String              `tfsdk:"service"`
	SetupStatus    types.String              `tfsdk:"setup_status"`
	SetupTests     []SetupTestResultResponse `tfsdk:"setup_tests"`
	TimeZoneOffset types.String              `tfsdk:"time_zone_offset"`
}

func (r *DestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination"
}

func (r *DestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Destination Resource",

		Attributes: map[string]schema.Attribute{
			"group_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"GCP_US_EAST4",
						"GCP_US_WEST1",
						"GCP_EUROPE_WEST3",
						"GCP_AUSTRALIA_SOUTHEAST1",
						"GCP_NORTHAMERICA_NORTHEAST1",
						"GCP_EUROPE_WEST2",
						"GCP_ASIA_SOUTHEAST1",
						"AWS_US_EAST_1",
						"AWS_US_EAST_2",
						"AWS_US_WEST_2",
						"AWS_AP_SOUTHEAST_2",
						"AWS_EU_CENTRAL_1",
						"AWS_EU_WEST_1",
						" AWS_EU_WEST_2",
						"AZURE_EASTUS2",
						"AZURE_AUSTRALIAEAST",
						"GCP_ASIA_SOUTH1",
					),
				},
				Description: `Data processing location. This is where Fivetran will operate and run computation on data.`,
			},
			"service": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"setup_status": schema.StringAttribute{
				Computed: true,
			},
			"setup_tests": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"details": schema.MapAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
							Validators: []validator.Map{
								mapvalidator.ValueStringsAre(validators.IsValidJSON()),
							},
						},
						"message": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"title": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"time_zone_offset": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *DestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationRequest := data.ToSDKType()
	request := operations.CreateDestinationRequest{
		DestinationRequest: destinationRequest,
	}
	res, err := r.client.DestinationManagement.CreateDestination(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	destinationID := data.ID.ValueString()
	getRequest := operations.DestinationDetailsRequest{
		DestinationID: destinationID,
	}
	getResponse, err := r.client.DestinationManagement.DestinationDetails(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.DestinationDetails200ApplicationJSONObject.Data == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSDKType(getResponse.DestinationDetails200ApplicationJSONObject.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.ID.ValueString()
	request := operations.DestinationDetailsRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.DestinationManagement.DestinationDetails(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationDetails200ApplicationJSONObject.Data == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSDKType(res.DestinationDetails200ApplicationJSONObject.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationRequest := data.ToSDKType()
	destinationID := data.ID.ValueString()
	request := operations.ModifyDestinationRequest{
		DestinationRequest: destinationRequest,
		DestinationID:      destinationID,
	}
	res, err := r.client.DestinationManagement.ModifyDestination(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	destinationId1 := data.ID.ValueString()
	getRequest := operations.DestinationDetailsRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.DestinationManagement.DestinationDetails(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.DestinationDetails200ApplicationJSONObject.Data == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSDKType(getResponse.DestinationDetails200ApplicationJSONObject.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.ID.ValueString()
	request := operations.DeleteDestinationRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.DestinationManagement.DeleteDestination(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}