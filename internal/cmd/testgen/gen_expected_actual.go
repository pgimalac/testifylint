package main

import (
	"strings"
	"text/template"
)

type ExpectedActualCasesGenerator struct{}

func (g ExpectedActualCasesGenerator) Data() any {
	const (
		report = "expected-actual: need to reverse actual and expected values"
	)

	type test struct {
		Name          string
		VarSets       [][]string
		InvalidChecks []Check
		ValidChecks   []Check
	}

	return struct {
		Pkgs, Objs     []string
		SuiteSelectors []string
		Tests          []test
	}{
		Pkgs:           []string{"assert", "require"},
		Objs:           []string{"assObj", "reqObj"},
		SuiteSelectors: []string{"s", "s.Assert()", "assObj", "s.Require()", "reqObj"},
		Tests: []test{
			{
				Name: "Basic",
				VarSets: [][]string{
					{"uint(11)"}, {"uint8(12)"}, {"uint16(13)"}, {"uint32(14)"}, {"uint64(15)"},
					{"21"}, {"int8(22)"}, {"int16(23)"}, {"int32(24)"}, {"int64(25)"},
					{"float32(31.)"}, {"float64(32.)"},
					{"complex64(41 - 0.707i)"}, {"complex128(42 - 0.707i)"},
					{`"string"`}, {"'r'"},

					{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}, {"g"}, {"h"},
					{"i"}, {"j"}, {"k"}, {"l"}, {"m"}, {"n"}, {"o"}, {"p"},

					{"aa"}, {"bb"}, {"cc"}, {"dd"}, {"ee"}, {"ff"}, {"gg"}, {"hh"},
					{"ii"}, {"jj"}, {"kk"}, {"ll"}, {"mm"}, {"nn"}, {"oo"}, {"pp"},

					{"Monday"}, {"DayMonday"},

					{"expected"}, {"tt.expected"}, {"ttp.expected"},
				},
				InvalidChecks: []Check{
					{Fn: "Equal", Argsf: "result, %s", ReportMsgf: report, ProposedArgsf: "%s, result"},
					{Fn: "NotEqual", Argsf: "result, %s", ReportMsgf: report, ProposedArgsf: "%s, result"},
				},
				ValidChecks: []Check{
					{Fn: "Equal", Argsf: "%s, result"},
					{Fn: "NotEqual", Argsf: "%s, result"},
				},
			},
			{
				Name: "String",
				VarSets: [][]string{
					{`"string"`}, {"o"}, {"oo"}, {"string(DayNameMonday)"},
					{"expected"}, {"tt.expected"}, {"ttp.expected"},
				},
				InvalidChecks: []Check{
					{Fn: "JSONEq", Argsf: "result, %s", ReportMsgf: report, ProposedArgsf: "%s, result"},
					{Fn: "YAMLEq", Argsf: "result, %s", ReportMsgf: report, ProposedArgsf: "%s, result"},
				},
				ValidChecks: []Check{
					{Fn: "JSONEq", Argsf: "%s, result"},
					{Fn: "YAMLEq", Argsf: "%s, result"},
				},
			},
		},
	}
}

func (g ExpectedActualCasesGenerator) ErroredTemplate() *template.Template {
	return template.Must(template.New("ExpectedActualCasesGenerator.ErroredTemplate").
		Funcs(fm).
		Parse(expectedActualCasesTmplText))
}

func (g ExpectedActualCasesGenerator) GoldenTemplate() *template.Template {
	return template.Must(template.New("ExpectedActualCasesGenerator.GoldenTemplate").
		Funcs(fm).
		Parse(strings.ReplaceAll(expectedActualCasesTmplText, "NewCheckerExpander", "NewCheckerExpander.AsGolden")))
}

const expectedActualCasesTmplText = header + `

package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestConfusedWithExpectedActual(t *testing.T) {
	{{- block "vars" . }} 
	var result string

	const (
		a = uint(11)
		b = uint8(12)
		c = uint16(13)
		d = uint32(14)
		e = uint64(15)

		f = int(21)
		g = int8(22)
		h = int16(23)
		i = int32(24)
		j = int64(25)

		k = float32(31.)
		l = float64(32.)

		m = complex64(41 - 0.707i)
		n = complex128(42 - 0.707i)

		o = "string"
		p = 'r'
	)

	const (
		aa uint   = 11
		bb uint8  = 12
		cc uint16 = 13
		dd uint32 = 14
		ee uint64 = 15

		ff int   = 21
		gg int8  = 22
		hh int16 = 23
		ii int32 = 24
		jj int64 = 25

		kk float32 = 31.
		ll float64 = 32.

		mm complex64  = 41 - 0.707i
		nn complex128 = 42 - 0.707i

		oo string = "string"
		pp rune   = 'r'
	)

	const (
		Sunday = iota
		Monday
	)

	type Day int
	const (
		DaySunday = iota
		DayMonday
	)

	type DayName string
	const DayNameMonday DayName = "Monday"

	var expected string
	var tt struct{ expected string }
	ttp := &struct{ expected string }{}
	{{- end }}

	{{ range $pi, $pkg := $.Pkgs }}
	t.Run("{{ $pkg }}", func(t *testing.T) {
		{{- range $ti, $test := $.Tests }}
		// {{ $test.Name }}.
		{
			{{- range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.InvalidChecks }}
				{{ NewCheckerExpander.Expand $check $pkg $vars }}
				{{ end -}}
			}
			{{ end }}
			// Valid.
			{{ range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.ValidChecks }}
				{{ NewCheckerExpander.Expand $check $pkg $vars }}
				{{ end -}}
			}
			{{ end -}}
		}
		{{ end -}}
	})
	{{ end }}

	assObj, reqObj := assert.New(t), require.New(t)

	{{ range $pi, $obj := $.Objs }}
	t.Run("{{ $obj }}", func(t *testing.T) {
		{{- range $ti, $test := $.Tests }}
		// {{ $test.Name }}.
		{
			{{- range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.InvalidChecks }}
				{{ NewCheckerExpander.WithoutTArg.Expand $check $obj $vars }}
				{{ end -}}
			}
			{{ end }}
			// Valid.
			{{ range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.ValidChecks }}
				{{ NewCheckerExpander.WithoutTArg.Expand $check $obj $vars }}
				{{ end -}}
			}
			{{ end -}}
		}
		{{ end -}}
	})
	{{ end -}}
}

type ConfusedWithExpectedActualSuite struct {
	suite.Suite
}

func TestConfusedWithExpectedActualSuite(t *testing.T) {
	suite.Run(t, new(ConfusedWithExpectedActualSuite))
}

func (s *ConfusedWithExpectedActualSuite) TestAll() {
		{{ template "vars" .}}

	assObj, reqObj := s.Assert(), s.Require()

	{{- range $ti, $test := $.Tests }}
	// {{ $test.Name }}.
	{
		{{- range $si, $sel := $.SuiteSelectors }}
		{
			{{- range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.InvalidChecks }}
				{{ NewCheckerExpander.WithoutTArg.Expand $check $sel $vars }}
				{{ end -}}
			}
			{{ end }}
			// Valid.
			{{ range $vi, $vars := $test.VarSets }}
			{
				{{- range $ci, $check := $test.ValidChecks }}
				{{ NewCheckerExpander.WithoutTArg.Expand $check $sel $vars }}
				{{ end -}}
			}
			{{ end -}}
		}
		{{ end -}}
	}
	{{ end -}}
}
`
