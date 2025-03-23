package entities

type EventPlane struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	WorkPlan string `json:"work_plane"`
	Work     string `json:"work"`
}
