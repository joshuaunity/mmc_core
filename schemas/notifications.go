package schemas

type NotifyType string

const (
	EMAIL      NotifyType = "EMAIL"
	WEBPUSH    NotifyType = "WEBPUSH"
	MOBILEPUSH NotifyType = "MOBILEPUSH"
)

type MailTmpType string

const (
	likedPost       MailTmpType = "likedPost"
	PostComment     MailTmpType = "postComment"
	NewPost         MailTmpType = "newPost"
	NewFollower     MailTmpType = "newFollower"
	RankingReminder MailTmpType = "createRanking"
)

type PushNotificationType string

const (
	LIKE           PushNotificationType = "like"
	COMMENT        PushNotificationType = "comment"
	FOLLOW         PushNotificationType = "follow"
	RECOMMENDATION PushNotificationType = "recommendation"
)

type WebNotificationData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}

type RecordData struct {
	Receiver string
	Sender   string
	Type     string
	Content  string
	Action   string
}

type NotificationEvent struct {
	NotifyType          NotifyType           `json:"notify_type"`
	RecordData          *RecordData          `json:"record_data,omitempty"`
	EmailData           *EmailData           `json:"email_data,omitempty"`
	WebNotificationData *WebNotificationData `json:"web_notification_data,omitempty"`
}

type EmailData struct {
	To           string
	Subject      string
	TemplateFile string
	Data         interface{}
}
