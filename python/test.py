import math

try:
    num1 = 7
    num2 = int(input("masukan angka pembagi: "))
    num3 = int(input("masukan angka pengurang: "))
    print (math.sqrt((num1 / num2)-num3))
    print("Done calculation")
# except ZeroDivisionError:
#     print("pembagi ga boleh 0")
# except ValueError:
#     print('angka 3 ga boleh bikin minus')
except:
    print('Yang pasti error aja hasilnya')
    
print("Done calculation")

print("Program selesai")