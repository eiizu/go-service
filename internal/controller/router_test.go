package controller

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
)

func TestNew(t *testing.T) {
	type args struct {
		somethingC SomethingController
		statusC    StatusController
	}
	tests := []struct {
		name string
		args args
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.somethingC, tt.args.statusC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
