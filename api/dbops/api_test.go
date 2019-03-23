package dbops

import (
	"fmt"
	"testing"
)


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
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)

}

//测试增加用户
func testAddUser(t *testing.T) {
	err := AddUserCredential("clarkrao","123")
	if err != nil {
		t.Errorf("Error of AddUser: %v",err)
	}
}
//测试获取用户
func testGetUser(t *testing.T) {
	pwd ,err := GetUserCredential("clarkrao")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v",err)
	}

	fmt.Println("pwd : ",pwd)
}
//测试删除用户
func testDeleteUser(t *testing.T) {
	err := DeleteUser("clarkrao","123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v",err)
	}
}
//测试重新获取用户
func testRegetUser(t *testing.T) {
	pwd,err := GetUserCredential("clarkrao")
	if err != nil {
		t.Errorf("Error of RegetUser: %v",err)
	}

	if pwd != "" {
		 t.Errorf("Deleting user test failed")
	}
}