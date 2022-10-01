package build

import (
	"LinkShortening/config"
	errPkg "LinkShortening/internals/myerror"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDb(configDB config.Database) (*pgxpool.Pool, error) {
	var err error
	addressPostgres := "postgres://" + configDB.UserName + ":" + configDB.Password +
		"@" + configDB.Host + ":" + configDB.Port + "/" + configDB.SchemaName

	conn, err := pgxpool.Connect(context.Background(), addressPostgres)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.MCreateDBNotConnect,
		}
	}
	return conn, nil
}
