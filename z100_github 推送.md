## github 推送


首先，在 GitHub 上 fork 到自己的仓库，如 docker_user/blockchain_guide，然后 clone 到本地，并设置用户信息。

```
$ git clone git@github.com:docker_user/blockchain_guide.git
$ cd blockchain_guide
$ #do some change on the content
$ git commit -am "Fix issue #1: change helo to hello"
$ git push
```



### [remote rejected] master -> master (permission denied)

```
Type command:

git config --global --edit
Add these lines of configuration at the end of file:

[credential]
  helper = osxkeychain
  useHttpPath = true
  ```