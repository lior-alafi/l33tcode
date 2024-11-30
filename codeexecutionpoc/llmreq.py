import requests


inputs = '''strs = ["flower","flow","flight"]'''
code = ["""class Solution:
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
            """,

        """class Solution:
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
                rs += strs[0][i]+"a"
            else:
                break
        return rs
            """
]

for c in code:
    payload = {
        "model": "qwen2.5-coder-14b-instruct",
        "messages": [ 
        { "role": "system", "content": "given a code and inputs from the user:\n1. check the language the entire code was written if it's not python, reply to user {\"error\":\"WRONG LANGUAGE\"}, stop all executions and do not reply alternative solution!2. Simulate the given code instruction by instruction. do not fix it\n3. Report the trace of the given code and inputs at the end of each iteration.\n4. Think step by step and reply with the output of the code for the given input\n5. reply the simulated output in a json {\"input\": <put the inputs here>, \"output\": <put the fianl output of the simulated run on the given inputs here even if it's incorrect>}" },
        { "role": "user", "content":  f"@code@\n{c}\n@inputs@\n{inputs}"}
        ], 
        "temperature": 0.3, 
        "max_tokens": -1,
        "stream": False
    }
    print(c)
    print(inputs)
    r = requests.post("http://localhost:1234/v1/chat/completions", json=payload)


    print(r.status_code)
    response = r.json()
    print(response['choices'][0]['message']['content'])
    print("#"*5)



# LLM GOOD RESPONSE##
#
# class Solution:
#     def longestCommonPrefix(self, strs):
#         s = ""
#         if len(strs) == 1:
#             return strs[0] 
#         rs = ''
#         flag = True
#         mini = strs[0]
#         for i in strs:
#             if len(i)<len(mini):
#                 mini = i

#         for i in range(len(mini)):
#             for j in range(len(strs)):
#                 if strs[0][i] != strs[j][i]:
#                     flag = False
#                     break
#             if flag == True:
#                 rs += strs[0][i]
#             else:
#                 break
#         return rs

# strs = ["flower","flow","flight"]
# 200
# ```json
# {
#   "input": ["flower", "flow", "flight"],
#   "output": "fl"
# }
# ```
# ##### BAD CODE
# class Solution:
#     def longestCommonPrefix(self, strs):
#         s = ""
#         if len(strs) == 1:
#             return strs[0]
#         rs = ''
#         flag = True
#         mini = strs[0]
#         for i in strs:
#             if len(i)<len(mini):
#                 mini = i

#         for i in range(len(mini)):
#             for j in range(len(strs)):
#                 if strs[0][i] != strs[j][i]:
#                     flag = False
#                     break
#             if flag == True:
#                 rs += strs[0][i]+"a"
#             else:
#                 break
#         return rs

# strs = ["flower","flow","flight"]
# 200
# ```json
# {
#   "input": ["flower", "flow", "flight"],
#   "output": "fla"
# }
# ```
#####