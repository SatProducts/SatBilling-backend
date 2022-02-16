# PodBilling backend / API usage


## Authentication service

### `POST /auth/login/`

Generate JWT token by login and password

Example request content:
```json
{
	"login": "<username>",
	"password": "<password>"
}
```

Example response content:
```json
{
	"token": "<your-jwt-token>"
} 
```


## Users service

For using users service you need to get JWT token
and insert it to the request header like this:
```json
{
	"Authorization": "Bearer <your-jwt-token>"
}
```

### `GET /user/me/`

Getting a user record by himself\
Returns user record

### `GET /user/<id>/` `*Only for admin users`
Get user by id\
Returns user record

### `POST /user/` `*Only for admin users`
Creates user 

Example request content:
```json
{
	"login": "<user-login>",
	"password": "<user-password>",
	"permissions": "<user-permissions>"
}
```

If the user is created, the server will return 201 status

### `PUT /user/<id>/` `*Only for admin users`
Update user info by id

Example request content:
```json
{
    "vacation": true
}
```
If the user is updated, the server will return 200 status

### `DELETE /user/<id>/` `*Only for admin users`
Delete user by id\
If the user is deleted, the server will return 200 status

### `GET /user/workers/` `*Only for admin users`
Get all non-admin users


## Tasks service

For using users service you need to get JWT token
and insert it to the request header like this:
```json
{
	"Authorization": "Bearer <your-jwt-token>"
}
```

### `GET /task/<id>/` `*Only for specialist, admin and task creator`
Get task by id\
Returns task record

### `POST /task/<id>/` `*Only for operators`
Creates task

Example request content:
```json
	"title": "Fix router",
	"text": "Please fix fucking router",
	"address": "Lenina street, 1917"
```

If the task is created, the server will return 201 status

### `PUT /task/<id>/` `*Only for task creator`
Update user info by id

Example request content:
```json
{
    	"text": "Please fix fucking router, but this is new text"
}
```

If the user is updated, the server will return 200 status

### `DELETE /task/<id>/` `*Only for task creator and admin users`
Delete task by id\
If the task is deleted, the server will return 200 status

### `GET /task/for/` `*Only for specialists`
Get all tasks for you
Returns list of tasks records

### `GET /task/from/` `*Only for operators`
Get all the tasks created by you
Returns list of tasks records