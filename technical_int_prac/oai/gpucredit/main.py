import heapq

class GPUCredit:

    def __init__(self):
        self.credit = []

    def addCredit(self, creditID: str, amount: int, timestamp: int, expiration: int) -> None:
        self.credit.append([timestamp, creditID, amount, timestamp + expiration])
    
    def getBalance(self, timestamp: int) -> int | None:
        
        overflow = 0
        pq = []
        
        output = 0

        
            
        return output - overflow if output - overflow >= 0 else None

 
    def useCredit(self, timestamp: int, amount: int) -> None:

        #O(1)

        self.credit.append([timestamp, "", amount, -1])
def main():
    gpuCredit = GPUCredit()
    gpuCredit.addCredit('amazon', 40, 10, 50)
    gpuCredit.useCredit(30, 30)
    t = gpuCredit.getBalance(40)
    print(t)  # 10
    gpuCredit.addCredit('google', 20, 60, 10)
    t = gpuCredit.getBalance(60)
    print(t)  # 30
    t = gpuCredit.getBalance(61)
    print(t)  # 20
    t = gpuCredit.getBalance(70)
    print(t)  # 20
    t = gpuCredit.getBalance(71)
    print(t)  # None

if __name__ == "__main__":
    main()