package dbops

import "testing"

// init(loginDB -> truncate tables)
// run tests
// clear tables and truncate tables

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate comments")
}
func TestMain(m *testing.M) {
	// Test main for sub tests.
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("GetUser", testGetUser)
	t.Run("DeleteUser", testDelUser)
	t.Run("ReGetUser", testReGetUser)
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.run("Prepare", testAddUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("Joey", "food")
	if err != nil {
		t.Errorf("Error of add user %v", err)
	}
}
func testGetUser(t *testing.T) {
	password, err := GetUserCredential("Joey")
	if password != "food" || err != nil {
		t.Errorf("Error of get user")
	}

}
func testDelUser(t *testing.T) {
	err := DelUser("Joey", "food")
	if err != nil {
		t.Errorf("Error of delete user %v ", err)
	}
}
func testReGetUser(t *testing.T) {
	password, err := GetUserCredential("Joey")
	if err != nil {
		t.Errorf("Error of Re-get User %v", err)
	}
	if password != "" {
		t.Errorf("Deleteing user error")
	}
}
