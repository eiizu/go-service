package controller

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/unrolled/render"
)

func TestNewStatus(t *testing.T) {
	type args struct {
		uc StatusUseCase
		r  *render.Render
	}
	tests := []struct {
		name string
		args args
		want *Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatus(tt.args.uc, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_HandlerStatusz(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Status
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.HandlerStatusz(tt.args.w, tt.args.r)
		})
	}
}

func TestStatus_HandlerHealthz(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Status
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.HandlerHealthz(tt.args.w, tt.args.r)
		})
	}
}
