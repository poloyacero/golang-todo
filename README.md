Protected routes - Bearer Token
http://localhost:8080/tasks - POST
http://localhost:8080/tasks - GET
http://localhost:8080/tasks/{id} - GET
http://localhost:8080/tasks/{id} - DELETE
http://localhost:8080/tasks/{id} - PUT


Unprotected routes
http://localhost:8080/auth/signin - POST
{
    email,
    password
}
returns ACCESS TOKEN

http://localhost:8080/auth/signup - POST
{
    email,
    password
}