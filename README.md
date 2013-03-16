gorpc
=====

一个基于golang的简单rpc示例

- service
包含了服务注册与启动、客户端远程调用的方法

- calculate
包含Add方法的服务接口，封装了service提供的远程调用方法

- server
发布服务的示例，是一个独立的服务器端。依赖于service，同时提供calculate中定义的服务

- client
服务使用端，直接使用caculate提供的接口

