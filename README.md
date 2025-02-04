# Reciept Processor

A simple webservice that has two RESTApi endpoints for submitting a reciept and for looking up the points earned for the reciept.

I wrote this in Golang first and formost, because I love writng in Go. Also, the challenger stated they use Golang quite a bit themsleves. I figured this would be the best language to use. At least we speak the same language. 

The server listens on port :8080. That's static, I didn't feel it was worth it to make it more dynamic for this exersize.

The first endpoint being `/reciepts/process` takes a "reciept" as json, saves, processes and returns an id like below. 

`{"id": adb6b560-0eef-42bc-9d16-df48f30e89b2}`


The second endpoint `/receipts/{id}/points` enables the user to lookup points awarded for the reciept with the Id returned from the first endopoint. The expected return value with look like below.

`{ "points": 50 }`


# Prequisites
* Lastest version of Golang (1.22.3)
* Latest version of Docker ( Optional )

# Installation

### Native installation
Make sure you have the lastest version of Golang installed. At the time of this writing I am using version 1.22.3

Clone this repository 

`git clone https://github.com/arabenjamin/fetch.git`

Move into repo directory

`cd fetch/`

Install Golang depandancies

`go mod download`

Run it 

`go run` 


The server should start running on port :8080

You should see the below as output in the termnial

```
Running Fetch App
http: 2025/02/03 23:45:14 Starting receipt processor service ...
```


### Run with Docker
To run with Docker run this cmd from the repo directory

`docker build -t fetch-app .`

Next run

`docker run -p 8080:8080 fetch-app` 

The server should start displaying the same output as if we had ran it natively.


# Usage 
Once the server is running, one should be able to send requests to the server. Documentation is in the openApi.yaml file included in the repo.

* `/ping GET`
* `/reciepts/process POST`
* `/receipts/{id}/points`


## Examples
For example, if you were to reachout to :

`http://localhost:8080/receipts/process`

Perhaps from the cmd or your favorite RESTFULApi client
```
curl -X 'POST' \
  'http://localhost:8080/receipts/process' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
  ],
  "total": "6.49"
}'
```
You would expect to recieve this as a json response
```
{
	"id": "af993bf0-11de-4f59-a6fc-3692ebdb8649"
}
```
Using that ID returned, you can expect to use it at the endpoint
`http://localhost:8080/receipts/af993bf0-11de-4f59-a6fc-3692ebdb8649/points`

```
curl -X 'GET' \
  'http://localhost:8080/receipts/af993bf0-11de-4f59-a6fc-3692ebdb8649/points' \
  -H 'accept: application/json'
```
You should expect to get back something that looks this.
```
{
	"points": 23
}
```

# Problems
I'm very certain you could crash this server. Though I dont think that was the point of the exersize. 

Also I'm only kinda confident that it calculates the points correctly. It may for somethings and not for others. But for now I'm calling it good enogh.

