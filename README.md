### 参考

大規模サービス技術入門

### 概要

圧縮アルゴリズムであるVBCodeの疑似コードを参考にGolangで独自実装したもの。

### 課題

1. サンプルをギャップ＋VBCodeで符号化したものを書き出す
2. 書き出したバイナリを復元するプログラムを作成
3. オリジナルのテキストでどの程度圧縮できるか試す
4. VBCode以外の圧縮方法を試す

### 実行結果

```shell
$ ls -lh vbcode | grep eid
-rw-rw-r--@ 1 uenokensuke  staff   172M  6 25  2010 eid_tags.txt
$ ls -lh | grep output
-rw-r--r--@ 1 uenokensuke  staff    55K 10 18 23:19 output.txt
```
