package domainutil

import "testing"

func TestSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantHost string
		wantSld  string
		wantTld  string
		wantErr  bool
	}{
		{
			args:     args{"www.abc.com"},
			wantHost: "www",
			wantSld:  "abc",
			wantTld:  "com",
			wantErr:  false,
		},
		{
			name:     "empty host",
			args:     args{"abc.com"},
			wantHost: "",
			wantSld:  "abc",
			wantTld:  "com",
			wantErr:  false,
		},
		{
			name:     "chinese char",
			args:     args{"hello.你好.中国"},
			wantHost: "hello",
			wantSld:  "你好",
			wantTld:  "中国",
			wantErr:  false,
		},
		{
			name:     "unknown suffix",
			args:     args{"abc.unknown"},
			wantHost: "",
			wantSld:  "",
			wantTld:  "unknown",
			wantErr:  true,
		},
		{
			name:     "mutil level host",
			args:     args{"a.b.c.d.e.f.com"},
			wantHost: "a.b.c.d.e",
			wantSld:  "f",
			wantTld:  "com",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHost, gotSld, gotTld, err := Split(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHost != tt.wantHost {
				t.Errorf("Split() gotHost = %v, want %v", gotHost, tt.wantHost)
			}
			if gotSld != tt.wantSld {
				t.Errorf("Split() gotSld = %v, want %v", gotSld, tt.wantSld)
			}
			if gotTld != tt.wantTld {
				t.Errorf("Split() gotTld = %v, want %v", gotTld, tt.wantTld)
			}
		})
	}
}
