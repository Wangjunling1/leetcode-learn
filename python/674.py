a=[1,2,3,4,3,6]
flag=1
count=0
for i in range(a.__len__()-1):
    if a[i] < a[i+1]:
        flag+=1
    else:
        count=max(count,flag)
        flag=1
print(count)