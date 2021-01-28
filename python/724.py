"""
# leetcode 724题
# 解题思路 通过计算
# muns 列表全部总和为 total
# 右边的总和 total-sum-nums[i]
# 由此可推断 左边的综合为 sum = total-sum-nums[i]
# 最终可得公式 sum+sum =total=nums[i]
"""

def main(nums):
    total=sum(nums)
    sum=0
    for i,v in enumerate(nums):
        if 2*sum==total-v:
            return i
        sum+=v
    return -1
if __name__ == '__main__':
    print(main([1, 7, 3, 6, 5, 6]))
    print(main([1, 2, 3]))
    print(main([2, 1, -1]))
    print(main([0, 0, 0, 0, 1]))
    print(main([-1,-1,-1,-1,-1,-1]))