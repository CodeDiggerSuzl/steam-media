package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var testVideoID string

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

func testAddUser(t *testing.T) {
	err := AddUserCredential("Joey", "123")
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

// unit tests for video

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	testVideoID = vi.ID
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(testVideoID)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DelVideoInfo(testVideoID)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(testVideoID)
	if err != nil || vi != nil {
		t.Errorf("Error of ReGetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddCommnets", testAddComment)
	t.Run("ListComments", testListComments)
}

func testAddComment(t *testing.T) {
	vID := "12345"
	authorID := 1
	content := "love it"

	err := AddNewComments(vID, authorID, content)
	if err != nil {
		t.Errorf("Error of adding comment: %v", err)
	}
}

func testListComments(t *testing.T) {
	vID := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vID, from, to)
	if err != nil {
		t.Errorf("Error of listing comments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("Comments: %d, %v \n", i, ele)
	}
}
