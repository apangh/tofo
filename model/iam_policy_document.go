package model

type IamPolicyDocument struct {
	Version    string
	Id         string
	Statements []*IamPolicyStatement
}

type IamPolicyStatement struct {
	Sid           string
	Effect        string
	Actions       IamPolicyActions
	NotActions    IamPolicyActions
	Resources     IamPolicyResources
	NotResources  IamPolicyResources
	Principals    IamPolicyStatementPrincipals
	NotPrincipals IamPolicyStatementPrincipals
	Conditions    IamPolicyStatementConditions
}

type IamPolicyStatementPrincipal struct {
	Type        string
	IsString    bool
	Identifiers []string
}
type IamPolicyStatementPrincipals []IamPolicyStatementPrincipal

type IamPolicyStatementCondition struct {
	Test     string
	Variable string
	Values   []string
}
type IamPolicyStatementConditions []IamPolicyStatementCondition

type IamPolicyActions struct {
	IsString bool
	Values   []string
}

type IamPolicyResources struct {
	IsString bool
	Values   []string
}
