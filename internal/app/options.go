package app

// Flags stores the program options.
type Flags struct {
	SubdomainFile string
	DigRepo       bool
	DigCommits    bool
	RegexFile     string
	LanguageFile  string
	ConfigFile    string
	Pages         int
	GithubRepo    bool
	ResultsOnly   bool
	NoAPIKeys     bool
	NoScoring     bool
	NoFiles       bool
	NoKeywords    bool
	OnlyFiltered  bool
	Threads       int
	Debug         bool
	LegacySearch  bool
	NoGists       bool
	NoRepos       bool
	ManyResults   bool
	JsonOutput    bool
	OTPCode       string
	TextRegexes   config
}

var flags Flags

// GetFlags is a singleton that returns the program flags.
func GetFlags() *Flags {
	return &flags
}
