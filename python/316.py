"""
去重复并记录 字典序
"""


def main(s: str):
    dict_data={}

    str_list = [0 for _ in range(26)]
    new_list = []
    for i,k in enumerate(s):
        flag = ord(k) - 97

        if str_list[flag]:
            index_flag=new_list.index(k)
            for i in range(index_flag+1,len(new_list)):
                if new_list[index_flag]>new_list[i]:
                    new_list.remove(k)
                    new_list.append(k)
                    break
        else:
            str_list[flag] = 1
            new_list.append(k)
        # elif

    print(new_list)


if __name__ == '__main__':
    s = "cbacdcbc"
    # s="bcabc"

    main(s)
