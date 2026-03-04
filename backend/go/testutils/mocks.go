package testutils

import (
	"bytes"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

// MockPgQuery for pg.Query
type MockPgQuery struct {
	Err        error
	modelDest  any
	WhereID    string
	SelectFunc func(modelDest any, dest ...interface{}) error              // Function to customize Select behavior
	InsertFunc func(modelDest any, dest ...interface{}) (pg.Result, error) // Function to customize Insert behavior
	UpdateFunc func(modelDest any, dest ...interface{}) (pg.Result, error) // Function to customize Update behavior
}

type MockPgResult struct {
	NumRowsAffected int
	Err             error
}

func (r *MockPgResult) RowsAffected() int {
	return r.NumRowsAffected
}

func (r *MockPgResult) RowsReturned() int {
	return 0
}

func (r *MockPgResult) Error() error {
	return r.Err
}

func (r *MockPgResult) Model() orm.Model {
	return nil
}

func (mq *MockPgQuery) Column(columns ...string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) ColumnExpr(expr string, params ...interface{}) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Relation(name string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Limit(count int) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Offset(offset int) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Where(query string, params ...interface{}) interfaces.PgxQuerySeter {
	if (query == "id = ?" || query == "blog_post.id = ?" || query == "home_content.id = ?" || query == "groove_jr_content.id = ?" || query == "about_content.id = ?") && len(params) > 0 {
		if id, ok := params[0].(string); ok {
			mq.WhereID = id
		}
	}
	return mq
}

func (mq *MockPgQuery) Order(orders ...string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Group(columns ...string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Join(join string, params ...interface{}) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) OnConflict(s string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Set(s string) interfaces.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Select(dest ...interface{}) error {
	if mq.Err != nil {
		return mq.Err
	}

	if mq.WhereID != "" && mq.WhereID == "not-found" {
		return pg.ErrNoRows
	}

	if mq.SelectFunc != nil {
		return mq.SelectFunc(mq.modelDest, dest...)
	}

	// Default behavior if SelectFunc is not set
	return nil
}

func (mq *MockPgQuery) SelectAndCount(dest ...interface{}) (int, error) {
	if mq.Err != nil {
		return 0, mq.Err
	}
	if mq.SelectFunc != nil {
		err := mq.SelectFunc(mq.modelDest, dest...)
		return 0, err
	}
	return 0, nil
}

func (mq *MockPgQuery) Insert(dest ...interface{}) (pg.Result, error) {
	if mq.Err != nil {
		return &MockPgResult{Err: mq.Err}, mq.Err
	}
	if mq.InsertFunc != nil {
		return mq.InsertFunc(mq.modelDest, dest...)
	}
	return &MockPgResult{NumRowsAffected: 1}, nil
}

func (mq *MockPgQuery) Update(dest ...interface{}) (pg.Result, error) {
	if mq.Err != nil {
		return &MockPgResult{Err: mq.Err}, mq.Err
	}
	if mq.UpdateFunc != nil {
		return mq.UpdateFunc(mq.modelDest, dest...)
	}
	return &MockPgResult{NumRowsAffected: 1}, nil
}

func (mq *MockPgQuery) Delete(dest ...interface{}) (pg.Result, error) {
	if mq.Err != nil {
		return &MockPgResult{Err: mq.Err}, mq.Err
	}
	return &MockPgResult{NumRowsAffected: 1}, nil
}

// MockPgDB for pg.DB
type MockPgDB struct {
	MockQuery *MockPgQuery
}

// NewMockPgDB returns a new MockPgDB instance.
func NewMockPgDB() *MockPgDB {
	return &MockPgDB{}
}

func (m *MockPgDB) Model(model ...interface{}) interfaces.PgxQuerySeter {
	if m.MockQuery != nil {
		m.MockQuery.modelDest = model[0]
		return m.MockQuery
	}
	return &MockPgQuery{modelDest: model[0]}
}

func (m *MockPgDB) RunInTransaction(fn func(interfaces.PgxDB) error) error {
	return fn(m)
}
