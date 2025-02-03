# Reciept Processor

A simple webservice that has two RESTApi endpoints for submitting a reciept and for looking up the points earned for the reciept. 

The first endpoint being `/reciepts/process` takes a "reciept" as json, saves, processes and returns an Id like below. 

`{"id": adb6b560-0eef-42bc-9d16-df48f30e89b2}`

The second endpoint `/receipts/{id}/points` enables the user to lookup points awarded for the reciept with the Id returned from the first endopoint. The expected return value with look like below.

`{ "points": 50 }`


# Prequisites
* Lastest version of Golang (1.22.3)
* Latest version of Docker ( Optional )

# Installation

### Native installation
Make sure you have Golang installed.

Clone this repository 

`git clone https://github.com/arabenjamin/fetch.git`

Move into repo directory

`cd fetch/`

Install Golang depandancies
`go mod download`

Run it 

`go run` 


The server should start running on port :8080


### Run with Docker



# Usage 
Once the server is running one should be able to send requests to the server. Documentation is in the openApi.yaml file included in the repo

## Examples
