// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"kratos-uba/app/core/service/internal/data/ent/migrate"

	"kratos-uba/app/core/service/internal/data/ent/application"
	"kratos-uba/app/core/service/internal/data/ent/attribute"
	"kratos-uba/app/core/service/internal/data/ent/debugdevice"
	"kratos-uba/app/core/service/internal/data/ent/metaevent"
	"kratos-uba/app/core/service/internal/data/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Application is the client for interacting with the Application builders.
	Application *ApplicationClient
	// Attribute is the client for interacting with the Attribute builders.
	Attribute *AttributeClient
	// DebugDevice is the client for interacting with the DebugDevice builders.
	DebugDevice *DebugDeviceClient
	// MetaEvent is the client for interacting with the MetaEvent builders.
	MetaEvent *MetaEventClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Application = NewApplicationClient(c.config)
	c.Attribute = NewAttributeClient(c.config)
	c.DebugDevice = NewDebugDeviceClient(c.config)
	c.MetaEvent = NewMetaEventClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:         ctx,
		config:      cfg,
		Application: NewApplicationClient(cfg),
		Attribute:   NewAttributeClient(cfg),
		DebugDevice: NewDebugDeviceClient(cfg),
		MetaEvent:   NewMetaEventClient(cfg),
		User:        NewUserClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Application: NewApplicationClient(cfg),
		Attribute:   NewAttributeClient(cfg),
		DebugDevice: NewDebugDeviceClient(cfg),
		MetaEvent:   NewMetaEventClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Application.
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
	c.Application.Use(hooks...)
	c.Attribute.Use(hooks...)
	c.DebugDevice.Use(hooks...)
	c.MetaEvent.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Application.Intercept(interceptors...)
	c.Attribute.Intercept(interceptors...)
	c.DebugDevice.Intercept(interceptors...)
	c.MetaEvent.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ApplicationMutation:
		return c.Application.mutate(ctx, m)
	case *AttributeMutation:
		return c.Attribute.mutate(ctx, m)
	case *DebugDeviceMutation:
		return c.DebugDevice.mutate(ctx, m)
	case *MetaEventMutation:
		return c.MetaEvent.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ApplicationClient is a client for the Application schema.
type ApplicationClient struct {
	config
}

// NewApplicationClient returns a client for the Application from the given config.
func NewApplicationClient(c config) *ApplicationClient {
	return &ApplicationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `application.Hooks(f(g(h())))`.
func (c *ApplicationClient) Use(hooks ...Hook) {
	c.hooks.Application = append(c.hooks.Application, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `application.Intercept(f(g(h())))`.
func (c *ApplicationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Application = append(c.inters.Application, interceptors...)
}

// Create returns a builder for creating a Application entity.
func (c *ApplicationClient) Create() *ApplicationCreate {
	mutation := newApplicationMutation(c.config, OpCreate)
	return &ApplicationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Application entities.
func (c *ApplicationClient) CreateBulk(builders ...*ApplicationCreate) *ApplicationCreateBulk {
	return &ApplicationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ApplicationClient) MapCreateBulk(slice any, setFunc func(*ApplicationCreate, int)) *ApplicationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ApplicationCreateBulk{err: fmt.Errorf("calling to ApplicationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ApplicationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ApplicationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Application.
func (c *ApplicationClient) Update() *ApplicationUpdate {
	mutation := newApplicationMutation(c.config, OpUpdate)
	return &ApplicationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApplicationClient) UpdateOne(a *Application) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplication(a))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApplicationClient) UpdateOneID(id uint32) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplicationID(id))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Application.
func (c *ApplicationClient) Delete() *ApplicationDelete {
	mutation := newApplicationMutation(c.config, OpDelete)
	return &ApplicationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ApplicationClient) DeleteOne(a *Application) *ApplicationDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ApplicationClient) DeleteOneID(id uint32) *ApplicationDeleteOne {
	builder := c.Delete().Where(application.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApplicationDeleteOne{builder}
}

// Query returns a query builder for Application.
func (c *ApplicationClient) Query() *ApplicationQuery {
	return &ApplicationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeApplication},
		inters: c.Interceptors(),
	}
}

// Get returns a Application entity by its id.
func (c *ApplicationClient) Get(ctx context.Context, id uint32) (*Application, error) {
	return c.Query().Where(application.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApplicationClient) GetX(ctx context.Context, id uint32) *Application {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ApplicationClient) Hooks() []Hook {
	return c.hooks.Application
}

// Interceptors returns the client interceptors.
func (c *ApplicationClient) Interceptors() []Interceptor {
	return c.inters.Application
}

func (c *ApplicationClient) mutate(ctx context.Context, m *ApplicationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ApplicationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ApplicationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ApplicationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Application mutation op: %q", m.Op())
	}
}

// AttributeClient is a client for the Attribute schema.
type AttributeClient struct {
	config
}

// NewAttributeClient returns a client for the Attribute from the given config.
func NewAttributeClient(c config) *AttributeClient {
	return &AttributeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attribute.Hooks(f(g(h())))`.
func (c *AttributeClient) Use(hooks ...Hook) {
	c.hooks.Attribute = append(c.hooks.Attribute, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `attribute.Intercept(f(g(h())))`.
func (c *AttributeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Attribute = append(c.inters.Attribute, interceptors...)
}

// Create returns a builder for creating a Attribute entity.
func (c *AttributeClient) Create() *AttributeCreate {
	mutation := newAttributeMutation(c.config, OpCreate)
	return &AttributeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Attribute entities.
func (c *AttributeClient) CreateBulk(builders ...*AttributeCreate) *AttributeCreateBulk {
	return &AttributeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AttributeClient) MapCreateBulk(slice any, setFunc func(*AttributeCreate, int)) *AttributeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AttributeCreateBulk{err: fmt.Errorf("calling to AttributeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AttributeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AttributeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Attribute.
func (c *AttributeClient) Update() *AttributeUpdate {
	mutation := newAttributeMutation(c.config, OpUpdate)
	return &AttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttributeClient) UpdateOne(a *Attribute) *AttributeUpdateOne {
	mutation := newAttributeMutation(c.config, OpUpdateOne, withAttribute(a))
	return &AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttributeClient) UpdateOneID(id uint32) *AttributeUpdateOne {
	mutation := newAttributeMutation(c.config, OpUpdateOne, withAttributeID(id))
	return &AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Attribute.
func (c *AttributeClient) Delete() *AttributeDelete {
	mutation := newAttributeMutation(c.config, OpDelete)
	return &AttributeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AttributeClient) DeleteOne(a *Attribute) *AttributeDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AttributeClient) DeleteOneID(id uint32) *AttributeDeleteOne {
	builder := c.Delete().Where(attribute.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttributeDeleteOne{builder}
}

// Query returns a query builder for Attribute.
func (c *AttributeClient) Query() *AttributeQuery {
	return &AttributeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAttribute},
		inters: c.Interceptors(),
	}
}

// Get returns a Attribute entity by its id.
func (c *AttributeClient) Get(ctx context.Context, id uint32) (*Attribute, error) {
	return c.Query().Where(attribute.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttributeClient) GetX(ctx context.Context, id uint32) *Attribute {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AttributeClient) Hooks() []Hook {
	return c.hooks.Attribute
}

// Interceptors returns the client interceptors.
func (c *AttributeClient) Interceptors() []Interceptor {
	return c.inters.Attribute
}

func (c *AttributeClient) mutate(ctx context.Context, m *AttributeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AttributeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AttributeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Attribute mutation op: %q", m.Op())
	}
}

// DebugDeviceClient is a client for the DebugDevice schema.
type DebugDeviceClient struct {
	config
}

// NewDebugDeviceClient returns a client for the DebugDevice from the given config.
func NewDebugDeviceClient(c config) *DebugDeviceClient {
	return &DebugDeviceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `debugdevice.Hooks(f(g(h())))`.
func (c *DebugDeviceClient) Use(hooks ...Hook) {
	c.hooks.DebugDevice = append(c.hooks.DebugDevice, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `debugdevice.Intercept(f(g(h())))`.
func (c *DebugDeviceClient) Intercept(interceptors ...Interceptor) {
	c.inters.DebugDevice = append(c.inters.DebugDevice, interceptors...)
}

// Create returns a builder for creating a DebugDevice entity.
func (c *DebugDeviceClient) Create() *DebugDeviceCreate {
	mutation := newDebugDeviceMutation(c.config, OpCreate)
	return &DebugDeviceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DebugDevice entities.
func (c *DebugDeviceClient) CreateBulk(builders ...*DebugDeviceCreate) *DebugDeviceCreateBulk {
	return &DebugDeviceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DebugDeviceClient) MapCreateBulk(slice any, setFunc func(*DebugDeviceCreate, int)) *DebugDeviceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DebugDeviceCreateBulk{err: fmt.Errorf("calling to DebugDeviceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DebugDeviceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DebugDeviceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DebugDevice.
func (c *DebugDeviceClient) Update() *DebugDeviceUpdate {
	mutation := newDebugDeviceMutation(c.config, OpUpdate)
	return &DebugDeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DebugDeviceClient) UpdateOne(dd *DebugDevice) *DebugDeviceUpdateOne {
	mutation := newDebugDeviceMutation(c.config, OpUpdateOne, withDebugDevice(dd))
	return &DebugDeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DebugDeviceClient) UpdateOneID(id uint32) *DebugDeviceUpdateOne {
	mutation := newDebugDeviceMutation(c.config, OpUpdateOne, withDebugDeviceID(id))
	return &DebugDeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DebugDevice.
func (c *DebugDeviceClient) Delete() *DebugDeviceDelete {
	mutation := newDebugDeviceMutation(c.config, OpDelete)
	return &DebugDeviceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DebugDeviceClient) DeleteOne(dd *DebugDevice) *DebugDeviceDeleteOne {
	return c.DeleteOneID(dd.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DebugDeviceClient) DeleteOneID(id uint32) *DebugDeviceDeleteOne {
	builder := c.Delete().Where(debugdevice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DebugDeviceDeleteOne{builder}
}

// Query returns a query builder for DebugDevice.
func (c *DebugDeviceClient) Query() *DebugDeviceQuery {
	return &DebugDeviceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDebugDevice},
		inters: c.Interceptors(),
	}
}

// Get returns a DebugDevice entity by its id.
func (c *DebugDeviceClient) Get(ctx context.Context, id uint32) (*DebugDevice, error) {
	return c.Query().Where(debugdevice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DebugDeviceClient) GetX(ctx context.Context, id uint32) *DebugDevice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DebugDeviceClient) Hooks() []Hook {
	return c.hooks.DebugDevice
}

// Interceptors returns the client interceptors.
func (c *DebugDeviceClient) Interceptors() []Interceptor {
	return c.inters.DebugDevice
}

func (c *DebugDeviceClient) mutate(ctx context.Context, m *DebugDeviceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DebugDeviceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DebugDeviceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DebugDeviceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DebugDeviceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown DebugDevice mutation op: %q", m.Op())
	}
}

// MetaEventClient is a client for the MetaEvent schema.
type MetaEventClient struct {
	config
}

// NewMetaEventClient returns a client for the MetaEvent from the given config.
func NewMetaEventClient(c config) *MetaEventClient {
	return &MetaEventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `metaevent.Hooks(f(g(h())))`.
func (c *MetaEventClient) Use(hooks ...Hook) {
	c.hooks.MetaEvent = append(c.hooks.MetaEvent, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `metaevent.Intercept(f(g(h())))`.
func (c *MetaEventClient) Intercept(interceptors ...Interceptor) {
	c.inters.MetaEvent = append(c.inters.MetaEvent, interceptors...)
}

// Create returns a builder for creating a MetaEvent entity.
func (c *MetaEventClient) Create() *MetaEventCreate {
	mutation := newMetaEventMutation(c.config, OpCreate)
	return &MetaEventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MetaEvent entities.
func (c *MetaEventClient) CreateBulk(builders ...*MetaEventCreate) *MetaEventCreateBulk {
	return &MetaEventCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MetaEventClient) MapCreateBulk(slice any, setFunc func(*MetaEventCreate, int)) *MetaEventCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MetaEventCreateBulk{err: fmt.Errorf("calling to MetaEventClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MetaEventCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MetaEventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MetaEvent.
func (c *MetaEventClient) Update() *MetaEventUpdate {
	mutation := newMetaEventMutation(c.config, OpUpdate)
	return &MetaEventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MetaEventClient) UpdateOne(me *MetaEvent) *MetaEventUpdateOne {
	mutation := newMetaEventMutation(c.config, OpUpdateOne, withMetaEvent(me))
	return &MetaEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MetaEventClient) UpdateOneID(id uint32) *MetaEventUpdateOne {
	mutation := newMetaEventMutation(c.config, OpUpdateOne, withMetaEventID(id))
	return &MetaEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MetaEvent.
func (c *MetaEventClient) Delete() *MetaEventDelete {
	mutation := newMetaEventMutation(c.config, OpDelete)
	return &MetaEventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MetaEventClient) DeleteOne(me *MetaEvent) *MetaEventDeleteOne {
	return c.DeleteOneID(me.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MetaEventClient) DeleteOneID(id uint32) *MetaEventDeleteOne {
	builder := c.Delete().Where(metaevent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MetaEventDeleteOne{builder}
}

// Query returns a query builder for MetaEvent.
func (c *MetaEventClient) Query() *MetaEventQuery {
	return &MetaEventQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMetaEvent},
		inters: c.Interceptors(),
	}
}

// Get returns a MetaEvent entity by its id.
func (c *MetaEventClient) Get(ctx context.Context, id uint32) (*MetaEvent, error) {
	return c.Query().Where(metaevent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MetaEventClient) GetX(ctx context.Context, id uint32) *MetaEvent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MetaEventClient) Hooks() []Hook {
	return c.hooks.MetaEvent
}

// Interceptors returns the client interceptors.
func (c *MetaEventClient) Interceptors() []Interceptor {
	return c.inters.MetaEvent
}

func (c *MetaEventClient) mutate(ctx context.Context, m *MetaEventMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MetaEventCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MetaEventUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MetaEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MetaEventDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MetaEvent mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Application, Attribute, DebugDevice, MetaEvent, User []ent.Hook
	}
	inters struct {
		Application, Attribute, DebugDevice, MetaEvent, User []ent.Interceptor
	}
)
