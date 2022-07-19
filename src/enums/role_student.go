package enums

type RoleStudent int

const (
	OnlyViewMark         RoleStudent = 1
	ViewMarkAndTrueFalse             = 2
	PreventViewMark                  = 3
	ViewMarkAndAnswer                = 4
)
