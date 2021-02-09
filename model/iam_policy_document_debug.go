package model

import "fmt"

func (i *IamPolicyDocument) Dump() {
	fmt.Printf("Version: %s\n", i.Version)
	fmt.Printf("Id: %s\n", i.Id)
	for _, s := range i.Statements {
		s.Dump()
	}
}

func (s *IamPolicyStatement) Dump() {
	fmt.Printf("\tSid: %s\n", s.Sid)
	fmt.Printf("\tEffect: %s\n", s.Effect)
	s.Actions.Dump("Action")
	s.NotActions.Dump("NotAction")
	s.Resources.Dump("Resource")
	s.NotResources.Dump("NotResource")
	for _, p := range s.Principals {
		p.Dump("Principal")
	}
	for _, np := range s.NotPrincipals {
		np.Dump("NotPrincipal")
	}
	for _, c := range s.Conditions {
		c.Dump("Condition")
	}
}

func (a *IamPolicyActions) Dump(t string) {
	if a.IsString {
		fmt.Printf("\t\t%s: %s\n", t, a.Values[0])
	} else if len(a.Values) != 0 {
		fmt.Printf("\t\t%s: ", t)
		for _, v := range a.Values {
			fmt.Printf("\n\t\t\t%s ", v)
		}
		fmt.Printf("\n")
	}
}

func (r *IamPolicyResources) Dump(t string) {
	if r.IsString {
		fmt.Printf("\t\t%s: %s\n", t, r.Values[0])
	} else if len(r.Values) != 0 {
		fmt.Printf("\t\t%s: ", t)
		for _, v := range r.Values {
			fmt.Printf("\n\t\t\t%s ", v)
		}
		fmt.Printf("\n")
	}
}

func (p *IamPolicyStatementPrincipal) Dump(t string) {
	if p.IsString {
		fmt.Printf("\t\t%s: %s - %s\n", t, p.Type, p.Identifiers[0])
	} else {
		fmt.Printf("\t\t%s: %s -", t, p.Type)
		for _, v := range p.Identifiers {
			fmt.Printf("\n\t\t\t%s ", v)
		}
		fmt.Printf("\n")
	}
}

func (p *IamPolicyStatementCondition) Dump(t string) {
	fmt.Printf("\t\t%s %s %s", t, p.Test, p.Variable)
	for _, v := range p.Values {
		fmt.Printf("\n\t\t\t%s ", v)
	}
	fmt.Printf("\n")
}
