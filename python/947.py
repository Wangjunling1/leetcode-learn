"""
并查集解题
"""

# n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。
#
# 如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。
#
# 给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。

# 解题
# 使用并查集实现该题解释
# 并查集模版
class bing:
    def __init__(self,len_list):
        self.queue=list(range(len_list))
        self.root=set()
    def find(self,index):
        if self.queue[index]!=index:
            return self.find(self.queue[index])
        else:
            return index
    def union(self,index1,index2):
        x,y=self.find(index1),self.find(index2)
        if x!=y:
            self.queue[x] = y
        else:

            return "父节点相同"
a=[[0,1],[1,0]]
disjoint_set=bing(len(a))
for x,y in a:
    print(disjoint_set.union(x,y))
print(disjoint_set.queue)
print(disjoint_set.root)


for i,k in a:
    print(disjoint_set.find(i))

if __name__ == '__main__':
    stones=[[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
    parent = list(range(len(stones)))

    n = 10010 # 这里设置这么多是为了区分 两个父节点，防止 合并时连成一片
    parent = list(range(2 * n))

    co=0
    # 并查集查找
    def findset(x):
        global co
        if parent.index(x)==parent[-1]:
            co+=1
        if x != parent[x]:
            return findset(parent[x])
        return parent[x]


    # 合并
    def union(i, j):

        parent[findset(i)] = findset(j)
        # 连通横纵坐标


    for i, j in stones:
        union(i, j + n)
    # 获取连通区域的根节点
    root = set()
    for i, j in stones:
        root.add(findset(i))
    print( len(stones) - len(root))
    print(co)

