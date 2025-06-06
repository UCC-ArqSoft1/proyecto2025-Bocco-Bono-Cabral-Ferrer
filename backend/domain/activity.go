package domain

type Activity struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Capacity    int                `json:"capacity"`
	Category    string             `json:"category"`
	Profesor    string             `json:"profesor"`
	ImageUrl    string             `json:"image_url"`
	Schedules   []ActivitySchedule `json:"schedules"`
}

type ActivitySchedule struct {
	Id         int    `json:"id"`
	ActivityId int    `json:"activity_id"`
	Day        string `json:"day"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type Activities []Activity
