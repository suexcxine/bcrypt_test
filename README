## 目标
验证golang.org/x/crypto/bcrypt是否可用

## 思路
测试用该包生成的hash是否可以用python verify,
以及使用python的passlib.hash.encrypt生成的hash是否可以用该包verify

## 使用说明

确保gen.py和verify.py有可执行权限
```
chmod +x gen.py verify.py
```

使用python生成20个随机字符串并hash,使用go验证
```
./gen.py 20 | go run verify.g
```

输出如下:
<pre>
0QYIEeJd $2a$04$E4kFvo8bZl1Z7meQAA34WedNPE2suYExfO5JNO6me96Oov0H.huUi verify ok
UfueNvpM $2a$04$gk4BY6inumTf5ziGTjt72uDeCQA.Elfhh5HotMH6T1nz7odvQDTS6 verify ok
OiaGqoDy $2a$04$1P0iPQPQ6BzYzms6eoULDeNoudJWaPGxUaablJgUJm3ni6ynvWbtW verify ok
Ua45W057 $2a$04$fRc7YXmrS0NyM41H9mHt8.BI7UmHTVJBReiXkBjK5PVgAr25TKC/C verify ok
noZShaOh $2a$04$M7bcvJXS81MB77XBS/MSf.MFwtXw2yJ3zYFs3W/n2n3yU3tGKpOR2 verify ok
RQfY83Eh $2a$04$/ikPG/o49c3IvXcGfiPCb.RrvUfnDgd9II2Y2VzM/2lkETgBR/I/K verify ok
MBCb3ibl $2a$04$uf9lJRSHN8xHFlZB5QNnceFOQ1WM1HaQhlnFN3IC6BGPXbhpvrVAC verify ok
RTDraMft $2a$04$VWjFq2hyYYkK6O4tuKbauuvdIZK5OQ19QtdoNKixHTFl6a0jQqF0u verify ok
WeTjQmOI $2a$04$.yYOn7qLsij0kRlJZvEAxeVc83RY6J8Z4z0pQNamUn9XirFOl5ElW verify ok
qmb2k6Eh $2a$04$qmpkna93/ic6T6F/26Jv7.J9Mufaz4RBDYRsCF72dF8OFtxNEicWq verify ok
5gbmNG4N $2a$04$FPq3DKINFSnZwq6YsHNXSuP3NUrqKrbsvf.8otAAjz95zMBgDYTYG verify ok
p4iqvngB $2a$04$4Kr7AFWY9VUCFUF1gOq1guw6tPYtlXDc8YVHPHqpMBz4lK8JvPFu6 verify ok
UCGlomWW $2a$04$n.VSi3mnwoSfCZgh8JBLt.lzzZF./WM3HQrXG9Csp5v0sFsQHo9Cy verify ok
X2570r2M $2a$04$CSx/iERY34fGqBIwgz48K.DcZ09YdNmGtTZX5ok6YBWVZAiJRdh0e verify ok
aDSYtFmx $2a$04$2pg0MvH0M7Ox31j5CVE1z.3ZRyBXGfUToFSLknjWL3jj6d0x078Va verify ok
eyNQ1wAK $2a$04$n5MzZQrPQcWhhgA1aUDtDedIaoGkbFZEf5pSR25lN4gve575BaJWC verify ok
iA3MIWvH $2a$04$d0AK023dU0TJ.xz1oRnfjeCzdhZF9MD7.JRQshJy3N7eTH/ps1J82 verify ok
8LwMrrtH $2a$04$rfLln50Ww32Ye3oAFcuCueAH0t4.Yhn9zCtPY2sJv4goryCUqBOcS verify ok
JixXgEr6 $2a$04$aINpktGhGH7BKTusX.2rzOwCeZb3hEjyC/uKNoUJN61rEFFcG1tZm verify ok
BcPLPMgD $2a$04$BTiuUvi4YyXL6zYcwGPtMet.D6qYbHdfmR7n9yTJFX4ANtfSQCNGq verify ok
</pre>

使用go生成,使用python验证
```
go run gen.go 20 | ./verify.py
```

使用python生成,使用python验证
```
./gen.py | ./verify.py
```

使用go生成,使用go验证
```
go run gen.go 20 | go run verify.go
```

