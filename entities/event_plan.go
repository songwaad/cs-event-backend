package entities

type EventPlan struct {
	EventPlanID uint   `json:"event_plan_id" gorm:"primaryKey"`
	WorkPlan    string `json:"work_plane"`
	Work        string `json:"work"`
}
