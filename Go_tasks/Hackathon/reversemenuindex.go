package piscine

func ReverseMenuIndex(menu []string) []string {
	revmenu := make([]string, len(menu))

	for i, arg := range menu {
		revmenu[len(menu)-1-i] = arg
	}
	return revmenu
}
