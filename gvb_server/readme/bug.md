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


# 3.

