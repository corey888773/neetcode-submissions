class Solution:
    def calculate(self, s: str) -> int:
        s += "+" # sentinel
        sign = "+"
        num = 0
        stack = []
        

        for i, c in enumerate(s):
            if c.isdigit():
                num = num * 10 + int(c)

            if c in ["+", "-", "*", "/"]:
                if sign == "+":
                    stack.append(num)
                if sign == "-":
                    stack.append(-num)
                if sign == "*":
                    stack.append(stack.pop() * num)
                if sign == "/":
                    stack.append(int(stack.pop() / num))
                
                num = 0
                sign = c

        return sum(stack)

