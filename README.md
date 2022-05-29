# fiber1

1. Web page
2. Form
3. API / CRUD 
4. MySQL
5. Fiber without ORM


# Dependencies

```shell
docker-compose up

air
```

# Setup Database

```shell
mysql -u root -h 127.0.0.1 -p

CREATE DATABASE test1;
```

# Updating test

```shell
cd business

go test -update .
```

# How to test manually

```shell
curl -X POST -d 'email=b@gmail.com&password=test&name=Kis' http://127.0.0.1:3000/guest/register
curl -X POST -H 'content-type: application/json' -d '{"email":"c@gmail.com","password":"test","name":"Kis"}' http://127.0.0.1:3000/guest/register
```
