package main

import (
	"text/template"
)

type FloatCompareCasesGenerator struct{}

func (g FloatCompareCasesGenerator) Data() any {
	const (
		report     = "float-compare: use %s.%s"
		proposedFn = "InDelta"
	)

	return struct {
		Bits           []string
		Pkgs, Objs     []string
		SuiteSelectors []string
		InvalidChecks  []Check
		ValidChecks    []Check
	}{
		Bits:           []string{"32", "64"},
		Pkgs:           []string{"assert", "require"},
		Objs:           []string{"assObj", "reqObj"},
		SuiteSelectors: []string{"s", "s.Assert()", "assObj", "s.Require()", "reqObj"},
		InvalidChecks: []Check{
			{Fn: "Equal", Argsf: "42.42, a", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "NotEqual", Argsf: "b, cc.c", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "Greater", Argsf: "d, e", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "GreaterOrEqual", Argsf: "(*f).c, *g", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "Less", Argsf: "h.Calculate(), floatOp()", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "LessOrEqual", Argsf: "42.42, a", ReportMsgf: report, ProposedFn: proposedFn},

			{Fn: "True", Argsf: "a == d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "True", Argsf: "a != d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "True", Argsf: "a > d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "True", Argsf: "a >= d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "True", Argsf: "a < d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "True", Argsf: "a <= d", ReportMsgf: report, ProposedFn: proposedFn},

			{Fn: "False", Argsf: "a == d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "False", Argsf: "a != d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "False", Argsf: "a > d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "False", Argsf: "a >= d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "False", Argsf: "a < d", ReportMsgf: report, ProposedFn: proposedFn},
			{Fn: "False", Argsf: "a <= d", ReportMsgf: report, ProposedFn: proposedFn},
		},
		ValidChecks: []Check{
			{Fn: "InDelta", Argsf: "42.42, a, 0.0001"},
			{Fn: "InDelta", Argsf: "b, cc.c, 0.0001"},
			{Fn: "InDelta", Argsf: "(*f).c, *g, 0.0001"},
			{Fn: "InDelta", Argsf: "h.Calculate(), floatOp(), 0.0001"},
			{Fn: "InDelta", Argsf: "42.42, a, 0.0001"},
		},
	}
}

func (g FloatCompareCasesGenerator) ErroredTemplate() *template.Template {
	return template.Must(template.New("FloatCompareCasesGenerator.ErroredTemplate").
		Funcs(fm).
		Parse(floatCompareCasesTmplText))
}

func (g FloatCompareCasesGenerator) GoldenTemplate() *template.Template {
	return nil
}

const floatCompareCasesTmplText = header + `

package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)
{{ range $bi, $bits := .Bits }}
func TestFloat{{ $bits }}Compare(t *testing.T) {
	type number float{{ $bits }}
	type withFloat{{ $bits }} struct{ c float{{ $bits }} }
	floatOp := func() float{{ $bits }} { return 0. }

	var a float{{ $bits }}
	var b number
	var cc withFloat{{ $bits }}
	d := float{{ $bits }}(1.01)
	const e = float{{ $bits }}(2.02)
	f := new(withFloat{{ $bits }})
	var g *float{{ $bits }}
	var h withFloat{{ $bits }}Method

	{{ range $pi, $pkg := $.Pkgs }}
	t.Run("{{ $pkg }}", func(t *testing.T) {
		{
			{{- range $ci, $check := $.InvalidChecks }}
			{{ NewCheckerExpander.Expand $check $pkg nil }}
			{{ end -}}
		}

		// Valid.

		{
			{{- range $ci, $check := $.ValidChecks }}
			{{ NewCheckerExpander.Expand $check $pkg nil }}
			{{ end -}}
		}
	})
	{{ end }}

	assObj, reqObj := assert.New(t), require.New(t)

	{{ range $oi, $obj := $.Objs }}
	t.Run("{{ $obj }}", func(t *testing.T) {
		{
			{{- range $ci, $check := $.InvalidChecks }}
			{{ NewCheckerExpander.WithoutTArg.Expand $check $obj nil }}
			{{ end -}}
		}

		// Valid.

		{
			{{- range $ci, $check := $.ValidChecks }}
			{{ NewCheckerExpander.WithoutTArg.Expand $check $obj nil }}
			{{ end -}}
		}
	})
	{{ end -}}
}

type Float{{ $bits }}CompareSuite struct {
	suite.Suite
}

func TestFloat{{ $bits }}CompareSuite(t *testing.T) {
	suite.Run(t, new(Float{{ $bits }}CompareSuite))
}

func (s *Float{{ $bits }}CompareSuite) TestAll() {
	type number float{{ $bits }}
	type withFloat{{ $bits }} struct{ c float{{ $bits }} }
	floatOp := func() float{{ $bits }} { return 0. }

	var a float{{ $bits }}
	var b number
	var cc withFloat{{ $bits }}
	d := float{{ $bits }}(1.01)
	const e = float{{ $bits }}(2.02)
	f := new(withFloat{{ $bits }})
	var g *float{{ $bits }}
	var h withFloat{{ $bits }}Method

	assObj, reqObj := s.Assert(), s.Require()

	{{ range $si, $sel := $.SuiteSelectors }}
	{
		{
			{{- range $ci, $check := $.InvalidChecks }}
			{{ NewCheckerExpander.WithoutTArg.Expand $check $sel nil }}
			{{ end -}}
		}

		// Valid.

		{
			{{- range $ci, $check := $.ValidChecks }}
			{{ NewCheckerExpander.WithoutTArg.Expand $check $sel nil }}
			{{ end -}}
		}
	}
	{{ end -}}
}

type withFloat{{ $bits }}Method struct{}

func (withFloat{{ $bits }}Method) Calculate() float{{ $bits }} { return 0. }
{{ end }}
`
