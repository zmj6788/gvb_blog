# 1.更新用户权限或用户昵称后，是否需要实时强制用户重新登陆，刷新token?

不然用户在token刷新时间期限内，可以使用之前的权限和昵称保留登陆状态，信息有错误

## 对应接口：
``
```
router.PUT("/user_role",middleware.JwtAdmin(),userApis.UserUpdateRoleView)
```

## 解决方式：将更改用户的token失效


# 2.管理员修改用户权限后用户无法调用退出登录接口

管理员修改用户权限后，将用户token注销失效，达到强制用户重新登陆，更新用户信息的目的

但也使得用户无法成功调用用户退出登录接口，无法正常退出

## 对应接口：

```
router.PUT("/user_role",middleware.JwtAdmin(),userApis.UserUpdateRoleView)

router.POST("/logout",middleware.JwtAuth(),userApis.UserLogoutView)
```

## 解决方法：退出登录接口不使用中间件


# 3.用户删除接口无法正常调用

## 错误提示：

```
Error 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```
## 对应接口：

```
router.DELETE("/users",middleware.JwtAdmin(),userApis.UserRemoveView)
```

## 解决方法：

[杀掉线程](https://blog.csdn.net/herry16354/article/details/141224846?spm=1001.2101.3001.6650.4&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EYuanLiJiHua%7EPosition-4-141224846-blog-76186661.235%5Ev43%5Epc_blog_bottom_relevance_base1&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EYuanLiJiHua%7EPosition-4-141224846-blog-76186661.235%5Ev43%5Epc_blog_bottom_relevance_base1&utm_relevant_index=6)

```

//查询未提交事务,查到一个一直没有提交的只读事务（trx_state=”[LOCK]WAIT”）
//找到对应线程kill 它
SELECT * FROM information_schema.INNODB_TRX;

//kill 线程ID。线程id为表中的trx_mysql_thread_id字段。
kill 8; 
```