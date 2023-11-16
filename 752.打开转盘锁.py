#
# @lc app=leetcode.cn id=752 lang=python
#
# [752] 打开转盘锁
#

# @lc code=start
class Solution(object):
    def openLock(self, deadends, target):
        """
        :type deadends: List[str]
        :type target: str
        :rtype: int
        """
        init_status = "0000"
        if target == init_status:
            return 0
        dead = set(deadends)
        if init_status in dead:
            return -1

        # 旋转一次得到的数字
        def next_num(num):
            if num == "9":
                return "0"
            return str(int(num) + 1)

        def pre_num(num):
            if num == "0":
                return "9"
            return str(int(num) - 1)

        def next_state(state):
            '''
            广度优先算法
            返回一个生成器, 每次返回某个轮子转动一次得到的状态, 可返回8个
            '''
            s = list(state)
            for i in range(4):
                num = s[i]
                s[i] = next_num(num)
                yield "".join(s)
                s[i] = pre_num(num)
                yield "".join(s)
                s[i] = num # 改回原来的数字
                
        q = deque([(init_status, 0)]) # 每个元素包含当前状态和已拨动的次数
        seen = set(init_status) # 已试过的状态
        while q:
            state, steps = q.popleft()
            for next in next_state(state):
                if next not in dead and next not in seen:
                    if next == target:
                        # 仅有正确的情况才会返回, 错误的会结束循环
                        return steps + 1
                    q.append((next, steps + 1))
                    seen.add(next)

        return -1
# @lc code=end

