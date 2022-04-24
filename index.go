package data_postgres

import (
	"github.com/chefsgo/data"
	dp "github.com/chefsgo/data-postgres"
)

func init() {
	driver := dp.Driver()
	data.Register("timescale", driver)
	data.Register("timescaledb", driver)
	data.Register("tsdb", driver)
}
