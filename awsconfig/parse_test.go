package awsconfig

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		f, err := os.Open("parser/testdata/example1")
		if err != nil {
			t.Fatalf("could not open: %s", err)
		}
		defer f.Close()
		got, err := Parse(f)
		if err != nil {
			t.Fatalf("Parse() error: %s", err)
		}
		want := &Config{
			Profiles: []*Profile{
				{Name: "default", Region: "us-west-2"},
				{Name: "development", Region: "us-east-1"},
			},
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("mismatch (-got +want):\n%s", diff)
		}
	})
}
