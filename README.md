# go_login

func: CreateUser
	POST(http://localhost:8080/user-api/creatuser)
	Body -> JSON
	 {
	 	"ID": 3,
	 	"Name": "test2",
	 	"Age": 19
	 }
	 
	-----------------------------------------------
func: GetUsers
	GET(http://localhost:8080/user-api/getuserall)
	paramiter:

	-----------------------------------------------
func: GetUserByID
	GET(http://localhost:8080/user-api/getuserbyid?id=1)
	paramiter:
	id