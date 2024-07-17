package requests

type NewCampaign struct {
	Name    string
	Content string
	Emails  []string
}
