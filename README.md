# Fiber1 Demo

This is example of [how to structure your golang project](https://kokizzu.blogspot.com/2022/05/how-to-structure-layer-your-golang-project.html) article with fiber and standard mysql (you can change it to whatever framework and persistence libraries you like, the structure should still be similar).

This is example how to do these things:

1. Web page
2. Form
3. API / CRUD 
4. MySQL
5. Fiber without ORM
6. DockerTest
7. AutoGold

without clean architecture, but the business logic is pure from serialization.
the model itself depends on real persistence and tested using dockertest, so it should always works.

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
