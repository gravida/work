package models

import ()

// WorkLine -
type WorkLine struct {
	Id      int64  `json:"id"`
	Wid     int64  `xorm:"unique(name)"json:"wid"`
	Name    string `xorm:"unique(name) NOT NULL" json:"name"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// Add
func AddWorkLine(w *WorkLine) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(w); err != nil {
		return err
	}

	return sess.Commit()
}

// Del

// validate name -
func ExistWorkLineByName(wid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Get(&WorkLine{Wid: wid, Name: name})
}

// count
func CountWorkLines() (total int64, err error) {
	w := new(WorkLine)
	total, err = x.Count(w)
	return total, err
}

// query
func QueryWorkLineByID(id int64) (*WorkLine, bool, error) {
	w := new(WorkLine)
	has, err := x.Id(id).Get(w)
	return w, has, err
}

func QueryAllWorkLines(page, pageSize int) ([]*WorkLine, error) {
	works := make([]*WorkLine, 0, pageSize)
	return works, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&works)
}
