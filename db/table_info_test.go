package db

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

type Dict struct {
	ID    string
	Label string
}

func (P Dict) TableName() string {
	return "sys_dict"
}
func TestNormal(t *testing.T) {
	var dicts []Dict
	tx:=DB
	//tx.Select("id")
	//tx.Order("type")
	//tx.Limit(10)
	tx = tx.Where("id=?", "asdf")
	tx.Where("label like ?",0).Find(&dicts)
	log.Info(dicts)
}

func TestGetTableComment(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name: "case1",
			args: args{
				tableName: "test1",
			},
			wantResult: "",
			wantErr:    true,
		},
		{
			name: "case2",
			args: args{
				tableName: "test",
			},
			wantResult: "测试",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetTableComment(tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTableComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("GetTableComment() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
