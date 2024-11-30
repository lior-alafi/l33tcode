import requests

inputs = '''strs = ["flower","flow","flight"]
        expected = "fl"'''
code = "class Solution:\n    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:\n        n1 = len(nums1)\n        n2 = len(nums2)\n\n        return  self.__simple_median(sorted(nums1+nums2),n2+n1)\n\n     \n    def __simple_median(self,nums: List[int],n:int) -> float:\n        if n == 0:\n            return 0.0\n\n        if n == 1:\n            return nums[0]\n\n        if n % 2 == 0:\n            print(f'{n} {n//2} {n//2+1}')\n            return (nums[n//2 - 1]+nums[n//2+1 -1])/2\n        else:\n            print((n+1)/2)\n            return nums[(n+1)//2 -1]"
payload = {
    "model": "qwen2.5-coder-14b-instruct",
    "messages": [ 
      { "role": "system", "content": "given a code and inputs from the user:\n1. check the language the entire code was written if it's not python, reply to user {\"error\":\"WRONG LANGUAGE\"}, stop all executions and do not reply alternative solution!2. Simulate the code instruction by instruction.\n3. Report the trace of the code at the end of each iteration.\n4. Think step by step and reply with the output of the function for the given input\n5. reply the output in a json {\"input\": <put the inputs here>, \"output\": <put the output here>}" },
      { "role": "user", "content":  f"@code@\n{code}\n@inputs@\n{inputs}"}
    ], 
    "temperature": 0.7, 
    "max_tokens": -1,
    "stream": False
}
print(code)
print(inputs)
r = requests.post("http://localhost:1234/v1/chat/completions", json=payload)


print(r.status_code)
response = r.json()
print(response['choices'][0]['message']['content'])
