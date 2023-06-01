// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adshao/ordinals-indexer/internal/data/ent/migrate"

	"github.com/adshao/ordinals-indexer/internal/data/ent/collection"
	"github.com/adshao/ordinals-indexer/internal/data/ent/inscription"
	"github.com/adshao/ordinals-indexer/internal/data/ent/token"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Collection is the client for interacting with the Collection builders.
	Collection *CollectionClient
	// Inscription is the client for interacting with the Inscription builders.
	Inscription *InscriptionClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
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
	c.Collection = NewCollectionClient(c.config)
	c.Inscription = NewInscriptionClient(c.config)
	c.Token = NewTokenClient(c.config)
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
		ctx:         ctx,
		config:      cfg,
		Collection:  NewCollectionClient(cfg),
		Inscription: NewInscriptionClient(cfg),
		Token:       NewTokenClient(cfg),
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
		Collection:  NewCollectionClient(cfg),
		Inscription: NewInscriptionClient(cfg),
		Token:       NewTokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Collection.
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
	c.Collection.Use(hooks...)
	c.Inscription.Use(hooks...)
	c.Token.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Collection.Intercept(interceptors...)
	c.Inscription.Intercept(interceptors...)
	c.Token.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CollectionMutation:
		return c.Collection.mutate(ctx, m)
	case *InscriptionMutation:
		return c.Inscription.mutate(ctx, m)
	case *TokenMutation:
		return c.Token.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CollectionClient is a client for the Collection schema.
type CollectionClient struct {
	config
}

// NewCollectionClient returns a client for the Collection from the given config.
func NewCollectionClient(c config) *CollectionClient {
	return &CollectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `collection.Hooks(f(g(h())))`.
func (c *CollectionClient) Use(hooks ...Hook) {
	c.hooks.Collection = append(c.hooks.Collection, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `collection.Intercept(f(g(h())))`.
func (c *CollectionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Collection = append(c.inters.Collection, interceptors...)
}

// Create returns a builder for creating a Collection entity.
func (c *CollectionClient) Create() *CollectionCreate {
	mutation := newCollectionMutation(c.config, OpCreate)
	return &CollectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Collection entities.
func (c *CollectionClient) CreateBulk(builders ...*CollectionCreate) *CollectionCreateBulk {
	return &CollectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Collection.
func (c *CollectionClient) Update() *CollectionUpdate {
	mutation := newCollectionMutation(c.config, OpUpdate)
	return &CollectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CollectionClient) UpdateOne(co *Collection) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollection(co))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CollectionClient) UpdateOneID(id int) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollectionID(id))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Collection.
func (c *CollectionClient) Delete() *CollectionDelete {
	mutation := newCollectionMutation(c.config, OpDelete)
	return &CollectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CollectionClient) DeleteOne(co *Collection) *CollectionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CollectionClient) DeleteOneID(id int) *CollectionDeleteOne {
	builder := c.Delete().Where(collection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CollectionDeleteOne{builder}
}

// Query returns a query builder for Collection.
func (c *CollectionClient) Query() *CollectionQuery {
	return &CollectionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCollection},
		inters: c.Interceptors(),
	}
}

// Get returns a Collection entity by its id.
func (c *CollectionClient) Get(ctx context.Context, id int) (*Collection, error) {
	return c.Query().Where(collection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CollectionClient) GetX(ctx context.Context, id int) *Collection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTokens queries the tokens edge of a Collection.
func (c *CollectionClient) QueryTokens(co *Collection) *TokenQuery {
	query := (&TokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, id),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, collection.TokensTable, collection.TokensColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CollectionClient) Hooks() []Hook {
	return c.hooks.Collection
}

// Interceptors returns the client interceptors.
func (c *CollectionClient) Interceptors() []Interceptor {
	return c.inters.Collection
}

func (c *CollectionClient) mutate(ctx context.Context, m *CollectionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CollectionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CollectionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CollectionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Collection mutation op: %q", m.Op())
	}
}

// InscriptionClient is a client for the Inscription schema.
type InscriptionClient struct {
	config
}

// NewInscriptionClient returns a client for the Inscription from the given config.
func NewInscriptionClient(c config) *InscriptionClient {
	return &InscriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `inscription.Hooks(f(g(h())))`.
func (c *InscriptionClient) Use(hooks ...Hook) {
	c.hooks.Inscription = append(c.hooks.Inscription, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `inscription.Intercept(f(g(h())))`.
func (c *InscriptionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Inscription = append(c.inters.Inscription, interceptors...)
}

// Create returns a builder for creating a Inscription entity.
func (c *InscriptionClient) Create() *InscriptionCreate {
	mutation := newInscriptionMutation(c.config, OpCreate)
	return &InscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Inscription entities.
func (c *InscriptionClient) CreateBulk(builders ...*InscriptionCreate) *InscriptionCreateBulk {
	return &InscriptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Inscription.
func (c *InscriptionClient) Update() *InscriptionUpdate {
	mutation := newInscriptionMutation(c.config, OpUpdate)
	return &InscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InscriptionClient) UpdateOne(i *Inscription) *InscriptionUpdateOne {
	mutation := newInscriptionMutation(c.config, OpUpdateOne, withInscription(i))
	return &InscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InscriptionClient) UpdateOneID(id int) *InscriptionUpdateOne {
	mutation := newInscriptionMutation(c.config, OpUpdateOne, withInscriptionID(id))
	return &InscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Inscription.
func (c *InscriptionClient) Delete() *InscriptionDelete {
	mutation := newInscriptionMutation(c.config, OpDelete)
	return &InscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InscriptionClient) DeleteOne(i *Inscription) *InscriptionDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *InscriptionClient) DeleteOneID(id int) *InscriptionDeleteOne {
	builder := c.Delete().Where(inscription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InscriptionDeleteOne{builder}
}

// Query returns a query builder for Inscription.
func (c *InscriptionClient) Query() *InscriptionQuery {
	return &InscriptionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeInscription},
		inters: c.Interceptors(),
	}
}

// Get returns a Inscription entity by its id.
func (c *InscriptionClient) Get(ctx context.Context, id int) (*Inscription, error) {
	return c.Query().Where(inscription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InscriptionClient) GetX(ctx context.Context, id int) *Inscription {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *InscriptionClient) Hooks() []Hook {
	return c.hooks.Inscription
}

// Interceptors returns the client interceptors.
func (c *InscriptionClient) Interceptors() []Interceptor {
	return c.inters.Inscription
}

func (c *InscriptionClient) mutate(ctx context.Context, m *InscriptionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&InscriptionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&InscriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&InscriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&InscriptionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Inscription mutation op: %q", m.Op())
	}
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `token.Intercept(f(g(h())))`.
func (c *TokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.Token = append(c.inters.Token, interceptors...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id int) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id int) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeToken},
		inters: c.Interceptors(),
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id int) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id int) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCollection queries the collection edge of a Token.
func (c *TokenClient) QueryCollection(t *Token) *CollectionQuery {
	query := (&CollectionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(collection.Table, collection.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, token.CollectionTable, token.CollectionColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// Interceptors returns the client interceptors.
func (c *TokenClient) Interceptors() []Interceptor {
	return c.inters.Token
}

func (c *TokenClient) mutate(ctx context.Context, m *TokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Token mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Collection, Inscription, Token []ent.Hook
	}
	inters struct {
		Collection, Inscription, Token []ent.Interceptor
	}
)