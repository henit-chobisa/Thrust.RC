package enums

type Task int

const (
	Check_Initial_Configuration Task = iota + 1
	Pull_Containers
	Run_containers
	Setup_Project_Environment
	Start_Companion_Container
	Show_Companion_Logs
)

func (t Task) String() string {
	return [...]string{"Checking Initial Configuration", "Pull Containers", "Run the containers", "Setup Project Environment", "Pull Companion Containers", "Companion Logs"}[t-1]
}
