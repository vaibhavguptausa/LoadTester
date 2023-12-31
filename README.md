A simple load tester written in go

Currently supports uniform distribution of the load. 

set up the project and run it in local, this will be hosted on 8090
you need to define the total number of requests to make within a specified time. 

**sample cURL** -

curl --location 'http://localhost:8090/load' \
--header 'Content-Type: application/json' \
--data '{
    "req_total": 1000,
    "strategy": "UNIFORM",
    "time": "10s",
    "url": <your URL>,
    "method": "GET",
    "body": null
}'

Currently just GET is supported, however POST should be fairly simple to extend.
