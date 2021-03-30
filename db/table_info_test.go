package db

import "testing"

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
			wantErr: true,
		},
		{
			name: "case2",
			args: args{
				tableName: "test",
			},
			wantResult: "测试",
			wantErr: false,
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