package data_postgres

import (
	"github.com/chefsgo/chef"
	dp "github.com/chefsgo/data-postgres"
)

func init() {
	driver := dp.Driver()
	chef.Register("timescale", driver)
	chef.Register("timescaledb", driver)
	chef.Register("tsdb", driver)
}
