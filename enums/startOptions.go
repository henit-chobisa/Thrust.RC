package enums

type StartOption int

const (
	config StartOption = iota + 1
	watcher
	username
	email
	password
	name
	virtual
	deps
	composeFilePath
)

func (s StartOption) String() string {
	return [...]string{"config", "watcher", "username", "email", "password", "name", "virtual", "deps", "composeFilePath"}[s-1]
}
