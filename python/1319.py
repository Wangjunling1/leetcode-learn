"""
并查集思路解题
"""


# 并查集模板
class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        # self.size = [1] * n
        # self.n = n
        # # 当前连通分量数目
        self.setCount = []

    def findset(self, x: int) -> int:
        if self.parent[x] != x:
            self.findset(self.parent[x])
        return self.parent[x]

    def unite(self, x: int, y: int):
        x, y = self.findset(x), self.findset(y)
        if x != y:
            if




class Solution:
    def makeConnected(self, connections) :

        set1 = dict()
        a=0
        uf = UnionFind(len(connections))
        for x, y in connections:
            uf.unite(x, y)
            if set1.get(str(x)+"x") or set1.get(str(y)+"y"):
                a+=1
                print(x,y)
            else:
                set1[str(x)+"x"]=True
                set1[str(y)+"y"]=True
            print(set1)
        print(a)
        print(uf.parent)
        print(uf.setCount)
        # return uf.setCount - 1
if __name__ == '__main__':
    a=Solution()
    a.makeConnected([[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]])
