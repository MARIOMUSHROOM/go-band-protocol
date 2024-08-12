How to run the project
go run cmd/main.go

Problem 1: Boss Baby's Revenge
[GET] {endpoint}/public/boss-baby/SSRSRR

Problem 2: Superman's Chicken Rescue
[POST] {endpoint}/public/superman-chicken
body : 
{
    "n": 6,
    "k": 10,
    "position": [1, 11, 30, 34, 35, 37]
}

Problem 3: Transaction Broadcasting and Monitoring Client
[POST] {endpoint}/public/transaction
body : 
{
    "symbol": "ETH",
    "price": 4500,
    "timestamp": 1678912345
}