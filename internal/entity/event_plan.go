package entity

type EventPlan struct {
	EventPlanID uint   `json:"event_plan_id" gorm:"primaryKey"`
	WorkPlan    string `json:"work_plan"`
	Work        string `json:"work"`
}
