package task

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

/*
- Pending: initial state, starting point for every task
- Scheduled: a task moves to this state once the manager has scheduled it onto a worker
- Running: a task moves to this state when a worker sucessfully start the task
- Completed: a task moves to this state when it completes its work in a normal way (not Failed)
- Failed: if a task does fail, it moves to this states
*/

func ValidStateTransition(src State, dst State) bool {
	// 0 1
	if (src == Pending && dst == Scheduled) { return true }
	// 0 4
	if (src == Pending && dst == Failed) { return true }
	// 1 2
	if (src == Scheduled && dst == Running) { return true }
	// 1 4
	if (src == Scheduled && dst == Failed) { return true }
	// 2 3
	if (src == Running && dst == Completed) { return true } else { return false}
	// 1 3 => Scheduled => Completed: invalid
}