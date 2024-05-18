﻿# GO_REST
This is the GO_REST API which I have create by following a tutorial.
The main purpose is to create a REST API with GO using GIN framework.
The API support events, where user can signup, login and create or modify (update/delete) events.
It provides feature for user to register for particular event.




if you run the server in local, the urls to access these endpoints will be

http://localhost:8080/events  (get all events)
http://localhost:8080/events/<id> (get event for specific id)
http://localhost:8080/events   (to postevent, you need to be logged in and pass jwt token in request which you get on login )
http://localhost:8080/events/<id>    (to edit/delete event, you need to be logged in and pass jwt token in request which you get on login )
http://localhost:8080/signup/  (to sign up, need to pass email and password)	
http://localhost:8080/login/  (to log in, need to pass email and password)
http://localhost:8080/events/<id>/register  (to register or cancel registration)
