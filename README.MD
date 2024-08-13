<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">冷空CMS</h1>
<h4 align="center">基于结构化文本文件、服务器消耗小的内容管理系统</h4>

## 一、整体设计
冷空是一款模块化、结构化的CMS系统，其特点是后端服务器的消耗小。通过以下三种方式来降低消耗：
1. 直接让前端处理和渲染结构化的数据
2. 切割管理功能到单独的项目
3. 切割文件功能到单独的项目
### 1.模块化设计（工程目录）
本项目共分4个子项目，其各自关系如下图：
![1.png](doc/img/1.png)
* 后端项目：详见 server/README.MD，负责与数据库交互、处理业务逻辑，并提供api接口。
* 前端项目：详见 ui_react/README.MD，负责渲染结构化文件、展示界面，并提供用户交互。
* 文件项目：详见 file_server/README.MD，提供文件上传接口和文件读取接口。
* 管理端项目：详见 admin_gui/README.MD，负责整个CMS系统的管理，包含内容的增删改查等。
### 2.结构化文件的设计
本项目将内容以结构化文件的方式存储。在这里，内容，比如说一篇文章，就是一个类似如下的文本文件：
```
{"id":1,"lineType":"header","fontSize":14,"color":"black","bgColor":"white","value":"关于冷空CMS"}
{"id":2,"lineType":"title1","fontSize":12,"color":"black","bgColor":"white","value":"一级标题"}
{"id":3,"lineType":"title2","fontSize":12,"color":"black","bgColor":"white","value":"二级标题"}
{"id":4,"lineType":"body","fontSize":10,"color":"black","bgColor":"white","value":"正文"}
{"id":5,"lineType":"img","value":"1.png"}
```
前端项目将直接读取该文件并进行渲染，结构式数据解析如下：

id ：即文本文件的第几行，界面的第几个段落

lineType：该行的类型，共分7个：
* header ： 大标题，即内容名。
* title1 至 title4 ：一级标题到四级标题。
* body：正文。
* img：图片。

fontSize：字体大小

color：前景色

bgColor：背景色

value： 内容
### 二、快速开始
代码重构中，待全部完成后补充....... s山代码太多,不知道啥时候能完成