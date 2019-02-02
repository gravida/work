package models

import ()

// CountUnit -
type CountUnit struct {
	Id      int64   `json:"id"`
	Tid     int64   `xorm:"unique(name)"json:"tid"`
	Name    string  `xorm:"unique(name) NOT NULL" json:"name"`
	Desc    string  `xorm:"varchar(32)" json:"desc"`
	Price   float32 `json:"price"`
	Created int64   `xorm:"created" json:"created"`
	Updated int64   `xorm:"updated" json:"updated"`
}

type CountUnitLine struct {
	Id    int64 `json:"id"`
	Lid   int64 `json:"lid"`
	Cid   int64 `json:"cid"`
	Count int   `json:"count"`
}

// Add
func AddCountUnit(c *CountUnit) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(c); err != nil {
		return err
	}

	return sess.Commit()
}

// Del

// validate name -
func ExistCountUnitByName(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Get(&Work{Uid: uid, Name: name})
}

// count
func CountWorks() (total int64, err error) {
	w := new(Work)
	total, err = x.Count(w)
	return total, err
}

// query
func QueryWorkByID(id int64) (*Work, bool, error) {
	w := new(Work)
	has, err := x.Id(id).Get(w)
	return w, has, err
}

func QueryAllWorks(page, pageSize int) ([]*Work, error) {
	works := make([]*Work, 0, pageSize)
	return works, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&works)
}
