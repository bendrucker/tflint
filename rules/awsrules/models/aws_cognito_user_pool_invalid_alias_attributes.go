// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoUserPoolInvalidAliasAttributesRule checks the pattern is valid
type AwsCognitoUserPoolInvalidAliasAttributesRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCognitoUserPoolInvalidAliasAttributesRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidAliasAttributesRule() *AwsCognitoUserPoolInvalidAliasAttributesRule {
	return &AwsCognitoUserPoolInvalidAliasAttributesRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "alias_attributes",
		enum: []string{
			"phone_number",
			"email",
			"preferred_username",
		},
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidAliasAttributesRule) Name() string {
	return "aws_cognito_user_pool_invalid_alias_attributes"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidAliasAttributesRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCognitoUserPoolInvalidAliasAttributesRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidAliasAttributesRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidAliasAttributesRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

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
					`alias_attributes is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}