#AUTHORIZATION SERVICE
This service will authenticate the user according to its specific role and return HTTP status code 200 whent he user is authorized and a JWT token with a configurable validity. It should return 401 for unauthorized users.

      POST /user/signup

This Api should create a user with the specified role. A user with role Station or Command may only be created if the requesting user has the role Command. Creating a user with role Ship does not require authentication.

      POST /auth/login

This Api should authenticate the user and return a JWT token - the token should have the username, userid (from database) and the user role.

# DATABASE SERVICE
The system will require databases to store users, stations, command center details, etc. None of the other three microservice should communicate directly with the database, but use this service to talk to the database instead.

# CENTRAL COMMAND SERVICE
All shipping stations and spaceships must register themselves with the central command. An unregistered shipping station can not accept spaceships, and an unregistered spaceship can not land on a shipping station. The central command has services to find out the current status of shipping stations and spaceships.

      POST /centcom/station/register

This Api should be called only once by a newly created shipping station. A second call to this Api should return 400: Bad Request. A valid call should return 200 with the following response

Only users with role Ship may call this Api. Any other role should result in 400: Bad Request.

      GET /centcom/station/all

This Api should return all the registered shipping stations and their current state

Only users with role Command and Ship may access this Api. Any other role should result in 400: Bad Request.

If the requesting user's role is Ship.

      POST /centcom/ship/register

This Api should be called only once by each newly created ship. A second call by the same ship should return 400: Bad Request. A valid call should return 200 with an empty body.

Only users with role Ship may call this Api, any other role should result in 400: Bad Request.

      GET /centcom/ship/all

the role Command is required to call this Api.

# SHIPPING STATION SERVICE
Spaceships need to communicate with this Api to land on the shipping station.

      POST /shipping-station/request-landing

      POST /shipping-station/land
