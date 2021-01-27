from functools import lru_cache
def func(a):
    a.append(a[-1]+a[-2])
if __name__ == '__main__':
    a = [0,1]
    for i in range(3-1):
        a.append(a[-1] + a[-2])

    print(a[-1])