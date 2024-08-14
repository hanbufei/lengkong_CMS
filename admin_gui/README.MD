不纠结了，放弃electron了，转而使用gpui（https://github.com/zed-industries/zed/blob/main/crates/gpui/README.md）

不是因为electron不好，而是打算通过这个子项目来学习rust......可惜了我已经完成的electron代码。

## 一、边写边学

crates内可以理解为所有的包，其中app为默认的主包，是在根目录下的Cargo.toml中定义的。同时，该toml文件中要定义工作空间内所有的依赖。

类似crates/ui这种为lib依赖，在crates/ui/Cargo.toml依赖及相关参数。 这样，在crates/app主包中就可以通过use ui::xxx复用相关代码。

## 二、打包

```
cd crates/app
cargo bundle --release
```

采用cargo-bundle进行打包，详见https://crates.io/crates/cargo-bundle。
