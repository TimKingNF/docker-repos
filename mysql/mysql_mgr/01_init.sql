/* 没有数据库数据的时候才会执行, 并且创建数据库之后有部分数据属于其他组的事务, 需要手动调整过来

select @@global.gtid_executed;

SET GTID_NEXT='a1c5e25e-2715-11e7-bbe4-0800273fb9a2:1';
BEGIN;
COMMIT;
SET GTID_NEXT='AUTOMATIC';

START GROUP_REPLICATION;

单个节点挂掉不会重新加入组复制, 需要手动开启
如果所有节点都挂了需要原来启动的顺序重启开启
单主模式下, 只有 PRIMARY 节点可写
多主模式下, 如果事务修改了同一行数据,则事务本身不会出现阻塞,后提交的修改因为数据不一致而会触发事务回滚,需要重新更新提交
*/

/* 创建用户操作无需生成bin_log */
SET SQL_LOG_BIN=0;

/* 为group_replication_recovery通道设置组复制用户
 MySQL组复制在group_replication_recovery通道上工作，以在成员之间传输事务。因此，我们必须在每个服务器上设置具有REPLICATION-SLAVE权限的复制用户。
 */

/* 创建用户 */
CREATE USER repl@'%' IDENTIFIED BY 'repl';

GRANT REPLICATION SLAVE ON *.* TO repl@'%';

FLUSH PRIVILEGES;

SET SQL_LOG_BIN=1;

/* 告诉MySQL服务器使用我们为group_replication_recovery通道创建的复制用户 */
CHANGE MASTER TO MASTER_USER='repl', MASTER_PASSWORD='repl' FOR CHANNEL 'group_replication_recovery';

/* 安装MGR插件 */
INSTALL PLUGIN group_replication SONAME 'group_replication.so';


