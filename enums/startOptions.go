package enums

type StartOption int

const (
	Config StartOption = iota + 1
	Watcher
	Username
	Email
	Password
	Name
	Virtual
	Deps
	ComposeFilePath
)

func (s StartOption) String() string {
	return [...]string{"config", "watcher.watcher", "admin.username", "admin.email", "admin.password", "admin.name", "virtual", "deps", "composeFilePath"}[s-1]
}
