# tableReservationProject


This is a simple web server for table reserving - written with GoLang. 

Restaurant assumed has 30 tables. You can see if table reserved or not, after reservation request.

There are no reserved tables at the beggining. 

For reserving this rest endpoint should be used : http://localhost:8080/
Only variable is table. 

Example body : 

```
{
  "table" : 10
}
```
