use mysql;
select host, user from user;
-- 因为mysql版本是5.7，因此新建用户为如下命令：
create user moneydodo identified by 'moneydodo';
-- 将MoneyDodo数据库的权限授权给创建的moneydodo用户，密码为moneydodo：
grant all on MoneyDodo.* to moneydodo@'%' identified by 'moneydodo' with grant option;
-- 这一条命令一定要有：
flush privileges;