package main

import (
	"github.com/thrasher-corp/sqlboiler/drivers"
	"github.com/thrasher-corp/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
