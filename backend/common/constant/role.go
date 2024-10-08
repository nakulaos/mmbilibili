package constant

var (
	NormalRole int = 1
	AdminRole  int = 2
	RootRole   int = 4
)

func HasRole(userRole int, role int) bool {
	ret := (userRole&role == role)
	return ret
}
