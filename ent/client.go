// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/tonimaru/try-ent/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/tonimaru/try-ent/ent/fulltxt"
	entgeo "github.com/tonimaru/try-ent/ent/geo"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Fulltxt is the client for interacting with the Fulltxt builders.
	Fulltxt *FulltxtClient
	// Geo is the client for interacting with the Geo builders.
	Geo *GeoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Fulltxt = NewFulltxtClient(c.config)
	c.Geo = NewGeoClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Fulltxt: NewFulltxtClient(cfg),
		Geo:     NewGeoClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Fulltxt: NewFulltxtClient(cfg),
		Geo:     NewGeoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Fulltxt.
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
	c.Fulltxt.Use(hooks...)
	c.Geo.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Fulltxt.Intercept(interceptors...)
	c.Geo.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *FulltxtMutation:
		return c.Fulltxt.mutate(ctx, m)
	case *GeoMutation:
		return c.Geo.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// FulltxtClient is a client for the Fulltxt schema.
type FulltxtClient struct {
	config
}

// NewFulltxtClient returns a client for the Fulltxt from the given config.
func NewFulltxtClient(c config) *FulltxtClient {
	return &FulltxtClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fulltxt.Hooks(f(g(h())))`.
func (c *FulltxtClient) Use(hooks ...Hook) {
	c.hooks.Fulltxt = append(c.hooks.Fulltxt, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `fulltxt.Intercept(f(g(h())))`.
func (c *FulltxtClient) Intercept(interceptors ...Interceptor) {
	c.inters.Fulltxt = append(c.inters.Fulltxt, interceptors...)
}

// Create returns a builder for creating a Fulltxt entity.
func (c *FulltxtClient) Create() *FulltxtCreate {
	mutation := newFulltxtMutation(c.config, OpCreate)
	return &FulltxtCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Fulltxt entities.
func (c *FulltxtClient) CreateBulk(builders ...*FulltxtCreate) *FulltxtCreateBulk {
	return &FulltxtCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Fulltxt.
func (c *FulltxtClient) Update() *FulltxtUpdate {
	mutation := newFulltxtMutation(c.config, OpUpdate)
	return &FulltxtUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FulltxtClient) UpdateOne(f *Fulltxt) *FulltxtUpdateOne {
	mutation := newFulltxtMutation(c.config, OpUpdateOne, withFulltxt(f))
	return &FulltxtUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FulltxtClient) UpdateOneID(id int) *FulltxtUpdateOne {
	mutation := newFulltxtMutation(c.config, OpUpdateOne, withFulltxtID(id))
	return &FulltxtUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Fulltxt.
func (c *FulltxtClient) Delete() *FulltxtDelete {
	mutation := newFulltxtMutation(c.config, OpDelete)
	return &FulltxtDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FulltxtClient) DeleteOne(f *Fulltxt) *FulltxtDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FulltxtClient) DeleteOneID(id int) *FulltxtDeleteOne {
	builder := c.Delete().Where(fulltxt.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FulltxtDeleteOne{builder}
}

// Query returns a query builder for Fulltxt.
func (c *FulltxtClient) Query() *FulltxtQuery {
	return &FulltxtQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFulltxt},
		inters: c.Interceptors(),
	}
}

// Get returns a Fulltxt entity by its id.
func (c *FulltxtClient) Get(ctx context.Context, id int) (*Fulltxt, error) {
	return c.Query().Where(fulltxt.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FulltxtClient) GetX(ctx context.Context, id int) *Fulltxt {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FulltxtClient) Hooks() []Hook {
	return c.hooks.Fulltxt
}

// Interceptors returns the client interceptors.
func (c *FulltxtClient) Interceptors() []Interceptor {
	return c.inters.Fulltxt
}

func (c *FulltxtClient) mutate(ctx context.Context, m *FulltxtMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FulltxtCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FulltxtUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FulltxtUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FulltxtDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Fulltxt mutation op: %q", m.Op())
	}
}

// GeoClient is a client for the Geo schema.
type GeoClient struct {
	config
}

// NewGeoClient returns a client for the Geo from the given config.
func NewGeoClient(c config) *GeoClient {
	return &GeoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entgeo.Hooks(f(g(h())))`.
func (c *GeoClient) Use(hooks ...Hook) {
	c.hooks.Geo = append(c.hooks.Geo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entgeo.Intercept(f(g(h())))`.
func (c *GeoClient) Intercept(interceptors ...Interceptor) {
	c.inters.Geo = append(c.inters.Geo, interceptors...)
}

// Create returns a builder for creating a Geo entity.
func (c *GeoClient) Create() *GeoCreate {
	mutation := newGeoMutation(c.config, OpCreate)
	return &GeoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Geo entities.
func (c *GeoClient) CreateBulk(builders ...*GeoCreate) *GeoCreateBulk {
	return &GeoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Geo.
func (c *GeoClient) Update() *GeoUpdate {
	mutation := newGeoMutation(c.config, OpUpdate)
	return &GeoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GeoClient) UpdateOne(ge *Geo) *GeoUpdateOne {
	mutation := newGeoMutation(c.config, OpUpdateOne, withGeo(ge))
	return &GeoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GeoClient) UpdateOneID(id int) *GeoUpdateOne {
	mutation := newGeoMutation(c.config, OpUpdateOne, withGeoID(id))
	return &GeoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Geo.
func (c *GeoClient) Delete() *GeoDelete {
	mutation := newGeoMutation(c.config, OpDelete)
	return &GeoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GeoClient) DeleteOne(ge *Geo) *GeoDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GeoClient) DeleteOneID(id int) *GeoDeleteOne {
	builder := c.Delete().Where(entgeo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GeoDeleteOne{builder}
}

// Query returns a query builder for Geo.
func (c *GeoClient) Query() *GeoQuery {
	return &GeoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGeo},
		inters: c.Interceptors(),
	}
}

// Get returns a Geo entity by its id.
func (c *GeoClient) Get(ctx context.Context, id int) (*Geo, error) {
	return c.Query().Where(entgeo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GeoClient) GetX(ctx context.Context, id int) *Geo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GeoClient) Hooks() []Hook {
	return c.hooks.Geo
}

// Interceptors returns the client interceptors.
func (c *GeoClient) Interceptors() []Interceptor {
	return c.inters.Geo
}

func (c *GeoClient) mutate(ctx context.Context, m *GeoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GeoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GeoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GeoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GeoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Geo mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Fulltxt, Geo []ent.Hook
	}
	inters struct {
		Fulltxt, Geo []ent.Interceptor
	}
)