##
step 1 同步主机数据，目前只支持aws,支持env

go run getHosts.go -env -key -secret
env 必须为test，或者prod 

####
step 2 运行./se 命令

./se strings

####
step 3 修改./se文件中本地ssh的登录私钥

cmd="ssh -l $USER  -p {0} -i //.ssh/private_key {1}".format(Port,IP)

####
step 4 修改 ./se中以下2个文件的路径为绝对路径

ct = open('./.cloudhosts_test',"r").readlines()
cp = open('./.cloudhosts_prod',"r").readlines()


####
step 5 将se 文件中放入环境变量PATH中
