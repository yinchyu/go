class a():
    def __init__(self):
        print("init")
    def showa(self):
        print("show a")
        self.showb()
    def showb(self):
        print("show b")




class b(a):
    def __init__(self):
        # 继承的时候如果不主动调用就不会初始化
        super(b, self).__init__()
        print("init teacher")
    def showb(self):
        print("show teacher")



if __name__ =="__main__":
    c=b()
    c.showa()
    c.showb()
    # init
    # init teacher
    # show a
    # show teacher
    # show teacher