package policy

import (
	"github.com/totoval/framework/auth"
)

type key = string
type value = string
type Policier interface {
	Before(IUser auth.IUser, routeParamMap map[key]value) *bool
	Create(IUser auth.IUser, routeParamMap map[key]value) bool
	Update(IUser auth.IUser, routeParamMap map[key]value) bool
	Delete(IUser auth.IUser, routeParamMap map[key]value) bool
	ForceDelete(IUser auth.IUser, routeParamMap map[key]value) bool
	View(IUser auth.IUser, routeParamMap map[key]value) bool
	Restore(IUser auth.IUser, routeParamMap map[key]value) bool
}

type Action byte

const (
	ActionCreate Action = iota
	ActionUpdate
	ActionDelete
	ActionForceDelete
	ActionView
	ActionRestore
)

type Authorization struct {
	auth.RequestUser
}

func (a *Authorization) Authorize(c auth.Context, policies Policier, action Action) (permit bool, user auth.IUser) {
	if a.RequestUser.Scan(c) {
		return false, nil
	}
	user = a.RequestUser.User()

	// if use Authorize func, routeParamMap is nil
	return policyValidate(user, policies, action, nil), user
}

func policyValidate(user auth.IUser, policies Policier, action Action, routeParamMap map[key]value) bool {
	if user == nil {
		return true
	}
	if policies == nil {
		return true
	}

	if beforeResult := policies.Before(user, routeParamMap); beforeResult != nil {
		return *beforeResult
	}

	switch action {
	case ActionCreate:
		return policies.Create(user, routeParamMap)
	case ActionUpdate:
		return policies.Update(user, routeParamMap)
	case ActionDelete:
		return policies.Delete(user, routeParamMap)
	case ActionForceDelete:
		return policies.ForceDelete(user, routeParamMap)
	case ActionView:
		return policies.View(user, routeParamMap)
	case ActionRestore:
		return policies.Restore(user, routeParamMap)
	default:
		return false
	}
}
