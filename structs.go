package harvest

type Har struct {
	Log struct {
		Entries []struct {
			ResourceType string `json:"_resourceType"`
			Request      struct {
				URL string `json:"url"`
			} `json:"request"`
			Response struct {
				Content struct {
					Encoding string `json:"encoding"`
					MimeType string `json:"mimeType"`
					Size     int64  `json:"size"`
					Text     string `json:"text"`
				} `json:"content"`
			} `json:"response"`
		} `json:"entries"`
	} `json:"log"`
}
