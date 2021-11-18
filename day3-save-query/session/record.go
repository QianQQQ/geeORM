package session

import (
	"geeORM/clause"
	"reflect"
)

//u1, u2 := &User{Name: "Tom", Age: 18}, &User{Name: "Sam", Age: 25}
//s.Insert(u1, u2, ...)
func (s *Session) Insert(values ...interface{}) (int64, error) {
	var recordValues []interface{}
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}
	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// var users []User
// s.Find(&users);
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 通过[]User类型拿到User类型
	destType := destSlice.Type().Elem()
	// 通过User类型来设置RefTable
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}
	for rows.Next() {
		// 创建一个User类型的玩意, 并且是可以修改的
		dest := reflect.New(destType).Elem()
		// Scan多个传入参数, 就是User各类型对应的指针
		var values []interface{}
		for _, name := range table.FieldNames {
			// 获取到User各类型对应的指针
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// 将该行记录的值赋值给values中每一个指针
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
