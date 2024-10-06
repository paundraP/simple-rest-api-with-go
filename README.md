# simple-rest-api-with-go

Using Gin as a framework and gorm as an orm.

# Description

This project implements a simple RESTful API for a ticketing system that allows users to register and log in for authentication. Utilizing JSON Web Tokens (JWT) for secure token creation and management. 

# Endpoints

(POST) /register : for user to registered their account

(POST) /login : to get their token to get into the main website. 

(POST) /tickets : to order a ticket 

(GET) /tickets : if the role is user, it will show the user's ticket details. if the role is admin, all tickets in database is showed. 

(GET, PATCH, DELETE) /ticket/:id : only for admin. use for search user's ticket by their id, update the status to verification their payment, and delete ticket. 

