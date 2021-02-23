# 该题很好的使用了FIFO原理。
# 通过X定义一个栈大小

class Solution:
    def maxSatisfied(self, customers, grumpy, X):
        total = 0
        sum_val = 0
        for i in range(len(customers)):
            if grumpy[i] == 0:
                total += customers[i]
            if i < X and grumpy[i] == 1:
                sum_val += customers[i]

        max_val = sum_val
        # 核心代码。判断进栈是否为数据是否为生气。如果生气将总数增加。
        # 反之如果出栈元素是生气则需要总数删减
        for i in range(X, len(grumpy)):
            if grumpy[i]: sum_val += customers[i]
            if grumpy[i - X]: sum_val -= customers[i - X]
            max_val = max(sum_val, max_val)
        return max_val + total


if __name__ == '__main__':
    customers = [10, 1, 7]

    grumpy = [0, 0, 0]
    X = 2
    test = Solution()
    test.maxSatisfied(customers, grumpy, X)
