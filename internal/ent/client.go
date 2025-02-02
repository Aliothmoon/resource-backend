// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/MirrorChyan/resource-backend/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/MirrorChyan/resource-backend/internal/ent/latestversion"
	"github.com/MirrorChyan/resource-backend/internal/ent/resource"
	"github.com/MirrorChyan/resource-backend/internal/ent/storage"
	"github.com/MirrorChyan/resource-backend/internal/ent/version"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LatestVersion is the client for interacting with the LatestVersion builders.
	LatestVersion *LatestVersionClient
	// Resource is the client for interacting with the Resource builders.
	Resource *ResourceClient
	// Storage is the client for interacting with the Storage builders.
	Storage *StorageClient
	// Version is the client for interacting with the Version builders.
	Version *VersionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.LatestVersion = NewLatestVersionClient(c.config)
	c.Resource = NewResourceClient(c.config)
	c.Storage = NewStorageClient(c.config)
	c.Version = NewVersionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		LatestVersion: NewLatestVersionClient(cfg),
		Resource:      NewResourceClient(cfg),
		Storage:       NewStorageClient(cfg),
		Version:       NewVersionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		LatestVersion: NewLatestVersionClient(cfg),
		Resource:      NewResourceClient(cfg),
		Storage:       NewStorageClient(cfg),
		Version:       NewVersionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LatestVersion.
//		Query().
//		Count(ctx)
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
	c.LatestVersion.Use(hooks...)
	c.Resource.Use(hooks...)
	c.Storage.Use(hooks...)
	c.Version.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.LatestVersion.Intercept(interceptors...)
	c.Resource.Intercept(interceptors...)
	c.Storage.Intercept(interceptors...)
	c.Version.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LatestVersionMutation:
		return c.LatestVersion.mutate(ctx, m)
	case *ResourceMutation:
		return c.Resource.mutate(ctx, m)
	case *StorageMutation:
		return c.Storage.mutate(ctx, m)
	case *VersionMutation:
		return c.Version.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LatestVersionClient is a client for the LatestVersion schema.
type LatestVersionClient struct {
	config
}

// NewLatestVersionClient returns a client for the LatestVersion from the given config.
func NewLatestVersionClient(c config) *LatestVersionClient {
	return &LatestVersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `latestversion.Hooks(f(g(h())))`.
func (c *LatestVersionClient) Use(hooks ...Hook) {
	c.hooks.LatestVersion = append(c.hooks.LatestVersion, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `latestversion.Intercept(f(g(h())))`.
func (c *LatestVersionClient) Intercept(interceptors ...Interceptor) {
	c.inters.LatestVersion = append(c.inters.LatestVersion, interceptors...)
}

// Create returns a builder for creating a LatestVersion entity.
func (c *LatestVersionClient) Create() *LatestVersionCreate {
	mutation := newLatestVersionMutation(c.config, OpCreate)
	return &LatestVersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LatestVersion entities.
func (c *LatestVersionClient) CreateBulk(builders ...*LatestVersionCreate) *LatestVersionCreateBulk {
	return &LatestVersionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LatestVersionClient) MapCreateBulk(slice any, setFunc func(*LatestVersionCreate, int)) *LatestVersionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LatestVersionCreateBulk{err: fmt.Errorf("calling to LatestVersionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LatestVersionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LatestVersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LatestVersion.
func (c *LatestVersionClient) Update() *LatestVersionUpdate {
	mutation := newLatestVersionMutation(c.config, OpUpdate)
	return &LatestVersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LatestVersionClient) UpdateOne(lv *LatestVersion) *LatestVersionUpdateOne {
	mutation := newLatestVersionMutation(c.config, OpUpdateOne, withLatestVersion(lv))
	return &LatestVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LatestVersionClient) UpdateOneID(id int) *LatestVersionUpdateOne {
	mutation := newLatestVersionMutation(c.config, OpUpdateOne, withLatestVersionID(id))
	return &LatestVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LatestVersion.
func (c *LatestVersionClient) Delete() *LatestVersionDelete {
	mutation := newLatestVersionMutation(c.config, OpDelete)
	return &LatestVersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LatestVersionClient) DeleteOne(lv *LatestVersion) *LatestVersionDeleteOne {
	return c.DeleteOneID(lv.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LatestVersionClient) DeleteOneID(id int) *LatestVersionDeleteOne {
	builder := c.Delete().Where(latestversion.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LatestVersionDeleteOne{builder}
}

// Query returns a query builder for LatestVersion.
func (c *LatestVersionClient) Query() *LatestVersionQuery {
	return &LatestVersionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLatestVersion},
		inters: c.Interceptors(),
	}
}

// Get returns a LatestVersion entity by its id.
func (c *LatestVersionClient) Get(ctx context.Context, id int) (*LatestVersion, error) {
	return c.Query().Where(latestversion.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LatestVersionClient) GetX(ctx context.Context, id int) *LatestVersion {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResource queries the resource edge of a LatestVersion.
func (c *LatestVersionClient) QueryResource(lv *LatestVersion) *ResourceQuery {
	query := (&ResourceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := lv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(latestversion.Table, latestversion.FieldID, id),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, latestversion.ResourceTable, latestversion.ResourceColumn),
		)
		fromV = sqlgraph.Neighbors(lv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVersion queries the version edge of a LatestVersion.
func (c *LatestVersionClient) QueryVersion(lv *LatestVersion) *VersionQuery {
	query := (&VersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := lv.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(latestversion.Table, latestversion.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, latestversion.VersionTable, latestversion.VersionColumn),
		)
		fromV = sqlgraph.Neighbors(lv.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LatestVersionClient) Hooks() []Hook {
	return c.hooks.LatestVersion
}

// Interceptors returns the client interceptors.
func (c *LatestVersionClient) Interceptors() []Interceptor {
	return c.inters.LatestVersion
}

func (c *LatestVersionClient) mutate(ctx context.Context, m *LatestVersionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LatestVersionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LatestVersionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LatestVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LatestVersionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LatestVersion mutation op: %q", m.Op())
	}
}

// ResourceClient is a client for the Resource schema.
type ResourceClient struct {
	config
}

// NewResourceClient returns a client for the Resource from the given config.
func NewResourceClient(c config) *ResourceClient {
	return &ResourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `resource.Hooks(f(g(h())))`.
func (c *ResourceClient) Use(hooks ...Hook) {
	c.hooks.Resource = append(c.hooks.Resource, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `resource.Intercept(f(g(h())))`.
func (c *ResourceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Resource = append(c.inters.Resource, interceptors...)
}

// Create returns a builder for creating a Resource entity.
func (c *ResourceClient) Create() *ResourceCreate {
	mutation := newResourceMutation(c.config, OpCreate)
	return &ResourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Resource entities.
func (c *ResourceClient) CreateBulk(builders ...*ResourceCreate) *ResourceCreateBulk {
	return &ResourceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ResourceClient) MapCreateBulk(slice any, setFunc func(*ResourceCreate, int)) *ResourceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ResourceCreateBulk{err: fmt.Errorf("calling to ResourceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ResourceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ResourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Resource.
func (c *ResourceClient) Update() *ResourceUpdate {
	mutation := newResourceMutation(c.config, OpUpdate)
	return &ResourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResourceClient) UpdateOne(r *Resource) *ResourceUpdateOne {
	mutation := newResourceMutation(c.config, OpUpdateOne, withResource(r))
	return &ResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResourceClient) UpdateOneID(id string) *ResourceUpdateOne {
	mutation := newResourceMutation(c.config, OpUpdateOne, withResourceID(id))
	return &ResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Resource.
func (c *ResourceClient) Delete() *ResourceDelete {
	mutation := newResourceMutation(c.config, OpDelete)
	return &ResourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ResourceClient) DeleteOne(r *Resource) *ResourceDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ResourceClient) DeleteOneID(id string) *ResourceDeleteOne {
	builder := c.Delete().Where(resource.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResourceDeleteOne{builder}
}

// Query returns a query builder for Resource.
func (c *ResourceClient) Query() *ResourceQuery {
	return &ResourceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeResource},
		inters: c.Interceptors(),
	}
}

// Get returns a Resource entity by its id.
func (c *ResourceClient) Get(ctx context.Context, id string) (*Resource, error) {
	return c.Query().Where(resource.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResourceClient) GetX(ctx context.Context, id string) *Resource {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVersions queries the versions edge of a Resource.
func (c *ResourceClient) QueryVersions(r *Resource) *VersionQuery {
	query := (&VersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, resource.VersionsTable, resource.VersionsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLatestVersions queries the latest_versions edge of a Resource.
func (c *ResourceClient) QueryLatestVersions(r *Resource) *LatestVersionQuery {
	query := (&LatestVersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, id),
			sqlgraph.To(latestversion.Table, latestversion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, resource.LatestVersionsTable, resource.LatestVersionsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ResourceClient) Hooks() []Hook {
	return c.hooks.Resource
}

// Interceptors returns the client interceptors.
func (c *ResourceClient) Interceptors() []Interceptor {
	return c.inters.Resource
}

func (c *ResourceClient) mutate(ctx context.Context, m *ResourceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ResourceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ResourceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ResourceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Resource mutation op: %q", m.Op())
	}
}

// StorageClient is a client for the Storage schema.
type StorageClient struct {
	config
}

// NewStorageClient returns a client for the Storage from the given config.
func NewStorageClient(c config) *StorageClient {
	return &StorageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `storage.Hooks(f(g(h())))`.
func (c *StorageClient) Use(hooks ...Hook) {
	c.hooks.Storage = append(c.hooks.Storage, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `storage.Intercept(f(g(h())))`.
func (c *StorageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Storage = append(c.inters.Storage, interceptors...)
}

// Create returns a builder for creating a Storage entity.
func (c *StorageClient) Create() *StorageCreate {
	mutation := newStorageMutation(c.config, OpCreate)
	return &StorageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Storage entities.
func (c *StorageClient) CreateBulk(builders ...*StorageCreate) *StorageCreateBulk {
	return &StorageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StorageClient) MapCreateBulk(slice any, setFunc func(*StorageCreate, int)) *StorageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StorageCreateBulk{err: fmt.Errorf("calling to StorageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StorageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StorageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Storage.
func (c *StorageClient) Update() *StorageUpdate {
	mutation := newStorageMutation(c.config, OpUpdate)
	return &StorageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StorageClient) UpdateOne(s *Storage) *StorageUpdateOne {
	mutation := newStorageMutation(c.config, OpUpdateOne, withStorage(s))
	return &StorageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StorageClient) UpdateOneID(id int) *StorageUpdateOne {
	mutation := newStorageMutation(c.config, OpUpdateOne, withStorageID(id))
	return &StorageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Storage.
func (c *StorageClient) Delete() *StorageDelete {
	mutation := newStorageMutation(c.config, OpDelete)
	return &StorageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StorageClient) DeleteOne(s *Storage) *StorageDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StorageClient) DeleteOneID(id int) *StorageDeleteOne {
	builder := c.Delete().Where(storage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StorageDeleteOne{builder}
}

// Query returns a query builder for Storage.
func (c *StorageClient) Query() *StorageQuery {
	return &StorageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStorage},
		inters: c.Interceptors(),
	}
}

// Get returns a Storage entity by its id.
func (c *StorageClient) Get(ctx context.Context, id int) (*Storage, error) {
	return c.Query().Where(storage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StorageClient) GetX(ctx context.Context, id int) *Storage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVersion queries the version edge of a Storage.
func (c *StorageClient) QueryVersion(s *Storage) *VersionQuery {
	query := (&VersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(storage.Table, storage.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, storage.VersionTable, storage.VersionColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOldVersion queries the old_version edge of a Storage.
func (c *StorageClient) QueryOldVersion(s *Storage) *VersionQuery {
	query := (&VersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(storage.Table, storage.FieldID, id),
			sqlgraph.To(version.Table, version.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, storage.OldVersionTable, storage.OldVersionColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StorageClient) Hooks() []Hook {
	return c.hooks.Storage
}

// Interceptors returns the client interceptors.
func (c *StorageClient) Interceptors() []Interceptor {
	return c.inters.Storage
}

func (c *StorageClient) mutate(ctx context.Context, m *StorageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StorageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StorageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StorageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StorageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Storage mutation op: %q", m.Op())
	}
}

// VersionClient is a client for the Version schema.
type VersionClient struct {
	config
}

// NewVersionClient returns a client for the Version from the given config.
func NewVersionClient(c config) *VersionClient {
	return &VersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `version.Hooks(f(g(h())))`.
func (c *VersionClient) Use(hooks ...Hook) {
	c.hooks.Version = append(c.hooks.Version, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `version.Intercept(f(g(h())))`.
func (c *VersionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Version = append(c.inters.Version, interceptors...)
}

// Create returns a builder for creating a Version entity.
func (c *VersionClient) Create() *VersionCreate {
	mutation := newVersionMutation(c.config, OpCreate)
	return &VersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Version entities.
func (c *VersionClient) CreateBulk(builders ...*VersionCreate) *VersionCreateBulk {
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *VersionClient) MapCreateBulk(slice any, setFunc func(*VersionCreate, int)) *VersionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &VersionCreateBulk{err: fmt.Errorf("calling to VersionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*VersionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Version.
func (c *VersionClient) Update() *VersionUpdate {
	mutation := newVersionMutation(c.config, OpUpdate)
	return &VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VersionClient) UpdateOne(v *Version) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersion(v))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VersionClient) UpdateOneID(id int) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersionID(id))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Version.
func (c *VersionClient) Delete() *VersionDelete {
	mutation := newVersionMutation(c.config, OpDelete)
	return &VersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VersionClient) DeleteOne(v *Version) *VersionDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VersionClient) DeleteOneID(id int) *VersionDeleteOne {
	builder := c.Delete().Where(version.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VersionDeleteOne{builder}
}

// Query returns a query builder for Version.
func (c *VersionClient) Query() *VersionQuery {
	return &VersionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVersion},
		inters: c.Interceptors(),
	}
}

// Get returns a Version entity by its id.
func (c *VersionClient) Get(ctx context.Context, id int) (*Version, error) {
	return c.Query().Where(version.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VersionClient) GetX(ctx context.Context, id int) *Version {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStorages queries the storages edge of a Version.
func (c *VersionClient) QueryStorages(v *Version) *StorageQuery {
	query := (&StorageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, id),
			sqlgraph.To(storage.Table, storage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, version.StoragesTable, version.StoragesColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryResource queries the resource edge of a Version.
func (c *VersionClient) QueryResource(v *Version) *ResourceQuery {
	query := (&ResourceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(version.Table, version.FieldID, id),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, version.ResourceTable, version.ResourceColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VersionClient) Hooks() []Hook {
	return c.hooks.Version
}

// Interceptors returns the client interceptors.
func (c *VersionClient) Interceptors() []Interceptor {
	return c.inters.Version
}

func (c *VersionClient) mutate(ctx context.Context, m *VersionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VersionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VersionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Version mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		LatestVersion, Resource, Storage, Version []ent.Hook
	}
	inters struct {
		LatestVersion, Resource, Storage, Version []ent.Interceptor
	}
)
