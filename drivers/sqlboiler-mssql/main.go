package main

import (
	"github.com/thrasher-corp/sqlboiler/drivers"
	"github.com/thrasher-corp/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
