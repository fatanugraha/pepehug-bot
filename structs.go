package main

type Profile struct {
	DisplayName string `json:"display_name"`
	Image512    string `json:"image_512"`
}

type ProfileAPIResult struct {
	Ok      bool    `json:"ok"`
	Profile Profile `json:"profile"`
}

type Attachment struct {
	Title    string `json:"title"`
	Pretext  string `json:"pretext"`
	ImageURL string `json:"image_url"`
	ThumbURL string `json:"thumb_url"`
	Footer   string `json:"footer"`
}

type ResponsePayload struct {
	Attachments  *[]Attachment `json:"attachments"`
	Text         string        `json:"text"`
	ResponseType string        `json:"response_type"`
}
