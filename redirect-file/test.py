import time
while True:
    print(str(time.time())*1024)
#     fw=open(str(time.time())+".log","a+",encoding="utf8")
#     fw.write(str(time.time())+"\n")
#     fw.close()
    time.sleep(0.01)