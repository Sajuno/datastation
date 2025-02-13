package runner

import (
	"context"
)

func (ec EvalContext) testODBCConnection(ctx context.Context, db *DatabaseConnectorInfo) error {
	_, conn, err := ec.getConnectionString(db.Database)
	if err != nil {
		return err
	}

	driver, err := openODBCDriver(conn)
	if err != nil {
		return err
	}

	// Because odbc is for sql databases, this should work, else we have to use cgo and call the actual driver.
	if _, err := driver.QueryContext(ctx, "SELECT 1"); err != nil {
		return err
	}

	return nil
}
