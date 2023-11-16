#
# @lc app=leetcode.cn id=307 lang=python
#
# [307] 区域和检索 - 数组可修改
#

# @lc code=start
class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        """
        n = len(nums)
        self.n = n
        self.seg = [0] * (n * 4) # tree node的数量为原始数组的4倍
        # 从seg的root开始构建, 其seg_id为0, 负责的范围是(0, n-1)
        self.build(nums, 0, 0, n-1)

    def build(self, nums, seg_id, start, end):
        '''
        利用数组构造一个segment tree
            size = len(nums)
            middle = start + (end = start) // 2
            left child负责[start, middle]
            right child负责[middle+1, end]
        递归地构造出某个node的所有子node
        seg_id: 在segment数组里面的下标 
        start: 在nums数组的范围开始下标
        end: 在nums数组的范围结束下标
        '''
        if start == end:
            self.seg[seg_id] = nums[start]
            return
        middle = start + (end - start) // 2
        self.build(nums, 2*seg_id + 1, start, middle)
        self.build(nums, 2*seg_id + 2, middle+1, end)
        self.seg[seg_id] = self.seg[2*seg_id+1] + self.seg[2*seg_id+2]

    def update(self, index, val):
        def change(index, val, seg_id, start, end):
            if start == end:
                self.seg[seg_id] = val
                return
            middle = start + (end - start) // 2
            if index <= middle:
                change(index, val, 2*seg_id + 1, start, middle)
            else:
                change(index, val, 2*seg_id + 2, middle+1, end)
            self.seg[seg_id] = self.seg[2*seg_id+1] + self.seg[2*seg_id+2]
        change(index, val, 0, 0, self.n - 1)

    def sumRange(self, left, right):
        def range(left, right, seg_id, start, end):
            if left == start and right == end:
                return self.seg[seg_id]
            middle = start + (end - start) // 2
            if right <= middle:
                # 完全落在左区间
                return range(left, right, 2*seg_id + 1, start, middle)
            if left > middle:
                # 完全落在右区间
                return range(left, right, 2*seg_id + 2, middle+1, end)
            # 左右区间都有覆盖
            return range(left, middle, 2*seg_id + 1, start, middle) + range(middle+1, right, 2*seg_id + 2, middle+1, end)
        
        return range(left, right, 0, 0, self.n - 1)


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(index,val)
# param_2 = obj.sumRange(left,right)
# @lc code=end

method = ["NumArray","sumRange","sumRange","sumRange","update","update","update","sumRange","update","sumRange","update"]
data = [[0,9,5,7,3],[4,4],[2,4],[3,3],[4,5],[1,7],[0,8],[1,2],[1,9],[4,4],[3,4]]

for m, d in zip(method, data):
    match m:
        case "NumArray":
            exec(f"obj = {m}({d})")
        case "update":
            exec(f"obj.{m}(*{d})")
        case "sumRange":
            exec(f"print(obj.{m}(*{d}))")
