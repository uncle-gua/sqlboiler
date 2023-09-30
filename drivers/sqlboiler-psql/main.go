package main

import (
	"github.com/uncle-gua/sqlboiler/drivers"
	"github.com/uncle-gua/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
