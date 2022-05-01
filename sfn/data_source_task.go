package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SFNTask struct {
	Type       string                 `json:"Type"`
	Name       string                 `json:"Name"`
	Comment    string                 `json:"Comment,omitempty"`
	Resource   string									`json:"Resoutce"`
	Next       string                 `json:"Next"`
	InputPath  string                 `json:"InputPath,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty"`
	Result		 map[string]interface{} `json:"Result,omitempty"`
	ResultPath string                 `json:"ResultPath,omitempty"`
	OutputPath string                 `json:"OutputPath,omitempty"`
}

func dataSourceTask() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"comment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Pass Step",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"resource": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"next": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"inputpath": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"result": {
				Type: schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"resultpath": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"outputpath": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"step": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	taskStep := &SFNTask{}

	taskStep.Type = "Pass"

	if comment, ok := d.GetOk("comment"); ok {
		taskStep.Comment = comment.(string)
	}

	if name, ok := d.GetOk("name"); ok {
		taskStep.Name = name.(string)
	}

	if resource, ok := d.GetOk("resource"); ok {
		taskStep.Resource = resource.(string)
	}

	if next, ok := d.GetOk("next"); ok {
		taskStep.Next = next.(string)
	}

	if inputpath, ok := d.GetOk("inputpath"); ok {
		taskStep.InputPath = inputpath.(string)
	}

	if parameters, ok := d.GetOk("parameters"); ok {
		taskStep.Parameters = parameters.(map[string]interface{})
	}

	if result, ok := d.GetOk("result"); ok {
		taskStep.Result = result.(map[string]interface{})
	}

	if resultpath, ok := d.GetOk("resultpath"); ok {
		taskStep.ResultPath = resultpath.(string)
	}

	if outputpath, ok := d.GetOk("outputpath"); ok {
		taskStep.OutputPath = outputpath.(string)
	}

	jsonDoc, err := json.MarshalIndent(taskStep, "", "  ")

	if err != nil {
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	d.Set("step", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}