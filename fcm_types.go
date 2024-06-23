package fcm

type Notification struct {
	Title                string `json:"title,omitempty"`
	Body                 string `json:"body,omitempty"`
	Image                string `json:"image,omitempty"`
	Icon                 string `json:"icon,omitempty"`
	Sound                string `json:"sound,omitempty"`
	Tag                  string `json:"tag,omitempty"`
	Color                string `json:"color,omitempty"`
	ClickAction          string `json:"click_action,omitempty"`
	BodyLocKey           string `json:"body_loc_key,omitempty"`
	BodyLocArgs          string `json:"body_loc_args,omitempty"`
	TitleLocKey          string `json:"title_loc_key,omitempty"`
	NotificationPriority string `json:"notification_priority,omitempty"`
}

type APNAlert struct {
	Title        string `json:"title,omitempty"`
	Subtitle     string `json:"subtitle,omitempty"`
	Body         string `json:"body,omitempty"`
	TitleLocKey  string `json:"title-loc-key,omitempty"`
	TitleLocArgs string `json:"title-loc-args,omitempty"`
	ActionLocKey string `json:"action-loc-key,omitempty"`
	LocKey       string `json:"loc-key,omitempty"`
	LocArgs      string `json:"loc-args,omitempty"`
	LaunchImage  string `json:"launch-image,omitempty"`
}

type APNS struct {
	Alert            APNAlert `json:"alert,omitempty"`
	Badge            int      `json:"badge,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	ContentAvailable int      `json:"content_available,omitempty"`
	ThreadID         string   `json:"thread_id,omitempty"`
	Category         string   `json:"category,omitempty"`
}

type AndroidConfig struct {
	CollapseKey           string            `json:"collapse_key,omitempty"`
	Priority              string            `json:"priority,omitempty"`
	Ttl                   string            `json:"ttl,omitempty"`
	RestrictedPackageName string            `json:"restricted_package_name,omitempty"`
	Data                  map[string]string `json:"data,omitempty"`
	Notification          Notification      `json:"notification,omitempty"`
	FcmOptions            map[string]string `json:"fcm_options,omitempty"`
	DirectBootOk          bool              `json:"direct_boot_ok,omitempty"`
}

type WebpushConfig struct {
	Headers      map[string]string `json:"headers,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Notification Notification      `json:"notification,omitempty"`
	FcmOptions   map[string]string `json:"fcm_options,omitempty"`
}

type APNSFcmOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
	Image          string `json:"image,omitempty"`
}

type APNSPayload struct {
	Aps APNS `json:"aps,omitempty"`
}

type APNSConfig struct {
	Headers map[string]string `json:"headers,omitempty"`
	Payload APNSPayload       `json:"payload,omitempty"`
}

type Message struct {
	Token        string            `json:"token,omitempty"`
	Tokens       []string          `json:"tokens,omitempty"`
	Topic        string            `json:"topic,omitempty"`
	Notification Notification      `json:"notification,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Android      AndroidConfig     `json:"android,omitempty"`
	Webpush      WebpushConfig     `json:"webpush,omitempty"`
	APNS         APNSConfig        `json:"apns,omitempty"`
	Condition    string            `json:"condition,omitempty"`
}

type MessagePayload struct {
	Message Message `json:"message,omitempty"`
}
