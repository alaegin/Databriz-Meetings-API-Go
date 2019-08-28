package azure

type MembersList struct {
	Count   int      `json:"count"`
	Members []Member `json:"value"`
}

type Member struct {
	Identity struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       Links  `json:"_links"`
		ID          string `json:"id"`
		UniqueName  string `json:"uniqueName"`
		ImageURL    string `json:"imageUrl"`
		Descriptor  string `json:"descriptor"`
	} `json:"identity"`
}

type Links struct {
	Avatar Link `json:"avatar"`
}

type Link struct {
	Href string `json:"href"`
}
