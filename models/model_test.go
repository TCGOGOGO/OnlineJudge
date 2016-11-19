package models

import (
	"testing"
)

type MyModel struct {
	Model
}

func (this *MyModel) TestFunc(t *testing.T) error {
	t.Log("In TestFunc, try to open/close DB")
	if err := this.OpenDB(); err != nil {
		t.Errorf("Failed to open database: %s", err)
		return err
	}
	defer this.CloseDB()
	return nil
}

func TestModel(t *testing.T) {
	t.Log("TestModel")
	mymodel := MyModel{}
	mymodel.TestFunc(t)
}

func (this *MyModel) TestFuncGetAllFields(t *testing.T) error {
	if err := this.OpenDB(); err != nil {
		t.Errorf("Failed to open database: %s", err)
		return err
	}
	defer this.CloseDB()
	t.Log("Get all fields from a struct")
	type Foo struct {
		RunId      int `db:"run_id"`
		Status     string
		StatusCode string `db:"status_code"`
	}
	f := Foo{1, "a", "b"}
	sql, err := this.GenerateInsertSQL(f, "foo", nil, []string{"run_id"})
	t.Log(sql)
	if err != nil {
		t.Errorf("%s", err)
	}
	_, err = this.GenerateInsertSQL(f, "foo", []string{"status"}, []string{"run_id"})
	t.Log(err)
	if err == nil {
		t.Errorf("Supposed to be failed, %s", err)

	}
	return err
}

func (this *MyModel) TestGenerateSelectSQL(t *testing.T) error {
	if err := this.OpenDB(); err != nil {
		t.Errorf("Failed to open database: %s", err)
		return err
	}
	defer this.CloseDB()
	type Foo struct {
		RunId      int `db:"run_id"`
		Status     string
		StatusCode string `db:"status_code"`
	}
	f := Foo{1, "a", "b"}
	sql, err := this.GenerateSelectSQL(f, nil, nil)
	if err != nil {
		t.Errorf("Failed to generate select sql, %s", err)
		return err
	}
	t.Log(sql)
	return nil
}

func TestGetAllFields(t *testing.T) {
	mymodel := MyModel{}
	mymodel.TestFuncGetAllFields(t)
	mymodel.TestGenerateSelectSQL(t)
}
