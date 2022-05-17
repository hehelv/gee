# gee

- 参考自极客兔兔7天实现gee框架系列

## day1
- 基于net/http 包实现简易的一个web服务器。
- 实现了静态路由的功能


## day2
- 将与请求相关的信息，封装到Context中


## day3
- 基于Trie树实现动态路由的功能，给router添加roots属性，记录每一个方法对应的路由树
- 实现了类似“：name”， “*” 的动态路由

## day4
- 实现了分组路由， 即路由组
- 路由组支持 分组嵌套，通过prefix前缀来区分路由组， 并且通过engine来间接使用engine的操作，如addRoute


