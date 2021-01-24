def test(twoList):
    for i in range(len(twoList) // 2):
        for j in range(i, len(twoList[i]) - i - 1):
            twoList[i][j], twoList[len(twoList) - j - 1][i], twoList[-i - 1][len(twoList) - j - 1], twoList[j][
                len(twoList[i]) - i - 1] = twoList[len(twoList) - j - 1][i], twoList[-i - 1][len(twoList) - j - 1], \
                                           twoList[j][len(twoList[i]) - i - 1], twoList[i][j]
    return twoList


if __name__ == '__main__':
    twoList = [
        [5, 1, 9, 11],
        [2, 4, 8, 10],
        [13, 3, 6, 7],
        [15, 14, 12, 16]
    ]
    test(twoList)
