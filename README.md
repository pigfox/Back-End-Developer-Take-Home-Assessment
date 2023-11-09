# Back-End-Developer-Take-Home-Assessment
1. Clone this https://github.com/pigfox/Back-End-Developer-Take-Home-Assessment or just extract the zip file.
2. $ cd Back-End-Developer-Take-Home-Assessment
3. Copy .env.example to .env and enter a key
4. You might have to run "$ go mod tidy"
5. Start the server "$ go run *.go"

6. Register
Method:POST
Headers:
Content-Type:application/json
URL:http://localhost:9898/register
BODY:
{
"email": "xyz@gmail.com",
"password": "12348567yyy"
}

7. Login
Method:POST
Headers:
Content-Type:application/json
URL:http://localhost:9898/login
BODY:
{
"email": "xyz@gmail.com",
"password": "12348567yyy"
}

8. Create task(s)
Method:POST
Headers:
Content-Type:application/json
Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBnbWFpbC5jb20iLCJleHAiOjE2OTk1NTI1MjR9.WCd9OpNOsLgIfkoRImE9gLfqVJLn0vDCesaTXLrARnk
URL:http://localhost:9898/task
BODY:
{
"title": "Title",
"description": "description",
"due_date": "2023-11-23",
"status":"Pending" 
}

9. View tasks
Method:GET
Headers:
Content-Type:application/json
Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBnbWFpbC5jb20iLCJleHAiOjE2OTk1NTI1MjR9.WCd9OpNOsLgIfkoRImE9gLfqVJLn0vDCesaTXLrARnk
URL:http://localhost:9898/tasks/1

10. View task
Method:GET
Headers:
Content-Type:application/json
Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBnbWFpbC5jb20iLCJleHAiOjE2OTk1NTI1MjR9.WCd9OpNOsLgIfkoRImE9gLfqVJLn0vDCesaTXLrARnk
URL:http://localhost:9898/task/1

11. Update task
Method:PUT
Headers:
Content-Type:application/json
Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBnbWFpbC5jb20iLCJleHAiOjE2OTk1NTI1MjR9.WCd9OpNOsLgIfkoRImE9gLfqVJLn0vDCesaTXLrARnk
URL:http://localhost:9898/task/1
BODY:
{
"title": "Title2",
"description": "description2",
"due_date": "2023-11-23",
"status":"Pending" 
}

12. Delete task
Method:DELETE
Headers:
Content-Type:application/json
Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBnbWFpbC5jb20iLCJleHAiOjE2OTk1NTI1MjR9.WCd9OpNOsLgIfkoRImE9gLfqVJLn0vDCesaTXLrARnk
URL:http://localhost:9898/task/1