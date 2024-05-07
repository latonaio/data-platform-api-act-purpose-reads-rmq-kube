package requests

type Text struct {
	ActPurpose     		string  `json:"ActPurpose"`
	Language          	string  `json:"Language"`
	ActPurposeName		string  `json:"ActPurposeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
