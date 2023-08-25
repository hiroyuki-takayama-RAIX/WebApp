package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	// set env argument in test environment.
	t.Setenv("PORT", fmt.Sprint(wantPort))

	// this New() is in config package. Identify pakage with file name.
	// check which New() causes some kinds of error
	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	// check port
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}

	// check env
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}
