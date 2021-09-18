# Notice 
Sorry I forgot to add the Request Method Verifier in the API Service.

    func (s *Server) VerifyMethod(next http.HandlerFunc, method string) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            if r.Method == method {
                next(w, r)
            } else {
                PrepareResponse(w, http.StatusBadRequest, "")
                return
            }
        }
    }

Now It is added.

# Deployment
To deploy the Solution all we have to do is use the docker-compose.yml file existing in the project:

    $ docker-compose up

All the configuration data required for the application is in the docker compose file.

# Important Note
You need to provide the correct JSON BODY when performing requests like Shown in the Testing.

# Testing
The Default login included is:
    
    localhost:8080/auth/login

    {
        "username": "ayoub",
        "password" : "ayoub1111"
    }

Adding new User:

    localhost:8080/user/signup

    {
        "username": "ayoub3",
        "password" : "ayoub1111",
        "role" : "Ship"
    }

Station Register:

    localhost:8080/centcom/station/register

    {
        "capacity" : 123.3,
        "docks" : [{
                        "numDockingPorts" :  12
                    },{
                        "numDockingPorts" : 14
                    },{
                        "numDockingPorts" : 66
                    }]
    }

All Stations 'Command' Role: 

    localhost:8080/centcom/station/all

All Stations for a Specific Ship 'Ship' Role:

    localhost:8080/centcom/station/all

    {
        "shipId": "1"
    }

Register New Ship: 

    localhost:8080/centcom/ship/register

    {
        "weight" : 15.9
    }

All Ships:

    localhost:8080/centcom/ship/all

Request Landing in a Space Station by a Ship:

    localhost:8080/shipping-station/request-landing

    {
        "time" : 30,
        "idShip" : "1",
        "idStation" : "1"
    }

Landing in a Space Station by a Ship:

    localhost:8080/shipping-station/land

    {
        "time" : 30,
        "idShip" : "1",
        "idStation" : "1"
    }


# Clarifications
I used a secure (just for show, can be removed easily) communication between different services that used gRPC.
The Certificates are present in the project. 

Take into consideration to not modify the containers names or we will have to rebuild new Certificates for secured communication.

    