## 一、边写边学
### 1.简介
该子项目主要目的是我个人为了学习rust，故而从较简单的electron换到rust的gui库。一开始学习的是zed的gpui库，但是gpui还不成熟。

再次换到iced。iced是专注于简单性和类型安全的Rust跨平台GUI库。https://crates.io/crates/iced
### 2.crates结构
crates内可以理解为所有的包，其中app为默认的主包，是在根目录下的Cargo.toml中定义的。同时，该toml文件中要定义工作空间内所有的依赖。

类似crates/ui这种为lib依赖，在crates/ui/Cargo.toml依赖及相关参数。 这样，在crates/app主包中就可以通过use ui::xxx复用相关代码。

## 二、打包

```
cd crates/app
cargo bundle --release
```

采用cargo-bundle进行打包，详见https://crates.io/crates/cargo-bundle。
