# Retail
Prerequisites:
  1. golang env
  2. POSTMAN
  3. mySQL Workbench 

How to run?
  1. go build (it willl create one executable file)
  2. .\shopalyst_application.exe
  
POSTMAN guide:
  1. In file function initializeRoutes helps to route request. 
     you can get request string from 
     for e.g shop.Router.HandleFunc("/Categorys", shop.getCategorys).Methods("GET")
     the request in postman would be http://localhost:8080/Categorys type of request would be GET (common prefix : http://localhost:8080)
