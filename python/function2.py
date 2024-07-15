def terapkan_function(fungsi, deret):
    for i in range(len(deret)):
        deret[i] = fungsi(deret[i])
    return deret

def ubah_ke_string(x):
    return str(x)

print(terapkan_function(ubah_ke_string, [1,2,3]))