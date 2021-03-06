// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCloudformationStackSetInvalidExecutionRoleNameRule checks the pattern is valid
type AwsCloudformationStackSetInvalidExecutionRoleNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudformationStackSetInvalidExecutionRoleNameRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidExecutionRoleNameRule() *AwsCloudformationStackSetInvalidExecutionRoleNameRule {
	return &AwsCloudformationStackSetInvalidExecutionRoleNameRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "execution_role_name",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z_0-9+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidExecutionRoleNameRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_execution_role_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidExecutionRoleNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInvalidExecutionRoleNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidExecutionRoleNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidExecutionRoleNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"execution_role_name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"execution_role_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z_0-9+=,.@-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
