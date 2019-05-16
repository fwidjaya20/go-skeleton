package tx

import (
	"github.com/jmoiron/sqlx"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// Run is a function that used to run the service under tx
func Run(db *sqlx.DB, config fazzdb.Config, fn func(q *fazzdb.Query) (interface{}, error)) (interface{}, error) {
	tx, err := db.Beginx()
	if nil != err {
		return nil, err
	}

	q := fazzdb.QueryTx(tx, config)
	res, err := fn(q)
	if nil != err {
		_ = q.Tx.Rollback()
		return nil, err
	}

	_ = q.Tx.Commit()

	return res, nil
}

// RunDefault basic boiler plate to start the transaction
func RunDefault(db *sqlx.DB, fn func(q *fazzdb.Query) (interface{}, error)) (interface{}, error) {
	return Run(db, fazzdb.DEFAULT_QUERY_CONFIG, fn)
}
