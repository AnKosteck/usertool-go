package main

type User struct {
	uid                                                     uint
	gid                                                     int
	name                                                    string
	homePermissions                                         string
	homePath, shellPath, gecko, fullName, informtion, email string
	passwordHash                                            string
	slurmAccounts                                           string
	uid                                                     uint
	//otherGroups Group[] //group slice
}

func (User u) createUser() {

}
