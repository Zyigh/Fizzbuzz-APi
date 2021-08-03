# Fizzbuzz API

A simple API made in [Go](https://golang.org/) that allows any program to consume a [Fizzbuzz algorithm](https://en.wikipedia.org/wiki/Fizz_buzz)

## Routes

Two routes are exposed

### /

Returns a list of strings from 1 to `limit`, replacing multiples of two numbers by two string

**Allowed parameters**

* int1 [0-9] *default 3*
    
    A number which multiples will be replaced by `str1`
    
* int2 [0-9] *default 5*

    A number which multiples will be replaced by `str2`

* limit [0-9] *default 100*

    The maximum element on which to compute FizzBuzz algorithm

* str1 [*] *default Fizz*

    The string that will replace multiples `int1`

* str2 [*] *default Buzz*

    The string that will replace multiples `int2`

Returns an object with a key `data` containing an array of string. This array is the result of the computation of FizzBuzz algorithm in limit

```bash
curl https://fizzbuzz.localhost/?int2=8&limit=10
# {"data":["1","2","Fizz","4","5","Fizz","7","Buzz","Fizz","10"]}

curl https://fizzbuzz.localhost/?limit=15
# {"data":["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]}
```

### /stats/most-frequent-request

Returns the most requested query all time for the FizzBuzz algorithm

```bash
curl https://fizzbuzz.localhost/stats/most-frequent-request
# {"int1":3,"int2":5,"limit":100,"str1":"Fizz","str2":"Buzz"}
```

## Deployement

This application was made with `go1.16.6` on a darwin/amd64 architecture, and compiled and worked successfully in a 
docker `golang` container and a `scratch` container. It doesn't provide any warranty of working on other architectures, yet
it is highly probable that it is highly portable as it does't use any external librairy.

The database used for the application is `CouchDB 3.1.1`. Although the route [/](#/) works without database, it is impossible to access
the route [/stats/most-frequent-request](#/stats/most-frequent-request) as it is computed from database.

### Configuration

The server only handles HTTPS connections. In order to make it works, you will need to create a certificate file and a certificate key.
We highly recommend using [mkcert](https://mkcert.org/) for its simplicity and efficiency.

```bash
mkcert fizzbuzz.localhost
```

Don't forget to add the following line in your hosts config file 

```text
127.0.0.1       fizzbuzz.localhost
```

You can also setup a specific configuration for the app within two files in directory `conf`

#### application.go

    Settings relative to the application overview

    * **AppHost** : the host on which the app will be deployed
    * **AppHTTPSPort** : the port on which the app will be deployed
    * **AppHTTPSCertFile** : the tls cert file
    * **AppHTTPSKeyFile** : the tls cert key file

#### database.go

    Settings relative to the database connection

    * DBHost : the host to reach the database
    * DBPassword : the password to connect to the database
    * DBPort : the port to reach the database
    * DBUserName : the username to connect to the database
    * MaxAttempts : the number of attempts to reach the database before launching the app. If this number is overflowed, the app will crqsh

### Manualy

Make sure you have an instance of [CouchDb](https://couchdb.apache.org/) running somewhere and that you have correct access,
and all your local configuration constants are properly set.

After changing these configuration consts at your convenience, you can simply build a binary by running

```bash
go build -o fizzbuzz .
```

This should work as long as your go version is at least `1.16.6`. You won't need to add additionals librairies or tools as
this program doesn't use any other lib than the ones provided with go

You just have to execute the binary afterwards to start the HTTP server which will be accessible at 
[fizzbuzz.localhost](https://fizzbuzz.localhost/)

### Docker

A `Dockerfile` and a `docker-compose.yml` are provided if you'd rather build and run the app in a docker container. If you
already have an instance of CouchDb running somewhere, as long as your const are properly set, you would just have to build
and run the docker as follow.

```bash
docker build --rm -t fizzbuzzapi .
docker run --name FizzBuzzApi -d -p 443:443 fizzbuzzapi
```

If you don't have a CouchDb instance easily accessible, you can use the `docker-compose` configuration that will creates
a container that doesn't have any direct access to the outside.

```bash
docker compose up --build
```

#### Important

Note that the docker image doesn't create the tls certificates and you will have to create them before starting the docker instances

### License

This program is written under [Apache 2.0 license](./License)
