package controllers

import (
	"better_calendar_backend/db"
	"better_calendar_backend/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func CreateNewTemplate(c *gin.Context) {
	var newTemplate db.TemplateModel
	err := c.ShouldBindJSON(&newTemplate)
	fmt.Println(newTemplate)
	utils.ErrorHandler(err)
	utils.ErrorHandler(err)
	db := db.GetDB()
	trx, err := db.Begin()
	utils.ErrorHandler(err)
	defer trx.Rollback()

	_ = db.QueryRow("SET TIME ZONE $1", newTemplate.TimeZone)
	utils.ErrorHandler(err)
	var eventIds []int32
	var tagIds []int32
	for i := 0; i < len(newTemplate.Events); i++ {
		var eventId int32
		event := newTemplate.Events[i]
		err := trx.QueryRow(`INSERT INTO events (startTime,endTime,location,summary,description) VALUES ($1,$2,$3,$4,$5) RETURNING eventId`, event.StartTime, event.EndTime, event.Location, event.Summary, event.Description).Scan(&eventId)
		if err != nil {
			fmt.Println("Error from event insertion")
			fmt.Println("err")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eventIds = append(eventIds, eventId)
	}

	for i := 0; i < len(newTemplate.Tags); i++ {
		var tagId int32
		tag := newTemplate.Tags[i]
		err := trx.QueryRow(`INSERT INTO tags (name) VALUES ($1)  ON CONFLICT (name) DO NOTHING RETURNING tagId`, tag.Name).Scan(&tagId)
		if err != nil {
			fmt.Println("Error from tag insertion")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tagIds = append(tagIds, tagId)
	}
	var templateId int32
	userId := c.MustGet("userId")
	fmt.Println(tagIds)
	fmt.Println(eventIds)
	err = trx.QueryRow(`INSERT INTO templates (name,userId,timeZone) VALUES ($1,$2,$3) RETURNING templateId`, newTemplate.Name, userId, newTemplate.TimeZone).Scan(&templateId)
	utils.ErrorHandler(err)
	for _, tagIdElem := range tagIds {
		_, err := trx.Exec(`INSERT INTO tag_templates (tagId,templateId) VALUES ($1,$2)`, tagIdElem, templateId)
		utils.ErrorHandler(err)
	}
	for _, eventIdElem := range eventIds {
		_, err := trx.Exec(`INSERT INTO event_templates (eventId,templateId) VALUES ($1,$2)`, eventIdElem, templateId)
		utils.ErrorHandler(err)
	}
	if err := trx.Commit(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template created succesfully"})
}

func GetAllTemplates(c *gin.Context) {
	userId := c.MustGet("userId")
	Db := db.GetDB()
	templates, err := Db.Query("SELECT templateId,name FROM templates WHERE userId = $1", userId)
	utils.ErrorHandler(err)
	var templatesArray []db.TemplateModelWithIds
	for templates.Next() {
		var newTemplate db.TemplateModelWithIds
		err := templates.Scan(&newTemplate.TemplateId, &newTemplate.Name)
		utils.ErrorHandler(err)
		tagRows, err := Db.Query("SELECT tags.name AS tagName FROM templates JOIN tag_templates ON templates.templateId = tag_templates.templateId JOIN tags ON tag_templates.tagId = tags.tagId WHERE templates.templateId = $1", newTemplate.TemplateId)
		utils.ErrorHandler(err)
		eventRows, err := Db.Query("SELECT events.* FROM templates JOIN event_templates ON templates.templateId = event_templates.templateId JOIN events ON event_templates.eventId = events.eventId WHERE templates.templateId = $1", newTemplate.TemplateId)
		utils.ErrorHandler(err)
		var tags []db.TagModel
		var events []db.EventModelWithIds
		for tagRows.Next() {
			var newTag db.TagModel
			tagRows.Scan(&newTag.Name)
			tags = append(tags, newTag)
		}
		for eventRows.Next() {
			var newEvent db.EventModelWithIds
			eventRows.Scan(&newEvent.EventId, &newEvent.StartTime, &newEvent.EndTime, &newEvent.Location, &newEvent.Summary, &newEvent.Description)
			events = append(events, newEvent)
		}
		newTemplate.Events = events
		newTemplate.Tags = tags
		templatesArray = append(templatesArray, newTemplate)

	}

	c.IndentedJSON(http.StatusOK, gin.H{"templates": templatesArray})

}

func AddTemplateToDay(c *gin.Context) {
	Db := db.GetDB()
	type createTemplateRequestBody struct {
		TemplateId  int32  `json:"templateId"`
		DateString  string `json:"dateString"`
		AccessToken string `json:"accessToken"`
		IdToken     string `json:"idToken"`
	}
	var reqBody createTemplateRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	config := utils.SetupGoogleConfig()
	client := config.Client(context.Background(), &oauth2.Token{
		AccessToken: reqBody.AccessToken,
	})
	calendarService, err := calendar.New(client)
	utils.ErrorHandler(err)
	var eventIds []int32
	queryRows, err := Db.Query("SELECT eventId FROM event_templates WHERE templateId =$1", reqBody.TemplateId)
	utils.ErrorHandler(err)
	for queryRows.Next() {
		var eventId int32
		queryRows.Scan(&eventId)
		eventIds = append(eventIds, eventId)
	}

	query := "SELECT * FROM events WHERE eventId in ("
	// run a loop through eventId array and add the numbers to the query string in a space seperated fashion, once looping is done add a closing paranthesis
	for i, eventId := range eventIds {
		query += fmt.Sprintf("%d", eventId)
		if i != len(eventIds)-1 {
			query += ","
		} else {
			query += ")"
		}
	}
	eventRows, err := Db.Query(query)
	utils.ErrorHandler(err)
	var events []db.EventModelWithIds
	for eventRows.Next() {
		var event db.EventModelWithIds
		eventRows.Scan(&event.EventId, &event.StartTime, &event.EndTime, &event.Location, &event.Summary, &event.Description)
		events = append(events, event)
	}

	for _, event := range events {
		fmt.Println("Day string: ", reqBody.DateString)
		fmt.Println("Time string: ", event.StartTime)
		formattedStartTimeFromDB, err := utils.GetFormattedTimeString(event.StartTime)
		utils.ErrorHandler(err)
		formattedEndTimeFromDB, err := utils.GetFormattedTimeString(event.EndTime)
		utils.ErrorHandler(err)
		startTime, err := utils.GetDateTimeString(reqBody.DateString, formattedStartTimeFromDB, "Asia/Kolkata")
		fmt.Println("start Time:", startTime)
		utils.ErrorHandler(err)
		endTime, err := utils.GetDateTimeString(reqBody.DateString, formattedEndTimeFromDB, "Asia/Kolkata")
		fmt.Println("end Time:", endTime)
		utils.ErrorHandler(err)
		eventForGoogle := &calendar.Event{
			Summary:     event.Summary,
			Location:    event.Location,
			Description: event.Description,
			Start: &calendar.EventDateTime{
				DateTime: startTime,
				TimeZone: "Asia/Kolkata",
			},
			End: &calendar.EventDateTime{
				DateTime: endTime,
				TimeZone: "Asia/Kolkata",
			},
		}
		_, err = calendarService.Events.Insert("primary", eventForGoogle).Do()
		utils.ErrorHandler(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})

}
