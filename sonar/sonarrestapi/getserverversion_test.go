package sonarrestapi

import (
	"errors"
	"sonarci/net"
	"testing"
)

func Test_restApi_GetServerVersion(t *testing.T) {
	mockVersion := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		chVersion := make(chan []byte, 1)
		chVersion <- []byte("1.0")
		chErr := make(chan error, 1)
		chErr <- nil
		return chVersion, chErr
	}}
	mockError := &mockConnection{doGet: func(route string) (<-chan []byte, <-chan error) {
		chError := make(chan error, 1)
		chError <- errors.New("failure")
		return nil, chError
	}}

	type fields struct {
		Connection net.Connection
		Server     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"success", fields{Connection: mockVersion, Server: "http://server"}, "1.0", false},
		{"error", fields{Connection: mockError, Server: "http://server"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restApi := &restApi{
				Connection: tt.fields.Connection,
				Server:     tt.fields.Server,
			}
			got, err := restApi.GetServerVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetServerVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
