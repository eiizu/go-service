package usecase

import (
	"reflect"
	"testing"
)

func TestNewSomething(t *testing.T) {
	tests := []struct {
		name string
		want *Something
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSomething(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_DoSomething(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		op      *Something
		args    args
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.op.DoSomething(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Something.DoSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Something.DoSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
