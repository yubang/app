# coding:UTF-8


"""
自动化安装脚本
@author: yubang
"""


import os
import shutil
import sys
import subprocess

dir_path = os.path.dirname(os.path.realpath(__file__))

# 检测python版本
if sys.version[0] != "3":
    print("请使用python3")
    exit()

# 安装提示
print("欢迎安装paas小平台，禁止未授权用于商业用途！")
print("请注意请勿重复执行该安装脚本！")
print("请输入Y/y确认安装")
y = input()
if y != 'y' and y != 'Y':
    print("退出安装！")
    exit()


# 一点点说明
print("一点点说明：")
print("安装完成后请手动导入数据到数据库，sql在doc/db.txt")
print("版本号：1.0（bate）")
print("平台需要python3")
print("web控制台需要运行index.py")
print("任务调度需要运行script.py")
print("按任意键安装")
input()


# 创建必须的文件夹
dirs = (
    dir_path + "/data/code",
    dir_path + "/data/config",
    dir_path + "/data/db",
    dir_path + "/data/log",
    dir_path + "/data/nginx_config",
    dir_path + "/data/session",
)
for t in dirs:
    if not os.path.exists(t):
        os.makedirs(t)


# 读取配置信息
print("请输入登录凭证（建议使用随机字符串）：")
token = input()

print("请输入管理员用户名：")
username = input()

print("请输入管理员密码：")
password = input()

print("请输入数据库域名：")
db_host = input()

print("请输入数据库端口：")
db_port = input()

print("请输入数据库名称：")
db_name = input()

print("请输入数据库用户名：")
db_usrname = input()

print("请输入数据库密码：")
db_password = input()

# 写配置文件
fp = open(dir_path+'/install/data/account.conf', 'r')
data = fp.read()
fp.close()
d = data % (token, username, password, db_host, db_port, db_name, db_usrname, db_password)

fp = open(dir_path+'/data/config/account.conf', 'w')
fp.write(d)
fp.close()


# 复制必须的数据库文件
shutil.copy(dir_path+"/install/data/base.db", dir_path+"/data/db/base.db")


# 安装docker镜像
dockers = (
    (dir_path+"/docker/static", "docker build -t paas-static .", "正在安装静态资源镜像"),
    (dir_path+"/docker/python2", "docker build -t paas-python2 .", "正在安装python2镜像"),
    (dir_path+"/docker/python3", "docker build -t paas-python3 .", "正在安装python3镜像"),
    (dir_path+"/docker/go1_5", "docker build -t paas-go .", "正在安装go1.5镜像"),
    (dir_path+"/docker/nodejs4", "docker build -t paas-nodejs4 .", "正在安装nodejs4镜像"),
    (dir_path+"/docker/java8", "docker build -t paas-java8 .", "正在安装java8镜像"),
)
for docker in dockers:
    print(docker[2])
    subprocess.call("cd %s && %s" % (docker[0], docker[1]), shell=True)
    print("\n安装一个镜像完成！\n\n")


print("安装完成，如需帮助请登录：https://github.com/yubang/app")
