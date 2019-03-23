package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempVid string

//init(dblogin,truncate tables)->run tests->clear data(truncate tables)

//清除tables
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

//测试流程
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestVideoInfosWorkFlow(t *testing.T) {
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddNewComments)
	t.Run("ListComments", testListComments)
}

//测试增加用户
func testAddUser(t *testing.T) {
	err := AddUserCredential("clarkrao", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

//测试获取用户
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("clarkrao")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}

	fmt.Println("pwd : ", pwd)
}

//测试删除用户
func testDeleteUser(t *testing.T) {
	err := DeleteUser("clarkrao", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

//测试重新获取用户
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("clarkrao")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func testAddNewVideo(t *testing.T) {
	info, err := AddNewVideo(1, "CLARK_VIDEO")
	if err != nil {
		t.Errorf("Error of AddNewVideo: %v", err)
	}
	//赋值给全局变量
	tempVid = info.Id

	fmt.Println("tempVid : ", tempVid)
}

func testGetVideoInfo(t *testing.T) {
	info, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
	fmt.Println("info : ", info)
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	info, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of testRegetVideoInfo: %v", err)
	}

	if info != nil {
		t.Errorf("Deleting video_info test failed")
	}
}

func testAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I Like It"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddNewComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)

	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comments: %d , %v \n", i, ele)
	}
}
