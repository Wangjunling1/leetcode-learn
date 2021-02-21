def aaa(l):
    a=0
    l_nums=[]
    for i in range(1,len(l)-1):
        if l[i]>l[a] and l[i]>l[i+1]:
            l_nums.append([a,i])
            a=i+1
        elif l[i]<l[a]:
            a=i

    if l[-1] > l[a]:
        l_nums.append(l[-1]-l[a])
    # num = 0
    print(l_nums)
    for i1,i2 in l_nums:

        if l_nums:
            l_nums.sort()
            for i in range(1,len(l_nums)+1):

                num+=l_nums[-i]
                if i==2:
                    break
        print(num)
        print(a)
if __name__ == '__main__':
    # aaa([3,3,5,0,0,3,1,4])
    # aaa([3,2,5,0,0,3,1,4])
    # aaa([1,2,3,4,5])
    # aaa([7,6,4,3,1])
    aaa([1,2,4,2,5,7,2,4,9,0])
