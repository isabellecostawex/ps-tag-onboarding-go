## Recreating Java Example in Golang: 

### Goal:

Implement an API service for managing user data, including creation, getting data and validation of user records in Golang.

### Technical Description

#### POST /save

- Requirements:  

This endpoint stores the user data into the database, or updates when the "id" is given. 
-Accepts JSON payloads for user data in these formats:

json  
{   
    "first_name": "John",
    "last_name": "Payne",
    "email": "johnpayne@gmail.com",
    "age": 22
}  

or 

json 
''{   
    "id": "user-20"
    "first_name": "John",
    "last_name": "Payne",
    "email": "johnpayne@gmail.com",
    "age": 22
} ''

Fail request if it's missing any required fields;
Ensure that email addresses are properly formatted;
Validates that the user's age meets a minimum requirement;
Fail request if a user with the same first and last name already exists.

#### GET /find/{id}

Returns user data by ID;

- Requirements:  

http://localhost:8080/find/1234

Returns user data by ID;
Returns a error if the user is not found.

### Acceptance Criteria

AC#1 - Successfully Save a User:  
GIVEN I am a user
AND I am recording my register
WHEN I make a POST request to /save
AND all the user fields in the payload are valid
THEN the user record is registered in the database 
AND a 200 Ok response is returned

AC#1 - Successfully Update a User:  
GIVEN I am a user
AND I am updating my register
WHEN I make a POST request to /save
AND all the user fields in the payload are valid, and I insert the ID that I want to update
THEN the user record is updated in the database 
AND a 200 Ok response is returned

AC#2 - Fail to Save User with missing data:  
GIVEN I am a User  
AND I am recording my register
WHEN I make a POST request to /save missing required fields
THEN I receive a 400 bad request response code  
AND an error message

AC#3 - Fail to Save User with invalid Data:  
GIVEN I am a User  
AND I am recording my register
WHEN I make a POST request to /save with invalid user data
THEN I receive a 400 bad request response code  
AND an error message

AC#4 - Successfully Get a User:  
GIVEN I am a user  
AND I am finding a user by ID
WHEN I make a GET request to /find/{id}
AND the ID value is valid
THEN the corresponding user data is returned  
AND a 200 Ok response is returned

AC#5 - Fail to Get a User:  
GIVEN I am a user  
AND I am looking for a User by its ID
WHEN I make a GET request to /find/{id}
AND the ID value is invalid  or doesn't exist in the database
THEN I receive a 404 Not Found response code  
