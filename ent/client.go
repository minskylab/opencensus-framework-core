// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"opencensus/core/ent/migrate"

	"opencensus/core/ent/bedrecord"
	"opencensus/core/ent/district"
	"opencensus/core/ent/organization"
	"opencensus/core/ent/oxygenrecord"
	"opencensus/core/ent/province"
	"opencensus/core/ent/region"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// BedRecord is the client for interacting with the BedRecord builders.
	BedRecord *BedRecordClient
	// District is the client for interacting with the District builders.
	District *DistrictClient
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
	// OxygenRecord is the client for interacting with the OxygenRecord builders.
	OxygenRecord *OxygenRecordClient
	// Province is the client for interacting with the Province builders.
	Province *ProvinceClient
	// Region is the client for interacting with the Region builders.
	Region *RegionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.BedRecord = NewBedRecordClient(c.config)
	c.District = NewDistrictClient(c.config)
	c.Organization = NewOrganizationClient(c.config)
	c.OxygenRecord = NewOxygenRecordClient(c.config)
	c.Province = NewProvinceClient(c.config)
	c.Region = NewRegionClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		BedRecord:    NewBedRecordClient(cfg),
		District:     NewDistrictClient(cfg),
		Organization: NewOrganizationClient(cfg),
		OxygenRecord: NewOxygenRecordClient(cfg),
		Province:     NewProvinceClient(cfg),
		Region:       NewRegionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:       cfg,
		BedRecord:    NewBedRecordClient(cfg),
		District:     NewDistrictClient(cfg),
		Organization: NewOrganizationClient(cfg),
		OxygenRecord: NewOxygenRecordClient(cfg),
		Province:     NewProvinceClient(cfg),
		Region:       NewRegionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		BedRecord.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.BedRecord.Use(hooks...)
	c.District.Use(hooks...)
	c.Organization.Use(hooks...)
	c.OxygenRecord.Use(hooks...)
	c.Province.Use(hooks...)
	c.Region.Use(hooks...)
}

// BedRecordClient is a client for the BedRecord schema.
type BedRecordClient struct {
	config
}

// NewBedRecordClient returns a client for the BedRecord from the given config.
func NewBedRecordClient(c config) *BedRecordClient {
	return &BedRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bedrecord.Hooks(f(g(h())))`.
func (c *BedRecordClient) Use(hooks ...Hook) {
	c.hooks.BedRecord = append(c.hooks.BedRecord, hooks...)
}

// Create returns a create builder for BedRecord.
func (c *BedRecordClient) Create() *BedRecordCreate {
	mutation := newBedRecordMutation(c.config, OpCreate)
	return &BedRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BedRecord entities.
func (c *BedRecordClient) CreateBulk(builders ...*BedRecordCreate) *BedRecordCreateBulk {
	return &BedRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BedRecord.
func (c *BedRecordClient) Update() *BedRecordUpdate {
	mutation := newBedRecordMutation(c.config, OpUpdate)
	return &BedRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BedRecordClient) UpdateOne(br *BedRecord) *BedRecordUpdateOne {
	mutation := newBedRecordMutation(c.config, OpUpdateOne, withBedRecord(br))
	return &BedRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BedRecordClient) UpdateOneID(id int) *BedRecordUpdateOne {
	mutation := newBedRecordMutation(c.config, OpUpdateOne, withBedRecordID(id))
	return &BedRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BedRecord.
func (c *BedRecordClient) Delete() *BedRecordDelete {
	mutation := newBedRecordMutation(c.config, OpDelete)
	return &BedRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BedRecordClient) DeleteOne(br *BedRecord) *BedRecordDeleteOne {
	return c.DeleteOneID(br.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BedRecordClient) DeleteOneID(id int) *BedRecordDeleteOne {
	builder := c.Delete().Where(bedrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BedRecordDeleteOne{builder}
}

// Query returns a query builder for BedRecord.
func (c *BedRecordClient) Query() *BedRecordQuery {
	return &BedRecordQuery{config: c.config}
}

// Get returns a BedRecord entity by its id.
func (c *BedRecordClient) Get(ctx context.Context, id int) (*BedRecord, error) {
	return c.Query().Where(bedrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BedRecordClient) GetX(ctx context.Context, id int) *BedRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a BedRecord.
func (c *BedRecordClient) QueryOrganization(br *BedRecord) *OrganizationQuery {
	query := &OrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := br.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bedrecord.Table, bedrecord.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, bedrecord.OrganizationTable, bedrecord.OrganizationPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(br.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BedRecordClient) Hooks() []Hook {
	return c.hooks.BedRecord
}

// DistrictClient is a client for the District schema.
type DistrictClient struct {
	config
}

// NewDistrictClient returns a client for the District from the given config.
func NewDistrictClient(c config) *DistrictClient {
	return &DistrictClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `district.Hooks(f(g(h())))`.
func (c *DistrictClient) Use(hooks ...Hook) {
	c.hooks.District = append(c.hooks.District, hooks...)
}

// Create returns a create builder for District.
func (c *DistrictClient) Create() *DistrictCreate {
	mutation := newDistrictMutation(c.config, OpCreate)
	return &DistrictCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of District entities.
func (c *DistrictClient) CreateBulk(builders ...*DistrictCreate) *DistrictCreateBulk {
	return &DistrictCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for District.
func (c *DistrictClient) Update() *DistrictUpdate {
	mutation := newDistrictMutation(c.config, OpUpdate)
	return &DistrictUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DistrictClient) UpdateOne(d *District) *DistrictUpdateOne {
	mutation := newDistrictMutation(c.config, OpUpdateOne, withDistrict(d))
	return &DistrictUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DistrictClient) UpdateOneID(id int) *DistrictUpdateOne {
	mutation := newDistrictMutation(c.config, OpUpdateOne, withDistrictID(id))
	return &DistrictUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for District.
func (c *DistrictClient) Delete() *DistrictDelete {
	mutation := newDistrictMutation(c.config, OpDelete)
	return &DistrictDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DistrictClient) DeleteOne(d *District) *DistrictDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DistrictClient) DeleteOneID(id int) *DistrictDeleteOne {
	builder := c.Delete().Where(district.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DistrictDeleteOne{builder}
}

// Query returns a query builder for District.
func (c *DistrictClient) Query() *DistrictQuery {
	return &DistrictQuery{config: c.config}
}

// Get returns a District entity by its id.
func (c *DistrictClient) Get(ctx context.Context, id int) (*District, error) {
	return c.Query().Where(district.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DistrictClient) GetX(ctx context.Context, id int) *District {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a District.
func (c *DistrictClient) QueryOrganization(d *District) *OrganizationQuery {
	query := &OrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(district.Table, district.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, district.OrganizationTable, district.OrganizationPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvince queries the province edge of a District.
func (c *DistrictClient) QueryProvince(d *District) *ProvinceQuery {
	query := &ProvinceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(district.Table, district.FieldID, id),
			sqlgraph.To(province.Table, province.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, district.ProvinceTable, district.ProvincePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DistrictClient) Hooks() []Hook {
	return c.hooks.District
}

// OrganizationClient is a client for the Organization schema.
type OrganizationClient struct {
	config
}

// NewOrganizationClient returns a client for the Organization from the given config.
func NewOrganizationClient(c config) *OrganizationClient {
	return &OrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organization.Hooks(f(g(h())))`.
func (c *OrganizationClient) Use(hooks ...Hook) {
	c.hooks.Organization = append(c.hooks.Organization, hooks...)
}

// Create returns a create builder for Organization.
func (c *OrganizationClient) Create() *OrganizationCreate {
	mutation := newOrganizationMutation(c.config, OpCreate)
	return &OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Organization entities.
func (c *OrganizationClient) CreateBulk(builders ...*OrganizationCreate) *OrganizationCreateBulk {
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Organization.
func (c *OrganizationClient) Update() *OrganizationUpdate {
	mutation := newOrganizationMutation(c.config, OpUpdate)
	return &OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationClient) UpdateOne(o *Organization) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganization(o))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationClient) UpdateOneID(id int) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganizationID(id))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Organization.
func (c *OrganizationClient) Delete() *OrganizationDelete {
	mutation := newOrganizationMutation(c.config, OpDelete)
	return &OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OrganizationClient) DeleteOne(o *Organization) *OrganizationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OrganizationClient) DeleteOneID(id int) *OrganizationDeleteOne {
	builder := c.Delete().Where(organization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationDeleteOne{builder}
}

// Query returns a query builder for Organization.
func (c *OrganizationClient) Query() *OrganizationQuery {
	return &OrganizationQuery{config: c.config}
}

// Get returns a Organization entity by its id.
func (c *OrganizationClient) Get(ctx context.Context, id int) (*Organization, error) {
	return c.Query().Where(organization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationClient) GetX(ctx context.Context, id int) *Organization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRegion queries the region edge of a Organization.
func (c *OrganizationClient) QueryRegion(o *Organization) *RegionQuery {
	query := &RegionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.RegionTable, organization.RegionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvince queries the province edge of a Organization.
func (c *OrganizationClient) QueryProvince(o *Organization) *ProvinceQuery {
	query := &ProvinceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(province.Table, province.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.ProvinceTable, organization.ProvincePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDistrict queries the district edge of a Organization.
func (c *OrganizationClient) QueryDistrict(o *Organization) *DistrictQuery {
	query := &DistrictQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(district.Table, district.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.DistrictTable, organization.DistrictPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOxygenRecords queries the oxygenRecords edge of a Organization.
func (c *OrganizationClient) QueryOxygenRecords(o *Organization) *OxygenRecordQuery {
	query := &OxygenRecordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(oxygenrecord.Table, oxygenrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.OxygenRecordsTable, organization.OxygenRecordsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBedRecords queries the bedRecords edge of a Organization.
func (c *OrganizationClient) QueryBedRecords(o *Organization) *BedRecordQuery {
	query := &BedRecordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(bedrecord.Table, bedrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organization.BedRecordsTable, organization.BedRecordsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrganizationClient) Hooks() []Hook {
	return c.hooks.Organization
}

// OxygenRecordClient is a client for the OxygenRecord schema.
type OxygenRecordClient struct {
	config
}

// NewOxygenRecordClient returns a client for the OxygenRecord from the given config.
func NewOxygenRecordClient(c config) *OxygenRecordClient {
	return &OxygenRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `oxygenrecord.Hooks(f(g(h())))`.
func (c *OxygenRecordClient) Use(hooks ...Hook) {
	c.hooks.OxygenRecord = append(c.hooks.OxygenRecord, hooks...)
}

// Create returns a create builder for OxygenRecord.
func (c *OxygenRecordClient) Create() *OxygenRecordCreate {
	mutation := newOxygenRecordMutation(c.config, OpCreate)
	return &OxygenRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OxygenRecord entities.
func (c *OxygenRecordClient) CreateBulk(builders ...*OxygenRecordCreate) *OxygenRecordCreateBulk {
	return &OxygenRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OxygenRecord.
func (c *OxygenRecordClient) Update() *OxygenRecordUpdate {
	mutation := newOxygenRecordMutation(c.config, OpUpdate)
	return &OxygenRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OxygenRecordClient) UpdateOne(or *OxygenRecord) *OxygenRecordUpdateOne {
	mutation := newOxygenRecordMutation(c.config, OpUpdateOne, withOxygenRecord(or))
	return &OxygenRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OxygenRecordClient) UpdateOneID(id int) *OxygenRecordUpdateOne {
	mutation := newOxygenRecordMutation(c.config, OpUpdateOne, withOxygenRecordID(id))
	return &OxygenRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OxygenRecord.
func (c *OxygenRecordClient) Delete() *OxygenRecordDelete {
	mutation := newOxygenRecordMutation(c.config, OpDelete)
	return &OxygenRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OxygenRecordClient) DeleteOne(or *OxygenRecord) *OxygenRecordDeleteOne {
	return c.DeleteOneID(or.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OxygenRecordClient) DeleteOneID(id int) *OxygenRecordDeleteOne {
	builder := c.Delete().Where(oxygenrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OxygenRecordDeleteOne{builder}
}

// Query returns a query builder for OxygenRecord.
func (c *OxygenRecordClient) Query() *OxygenRecordQuery {
	return &OxygenRecordQuery{config: c.config}
}

// Get returns a OxygenRecord entity by its id.
func (c *OxygenRecordClient) Get(ctx context.Context, id int) (*OxygenRecord, error) {
	return c.Query().Where(oxygenrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OxygenRecordClient) GetX(ctx context.Context, id int) *OxygenRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a OxygenRecord.
func (c *OxygenRecordClient) QueryOrganization(or *OxygenRecord) *OrganizationQuery {
	query := &OrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := or.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(oxygenrecord.Table, oxygenrecord.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, oxygenrecord.OrganizationTable, oxygenrecord.OrganizationPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(or.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OxygenRecordClient) Hooks() []Hook {
	return c.hooks.OxygenRecord
}

// ProvinceClient is a client for the Province schema.
type ProvinceClient struct {
	config
}

// NewProvinceClient returns a client for the Province from the given config.
func NewProvinceClient(c config) *ProvinceClient {
	return &ProvinceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `province.Hooks(f(g(h())))`.
func (c *ProvinceClient) Use(hooks ...Hook) {
	c.hooks.Province = append(c.hooks.Province, hooks...)
}

// Create returns a create builder for Province.
func (c *ProvinceClient) Create() *ProvinceCreate {
	mutation := newProvinceMutation(c.config, OpCreate)
	return &ProvinceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Province entities.
func (c *ProvinceClient) CreateBulk(builders ...*ProvinceCreate) *ProvinceCreateBulk {
	return &ProvinceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Province.
func (c *ProvinceClient) Update() *ProvinceUpdate {
	mutation := newProvinceMutation(c.config, OpUpdate)
	return &ProvinceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProvinceClient) UpdateOne(pr *Province) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvince(pr))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProvinceClient) UpdateOneID(id int) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvinceID(id))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Province.
func (c *ProvinceClient) Delete() *ProvinceDelete {
	mutation := newProvinceMutation(c.config, OpDelete)
	return &ProvinceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProvinceClient) DeleteOne(pr *Province) *ProvinceDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProvinceClient) DeleteOneID(id int) *ProvinceDeleteOne {
	builder := c.Delete().Where(province.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProvinceDeleteOne{builder}
}

// Query returns a query builder for Province.
func (c *ProvinceClient) Query() *ProvinceQuery {
	return &ProvinceQuery{config: c.config}
}

// Get returns a Province entity by its id.
func (c *ProvinceClient) Get(ctx context.Context, id int) (*Province, error) {
	return c.Query().Where(province.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProvinceClient) GetX(ctx context.Context, id int) *Province {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a Province.
func (c *ProvinceClient) QueryOrganization(pr *Province) *OrganizationQuery {
	query := &OrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(province.Table, province.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, province.OrganizationTable, province.OrganizationPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRegion queries the region edge of a Province.
func (c *ProvinceClient) QueryRegion(pr *Province) *RegionQuery {
	query := &RegionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(province.Table, province.FieldID, id),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, province.RegionTable, province.RegionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDistrict queries the district edge of a Province.
func (c *ProvinceClient) QueryDistrict(pr *Province) *DistrictQuery {
	query := &DistrictQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(province.Table, province.FieldID, id),
			sqlgraph.To(district.Table, district.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, province.DistrictTable, province.DistrictPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProvinceClient) Hooks() []Hook {
	return c.hooks.Province
}

// RegionClient is a client for the Region schema.
type RegionClient struct {
	config
}

// NewRegionClient returns a client for the Region from the given config.
func NewRegionClient(c config) *RegionClient {
	return &RegionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `region.Hooks(f(g(h())))`.
func (c *RegionClient) Use(hooks ...Hook) {
	c.hooks.Region = append(c.hooks.Region, hooks...)
}

// Create returns a create builder for Region.
func (c *RegionClient) Create() *RegionCreate {
	mutation := newRegionMutation(c.config, OpCreate)
	return &RegionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Region entities.
func (c *RegionClient) CreateBulk(builders ...*RegionCreate) *RegionCreateBulk {
	return &RegionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Region.
func (c *RegionClient) Update() *RegionUpdate {
	mutation := newRegionMutation(c.config, OpUpdate)
	return &RegionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegionClient) UpdateOne(r *Region) *RegionUpdateOne {
	mutation := newRegionMutation(c.config, OpUpdateOne, withRegion(r))
	return &RegionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegionClient) UpdateOneID(id int) *RegionUpdateOne {
	mutation := newRegionMutation(c.config, OpUpdateOne, withRegionID(id))
	return &RegionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Region.
func (c *RegionClient) Delete() *RegionDelete {
	mutation := newRegionMutation(c.config, OpDelete)
	return &RegionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RegionClient) DeleteOne(r *Region) *RegionDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RegionClient) DeleteOneID(id int) *RegionDeleteOne {
	builder := c.Delete().Where(region.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegionDeleteOne{builder}
}

// Query returns a query builder for Region.
func (c *RegionClient) Query() *RegionQuery {
	return &RegionQuery{config: c.config}
}

// Get returns a Region entity by its id.
func (c *RegionClient) Get(ctx context.Context, id int) (*Region, error) {
	return c.Query().Where(region.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegionClient) GetX(ctx context.Context, id int) *Region {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a Region.
func (c *RegionClient) QueryOrganization(r *Region) *OrganizationQuery {
	query := &OrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(region.Table, region.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, region.OrganizationTable, region.OrganizationPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvince queries the province edge of a Region.
func (c *RegionClient) QueryProvince(r *Region) *ProvinceQuery {
	query := &ProvinceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(region.Table, region.FieldID, id),
			sqlgraph.To(province.Table, province.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, region.ProvinceTable, region.ProvincePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RegionClient) Hooks() []Hook {
	return c.hooks.Region
}
