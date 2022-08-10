package {{.pkg}}
{{if .withCache}}
import (
    "context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{else}}
import "github.com/zeromicro/go-zero/core/stores/sqlx"
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		FindAll(ctx context.Context,orderBy string) ([]*{{.upperStartCamelObject}}, error)
        FindSome(ctx context.Context,where string ,orderBy string) ([]*{{.upperStartCamelObject}}, error)
        FindPage(ctx context.Context,page int64,size int64, where string ,orderBy string) ([]*{{.upperStartCamelObject}}, error)
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

func (m *custom{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s", sysDeptRows, m.table,orderBy)
	var resp []*{{.upperStartCamelObject}}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *custom{{.upperStartCamelObject}}Model) FindSome(ctx context.Context, where string ,orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s", sysDeptRows, m.table,where,orderBy)
	var resp []*{{.upperStartCamelObject}}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *custom{{.upperStartCamelObject}}Model) FindPage(ctx context.Context, page int64,size int64, where string ,orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	offset := (page - 1) * size
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s LIMIT %d,%d", sysDeptRows, m.table,where,orderBy,offset,size)
	var resp []*{{.upperStartCamelObject}}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}