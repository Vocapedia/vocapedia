package stream

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func StartStream(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	user := token.User(r)
	db := database.Manager()

	// Check if stream exists
	var stream entities.Stream
	result := db.Where("room_id = ?", room).First(&stream)

	if result.Error != nil {
		// Create new stream
		userID, _ := strconv.ParseInt(user.UserID, 10, 64)
		stream = entities.Stream{
			RoomID:      room,
			Title:       "Video Stream",
			Description: "Interactive video session",
			Password:    token.GenerateDeterministicToken(room, 8),
			CreatedBy:   userID,
			IsActive:    true,
			ScheduledAt: time.Now(),
			Lang:        "en",
			TargetLang:  "tr",
		}

		if err := db.Create(&stream).Error; err != nil {
			http.Error(w, "Failed to create stream", http.StatusInternalServerError)
			return
		}
	}

	// Update stream as active if it was inactive
	if !stream.IsActive {
		db.Model(&stream).Update("is_active", true)
		db.Model(&stream).Update("start_time", time.Now())
	}

	// Return stream details
	render.JSON(w, r, map[string]any{
		"room_id":      stream.RoomID,
		"lang":         stream.Lang,
		"target_lang":  stream.TargetLang,
		"scheduled_at": stream.ScheduledAt,
		"password":     stream.Password,
		"title":        stream.Title,
		"description":  stream.Description,
		"is_active":    stream.IsActive,
	})
}

// CreateStream creates a new stream with schedule and languages
func CreateStream(w http.ResponseWriter, r *http.Request) {
	// Parse request body for schedule and languages
	var req struct {
		Lang        string    `json:"lang"`
		TargetLang  string    `json:"target_lang"`
		ScheduledAt time.Time `json:"scheduled_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Generate unique ID
	id := snowflake.GenerateID()
	roomID := strconv.FormatInt(id, 10)
	// Get authenticated user
	user := token.User(r)
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)
	// Build stream entity
	stream := entities.Stream{
		RoomID:      roomID,
		Lang:        req.Lang,
		TargetLang:  req.TargetLang,
		ScheduledAt: req.ScheduledAt,
		Title:       "Video Stream",
		Description: "Interactive video session",
		Password:    token.GenerateDeterministicToken(roomID, 8),
		CreatedBy:   userID,
		IsActive:    false, // will set active on start
	}
	db := database.Manager()
	if err := db.Create(&stream).Error; err != nil {
		http.Error(w, "Failed to create stream", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"room": roomID})
}

// GetActiveStreams returns streams that are currently live (scheduled_at <= now && scheduled_at >= now-20m)
func GetActiveStreams(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var streams []entities.Stream
	now := time.Now()
	past := now.Add(-20 * time.Minute)
	err := db.Where("scheduled_at <= ? AND scheduled_at >= ?", now, past).
		Order("scheduled_at DESC").
		Find(&streams).Error
	if err != nil {
		http.Error(w, "Failed to fetch active streams", http.StatusInternalServerError)
		return
	}
	for i := range streams {
		streams[i].Participants = 1 + (int(streams[i].ID) % 5)
	}
	render.JSON(w, r, streams)
}

// GetRecentStreams returns streams that have ended (scheduled_at < now-20m)
func GetRecentStreams(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var streams []entities.Stream
	now := time.Now()
	past := now.Add(-20 * time.Minute)
	err := db.Where("scheduled_at < ?", past).
		Order("scheduled_at DESC").
		Find(&streams).Error
	if err != nil {
		http.Error(w, "Failed to fetch recent streams", http.StatusInternalServerError)
		return
	}
	for i := range streams {
		streams[i].Participants = 1 + (int(streams[i].ID) % 8)
	}
	render.JSON(w, r, streams)
}

// GetUpcomingStreams returns streams scheduled within next 12h
func GetUpcomingStreams(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var streams []entities.Stream
	now := time.Now()
	future := now.Add(12 * time.Hour)
	err := db.Where("scheduled_at > ? AND scheduled_at <= ?", now, future).
		Order("scheduled_at ASC").
		Find(&streams).Error
	if err != nil {
		http.Error(w, "Failed to fetch upcoming streams", http.StatusInternalServerError)
		return
	}
	for i := range streams {
		streams[i].Participants = 1 + (int(streams[i].ID) % 5)
	}
	render.JSON(w, r, streams)
}

func EndStream(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	user := token.User(r)
	db := database.Manager()

	var stream entities.Stream
	result := db.Where("room_id = ?", room).First(&stream)

	if result.Error != nil {
		http.Error(w, "Stream not found", http.StatusNotFound)
		return
	}

	// Check if user is the creator
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)
	if stream.CreatedBy != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// End the stream
	now := time.Now()
	db.Model(&stream).Updates(map[string]interface{}{
		"is_active": false,
		"end_time":  &now,
	})

	render.JSON(w, r, map[string]any{
		"message": "Stream ended successfully",
	})
}
