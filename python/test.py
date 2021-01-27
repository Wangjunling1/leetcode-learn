
def func(nums,s):
    i=0
    j=1
    flag=len(nums)
    while j<=len(nums):
        if sum(nums[i:j])>=s:
            flag=min(len(nums[i:j]),flag)
            i+=1
        else:
            j+=1
    print(flag)
if __name__ == '__main__':
    func([2,3,4,1,5,6],7)