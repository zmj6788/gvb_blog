# 1.更新用户权限或用户昵称后，是否需要实时强制用户重新登陆，刷新token?

不然用户在token刷新时间期限内，可以使用之前的权限和昵称保留登陆状态，信息有错误

## 对应接口：
``
```
router.PUT("/user_role",middleware.JwtAdmin(),userApis.UserUpdateRoleView)
```
# 2.

