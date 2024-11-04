package constant

type TiktokQueueStatus string

const (
	TiktokQueueStatusNone      TiktokQueueStatus = "none"
	TiktokQueueStatusInitiated TiktokQueueStatus = "initiated"
	TiktokQueueStatusFinished  TiktokQueueStatus = "finished"
)
