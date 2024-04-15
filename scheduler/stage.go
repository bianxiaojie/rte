package scheduler

type SchedulerStage interface {
	Run(SchedulerLinker)
}
