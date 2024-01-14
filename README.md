# DOCUMENTATION

### LIST

**Starting**

- go run main.go

**Migrating**

- sudah auto migrate

**Base URL**

- URL : http://localhost:8080/api/v1

**Create List**

- URL : http://localhost:8080/api/v1/list

- Method: POST

- Request Body by multipart/form-data

- Response Body :

```json
{
	"status": 200,
	"message": "Success to create list.",
	"data": null
}
```

**Update List**

- URL : http://localhost:8080/api/v1/list/:id

- Method: PUT

- Request Body by multipart/form-data

- Response Body :

```json
{
	"status": 200,
	"message": "Success to Update list.",
	"data": null
}
```

**Delete List**

- URL : http://localhost:8080/api/v1/list/:id

- Method: DELETE

- Response Body :

```json
{
	"status": 200,
	"message": "Success to delete list : 2",
	"data": null
}
```

---

**GET ALL List**

- URL : http://localhost:8080/api/v1/lists

- Method: GET

- Response Body :

```json
{
	"status": 200,
	"message": "Success to get lists.",
	"data": [
		{
			"id": 1,
			"title": "test1",
			"description": "test1",
			"file": "CV_MUHAMMAD REZA FADILAH.pdf",
			"sub_lists": [
				{
					"id": 1,
					"title": "sub test 1",
					"description": "sub test 1",
					"file": "",
					"list_id": 1
				},
				{
					"id": 2,
					"title": "sub test 1",
					"description": "sub test 1",
					"file": "CV_MUHAMMAD REZA FADILAH.pdf",
					"list_id": 1
				}
			]
		}
	]
}
```

**GET BY ID List**

- URL : http://localhost:8080/api/v1/list/:id

- Method: GET

- Response Body :

```json
{
	"status": 200,
	"message": "Success to get list by Id.",
	"data": {
		"id": 1,
		"title": "test1",
		"description": "test1",
		"file": "CV_MUHAMMAD REZA FADILAH.pdf",
		"sub_lists": [
			{
				"id": 1,
				"title": "sub test 1",
				"description": "sub test 1",
				"file": "",
				"list_id": 1
			},
			{
				"id": 2,
				"title": "sub test 1",
				"description": "sub test 1",
				"file": "CV_MUHAMMAD REZA FADILAH.pdf",
				"list_id": 1
			}
		]
	}
}
```

---

### SUBLIST

**Create Sublist**

- URL : http://localhost:8080/api/v1/sublist

- Method: POST

- Request Body by multipart/form-data

- Response Body :

```json
{
	"status": 200,
	"message": "Success to Create Sub List",
	"data": null
}
```

**Update Sublist**

- URL : http://localhost:8080/api/v1/sublist/:id

- Method: PUT

- Response Body :

```json
{
	"status": 200,
	"message": "Success to Update Sub List",
	"data": null
}
```

**Delete Sublist**

- URL : http://localhost:8080/api/v1/sublist/:id

- Method: DELETE

- Response Body :

```json
{
	"status": 200,
	"message": "Success to delete Sub List : 2",
	"data": null
}
```

**GET ALL Sublist**

- URL : http://localhost:8080/api/v1/sublists

- Method: GET

- Response Body :

```json
{
	"status": 200,
	"message": "Success to get Sub Lists",
	"data": [
		{
			"id": 1,
			"title": "update test 1",
			"description": "update test 1",
			"file": "CV_MUHAMMAD REZA FADILAH.pdf",
			"list_id": 1
		}
	]
}
```

**GET BY ID Sublist**

- URL : http://localhost:8080/api/v1/sublist/:id

- Method: GET

- Response Body :

```json
{
	"status": 200,
	"message": "Success to get Sub List by Id",
	"data": {
		"id": 1,
		"title": "update test 1",
		"description": "update test 1",
		"file": "CV_MUHAMMAD REZA FADILAH.pdf",
		"list_id": 1
	}
}
```
