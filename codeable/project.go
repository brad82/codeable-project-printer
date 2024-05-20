package codeable

import (
	"time"
)

type AvatarDefinition struct {
	TinyURL   string `json:"tiny_url"`
	SmallURL  string `json:"small_url"`
	MediumURL string `json:"medium_url"`
	LargeURL  string `json:"large_url"`
}

type Client struct {
	ID                  int              `json:"id"`
	FullName            string           `json:"full_name"`
	ClientType          string           `json:"client_type"`
	URL                 string           `json:"url"`
	AverageReviewRating float64          `json:"average_review_rating"`
	TimezoneOffset      int              `json:"timezone_offset"`
	TotalProjectsValue  float64          `json:"total_projects_value"`
	TotalProjects       string           `json:"total_projects"`
	Avatar              AvatarDefinition `json:"avatar"`
}

type Project struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	State          string    `json:"state"`
	URL            string    `json:"url"`
	ProjectType    string    `json:"project_type"`
	ProjectSubject string    `json:"project_subject"`
	PublishedDate  time.Time `json:"published_date"`
	PostedDate     time.Time `json:"posted_date"`
	BudgetRange    struct {
		Min string `json:"min"`
	} `json:"budget_range"`
	Estimate              string `json:"estimate"`
	Budget                string `json:"budget"`
	PublicCommentsCount   int    `json:"public_comments_count"`
	EstimatesCount        int    `json:"estimates_count"`
	EngagedExpertsCount   int    `json:"engaged_experts_count"`
	ReferralBadgeName     string `json:"referral_badge_name"`
	SubmissionFormVersion string `json:"submission_form_version"`
	Client                Client `json:"client"`
	QualityReview         string `json:"quality_review"`
}
