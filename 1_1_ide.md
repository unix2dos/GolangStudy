# golang ide vscode + debugging



## vscode使用
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

### theme
    molokai 自带

### plugins:
    1. go
    2. vscode-icons
    3. code runner
    4. markdown preview github
    5. markdown auto-open
    
    
### vscode增加调试:
    1. xcode-select --install
    2. 创建钥匙链 dlv-cert 证书
    3. 证书签名
    cd $GOPATH/src/github.com/derekparker
    git clone https://github.com/derekparker/delve.git
    CERT=dlv-cert make install

