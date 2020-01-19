package mailchimp

import (
	"net/url"
	"strconv"
	"time"
)

// ListsResponse data
type ListsResponse struct {
	Lists       []List          `json:"lists"`
	TotalItems  int             `json:"total_items"`
	Constraints ListConstraints `json:"constraints"`
	Links       []Link          `json:"_links"`
}

// List data
type List struct {
	ID                   string            `json:"id"`
	WebID                int               `json:"web_id"`
	Name                 string            `json:"name"`
	Contact              *ListContact      `json:"contact"`
	PermissionReminder   string            `json:"permission_reminder"`
	UseArchiveBar        bool              `json:"use_archive_bar"`
	CampaignDefaults     *CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe    string            `json:"notify_on_subscribe"`
	NotifyOnUnsubscribe  string            `json:"notify_on_unsubscribe"`
	DateCreated          *time.Time        `json:"date_created"`
	ListRating           int               `json:"list_rating"`
	EmailTypeOption      bool              `json:"email_type_option"`
	SubscribeURLShort    string            `json:"subscribe_url_short"`
	SubscribeURLLong     string            `json:"subscribe_url_long"`
	BeamerAddress        string            `json:"beamer_address"`
	Visibility           string            `json:"visibility"`
	DoubleOptin          bool              `json:"double_optin"`
	HasWelcome           bool              `json:"has_welcome"`
	MarketingPermissions bool              `json:"marketing_permissions"`
	Modules              []interface{}     `json:"modules"`
	Stats                *ListStats        `json:"stats"`
	Links                []Link            `json:"_links"`
}

// ListContact data
type ListContact struct {
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
}

// CampaignDefaults data
type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

// ListStats data
type ListStats struct {
	MemberCount               int        `json:"member_count"`
	UnsubscribeCount          int        `json:"unsubscribe_count"`
	CleanedCount              int        `json:"cleaned_count"`
	MemberCountSinceSend      int        `json:"member_count_since_send"`
	UnsubscribeCountSinceSend int        `json:"unsubscribe_count_since_send"`
	CleanedCountSinceSend     int        `json:"cleaned_count_since_send"`
	CampaignCount             int        `json:"campaign_count"`
	CampaignLastSent          string     `json:"campaign_last_sent"`
	MergeFieldCount           int        `json:"merge_field_count"`
	AvgSubRate                int        `json:"avg_sub_rate"`
	AvgUnsubRate              int        `json:"avg_unsub_rate"`
	TargetSubRate             int        `json:"target_sub_rate"`
	OpenRate                  int        `json:"open_rate"`
	ClickRate                 int        `json:"click_rate"`
	LastSubDate               *time.Time `json:"last_sub_date"`
	LastUnsubDate             string     `json:"last_unsub_date"`
}

// Link data
type Link struct {
	Rel          string `json:"rel"`
	Href         string `json:"href"`
	Method       string `json:"method"`
	TargetSchema string `json:"targetSchema,omitempty"`
	Schema       string `json:"schema,omitempty"`
}

// ListConstraints data
type ListConstraints struct {
	MayCreate             bool `json:"may_create"`
	MaxInstances          int  `json:"max_instances"`
	CurrentTotalInstances int  `json:"current_total_instances"`
}

// Lists API client
type Lists struct {
	Client
}

// All retrieves all Lists
func (c Lists) All(listParams url.Values) (ListsResponse, error) {
	r := ListsResponse{}
	err := c.Client.Request("GET", "/lists", listParams, nil, &r)
	return r, err
}

// Get list by ID
func (c Lists) Get(id int, listParams url.Values) (List, error) {
	var r List
	err := c.Client.Request("GET", "/lists/"+strconv.Itoa(id), listParams, nil, &r)
	return r, err
}
