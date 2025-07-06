package stream

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func StartStream(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	db := database.Manager()

	// Check if stream exists
	var stream entities.Stream
	result := db.Preload("Creator").Where("room_id = ?", room).First(&stream)

	if result.Error != nil {
		// Stream doesn't exist, return error
		http.Error(w, "Stream not found", http.StatusNotFound)
		return
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

// GetStreamByID gets a specific stream by room ID
func GetStreamByID(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	db := database.Manager()
	var stream entities.Stream
	result := db.Preload("Creator").Where("room_id = ?", room).First(&stream)

	if result.Error != nil {
		http.Error(w, "Stream not found", http.StatusNotFound)
		return
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
		"participants": 1 + (int(stream.ID) % 5), // Mock participant count
	})
}

// CreateStream creates a new stream with schedule and languages
func CreateStream(w http.ResponseWriter, r *http.Request) {
	// Parse request body for schedule and languages
	var params _createStream
	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if params.Title == "" || params.Description == "" || params.Lang == "" || params.TargetLang == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Generate unique ID
	id := snowflake.GenerateID()
	roomID := strconv.FormatInt(id, 10)

	// Get authenticated user
	user := token.User(r)
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)

	if !user.IsTeacher {
		http.Error(w, "Only teachers can create streams", http.StatusForbidden)
		return
	}
	t, err := time.Parse(time.RFC3339, params.ScheduledAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Build stream entity
	stream := entities.Stream{
		RoomID:          roomID,
		Title:           params.Title,
		Description:     params.Description,
		Lang:            params.Lang,
		TargetLang:      params.TargetLang,
		ScheduledAt:     t,
		Duration:        params.Duration,
		MaxParticipants: params.MaxParticipants,
		Password:        token.GenerateDeterministicToken(roomID, 8),
		CreatedBy:       userID,
		IsActive:        false, // will set active on start
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
	err := db.Preload("Creator").Where("scheduled_at <= ? AND scheduled_at >= ?", now, past).
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
	err := db.Preload("Creator").Where("scheduled_at < ?", past).
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

	// SQL query to filter streams that haven't reached max participants
	err := db.Preload("Creator").
		Preload("Students").
		Where("scheduled_at > ? AND scheduled_at <= ?", now, future).
		Where("(SELECT COUNT(*) FROM stream_students WHERE stream_students.stream_id = streams.id) < max_participants").
		Order("scheduled_at ASC").
		Find(&streams).Error

	if err != nil {
		http.Error(w, "Failed to fetch upcoming streams", http.StatusInternalServerError)
		return
	}

	// Set actual participant count
	for i := range streams {
		streams[i].Participants = len(streams[i].Students)
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
	result := db.Preload("Creator").Where("room_id = ?", room).First(&stream)

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

func JoinStream(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "room is required"})
		return
	}

	userID := token.User(r).UserID

	db := database.Manager()
	var currentUser entities.User
	if err := db.Where("id = ?", userID).First(&currentUser).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		var stream entities.Stream
		if err := tx.Preload("Creator").Preload("Students").Where("room_id = ?", room).First(&stream).Error; err != nil {
			return fmt.Errorf("stream not found: %w", err)
		}

		// Check if stream is upcoming
		if stream.ScheduledAt.Before(time.Now()) {
			return fmt.Errorf("stream has already started or ended")
		}

		// Check if already registered
		for _, student := range stream.Students {
			if student.ID == currentUser.ID {
				return fmt.Errorf("already registered for this stream")
			}
		}

		// Check if max participants reached
		if len(stream.Students) >= stream.MaxParticipants+1 {
			return fmt.Errorf("stream has reached maximum participants")
		}

		// Check tokens
		if currentUser.Tokens < 2 {
			return fmt.Errorf("insufficient tokens")
		}

		// Deduct tokens
		if err := tx.Model(&currentUser).Update("tokens", gorm.Expr("tokens - ?", 2)).Error; err != nil {
			return fmt.Errorf("failed to update user tokens: %w", err)
		}

		if err := tx.Exec("INSERT INTO stream_students (stream_id, user_id) VALUES (?, ?) ON CONFLICT DO NOTHING", stream.ID, currentUser.ID).Error; err != nil {
			return fmt.Errorf("failed to join stream: %w", err)
		}

		return nil
	})

	// Handle transaction result
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "stream not found"):
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": err.Error()})
		case strings.Contains(err.Error(), "already registered"):
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, map[string]string{"error": err.Error()})
		case strings.Contains(err.Error(), "maximum participants"):
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, map[string]string{"error": err.Error()})
		case strings.Contains(err.Error(), "insufficient tokens"):
			render.Status(r, http.StatusPaymentRequired)
			render.JSON(w, r, map[string]string{"error": err.Error()})
		default:
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": err.Error()})
		}
		return
	}

	render.JSON(w, r, map[string]any{
		"message":          "Successfully joined stream",
		"tokens_remaining": currentUser.Tokens - 2,
	})
}

// GetUserJoinedStreams returns streams that the user has joined
func GetUserJoinedStreams(w http.ResponseWriter, r *http.Request) {
	user := token.User(r)
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)

	db := database.Manager()
	var streams []entities.Stream
	now := time.Now()

	// Get streams user has joined, exclude ones that ended more than 45 minutes ago
	cutoffTime := now.Add(-45 * time.Minute)

	err := db.Preload("Creator").
		Preload("Students").
		Joins("JOIN stream_students ON stream_students.stream_id = streams.id").
		Where("stream_students.user_id = ?", userID).
		Where("scheduled_at > ?", cutoffTime).
		Order("scheduled_at ASC").
		Find(&streams).Error

	if err != nil {
		http.Error(w, "Failed to fetch joined streams", http.StatusInternalServerError)
		return
	}

	// Add additional info for each stream
	for i := range streams {
		streams[i].Participants = len(streams[i].Students)

		// Add cancellation deadline info (2 hours before stream)
		cancellationDeadline := streams[i].ScheduledAt.Add(-2 * time.Hour)
		streams[i].CanCancel = now.Before(cancellationDeadline)
	}

	render.JSON(w, r, streams)
}

// CancelStreamRegistration allows user to cancel their registration
func CancelStreamRegistration(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	user := token.User(r)
	userID, _ := strconv.ParseInt(user.UserID, 10, 64)

	db := database.Manager()
	var stream entities.Stream
	result := db.Preload("Students").Where("room_id = ?", room).First(&stream)

	if result.Error != nil {
		http.Error(w, "Stream not found", http.StatusNotFound)
		return
	}

	// Check if cancellation is still allowed (2 hours before stream)
	now := time.Now()
	cancellationDeadline := stream.ScheduledAt.Add(-2 * time.Hour)
	if now.After(cancellationDeadline) {
		http.Error(w, "Cancellation deadline has passed", http.StatusBadRequest)
		return
	}

	// Check if user is actually registered
	var currentUser entities.User
	if err := db.Where("id = ?", userID).First(&currentUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	isRegistered := false
	for _, student := range stream.Students {
		if student.ID == userID {
			isRegistered = true
			break
		}
	}

	if !isRegistered {
		http.Error(w, "User is not registered for this stream", http.StatusBadRequest)
		return
	}

	// Remove user from stream participants
	if err := db.Model(&stream).Association("Students").Delete(&currentUser); err != nil {
		http.Error(w, "Failed to cancel registration", http.StatusInternalServerError)
		return
	}

	// Refund 2 tokens to user
	if err := db.Model(&currentUser).Update("tokens", gorm.Expr("tokens + ?", 2)).Error; err != nil {
		// Log error but don't fail the request since registration was already cancelled
		http.Error(w, "Registration cancelled but token refund failed", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]any{
		"message":         "Registration cancelled successfully",
		"tokens_refunded": 2,
	})
}
