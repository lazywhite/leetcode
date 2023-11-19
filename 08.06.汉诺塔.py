
class Solution(object):
    def hanota(self, A, B, C):
        """
        :type A: List[int]
        :type B: List[int]
        :type C: List[int]
        :rtype: None Do not return anything, modify C in-place instead.
        """
        def hano(n, A, B, C):
            if n == 1:
                C.append(A.pop())
            else:
                hano(n-1, A, C, B)
                C.append(A.pop()) # 每次只能从顶部拿数据, 而不是底部
                hano(n-1, B, A, C)
        hano(len(A), A, B, C)
