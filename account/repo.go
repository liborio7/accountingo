package account

import (
	"database/sql"
	"github.com/liborio7/accountingo/cache"
	"github.com/satori/go.uuid"
)

type Repo struct {
	db *sql.DB
	ch cache.Cache
}

func NewRepo(db *sql.DB, ch cache.Cache) *Repo {
	return &Repo{db: db, ch: ch}
}

func (*Repo) idScan(r *sql.Rows, m *Model) error {
	return r.Scan(&m.Id)
}

func (*Repo) fullScan(r *sql.Rows, m *Model) error {
	return r.Scan(&m.Id, &m.Name, &m.Surname, &m.Currency, &m.Amount, &m.InsertDate, &m.InsertMillis, &m.Version)
}

func (r *Repo) Insert(m *Model) error {
	if uuid.Equal(m.Id, uuid.Nil) {
		m.Id = uuid.NewV4()
	}
	rows, err := r.db.Query("insert into account "+
		"values($1, $2, $3, $4, $5) returning id",
		m.Id, m.Name, m.Surname, m.Currency, m.Amount)
	if err != nil {
		return err
	}
	rows.Next()
	return r.idScan(rows, m)
}

func (r *Repo) LoadById(m *Model) error {
	rows, err := r.db.Query("select * from account "+
		"where id = $1", m.Id)
	if err != nil {
		return err
	}
	rows.Next()
	return r.fullScan(rows, m)
}

func (r *Repo) Load(sa uint64, limit uint) ([]Model, error) {
	rows, err := r.db.Query("select * from account "+
		"where insert_millis > $1 "+
		"order by insert_millis "+
		"limit $2", sa, limit)
	if err != nil {
		return nil, err
	}
	var res []Model
	for rows.Next() {
		m := Model{}
		err := r.fullScan(rows, &m)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}
