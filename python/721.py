# test=[["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
test=[["David","David0@m.co","David1@m.co"],["David","David3@m.co","David4@m.co"],["David","David4@m.co","David5@m.co"],["David","David2@m.co","David3@m.co"],["David","David1@m.co","David2@m.co"]]
##   [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]

new_test=[]
for i in range(len(test)):
    test1=set(test[i][1:])
    test1_name=test[i][0]
    # new_test.append(test[i])
    j=i+1
    while j<len(test) and len(test1)>1:
        if test1&set(test[j][1:]):
            test1=test1|set(test[j][1:])
            test[j]=[test[j][0]]
            j=i+1
        else:
            j+=1
    if test1:

        new_test.append([test1_name,*sorted(test1)])

print(new_test)