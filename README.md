# go-shopping

## Running mongo

```
mkdir ~/data
sudo docker run -d -p 27017:27017 -v ~/data:/data/db mongo
docker exec -it cart-mongo bash
```

## Running postgres

```
docker run -p 5432:5432 --name usermgmt-postgres -e POSTGRES_PASSWORD=makethisanenvvariable -d postgres
psql -h localhost -U postgres
CREATE DATABASE usermgmt;
```

TODO add postgres init script
