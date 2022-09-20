快手
select id,consignee as name,address,mobile,remark from ks_order_info_v2_copy1 order by id
update ks_order_info_v2_copy1 set consignee=?,mobile=?,address=?,remark ='ns220901' where id=?

天猫
明文字段receiver_name，receiver_address，receiver_tel
加密后保存于encrypt_receiver_name，encrypt_receiver_tel，encrypt_receiver_address
后续明文字段进行了脱敏规则，脱敏后示例：杜*龙  159*****186	文笔镇****************
sql存储于天猫脱敏处理220908。

涉及sql
select count(*) from cust_tm_order_record;
select tid id,receiver_name as name,receiver_address as address,receiver_tel as mobile,remark from cust_tm_order_record order by order_time
update cust_tm_order_record set encrypt_receiver_name=?,encrypt_receiver_tel=?,encrypt_receiver_address=?,remark ='ns2295' where tid=?


抖音
明文字段receiver_name，receiver_address，receiver_tel
加密后保存于encrypt_receiver_name，encrypt_receiver_tel，encrypt_receiver_address
后续明文字段进行了脱敏规则，脱敏后示例：杜*龙  159*****186	文笔镇****************
sql存储于天猫脱敏处理220908。