package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource = (*timeNowDataSource)(nil)
)

func NewTimeNowDataSource() datasource.DataSource {
	return &timeNowDataSource{}
}

type timeNowDataSource struct{}

func (t timeNowDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_now"
}

func (t timeNowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Load and return the time at plan time.",
		Attributes: map[string]schema.Attribute{
			"day": schema.Int64Attribute{
				Description: "Number day of timestamp.",
				Computed:    true,
			},
			"hour": schema.Int64Attribute{
				Description: "Number hour of timestamp.",
				Computed:    true,
			},
			"minute": schema.Int64Attribute{
				Description: "Number minute of timestamp.",
				Computed:    true,
			},
			"month": schema.Int64Attribute{
				Description: "Number month of timestamp.",
				Computed:    true,
			},
			"rfc3339": schema.StringAttribute{
				Description: "Base timestamp in " +
					"[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format " +
					"(see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., " +
					"`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.",
				Computed: true,
			},
			"second": schema.Int64Attribute{
				Description: "Number second of timestamp.",
				Computed:    true,
			},
			"unix": schema.Int64Attribute{
				Description: "Number of seconds since epoch time, e.g. `1581489373`.",
				Computed:    true,
			},
			"year": schema.Int64Attribute{
				Description: "Number year of timestamp.",
				Computed:    true,
			},
		},
	}
}

func (t timeNowDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	timestamp := time.Now().UTC()
	state := timeDataModelV0{
		Year:    types.Int64Value(int64(timestamp.Year())),
		Month:   types.Int64Value(int64(timestamp.Month())),
		Day:     types.Int64Value(int64(timestamp.Day())),
		Hour:    types.Int64Value(int64(timestamp.Hour())),
		Minute:  types.Int64Value(int64(timestamp.Minute())),
		Second:  types.Int64Value(int64(timestamp.Second())),
		RFC3339: types.StringValue(timestamp.Format(time.RFC3339)),
		Unix:    types.Int64Value(timestamp.Unix()),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

type timeDataModelV0 struct {
	Day     types.Int64  `tfsdk:"day"`
	Hour    types.Int64  `tfsdk:"hour"`
	Minute  types.Int64  `tfsdk:"minute"`
	Month   types.Int64  `tfsdk:"month"`
	RFC3339 types.String `tfsdk:"rfc3339"`
	Second  types.Int64  `tfsdk:"second"`
	Unix    types.Int64  `tfsdk:"unix"`
	Year    types.Int64  `tfsdk:"year"`
}
