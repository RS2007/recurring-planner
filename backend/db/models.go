package db

type UserModel struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	AccentColor string `json:"accentColor" binding:"required"`
}

type UserModelWithId struct {
	UserId int32 `json:"uuserId" binding:"required"`
	UserModel
}

type EventModel struct {
	Location    string `json:"location" binding:"required"`
	Summary     string `json:"summary" binding:"required"`
	StartTime   string `json:"startTime" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type EventModelWithIds struct {
	EventId int32 `json:"uuserId" binding:"required"`
	EventModel
}

type TagModel struct {
	Name string `json:"name" binding:"required"`
}

type TemplateModel struct {
	Name     string              `json:"name" binding:"required"`
	Tags     []TagModel          `json:"tags" binding:"required"`
	Events   []EventModelWithIds `json:"events" binding:"required"`
	TimeZone string              `json:"timeZone" binding:"required"`
}

type TemplateModelWithIds struct {
	TemplateId int32 `json:"templateId" binding:"required"`
	TemplateModel
}
