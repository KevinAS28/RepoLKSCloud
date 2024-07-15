# argument/parameter: kebutuhan bahan mesin kue
# print -> output program
# return -> output function dan langsung mengakhiri function
# kalau "return" (return kosong) itu sama aja seperti "return None"

# parameter: kebutuhan bahan function/mesin
# argument: actual bahannya

def tambah(x,y): # function ini punya 2 parameter, x dan y
  return x+y

# print(tambah("10","20")) # memanggil function dengan 2 argument yaitu 10 dan 20
# print(tambah(10,20))
# print(tambah([10,20],[20]))

def rata_rata(list_bilangan):
    hasil_penjumlahan = sum(list_bilangan)
    jumlah_bilangan = len(list_bilangan)
    avg = hasil_penjumlahan/jumlah_bilangan
    return avg

def hasil_bagi_dua(angka):
  if angka%2!=0: # kalau ganjil
    return "ganjil bang"
  else:
    return angka

a = list(range(10))
b = list(range(20,30))
c = list(range(20,30))

print(hasil_bagi_dua(9))

# print(rata_rata(10,20))
# print(rata_rata([10,20]))
# print(rata_rata(a))
# print(rata_rata(b))




exit()


def sum_numbers(int_arg, float_arg):
    hasil = int_arg + float_arg
    return hasil

total = sum_numbers(5, 3.14)
print(f"The sum is: {total}")


# ===============================================

def convert_and_sum(data_list):
  """This function converts elements in a list to floats and returns the sum."""
  total_sum = 0.0
  for item in data_list:
    float_item = float(item)
    total_sum += float_item
  return total_sum

mixed_data = [10, 3.14, "5", 20]  
total = convert_and_sum(mixed_data)
print('total:', total)