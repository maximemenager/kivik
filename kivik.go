// Package kivik is an assemble-your-own CouchDB API, proxy, cache, and server..
package kivik

import (
	"fmt"

	"github.com/flimzy/kivik/driver"
)

// Client is a client connection handle to a CouchDB-like server.
type Client struct {
	driver       driver.Driver
	dsn          string
	driverClient driver.Client
}

// New creates a new client object specified by its database driver name
// and a driver-specific data source name.
func New(driverName, dataSourceName string) (*Client, error) {
	driversMu.RLock()
	driveri, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("kivik: unknown driver %q (forgotten import?)", driverName)
	}
	client, err := driveri.NewClient(dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Client{
		driver:       driveri,
		dsn:          dataSourceName,
		driverClient: client,
	}, nil
}

// Version returns the reported server version
func (c *Client) Version() (string, error) {
	si, err := c.driverClient.ServerInfo()
	return si.Version(), err
}

// DB returns a handle to the requested database. No validation is done at
// this stage.
func (c *Client) DB(name string) *DB {
	return &DB{}
}

// AllDBs returns a list of all databases.
func (c *Client) AllDBs() ([]string, error) {
	return c.driverClient.AllDBs()
}
