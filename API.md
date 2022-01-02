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
Delete user info by id\
If the user is deleted, the server will return 200 status

### `GET /user/workers/` `*Only for admin users`
Get all non-admin users