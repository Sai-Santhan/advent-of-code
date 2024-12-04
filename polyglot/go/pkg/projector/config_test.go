package projector_test

import (
	"github.com/Sai-Santhan/aoc/pkg/projector"
	"reflect"
	"testing"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
	return opts
}

func testConfig(t *testing.T, args []string, operation projector.Operation) {
	opts := getOpts([]string{})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}
	if !reflect.DeepEqual([]string{}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}
	//if config.Operation == operation {
	//	t.Errorf("operation expect was %v but got %v", opts.Args\\\)
	//}
}

func TestConfigPrint(t *testing.T) {
	opts := getOpts([]string{})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}
	if !reflect.DeepEqual([]string{}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}
}

func TestConfigPrintKey(t *testing.T) {
	opts := getOpts([]string{"foo"})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}
	if !reflect.DeepEqual([]string{"foo"}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}
}

func TestConfigAddKeyValue(t *testing.T) {
	opts := getOpts([]string{"add", "foo", "bar"})
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}
	if !reflect.DeepEqual([]string{"foo", "bar"}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}
}
