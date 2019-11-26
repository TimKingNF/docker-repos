SET GLOBAL group_replication_bootstrap_group=ON;

START GROUP_REPLICATION;

SET GLOBAL group_replication_bootstrap_group=OFF;

-- 查看MGR组信息
-- SELECT * FROM performance_schema.replication_group_members;

