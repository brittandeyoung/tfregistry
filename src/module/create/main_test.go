package main_test

import (
	"testing"

	"github.com/brittandeyoung/tfregistry/src/internal/client"
	"github.com/brittandeyoung/tfregistry/src/internal/resource/module/odm"
)

func TestModule_create(t *testing.T) {
	mc, _ := client.NewClient("https://judp069no2.execute-api.us-east-1.amazonaws.com")

	module := odm.Module{
		Namespace:   "testing",
		Provider:    "aws",
		Name:        "name",
		Description: "This is a test description.",
		Source:      "https://localhost",
	}

	resp, err := mc.CreateModule(module)

	if err != nil {
		t.Fatalf("there was an error: %s", err.Error())
	}
	if resp.Namespace != module.Namespace {
		t.Fatal(`Namespace not set properly.`)
	}
	if resp.Provider != module.Provider {
		t.Fatal(`Provider not set properly.`)
	}
}
