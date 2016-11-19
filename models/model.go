package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"OnlineJudge/base"
	"OnlineJudge/models/db"
	"github.com/jmoiron/sqlx"
)

type Model struct {
	DB    *sqlx.DB
	Table string
}

func (this *Model) OpenDB() error {
	db.Init()
	var err error
	this.DB, err = db.NewDB()
	return err
}

func (this *Model) CloseDB() {
	this.DB.Close()
}

func (this *Model) GetAllFields(v interface{}) ([]string, error) {
	mapping := this.DB.Mapper.TypeMap(reflect.TypeOf(v))
	if len(mapping.Paths) == 0 {
		return nil, errors.New("Empty struct definition.")
	}
	fields := []string{}
	for k, _ := range mapping.Paths {
		fields = append(fields, k)
	}
	return fields, nil
}

func (this *Model) FilterFields(all_fields []string, required []string, excepts []string) ([]string, error) {
	if (excepts != nil && len(excepts) > 0) && (required != nil && len(required) > 0) {
		return nil, errors.New("Parameters conficted, using only required or only excepts.")
	}
	fields := []string{}
	if required != nil && len(required) > 0 {
		for _, v := range required {
			if base.ArrayContains(all_fields, v) {
				fields = append(fields, v)
			}
		}
	} else {
		if excepts != nil && len(excepts) > 0 {
			for _, v := range all_fields {
				if !base.ArrayContains(excepts, v) {
					fields = append(fields, v)
				}
			}
		} else {
			fields = all_fields
		}
	}
	return fields, nil
}

func (this *Model) GenerateInsertSQL(st interface{}, table string, required []string, excepts []string) (string, error) {
	all_fields, err := this.GetAllFields(st)
	if err != nil {
		return "", err
	}
	reversed_fields, err := this.FilterFields(all_fields, required, excepts)
	if err != nil {
		return "", err
	}
	fields := []string{}
	values := []string{}
	for _, v := range reversed_fields {
		fields = append(fields, v)
		values = append(values, ":"+v)
	}
	cols := "(" + strings.Join(fields, ",") + ")"
	vals := "(" + strings.Join(values, ",") + ")"
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, cols, vals), nil
}

func (this *Model) GenerateSelectSQL(st interface{}, required []string, excepts []string) (string, error) {
	all_fields, err := this.GetAllFields(st)
	if err != nil {
		return "", err
	}
	fields, err := this.FilterFields(all_fields, required, excepts)
	if err != nil {
		return "", err
	}
	for k, v := range fields {
		var dft string
		val := this.DB.Mapper.FieldByName(reflect.ValueOf(st), v)
		switch val.Interface().(type) {
		case int:
			dft = "0"
		case int64:
			dft = "0"
		case string:
			dft = "''"
		case time.Time:
			dft = "CURRENT_TIMESTAMP"
		}
		fields[k] = fmt.Sprintf("COALESCE(%s, %s) AS %s", v, dft, v)
	}
	return strings.Join(fields, ","), nil
}

func (this *Model) InlineInsert(st interface{}, required []string, excepts []string) (int64, error) {
	sql_insert, err := this.GenerateInsertSQL(st, this.Table, required, excepts)
	if err != nil {
		return 0, err
	}
	tx, err := this.DB.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	res, err := tx.NamedExec(sql_insert, st)
	if err != nil {
		return 0, err
	}
	last_insert_id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *Model) InlineUpdate(st interface{}, pk string, required []string, excepts []string) error {
	all_fields, err := this.GetAllFields(st)
	if err != nil {
		return err
	}
	fmt.Println(all_fields)
	//`UPDATE TABLE this.Table SET `
	return nil
}
