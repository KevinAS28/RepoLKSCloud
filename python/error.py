import math, json

def function_a(fn):
    with open(fn, 'r') as file:
        content =  file.read()
    return json.loads(content)
    

def function_b(fun, x):
    fun(x)

def function_c(x):
    return x
    
if __name__=="__main__":
    content = function_c(function_b(function_a, 'forloop.py'))
    print(content)
    