## 全排列

def dfs(path_list,root_list,nums,num):
    if path_list.__len__()==num:
        nums.append(path_list.copy())
        return
    for i in range(len(root_list)):
        path_list.append(root_list[i])
        dfs(path_list,root_list[:i]+root_list[i+1:],nums,num)
        path_list.pop()
        continue
root_list=[1,2,3]
nums=[]
path_list=[]
dfs(path_list,root_list,nums,root_list.__len__())