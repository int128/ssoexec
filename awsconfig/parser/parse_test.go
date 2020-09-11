package parser

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		f, err := os.Open("testdata/example1")
		if err != nil {
			t.Fatalf("could not open: %s", err)
		}
		defer f.Close()
		got, err := Parse(f)
		if err != nil {
			t.Fatalf("Parse() error: %s", err)
		}
		want := []*Section{
			{Header: "default", Values: map[string]string{
				"region": "us-west-2",
				"output": "json",
			}},
			{Header: "profile development", Values: map[string]string{
				"region":                  "us-east-1",
				"role_arn":                "arn:aws:iam::123456789012:role/role-name",
				"s3":                      "",
				"max_concurrent_requests": "20",
				"max_bandwidth":           "50MB/s",
			}},
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("mismatch (-got +want):\n%s", diff)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		got, err := Parse(strings.NewReader(""))
		if err != nil {
			t.Fatalf("Parse() error: %s", err)
		}
		var want []*Section
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("mismatch (-got +want):\n%s", diff)
		}
	})
	t.Run("NoSectionError", func(t *testing.T) {
		_, err := Parse(strings.NewReader(`key = value`))
		if err == nil {
			t.Fatalf("Parse() wants error but nil")
		}
		t.Logf("expected error: %s", err)
	})
}
