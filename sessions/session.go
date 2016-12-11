package sessions

import ()

type Session interface {
	IsLogin() bool

	GetUsername() string
	GetUserId() int64
	GetPrivilege() string
	GetIPAddr() string
	Get(key interface{}) (interface{}, error)

	SetUsername(username string)
	SetUserId(user_id int64)
	SetPrivilege(privilege string)
	SetIPAddr(ip string)
	Set(key interface{}, val interface{}) error

	Flashes(vars ...string) []interface{}
	AddFlash(value interface{}, vars ...string)
}