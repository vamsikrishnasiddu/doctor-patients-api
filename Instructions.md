1. First install mysql in linux and configure the username and password

2. After setting the username and password please verify your status
 by using systemctl status mysql

 3. Iam using mysql workbench 

  Create a database doctors_db and table doctors with these fields id, doctor_name,date_created

  4.To run the application 
    go run main.go

  the server is running on localhost:8080

  you can perform various crud operations.

  POST:localhost:8080/doctors

   {
       name : "example doctor"
   }

  GET: locahost:8080/doctors/id

  PATCH/PUT: localhost:8080/doctors/id

  DELETE: localhost:8080/doctors/id

   