package jsoncheck_test

import (
	"testing"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git/jsoncheck"
)

func TestCheck(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{`{"a":1,"b":2,"c":1.2312,"d":true,"e":false}`}, false},
		{"dup key", args{`{"a":1,"b":2,"a":3}`}, true},
		{"ok nested key", args{`{"a":1,"b":2,"c":{"a":1,"b":2,"c":3}}`}, false},
		{"dup nested key", args{`{"a":1,"b":2,"c":{"a":1,"b":2,"a":3}}`}, true},
		{"ok array nested key", args{`{"a":1,"b":2,"c":[{"a":1,"b":2,"c":3},{"a":1,"b":2,"c":3}]}`}, false},
		{"dup array nested key", args{`{"a":1,"b":2,"c":[{"a":1,"b":2,"c":3},{"a":1,"b":2,"a":3}]}`}, true},
	}
	for _, tt := range tests {
		tt := tt // fix range scope
		t.Run(tt.name, func(t *testing.T) {
			if err := jsoncheck.Check([]byte(tt.args.data)); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
