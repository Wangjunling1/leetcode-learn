s1 = "ab"
s2 ="eidboaooo"



def main(s1, s2):
    len_1, len_2 = len(s1), len(s2)
    if len_1 > len_2:
        return False
    char_count = [0 for i in range(26)]
    ascii_a = ord('a')
    for i in range(len_1):
        char_count[ord(s1[i]) - ascii_a] += 1
        char_count[ord(s2[i]) - ascii_a] -= 1
    for i in range(len_1, len_2):
        if isEqual(char_count):
            return True
        char_count[ord(s2[i - len_1]) - ascii_a] += 1
        char_count[ord(s2[i]) - ascii_a] -= 1
    return isEqual(char_count)


def isEqual(char_count):
    return not any(char_count)


if __name__ == '__main__':
    print(main(s1, s2))
