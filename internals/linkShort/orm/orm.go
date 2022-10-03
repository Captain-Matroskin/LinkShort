package orm

import (
	errPkg "LinkShortening/internals/myerror"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type LinkShortWrapperInterface interface {
	CreateLinkShort(linkFull string, linkShort string) error
	TakeLinkShort(linkShort string) (string, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

type LinkShortWrapper struct {
	Conn ConnectionInterface
}

func (w *LinkShortWrapper) CreateLinkShort(linkFull string, linkShort string) error {
	contextTransaction := context.Background()
	tx, errBeginConn := w.Conn.Begin(contextTransaction)
	if errBeginConn != nil {
		return &errPkg.MyErrors{
			Text: errPkg.LSHCreateLinkShortTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, errExecTx := tx.Exec(contextTransaction,
		"INSERT INTO public.link (link, link_short) VALUES ($1, $2)", linkFull, linkShort)
	if errExecTx != nil {
		println(errExecTx.Error())
		println(errPkg.LSHCreateLinkShortNotInsertUniqueDB)
		switch errExecTx.Error() {
		case errPkg.LSHCreateLinkShortNotInsertUniqueDB:
			return &errPkg.MyErrors{
				Text: errPkg.LSHCreateLinkShortNotInsertUnique,
			}
		default:
			return &errPkg.MyErrors{
				Text: errPkg.LSHCreateLinkShortNotInsert,
			}
		}
	}

	errCommitTx := tx.Commit(contextTransaction)
	if errCommitTx != nil {
		return &errPkg.MyErrors{
			Text: errPkg.LSHCreateLinkShortNotCommit,
		}
	}

	return nil
}

func (w *LinkShortWrapper) TakeLinkShort(linkShort string) (string, error) {
	contextTransaction := context.Background()
	tx, errBeginConn := w.Conn.Begin(contextTransaction)
	if errBeginConn != nil {
		return "", &errPkg.MyErrors{
			Text: errPkg.LSHTakeLinkShortTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var linkFull string
	errQueryRow := tx.QueryRow(contextTransaction,
		"SELECT link FROM public.link WHERE link_short = $1",
		linkShort).Scan(linkFull)
	if errQueryRow != nil {
		if errQueryRow == pgx.ErrNoRows {
			return "", &errPkg.MyErrors{
				Text: errPkg.LSHTakeLinkShortNotFound,
			}
		}
		return "", &errPkg.MyErrors{
			Text: errPkg.LSHTakeLinkShortNotScan,
		}
	}

	errCommitTx := tx.Commit(contextTransaction)
	if errCommitTx != nil {
		return "", &errPkg.MyErrors{
			Text: errPkg.LSHTakeLinkShortNotCommit,
		}
	}

	return linkFull, nil
}
