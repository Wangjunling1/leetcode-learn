## 删除列表中重复的元素
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


test = ListNode(0)
cc = test

for i in [1, 1, 2]:
    cc.next = ListNode(i)
    cc = cc.next
sss=ListNode(0, test)
new_l = sss

while new_l.next and new_l.next.next:

    while new_l.next and new_l.next.next and new_l.next.val == new_l.next.next.val:
        new_l.next = new_l.next.next

    new_l = new_l.next


print(sss)