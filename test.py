# Python 示例代码

# 变量声明
example_string = "Hello, World!"
counter = 0

# 函数定义
def calculate_sum(a, b):
    return a + b

# Lambda 函数
multiply = lambda x, y: x * y

# 异步函数定义
import asyncio
import httpx

async def fetch_data(url):
    async with httpx.AsyncClient() as client:
        try:
            response = await client.get(url)
            data = response.json()
            print(data)
        except Exception as error:
            print('Error fetching data:', error)

# 类定义
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def greet(self):
        print(f"Hi, my name is {self.name} and I am {self.age} years old.")

# 对象和列表
person = Person("Alice", 30)
numbers = [1, 2, 3, 4, 5]

# 循环
for number in numbers:
    print(number)

# 条件判断
if counter >= 5:
    print("Counter is greater than 5.")
else:
    print("Counter is 5 or less.")

# 使用函数
print(calculate_sum(5, 7))
print(multiply(4, 3))
person.greet()

# 调用异步函数
asyncio.run(fetch_data('https://api.example.com/data'))

