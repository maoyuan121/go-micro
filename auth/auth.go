// Package auth provides authentication and authorization capability
package auth

import (
	"context"
	"errors"
	"time"
)

const (
	// BearerScheme used for Authorization header
	BearerScheme = "Bearer "
	// ScopePublic is the scope applied to a rule to allow access to the public
	ScopePublic = ""
	// ScopeAccount is the scope applied to a rule to limit to users with any valid account
	ScopeAccount = "*"
)

var (
	// ErrInvalidToken is when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrForbidden is when a user does not have the necessary scope to access a resource
	ErrForbidden = errors.New("resource forbidden")
)

// Auth 提供身份验证和授权
type Auth interface {
	// Init the auth
	Init(opts ...Option)
	// Options set for auth
	Options() Options
	// 生成一个新的 account
	Generate(id string, opts ...GenerateOption) (*Account, error)
	// 检查 token
	Inspect(token string) (*Account, error)
	// 使用刷新令牌或凭据生成的令牌
	Token(opts ...TokenOption) (*Token, error)
	// 返回实现名
	String() string
}

// Rules 管理对资源的访问
type Rules interface {
	// 使用规则验证帐户对资源的访问
	Verify(acc *Account, res *Resource, opts ...VerifyOption) error
	// 授予对资源的访问权限
	Grant(rule *Rule) error
	// 取消对资源的访问
	Revoke(rule *Rule) error
	// List 返回用于验证请求的所有规则
	List(...ListOption) ([]*Rule, error)
}

// 由认证提供者提供的 Account
type Account struct {
	// ID of the account e.g. email
	ID string `json:"id"`
	// Type of the account, e.g. service
	Type string `json:"type"`
	// Issuer of the account
	Issuer string `json:"issuer"`
	// Any other associated metadata
	Metadata map[string]string `json:"metadata"`
	// Scopes the account has access to
	Scopes []string `json:"scopes"`
	// Secret for the account, e.g. the password
	Secret string `json:"secret"`
}

// Token 的生命周期可长可短
type Token struct {
	// 用于访问资源的 token
	AccessToken string `json:"access_token"`
	// 用于生成一个新的 token
	RefreshToken string `json:"refresh_token"`
	// token 的创建时间
	Created time.Time `json:"created"`
	// token 的过期时间
	Expiry time.Time `json:"expiry"`
}

// Expired returns a boolean indicating if the token needs to be refreshed
func (t *Token) Expired() bool {
	return t.Expiry.Unix() < time.Now().Unix()
}

// Resource is an entity such as a user or
type Resource struct {
	// Name of the resource, e.g. go.micro.service.notes
	Name string `json:"name"`
	// Type of resource, e.g. service
	Type string `json:"type"`
	// Endpoint resource e.g NotesService.Create
	Endpoint string `json:"endpoint"`
}

// Access 定义规则授予的访问类型
type Access int

const (
	// AccessGranted to a resource
	AccessGranted Access = iota
	// AccessDenied to a resource
	AccessDenied
)

// Rule 用于验证对资源的访问
type Rule struct {
	// ID of the rule, e.g. "public"
	ID string
	// 规则要求的范围，空白范围表示对公众开放，* 表示规则适用于任何有效帐户
	Scope string
	// 规则应用于的资源
	Resource *Resource
	// Access 决定规则是否授予或拒绝对资源的访问
	Access Access
	// 当验证请求时，规则应该采取的优先级，值越高，规则将越快地被应用
	Priority int32
}

type accountKey struct{}

// AccountFromContext gets the account from the context, which
// is set by the auth wrapper at the start of a call. If the account
// is not set, a nil account will be returned. The error is only returned
// when there was a problem retrieving an account
func AccountFromContext(ctx context.Context) (*Account, bool) {
	acc, ok := ctx.Value(accountKey{}).(*Account)
	return acc, ok
}

// ContextWithAccount sets the account in the context
func ContextWithAccount(ctx context.Context, account *Account) context.Context {
	return context.WithValue(ctx, accountKey{}, account)
}
