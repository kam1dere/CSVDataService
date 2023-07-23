package dataFormCSV

type Item struct {
	Id                        int64
	Uid                       string
	Domain                    string
	Cn                        string
	Department                string
	Title                     string
	Who                       string
	LogonCount                int64
	NumLogons7                int64
	NumShare7                 int64
	NumFile7                  int64
	NumAd7                    int64
	NumN7                     int64
	NumLogons14               int64
	NumShare14                int64
	NumFile14                 int64
	NumAd14                   int64
	NumN14                    int64
	NumLogons30               int64
	NumShare30                int64
	NumFile30                 int64
	NumAd30                   int64
	NumN30                    int64
	NumLogons150              int64
	NumShare150               int64
	NumFile150                int64
	NumAd150                  int64
	NumN150                   int64
	NumLogons365              int64
	NumShare365               int64
	NumFile365                int64
	NumAd365                  int64
	NumN365                   int64
	HasUserPrincipalName      bool
	HasMail                   bool
	HasPhone                  bool
	FlagDisabled              bool
	FlagLockout               bool
	FlagPasswordNotRequired   bool
	FlagPasswordCantChange    bool
	FlagDontExpirePassword    bool
	OwnedFiles                int64
	NumMailboxes              int64
	NumMemberOfGroups         int64
	NumMemberOfIndirectGroups int64
	MemberOfIndirectGroupsIds []string
	MemberOfGroupsIds         []string
	IsAdmin                   bool
	IsService                 bool
}