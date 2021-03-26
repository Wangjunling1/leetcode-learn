def func(nums):
    n = len(nums)

    candidate_k = [nums[n - 1]]
    max_k = float("-inf")

    for i in range(n - 2, -1, -1):
        if nums[i] < max_k:
            return True
        while candidate_k and nums[i] > candidate_k[-1]:
            max_k = candidate_k[-1]
            candidate_k.pop()
        if nums[i] > max_k:
            candidate_k.append(nums[i])

    return False
if __name__ == '__main__':
    func([-1,3,2,3,3,3,3,0])