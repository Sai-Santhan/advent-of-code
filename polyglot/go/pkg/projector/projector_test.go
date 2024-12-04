package projector_test

import (
	"github.com/Sai-Santhan/aoc/pkg/projector"
	"testing"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello, Frontend Masters",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("Expected to find value '%v'", value)
	}
	if value != v {
		t.Errorf("Expected to find value '%v' but received %v", value, v)
	}
}

func TestProjector_GetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is_great")
}

func TestProjector_SetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	proj.SetValue("foo", "baz")
	test(t, proj, "fem", "is_great")
	proj.SetValue("fem", "is_super_great")
	test(t, proj, "fem", "is_super_great")

	proj = getProjector("/", data)
	test(t, proj, "fem", "is_great")
}

func TestProjector_RemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	test(t, proj, "fem", "is_great")

	proj.RemoveValue("fem")
	test(t, proj, "fem", "is_great")
}
