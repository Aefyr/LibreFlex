# PostgreSQL setup for LibreFlex

### Basically you just need to have an **ads** table in your database with the following columns schema:
```SQL
CREATE TABLE ads
(
    id serial NOT NULL,
    banner text NOT NULL,
    target text NOT NULL,
    PRIMARY KEY (id)
)
```
**Please note that columns declaration order and names must be exactly the same as above. And you can't have any additional columns.**

**banner** column should contain ad banner image link  
**target** column should contain the link that will be opened when user click on the ad


### Once you've set up your database, use "postgres" adsRepoProvider in your config.json file:
```json
"flexApiServerConfig":{
   ...
   "adsRepoProvider":"postgres",
   "adsRepoConfig":{
      "dbHost":"127.0.0.1",
      "dbName":"flexadsdb",
      "dbUser":"flexads",
      "dbPass":"1337",
      "dbSslMode":"disable"
   }
   ...
}
```
**The same adsRepoProviderConfig works for the flexAdsServerConfig entry as well.**  

### adsRepoConfig fields explanation:
**dbHost** is your PostgreSQL host  
**dbName** is the name of the database for FlexAds  
**dbUser** is the database user (role) for LibreFlex to use when connecting to the database  
**dbPass** is **dbUser**'s password  
**dbSslMode** is one of the following:
```
* disable - No SSL
* require - Always SSL (skip verification)
* verify-ca - Always SSL (verify that the certificate presented by the
  server was signed by a trusted CA)
* verify-full - Always SSL (verify that the certification presented by
  the server was signed by a trusted CA and the server host name
  matches the one in the certificate)
```


### That's pretty much it
