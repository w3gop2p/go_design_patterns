package cmd

type UserIterator struct {
	index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.index < len(u.Users) {
		return true
	}
	return false

}
func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.index]
		u.index++
		return user
	}
	return nil
}
