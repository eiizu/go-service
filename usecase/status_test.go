package usecase

import (
	"reflect"
	"testing"
)

func TestNewStatus(t *testing.T) {
	type args struct {
		AppName string
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
			if got := NewStatus(tt.args.AppName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Statusz(t *testing.T) {
	tests := []struct {
		name    string
		s       *Status
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Statusz()
			if (err != nil) != tt.wantErr {
				t.Errorf("Status.Statusz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Status.Statusz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Healthz(t *testing.T) {
	tests := []struct {
		name    string
		uc      *Status
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.Healthz()
			if (err != nil) != tt.wantErr {
				t.Errorf("Status.Healthz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Status.Healthz() = %v, want %v", got, tt.want)
			}
		})
	}
}
