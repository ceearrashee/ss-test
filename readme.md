# Simple API server for user accounting
This mini API server implements a small test program for user accounting.
This server, as a ready-made solution, is published in the cloud at http://blow.pp.ua and will be available until 2023-11-30.

This test server implements the following API functions:
* Function to get authentication token:
```http request
GET http://blow.pp.ua/api/v1/token/generate
```
* Function for getting a list of token users:
```http request
GET http://blow.pp.ua/api/v1/users
Authorization: Bearer {{insert token here}}
```
* New user creation function (requires authentication token):
```http request
POST /api/v1/user
host: http://blow.pp.ua/
Authorization: Bearer Authorization: Bearer {{insert token here}}
Content-Type: application/json

{ "name": "Eugene" }
```
* Function to change an already existing user (requires authentication token):
```http request
PUT /api/v1/user/1
host: http://blow.pp.ua/
Authorization: Bearer Authorization: Bearer {{insert token here}}
Content-Type: application/json

{
"id": 1,
"name": "ceearrashee"
}
```
* Function to get an existing user by ID (requires an authentication token):
```http request
GET /api/v1/user/2
host: http://blow.pp.ua/
Authorization: Bearer Authorization: Bearer {{insert token here}}
```

* Function to delete an existing user by ID (requires an authentication token):
```http request
DELETE /api/v1/user/1
host: http://blow.pp.ua/
Authorization: Bearer Authorization: Bearer {{insert token here}}
```

Request/response json example with full field list:
```json
{
  "id": 1,
  "createdAt": "2021-09-30T20:00:00Z",
  "updatedAt": "2021-09-30T20:00:00Z",
  "name": "Eugene",
  "surname": "Androsov",
  "phone": "+380999999999",
  "address": "Ukraine, Kyiv, 1st street, 1"
}
```
