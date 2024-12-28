def fibonnaci(n):
    if n <= 1:
        return 1
    else:
        return fibonnaci(n-1) + fibonnaci(n-2)
    
print(fibonnaci(0))
print(fibonnaci(1))
print(fibonnaci(2))
print(fibonnaci(3))
print(fibonnaci(4))