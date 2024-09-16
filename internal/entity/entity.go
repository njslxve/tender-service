package entity

import "time"

type Tender struct {
	ID              string
	Name            string
	Description     string
	ServiceType     string
	Status          string
	CreatorUsername string
	OrganizationID  string
	Version         int32
	CreatedAt       time.Time
}

type Bid struct {
	ID          string
	Name        string
	Description string
	TenderID    string
	Status      string
	AuthorType  string
	AuthorID    string
	Version     int32
	CreatedAt   time.Time
}

type BidFeedback struct {
	ID          string
	BidID       string
	Description string
	AuthorID    string
	CreatedAt   time.Time
}

type BidDecision struct {
	ID            string
	TenderID      string
	BidID         string
	Decision      string
	ApprovedCount int32
	RejectedCount int32
}
