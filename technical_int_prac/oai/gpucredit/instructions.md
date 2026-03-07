You're building a system to track balances for the OpenAI API credit purchasing system.

API users can purchase credit grants, specified by an ID, that are active at a timestamp and that let them use the API. Unused credit grant balances expire at a certain time.

However, there's a kink in our system: requests can arrive out of order because our system is built on a really unstable network. For example, a request to subtract credits can arrive at the system before the request to add credits, even though the request to add credits has a lower timestamp.

Your task is to implement the Credits class, which should support the following operations:

Granting credits, subtracting credits, and getting the credit balance for a user
The ability to handle requests that arrive out of order
Some guidelines/hints:

Do NOT worry about memory or performance concerns. Write simple, working code. No fancy data structures are needed here.
All timestamps can be represented as ints for simplicity and are unique (no two actions will happen at the same timestamp).
Subtract from grants expiring sonnest first.
class GPUCredit:
        def addCredit(creditID: str, amount: int, timestamp: int, expiration: int) -> None：
    # A credit is an offering of GPU balance that expires after some expiration-time. The credit can be used only during [timestamp, timestamp + expiration]. **Check with your interviewer whether this period is inclusive '[]' or exclusive '()'. Examples given were inclusive.** A credit can be repeatedly used until expiration.
       
        def getBalance(timestamp: int) -> int | None: # return the balance remaining on the account at the timestamp, return None if there are no credit left. Note, balance cannot be negative. See edge case below.
         
        def useCredit(timestamp: int, amount: int) -> None
Example 1:

gpuCredit = GPUCredit()
gpuCredit.addCredit('microsoft', 10, 10, 30)
gpuCredit.getBalance(0) # returns None
gpuCredit.getBalance(10) # returns 10
gpuCredit.getBalance(40) # returns 10
gpuCredit.getBalance(41) # returns None
Example 2:

gpuCredit = GPUCredit()
gpuCredit.addCredit('amazon', 40, 10, 50)
gpuCredit.useCredit(30, 30)
gpuCredit.getBalance(40) # returns 10
gpuCredit.addCredit('google', 20, 60, 10)
gpuCredit.getBalance(60) # returns 30
gpuCredit.getBalance(61) # returns 20
gpuCredit.getBalance(70) # returns 20
gpuCredit.getBalance(71) # returns None