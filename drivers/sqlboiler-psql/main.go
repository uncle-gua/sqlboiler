package main

import (
	"github.com/thrasher-corp/sqlboiler/drivers"
	"github.com/thrasher-corp/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
