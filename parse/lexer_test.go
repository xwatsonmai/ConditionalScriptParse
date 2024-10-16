package parse

import "testing"

type test struct {
	input   string
	wantErr string
	output  bool
}

func TestParser(t *testing.T) {
	testcases := []test{
		{
			input:   "1 = 1",
			wantErr: "",
			output:  true,
		},
		{
			input:   "1=1 or 2=2 or 3=3",
			wantErr: "",
			output:  true,
		},
		{
			input:   "(1 = 1)",
			wantErr: "",
			output:  true,
		},
		{
			input:   "(1=1)",
			wantErr: "",
			output:  true,
		},
		{
			input:   "1 > 1",
			wantErr: "",
			output:  false,
		},
		{
			input:   "1>1",
			wantErr: "",
			output:  false,
		},
		{
			input:   "1 + 1 = 2",
			wantErr: "",
			output:  true,
		},
		{
			input:   "(1 + 1) > 2",
			wantErr: "",
			output:  false,
		},
		{
			input:   "2 > -1",
			wantErr: "",
			output:  true,
		},
		{
			input:   "2 >= 2",
			wantErr: "",
			output:  true,
		},
		{
			input:   "2 <= 2",
			wantErr: "",
			output:  true,
		},
		{
			input:   "2 > (0-1)",
			wantErr: "",
			output:  true,
		},
		{
			input:   "2 < (0-1)",
			wantErr: "",
			output:  false,
		},
		{
			input:   "\"asd\" = \"zxc\"",
			wantErr: "",
			output:  false,
		},
		{
			input:   "\"asd\" != \"zxc\"",
			wantErr: "",
			output:  true,
		},
		{
			input:   "\"asd\" + \"b\" = \"asdb\"",
			wantErr: "",
			output:  true,
		},
		{
			input:   "\"禁言\" in (\"你被禁言了\")",
			wantErr: "",
			output:  true,
		},
		{
			input:   `"\"是" = "是"`,
			wantErr: "",
			output:  false,
		},
	}
	//parseHandler := &exprLex{}
	for _, tc := range testcases {
		lex := Parse([]byte(tc.input))
		err := lex.GetError()
		v := lex.GetResult()
		var gotErr string
		if err != nil {
			gotErr = err.Error()
		}
		if gotErr != tc.wantErr {
			t.Errorf("%s err: %v, want %v", tc.input, gotErr, tc.wantErr)
		}
		if v != tc.output {
			t.Errorf("%s: %v, want %v", tc.input, v, tc.output)
		}
	}
}
