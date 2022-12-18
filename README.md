# ETI-Assignment-1

Design Consideration:
As the user is allowed to create both the passenger and driver account. It would be better to separate the menus into Passenger and Driver account separately
such that each of them are their own Microservice and would then be further separated into what non logged-in and logged-in users can do. These Logged-in menu
would then have offer functionalites such as being able to request trip for logged-in passengers to checking assigned trips for logged-in Drivers.

The Passenger Microservice would contain basic information such as name, email address and mobile number which the mobile number will be used as login verification,
this will then be checked by the database to ensure that such a mobile number exists within itself then be prompted to the logged-in passenger menu, whereby
it will then be able to communicate and use the Trip microservice's functions. The Driver Microservice would contain the basic information of the 
passenger Microservice but also the included identification number and car license number. Both of the logged-in accounts would then be able to 
communicate to the Trip's Microservice.

Architecture Diagram:
[ETI Architecture Diagram.docx](https://github.com/BrianLimMJ/ETI-Assignment-1/files/10253858/ETI.Architecture.Diagram.docx)

Instructions:
use go run main.go and use the numbers to navigate through the menus
