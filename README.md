# Create User's Account

Create an account for the non-authenticated user if an account for that user does not already exist. Each user can only have one account.

**URL** : `/api/register`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Provide the data of user's account to be created.

```json
{
  "email": "[required,email format]",
  "username": "[required,min=3,alphanumeric]",
  "first_name": "[required,min=3,alphabet]",
  "last_name": "[alphabet]",
  "phone": "[required,min=6,max=15,numeric]",
  "company": "[required,min=3]",
  "business_relation": "[required,min=3]",
  "password": "[required,min=8]",
  "password_confirm": "[required,min=8]"
}
```

**Data example**

```json
{
  "email": "email@gmail.com",
  "username": "ahmadakbar",
  "first_name": "Ahmad",
  "last_name": "Akbar",
  "phone": "081234567890",
  "company": "PT Bank Negara Indonesia Tbk",
  "business_relation": "EDM Team",
  "password": "password",
  "password_confirm": "password"
}
```

## Success Response

**Condition** : If everything is OK and an account didn't exist for the user.

**Code** : `200 OK`

**Content example**

```json
{
  "status": 200,
  "message": "568c2f7e-531f-4302-a3d5-f92e52ef8bf2"
}
```

## Error Responses

**Condition** : If an account already exists for the user.

**Code** : `500 INTERNAL SERVER ERROR`

**Content example**

```json
"409 Conflict: User exists with same username"
```

### and

**Condition** : If fields are missed or invalid format.

**Code** : `400 BAD REQUEST`

**Content example**

```json
"code=400, message=payload is unknown"
```
