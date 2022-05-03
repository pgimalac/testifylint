package main

import "text/template"

type ComparisonsCasesGenerator struct{}

func (g ComparisonsCasesGenerator) Template() *template.Template {
	return comparisonsCasesTmpl
}

func (g ComparisonsCasesGenerator) Data() any {
	return struct {
		Pkgs          []string
		InvalidChecks []Check
		ValidChecks   []Check
	}{
		Pkgs: []string{"assert", "require"},
		InvalidChecks: []Check{
			{Fn: "True", ArgsTmpl: "t, a == b", ReportedMsgf: "use %s.%s", ProposedFn: "Equal"},
			{Fn: "True", ArgsTmpl: "t, a != b", ReportedMsgf: "use %s.%s", ProposedFn: "NotEqual"},
			{Fn: "True", ArgsTmpl: "t, a > b", ReportedMsgf: "use %s.%s", ProposedFn: "Greater"},
			{Fn: "True", ArgsTmpl: "t, a >= b", ReportedMsgf: "use %s.%s", ProposedFn: "GreaterOrEqual"},
			{Fn: "True", ArgsTmpl: "t, a < b", ReportedMsgf: "use %s.%s", ProposedFn: "Less"},
			{Fn: "True", ArgsTmpl: "t, a <= b", ReportedMsgf: "use %s.%s", ProposedFn: "LessOrEqual"},

			{Fn: "False", ArgsTmpl: "t, a == b", ReportedMsgf: "use %s.%s", ProposedFn: "NotEqual"},
			{Fn: "False", ArgsTmpl: "t, a != b", ReportedMsgf: "use %s.%s", ProposedFn: "Equal"},
			{Fn: "False", ArgsTmpl: "t, a > b", ReportedMsgf: "use %s.%s", ProposedFn: "LessOrEqual"},
			{Fn: "False", ArgsTmpl: "t, a >= b", ReportedMsgf: "use %s.%s", ProposedFn: "Less"},
			{Fn: "False", ArgsTmpl: "t, a < b", ReportedMsgf: "use %s.%s", ProposedFn: "GreaterOrEqual"},
			{Fn: "False", ArgsTmpl: "t, a <= b", ReportedMsgf: "use %s.%s", ProposedFn: "Greater"},
		},
		ValidChecks: []Check{
			{Fn: "Equal", ArgsTmpl: "t, a, b"},
			{Fn: "NotEqual", ArgsTmpl: "t, a, b"},
			{Fn: "Greater", ArgsTmpl: "t, a, b"},
			{Fn: "GreaterOrEqual", ArgsTmpl: "t, a, b"},
			{Fn: "Less", ArgsTmpl: "t, a, b"},
			{Fn: "LessOrEqual", ArgsTmpl: "t, a, b"},
		},
	}
}

var comparisonsCasesTmpl = template.Must(template.New("comparisonsCasesTmpl").
	Funcs(template.FuncMap{
		"ExpandCheck": ExpandCheck,
	}).
	Parse(`// Code generated by testifylint/internal/cmd/testgen. DO NOT EDIT.

package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComparisons(t *testing.T) {
	var a, b int
	{{ range $pi, $pkg := .Pkgs }}
	t.Run("{{ $pkg }}", func(t *testing.T) {
		{
			{{- range $ci, $check := $.InvalidChecks }}
			{{ ExpandCheck $check $pkg nil }}
			{{ end }}}

		// Valid {{ $pkg }}s.

		{
			{{- range $ci, $check := $.ValidChecks }}
			{{ ExpandCheck $check $pkg nil }}
			{{ end }}}
	})
	{{ end }}} 
`))