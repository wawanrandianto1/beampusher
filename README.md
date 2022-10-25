# beampusher
test beam pusher


# frontend
- open terminal
- go to nodejs folder
- run command "npm run local"
- to stop ctrl + c / close terminal

# curl terminal test
curl -H "Content-Type: application/json" \
     -H "Authorization: Bearer 02F9FF929D029107235E599916A6EAF3FE33D55BA5EFDECC27A8A292B25458DA" \
     -X POST "https://dc1dfd43-1815-4b60-9ea0-b59937f51dd7.pushnotifications.pusher.com/publish_api/v1/instances/dc1dfd43-1815-4b60-9ea0-b59937f51dd7/publishes" \
     -d '{"interests":["hello"],"web":{"notification":{"title":"Hello3","body":"Hello1, world2!"}}}'

- nama interest harus "hello" karena sama dengan yang di subscribe di frontend
- title dan body bebas bisa diganti

# backend
