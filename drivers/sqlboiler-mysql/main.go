package main

import (
	"github.com/uncle-gua/sqlboiler/drivers"
	"github.com/uncle-gua/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
