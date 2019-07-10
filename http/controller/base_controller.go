package controller

import (
	"github.com/totoval/framework/auth"
	"github.com/totoval/framework/policy"
	"github.com/totoval/framework/validator"
)

type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
}
