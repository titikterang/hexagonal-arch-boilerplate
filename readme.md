# Setup & Development
---------------------------------------------
1. Package structure
2. Define Primary Ports, Secondary Ports
3. Define Models & Data Structure 
4. Proto Contract
5. Generate Proto File
6. Define Ports & Models
7. Implement Primary & Secondary Adapter
8. Implement Core Business Logic
9. Testing

# Wallet Service
------------------------------------------------
1. Deposit
   - user_id
   - amount
   - hset into redis
2. Get Balance
    - user_id
    - return amount
    - hgetall from redis
    - sum all amount from redis
 

# 