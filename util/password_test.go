package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "OK",
			password:  RandomString(6),
			assertion: assert.NoError,
		},
		{
			name:      "NG",
			password:  RandomString(1000),
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.password)
			tt.assertion(t, err)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword, _ := HashPassword(password)
	type args struct {
		password       string
		hashedPassword string
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "OK",
			args:      args{password: password, hashedPassword: hashedPassword},
			assertion: assert.NoError,
		},
		{
			name:      "NG",
			args:      args{password: RandomString(6), hashedPassword: hashedPassword},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, CheckPassword(tt.args.password, tt.args.hashedPassword))
		})
	}
}
