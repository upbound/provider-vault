package database

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func diffWith(attrs map[string]*terraform.ResourceAttrDiff) *terraform.InstanceDiff {
	return &terraform.InstanceDiff{Attributes: attrs}
}

func TestPasswordCreateOnlyDiff(t *testing.T) {
	pwAttr := func() map[string]*terraform.ResourceAttrDiff {
		return map[string]*terraform.ResourceAttrDiff{
			"postgresql.0.password":             {Old: "", New: "s3cr3t"},
			"mssql.0.password":                  {Old: "", New: "s3cr3t"},
			"redis_elasticache.0.password":      {Old: "", New: "s3cr3t"},
			"postgresql.0.max_open_connections": {Old: "0", New: "5"},
		}
	}

	cases := []struct {
		name       string
		state      *terraform.InstanceState
		wantPwKept bool // postgresql.0.password still present after the hook
		wantNonPw  bool // non-password attr always preserved
	}{
		{
			name:       "create keeps password (empty state ID)",
			state:      &terraform.InstanceState{ID: ""},
			wantPwKept: true,
			wantNonPw:  true,
		},
		{
			name:       "create keeps password (nil state)",
			state:      nil,
			wantPwKept: true,
			wantNonPw:  true,
		},
		{
			name:       "update drops password (non-empty state ID)",
			state:      &terraform.InstanceState{ID: "database/config/db-foo"},
			wantPwKept: false,
			wantNonPw:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := diffWith(pwAttr())
			out, err := passwordCreateOnlyDiff(d, tc.state, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			_, pwKept := out.Attributes["postgresql.0.password"]
			if pwKept != tc.wantPwKept {
				t.Errorf("postgresql.0.password present=%v, want %v", pwKept, tc.wantPwKept)
			}
			// On update, ALL engine passwords must be gone.
			if !tc.wantPwKept {
				for _, k := range []string{"mssql.0.password", "redis_elasticache.0.password"} {
					if _, ok := out.Attributes[k]; ok {
						t.Errorf("%s should have been dropped on update", k)
					}
				}
			}
			if _, ok := out.Attributes["postgresql.0.max_open_connections"]; ok != tc.wantNonPw {
				t.Errorf("non-password attr present=%v, want %v", ok, tc.wantNonPw)
			}
		})
	}
}

func TestPasswordCreateOnlyDiff_NilAndDestroy(t *testing.T) {
	// nil diff → returned as-is, no panic
	if out, err := passwordCreateOnlyDiff(nil, &terraform.InstanceState{ID: "x"}, nil); err != nil || out != nil {
		t.Errorf("nil diff: got (%v, %v), want (nil, nil)", out, err)
	}
	// destroy diff → untouched
	d := &terraform.InstanceDiff{Destroy: true, Attributes: map[string]*terraform.ResourceAttrDiff{
		"postgresql.0.password": {Old: "x", New: ""},
	}}
	out, err := passwordCreateOnlyDiff(d, &terraform.InstanceState{ID: "x"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := out.Attributes["postgresql.0.password"]; !ok {
		t.Error("destroy diff must not be modified")
	}
}

func TestPasswordEnginesComplete(t *testing.T) {
	// Guard against drift: every engine known to carry a password must be listed.
	want := map[string]bool{
		"cassandra": true, "couchbase": true, "elasticsearch": true, "hana": true,
		"influxdb": true, "mongodb": true, "mssql": true, "mysql": true,
		"mysql_aurora": true, "mysql_legacy": true, "mysql_rds": true, "oracle": true,
		"postgresql": true, "redis": true, "redis_elasticache": true, "redshift": true,
		"snowflake": true,
	}
	got := map[string]bool{}
	for _, e := range passwordEngines {
		got[e] = true
	}
	if len(got) != len(want) {
		t.Fatalf("passwordEngines has %d entries, want %d", len(got), len(want))
	}
	for e := range want {
		if !got[e] {
			t.Errorf("missing engine %q in passwordEngines", e)
		}
	}
}
