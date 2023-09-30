package main

import (
	"github.com/uncle-gua/sqlboiler/drivers"
	"github.com/uncle-gua/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
