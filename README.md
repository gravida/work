# gcs

```
用户数据
    Id、OpenId、
```
```
数量型表
    Id、Name、Desc Price
数值型表
    Id、Name、Desc

模版表
    Id、Name、Desc

模版列表表
    Id、TemplateId、EntryId、EntryType    

```

api管理

    用户添加 /v1/users POST

```sql
Insert into api_group values('API组管理');
Insert into api_group values('API管理');
Insert into api_group values('应用管理');
Insert into api_group values('用户管理');
```


用户管理
    登录 注册 用户列表 用户信息 用户创建 用户更新

1 登录 /v1/login POST
1 用户列表 /users GET
1 用户信息 /users/:id GET
1 用户创建 /users POST
1 用户更新 /users/:id PUT

Doctor summary (to see all details, run flutter doctor -v):
[✓] Flutter (Channel stable, v1.0.0, on Mac OS X 10.13.1 17B1003, locale zh-Hans-CN)
[✓] Android toolchain - develop for Android devices (Android SDK 27.0.3)
[!] iOS toolchain - develop for iOS devices (Xcode 9.1)
    ✗ libimobiledevice and ideviceinstaller are not installed. To install with Brew, run:
        brew update
        brew install --HEAD usbmuxd
        brew link usbmuxd
        brew install --HEAD libimobiledevice
        brew install ideviceinstaller
    ✗ ios-deploy not installed. To install with Brew:
        brew install ios-deploy
    ✗ CocoaPods not installed.
        CocoaPods is used to retrieve the iOS platform side's plugin code that responds to your plugin usage on the Dart side.
        Without resolving iOS dependencies with CocoaPods, plugins will not work on iOS.
        For more info, see https://flutter.io/platform-plugins
      To install:
        brew install cocoapods
        pod setup
[✓] Android Studio (version 3.1)
[!] IntelliJ IDEA Community Edition (version 2016.3.4)
    ✗ Flutter plugin not installed; this adds Flutter specific functionality.
    ✗ Dart plugin not installed; this adds Dart specific functionality.
    ✗ This install is older than the minimum recommended version of 2017.1.0.
[!] VS Code (version 1.30.1)
[✓] Connected device (1 available)


q100909...    

Alais:zuoxiang the paper media system

password:zxzn1127 

first and last name:xingru song 

organinational unit:technical department 

organination:Shenzhen zuoxiang intelligent and technology co.LTD 

City:Shenzhen 

State or provinc:guangdong 

Country code:86


### 权限设计

操作表
	组id
	权限名
	关键字

操作组表
	id
	组名称


```

权限操作关键字格式
	添加：A
	修改：U
	删除：D
	查找：Q

```

```
列表返回格式
{
	“pager”: {
		"page":1,
		"pageSize": 10,
		"total": 11,
	},
	"data":数组
}
```

CREATE TABLE Permission
(
id int primary key,
t_id int,
type int
);

CREATE TABLE Operation
(
id int primary key,
name varchar(255),
desc1 varchar(255),
key1 varchar(255)
);


INSERT INTO Permission VALUES (1, 1, 1);
INSERT INTO Permission VALUES (2, 2, 1);
INSERT INTO Permission VALUES (3, 1, 2);


INSERT INTO Operation VALUES (1, "添加用户", "add user ....", "add-user");
INSERT INTO Operation VALUES (2, "删除用户", "del user ....", "del-user");


select *,Permission.id from Operation,Permission where Operation.id 
in (select Permission.t_id from Permission where Permission.type=1);
select Permission.*,Operation.* from Permission
left
join 
Operation 
ON Permission.t_id=Operation.id
where Permission.type=1 
order by Permission.id asc limit 2 ;

select * from Permission, Operation where Permission.t_id = Operation.id
and Permission.type = 1;


select rowid,t.* from ( select *,(select count(1) from Permission) count from Permission ) t;
