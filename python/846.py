"""
顺子牌

"""
import collections
def func(nums,count):
    dict_count=dict()

    dict_count=collections.Counter(nums)
    for i in sorted(dict_count):
        if dict_count[i] > 0:
            val = dict_count[i]
            for key in range(i, i + count):
                if dict_count[key] >= val:
                    dict_count[key] -= val
                else:
                    return False

    return True


if __name__ == '__main__':

    # print(func(nums=[1,2,3,4,2,3,4,5,1,2,3,4,6],count=4))
    print(func(nums=[7,1,3],count=4))

