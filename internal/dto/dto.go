package dto

type Error struct {
	Reason string `json:"reason"`
}

type TenderRequest struct {
	Name            string `json:"name" validate:"gt=0,lte=100"`
	Description     string `json:"description" validate:"gt=0,lte=500"`
	ServiceType     string `json:"serviceType" validate:"oneof=Construction Delivery Manufacture"`
	OrganizationId  string `json:"organizationId" validate:"uuid4"`
	CreatorUsername string `json:"creatorUsername" validate:"required"`
}

type TenderResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ServiceType string `json:"serviceType"`
	Version     int32  `json:"version"`
	CreatedAt   string `json:"createdAt"`
}

type BidRequest struct {
	Name        string `json:"name" validate:"gt=0,lte=100"`
	Description string `json:"description" validate:"gt=0,lte=500"`
	TendterID   string `json:"tenderId" validate:"uuid4"`
	AuthorType  string `json:"authorType" validate:"oneof=Organization User"`
	AuthorID    string `json:"authorId" validate:"uuid4"`
}

type BidResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	AuthorType string `json:"authorType"`
	AuthorID   string `json:"authorId"`
	Version    int32  `json:"version"`
	CreatedAt  string `json:"createdAt"`
}

type ReviewResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}
