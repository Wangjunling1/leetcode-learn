## 实现 类似 from functoaols import lru_cache 功能
from collections import deque, OrderedDict


class LruCache:
    """
    借助OrderedDict实现缓存。
    该对象使用有序字典进行操作。使用了OrderedDict的特性，有序队列。
    orderedDic，底层使用了双向链表和hashmap 进行数据有序操作。
    使用双向链表，记录有序性。hashmap实现唯一性。
    """

    def __init__(self, long_max=100):
        self.__long_max = long_max
        self.__cache_data = OrderedDict()

    def get(self, key):
        if self.__cache_data.get(key):
            self.__cache_data.move_to_end(key)
            return self.__cache_data[key]
        else:
            return -1

    def put(self, key, value):

        if self.__cache_data.__len__() >= self.__long_max and self.__cache_data.get(key) is None:
            self.__cache_data.popitem(last=False)
            self.__cache_data[key] = value
        else:
            if self.__cache_data.get(key) is None:
                self.__cache_data[key] = value
            else:
                self.__cache_data[key] = value
                self.__cache_data.move_to_end(key)

## TODO：待完善
class LinkNode():
    def __init__(self, val=None):
        self.prev = None
        self.next = None
        self.val = val


class LruCache_link(object):
    """
    使用双向列表实现。
    """

    def __init__(self, long_max,root):
        self.__longMax = long_max
        # self.__root = LinkNode()
        self.__root = root
        self.__map = dict()

    def del_key(self, key):
        root_ = self.__root
        prev = root_

        while not root_.next is None:
            if root_.val == key:
                prev.next=root_.next
                break
            prev = root_.prev
            root_ = root_.next


    def move_key(self, key, flag: bool):

        """
        flag:True = 移动到首
             False = 移动到末尾
        """
        if flag:
            pass
        else:
            pass

if __name__ == '__main__':

    test=LinkNode(0)
    test_node=test
    for  i in range(1,100):
        new=LinkNode(i)
        new.prev=test_node
        test_node.next=new
        test_node=test_node.next
    c = LruCache_link(100,test)
    # c.root=test
    c.del_key(3)


def test1():
    a = OrderedDict()
    c = LruCache(5)
    c.put(1, 2)
    c.put(2, 2)
    c.put(3, 2)
    c.put(4, 2)
    c.put(5, 2)
    c.put(6, 2)
    c.put(3, 4)
    c.put(7, 4)
    print(a)
