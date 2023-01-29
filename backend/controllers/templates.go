package controllers

import (
	"better_calendar_backend/db"
	"better_calendar_backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewTemplate(c *gin.Context) {
	var newTemplate db.TemplateModel
	err := c.ShouldBindJSON(&newTemplate)
	fmt.Println(newTemplate)
	utils.ErrorHandler(err)
	db := db.GetDB()
	trx, err := db.Begin()
	utils.ErrorHandler(err)
	defer trx.Rollback()

	var eventIds []int32
	var tagIds []int32
	for i := 0; i < len(newTemplate.Events); i++ {
		var eventId int32
		event := newTemplate.Events[i]
		err := trx.QueryRow(fmt.Sprintf(`INSERT INTO events (startDateTime,endDateTime,location,summary,description) VALUES (%d,%d,'%s','%s','%s') RETURNING eventId`, event.StartTime, event.EndTime, event.Location, event.Summary, event.Description)).Scan(&eventId)
		fmt.Println(fmt.Sprintf(`INSERT INTO events (startDateTime,endDateTime,location,summary,description) VALUES (%d,%d,'%s','%s','%s') RETURNING eventId`, event.StartTime, event.EndTime, event.Location, event.Summary, event.Description))
		if err != nil {
			fmt.Println("Error from event insertion")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eventIds = append(eventIds, eventId)
	}

	for i := 0; i < len(newTemplate.Tags); i++ {
		var tagId int32
		tag := newTemplate.Tags[i]
		err := trx.QueryRow(fmt.Sprintf(`INSERT INTO tags (name) VALUES ('%s')  ON CONFLICT (name) DO NOTHING RETURNING tagId`, tag.Name)).Scan(&tagId)
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
	err = trx.QueryRow(fmt.Sprintf(`INSERT INTO templates (name,userId) VALUES ('%s',%d) RETURNING templateId`, newTemplate.Name, userId)).Scan(&templateId)
	fmt.Println(fmt.Sprintf(`INSERT INTO templates (name,userId) VALUES ('%s',%d) RETURNING templateId`, newTemplate.Name, userId))
	utils.ErrorHandler(err)
	for _, tagIdElem := range tagIds {
		_, err := trx.Exec(fmt.Sprintf(`INSERT INTO tag_templates (tagId,templateId) VALUES (%d,%d)`, tagIdElem, templateId))
		utils.ErrorHandler(err)
	}
	for _, eventIdElem := range eventIds {
		_, err := trx.Exec(fmt.Sprintf(`INSERT INTO event_templates (eventId,templateId) VALUES (%d,%d)`, eventIdElem, templateId))
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
	templates, err := Db.Query(fmt.Sprintf("SELECT templateId,name FROM templates WHERE userId = %d", userId))
	utils.ErrorHandler(err)
	var templatesArray []db.TemplateModelWithIds
	for templates.Next() {
		var newTemplate db.TemplateModelWithIds
		err := templates.Scan(&newTemplate.TemplateId, &newTemplate.Name)
		utils.ErrorHandler(err)
		tagRows, err := Db.Query(fmt.Sprintf("SELECT tags.name AS tagName FROM templates JOIN tag_templates ON templates.templateId = tag_templates.templateId JOIN tags ON tag_templates.tagId = tags.tagId WHERE templates.templateId = %d", newTemplate.TemplateId))
		eventRows, err := Db.Query(fmt.Sprintf("SELECT events.* FROM templates JOIN event_templates ON templates.templateId = event_templates.templateId JOIN events ON event_templates.eventId = events.eventId WHERE templates.templateId = %d", newTemplate.TemplateId))
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
