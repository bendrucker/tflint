// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaFunctionInvalidRuntimeRule checks the pattern is valid
type AwsLambdaFunctionInvalidRuntimeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLambdaFunctionInvalidRuntimeRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidRuntimeRule() *AwsLambdaFunctionInvalidRuntimeRule {
	return &AwsLambdaFunctionInvalidRuntimeRule{
		resourceType:  "aws_lambda_function",
		attributeName: "runtime",
		enum: []string{
			"nodejs",
			"nodejs4.3",
			"nodejs6.10",
			"nodejs8.10",
			"nodejs10.x",
			"nodejs12.x",
			"java8",
			"java11",
			"python2.7",
			"python3.6",
			"python3.7",
			"python3.8",
			"dotnetcore1.0",
			"dotnetcore2.0",
			"dotnetcore2.1",
			"nodejs4.3-edge",
			"go1.x",
			"ruby2.5",
			"provided",
		},
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidRuntimeRule) Name() string {
	return "aws_lambda_function_invalid_runtime"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidRuntimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidRuntimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidRuntimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidRuntimeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`runtime is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
