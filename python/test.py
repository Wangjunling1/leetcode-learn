# CORS_ORIGIN_WHITELIST = (
#     '127.0.0.1:8000'
# )
# CORS_ALLOW_CREDENTIALS = True  # 指明在跨域访问中，后端是否支持对cookie的操作。
# CORS_ALLOW_METHODS = (
#     'DELETE',
#     'GET',
#     'OPTIONS',
#     'PATCH',
#     'POST',
#     'PUT',
#     'VIEW',
# )
#
# CORS_ALLOW_HEADERS = (
#     'XMLHttpRequest',
#     'X_FILENAME',
#     'accept-encoding',
#     'authorization',
#     'content-type',
#     'dnt',
#     'origin',
#     'user-agent',
#     'x-csrftoken',
#     'x-requested-with',
#     'Pragma',
# )
#
# test_list = [12,8,12,18,19,10,17,20,6,8,13,1,19,11,5]
# l=3
# m=6
# max_list1 = -1
# max_list1_list = {}
# max_list2 = -1
# # for i in range(test_list.__len__() - l + 1):
# #     max_int = sum(test_list[i:i + l])
# #     if max_list1 <= max_int:
# #         max_list1 = max_int
# #         max_list1_list = set(range(i, i + l))
# #
# #         print(max_list1, max_list1_list)
# #     max_int1 = sum(test_list[i:i + m])
# #     if max_list2 <= max_int1 and not max_list1_list & set(range(i, i + m)):
# #         max_list2 = max_int1
# #         print(max_list2, test_list[i:i + m])
#
# # for i in range(test_list.__len__()-m+1):
# #     max_int = sum(test_list[i:i + m])
# #     if max_list2 <= max_int and not max_list1_list&set(range(i,i+m)) :
# #        max_list2 = max_int
# #        print(max_list2,test_list[i:i + m])
# r1 = 0
# r2 = 0
#
# while r1 + l <= test_list.__len__() or r2 + m <= test_list.__len__():
#     max_int = sum(test_list[r1:r1 + l])
#     if max_list1 <= max_int:
#         max_list1 = max_int
#         max_list1_list = set(range(r1, r1 + l))
#     r1 += 1
#     if not set(range(r2, r2 + m)) & max_list1_list:
#         max_list2 = max(sum(test_list[r2:r2 + m]), max_list2)
#         r2 += 1
#     elif r1 + l > test_list.__len__():
#         r2 = max(max_list1_list) + 1
#
# print(max_list1, max_list2)
# print(r1, r2)
#
# max1 = 0
# a = 0
# hashmap = list()
#
# hashmap = list([[],[]])
# M=m
# L=l
# for i in range(test_list.__len__()):
#      hashmap[0].append(sum(test_list[i:i + M]))
#      hashmap[1].append(sum(test_list[i:i + L]))
# for a in range(test_list.__len__()):
#     a_max = sum(test_list[a:a + L])
#     if a+L<test_list.__len__():
#         max1 = max(max1, a_max + max(hashmap[0][a+L:]))
#     if a>
#         max1 = max(max1, a_max + max(hashmap[1][:a-L+1]))
#     if a + M < test_list.__len__():
#         max1 = max(max1, a_max + max(hashmap[1][a+M:]))
# print(max1)
def test(num:int):
    for i in range(num,0,-1):
        yield i

a=test(2)
print(next(a))
print(next(a))
# print(next(a))