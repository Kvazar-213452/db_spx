import os
import subprocess

let = input("Type ")

if int(let) == 0:
    os.system("git add -A")
    name = input("Name: ")
    os.system(f'git commit -m "{name}"')
    os.system("git push")
elif int(let) == 1:
    os.system(r'go build -o main.exe -ldflags "-H=windowsgui" main.go')
    os.system(r"main.exe 8080")
elif int(let) == 2:
    subprocess.run([r"test\main.bat"])