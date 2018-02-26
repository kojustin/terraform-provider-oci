// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"supports_pdb": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDbVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DbVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbVersionsResponse
}

func (s *DbVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbVersionsDataSourceCrud) Get() error {
	request := oci_database.ListDbVersionsRequest{}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemShape, ok := s.D.GetOk("db_system_shape"); ok {
		tmp := dbSystemShape.(string)
		request.DbSystemShape = &tmp
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOk("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	response, err := s.Client.ListDbVersions(context.Background(), request, getRetryOptions(false, "database")...)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DbVersionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbVersion := map[string]interface{}{}

		if r.SupportsPdb != nil {
			dbVersion["supports_pdb"] = *r.SupportsPdb
		}

		if r.Version != nil {
			dbVersion["version"] = *r.Version
		}

		resources = append(resources, dbVersion)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("db_versions", resources); err != nil {
		panic(err)
	}

	return
}