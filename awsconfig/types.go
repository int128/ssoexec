package awsconfig

type Config struct {
	Profiles []*Profile
}

type Profile struct {
	Name string

	Region       string
	SSOStartURL  string
	SSORegion    string
	SSOAccountID string
	SSORoleName  string
}
