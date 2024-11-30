import logging


class Solution:
    def longestCommonPrefix(self, strs):
        s = ""
        if len(strs) == 1:
            return strs[0] 
        rs = ''
        flag = True
        mini = strs[0]
        for i in strs:
            if len(i)<len(mini):
                mini = i
            
        for i in range(len(mini)):
            for j in range(len(strs)):
                if strs[0][i] != strs[j][i]:
                    flag = False
                    break
            if flag == True:
                rs += strs[0][i]
            else:
                break
        return rs
        


if __name__=="__main__":
    try:
        strs = ["flower","flow","flight"]
        expected = 'fl'
        s = Solution()
        status = 'PASSED'
        if expected != s.longestCommonPrefix(strs):
            status = 'FAILED'
        print(f'<output>\n{status} input: {s.longestCommonPrefix(strs)} expected: {expected}\n</output>')
    except Exception as ex:
        print(f"error: {ex}")
# docker run -v=.:/tmp/run  python python -u /tmp/run/pyrunner.py