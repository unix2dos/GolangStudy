# golang ide vscode + debugging



### 1.用户设置
```
{
    "workbench.iconTheme": "vscode-icons",
    "files.autoSave": "afterDelay",
        "go.buildOnSave": true,
        "go.lintOnSave": true,
        "go.vetOnSave": true,
        "go.buildFlags": [],
        "go.lintFlags": [],
        "go.vetFlags": [],
        "go.coverOnSave": false,
        "go.useCodeSnippetsOnFunctionSuggest": false,
        "go.formatOnSave": true,
        //goimports
        "go.formatTool": "goreturns",
        "go.goroot": "/usr/local/go",//你的Goroot
        "go.gopath": "/Users/liuwei/golang",//你的Gopath
}
```

### 2.theme
    molokai 自带


### 3.plugins:
    1. go
    2. vscode-icons
    3. code runner
    4. markdown preview github
    5. markdown auto-open
    
    
### 4.vscode增加debug:
    1. xcode-select --install
    2. 钥匙链创建证书 dlv-cert
    
    3. 证书签名
    
    cd $GOPATH/src/github.com/derekparker
    
    git clone https://github.com/derekparker/delve.git
    
    cd delve
    
    CERT=dlv-cert make install

    4. launch.json可以不配置
### 5.vscode snippets
https://github.com/Microsoft/vscode-go/blob/master/snippets/go.json


### vscode 遇到的问题

#### flag provided but not defined: -goversion

一个是版本原因, 一个是vscode也要修改配置gopath, 坑爹

>Thank you, I was able to solve this by running brew uninstall --force go and then downloading the latest installer. Anyone who reads this and wants to use brew you could probably just do brew install go after the forced uninstall. I had to restart my terminal and Gogland after doing this.

#### vscode not jump define

```
        "go.useLanguageServer": true,
        "go.docsTool": "gogetdoc",
```

#### vscode could not launch process: exec: "lldb-server": executable file not found in $PATH

```
        xcode-select --install
```


#### vscode jump slow
安装https://github.com/sourcegraph/go-langserver
源码安装 需要 go install
```
  "go.useLanguageServer": true,
```

#### vscode goto Definition, Hover, Signature Help do not work for Dot imported functions 

go 1.92 貌似不行, 已提交vscode issue


#### vscode output window hide go
~/.vscode/扩展包/package.json 找到显示的
```
    "showOutput": "never"
```

