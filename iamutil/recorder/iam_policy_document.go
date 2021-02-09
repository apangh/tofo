package recorder

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/apangh/tofo/model"
)

var (
	PolicyDocInvalid = errors.New("PolicyDocInvalid")
)

type JsonIamPolicyDocument struct {
	Version    string                    `json:",omitempty"`
	Id         string                    `json:",omitempty"`
	Statements []*JsonIamPolicyStatement `json:"Statement"`
}

type JsonIamPolicyStatement struct {
	Sid           string
	Effect        string                           `json:",omitempty"`
	Actions       JsonIamPolicyActions             `json:"Action,omitempty"`
	NotActions    JsonIamPolicyActions             `json:"NotAction,omitempty"`
	Resources     JsonIamPolicyResources           `json:"Resource,omitempty"`
	NotResources  JsonIamPolicyResources           `json:"NotResource,omitempty"`
	Principals    JsonIamPolicyStatementPrincipals `json:"Principal,omitempty"`
	NotPrincipals JsonIamPolicyStatementPrincipals `json:"NotPrincipal,omitempty"`
	Conditions    JsonIamPolicyStatementConditions `json:"Condition,omitempty"`
}

type JsonIamPolicyStatementPrincipals model.IamPolicyStatementPrincipals
type JsonIamPolicyStatementConditions model.IamPolicyStatementConditions
type JsonIamPolicyActions model.IamPolicyActions
type JsonIamPolicyResources model.IamPolicyResources

func ToIamPolicyDocument(s string) (*model.IamPolicyDocument, error) {
	res, e := url.QueryUnescape(s)
	if e != nil {
		return nil, e
	}
	var pd JsonIamPolicyDocument
	if e := json.Unmarshal([]byte(res), &pd); e != nil {
		return nil, e
	}

	var stmt []*model.IamPolicyStatement
	for _, s := range pd.Statements {
		stmt = append(stmt, &model.IamPolicyStatement{
			Sid:           s.Sid,
			Effect:        s.Effect,
			Actions:       model.IamPolicyActions(s.Actions),
			NotActions:    model.IamPolicyActions(s.NotActions),
			Resources:     model.IamPolicyResources(s.Resources),
			NotResources:  model.IamPolicyResources(s.NotResources),
			Principals:    model.IamPolicyStatementPrincipals(s.Principals),
			NotPrincipals: model.IamPolicyStatementPrincipals(s.NotPrincipals),
			Conditions:    model.IamPolicyStatementConditions(s.Conditions),
		})
	}

	doc := &model.IamPolicyDocument{
		Version:    pd.Version,
		Id:         pd.Version,
		Statements: stmt,
	}
	return doc, nil
}
func (pr *JsonIamPolicyResources) UnmarshalJSON(b []byte) error {
	var data interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var out JsonIamPolicyResources
	switch vt := data.(type) {
	case string:
		values := []string{data.(string)}
		out.IsString = true
		out.Values = values
	case []interface{}:
		values := []string{}
		for _, v := range data.([]interface{}) {
			values = append(values, v.(string))
		}
		out.Values = values
	default:
		return fmt.Errorf("%w: Unsupported data type %T for %T.Values",
			PolicyDocInvalid, vt, *pr)
	}
	*pr = out
	return nil
}

func (pa *JsonIamPolicyActions) UnmarshalJSON(b []byte) error {
	var data interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var out JsonIamPolicyActions
	switch vt := data.(type) {
	case string:
		values := []string{data.(string)}
		out.IsString = true
		out.Values = values
	case []interface{}:
		values := []string{}
		for _, v := range data.([]interface{}) {
			values = append(values, v.(string))
		}
		out.Values = values
	default:
		return fmt.Errorf("%w: Unsupported data type %T for %T.Values",
			PolicyDocInvalid, vt, *pa)
	}
	*pa = out
	return nil
}

func (cs *JsonIamPolicyStatementConditions) UnmarshalJSON(b []byte) error {

	var data map[string]map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var out JsonIamPolicyStatementConditions
	for test_key, test_value := range data {
		for var_key, var_values := range test_value {
			var values []string
			switch vt := var_values.(type) {
			case string:
				values = []string{var_values.(string)}
			case []interface{}:
				for _, v := range var_values.([]interface{}) {
					values = append(values, v.(string))
				}
			default:
				return fmt.Errorf("%w: Unsupported data type %T for %T.Values",
					PolicyDocInvalid, vt, out)
			}
			out = append(out, model.IamPolicyStatementCondition{
				Test:     test_key,
				Variable: var_key,
				Values:   values,
			})
		}
	}

	*cs = out
	return nil
}

func (ps *JsonIamPolicyStatementPrincipals) UnmarshalJSON(b []byte) error {

	var data interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var out JsonIamPolicyStatementPrincipals

	switch t := data.(type) {
	case string:
		values := []string{"*"}
		out = append(out, model.IamPolicyStatementPrincipal{
			Type:        "*",
			Identifiers: values})
	case map[string]interface{}:
		for key, value := range data.(map[string]interface{}) {
			switch vt := value.(type) {
			case string:
				values := []string{value.(string)}
				out = append(out, model.IamPolicyStatementPrincipal{
					Type:        key,
					Identifiers: values,
					IsString:    true,
				})
			case []interface{}:
				values := []string{}
				for _, v := range value.([]interface{}) {
					values = append(values, v.(string))
				}
				out = append(out, model.IamPolicyStatementPrincipal{
					Type:        key,
					Identifiers: values,
				})
			default:
				return fmt.Errorf("%w: Unsupported data type %T for %T.Identifiers",
					PolicyDocInvalid, vt, *ps)
			}
		}
	default:
		return fmt.Errorf("%w: Unsupported data type %T for %T",
			PolicyDocInvalid, t, *ps)
	}

	*ps = out
	return nil
}
