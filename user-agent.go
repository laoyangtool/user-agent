/*
* @Author:  Yang
* @Email:   574502276@qq.com
* @Date:    2022/9/10 21:30
 */

package userAgent

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"sync"
)

const (
	userAgentPc = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4263.3 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.59 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4143.7 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.123 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4143.7 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.55 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.9 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4209.2 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.20 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4021.2 Safari/537.36
Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4257.0 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.95 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36
Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.123 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_16_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/82.0.4085.2 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4021.2 Safari/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.86 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4209.2 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4260.0 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4259.3 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/82.0.4050.12 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3983.2 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36 Edg/80.0.361.62
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36 Edg/80.0.361.111
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4096.0 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4250.0 Iron Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.67
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 Edg/86.0.622.51
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 Edg/86.0.622.56
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 Edg/86.0.622.58
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 Edg/86.0.622.69
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36 Edg/86.0.622.38
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.20 Safari/537.36 Edg/87.0.664.12
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.27 Safari/537.36 Edg/87.0.664.18
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36 Edg/87.0.664.41
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36 Edg/87.0.664.47
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4287.0 Safari/537.36 Edg/88.0.673.0
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4292.2 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4295.0 Safari/537.36 Edg/88.0.680.1
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4310.3 Safari/537.36
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4321.0 Safari/537.36 Edg/88.0.702.0
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4331.0 Safari/537.36 Edg/89.0.713.0
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4238.0 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36
Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36
`
	userAgentAndroid = `Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3404.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A3010 Build/OPM5.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3402.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36(KHTML, like Gecko) Chrome/61.0.3116.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3410.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo K50a40 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3417.3 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LG-H930 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3433.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Vega Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3430.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; P5W Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3430.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3434.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4 Pro Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3417.3 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3417.3 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MOB31V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3419.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3410.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H930 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3410.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3404.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700H Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 5 Plus Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-43-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G928F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1032 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 6.1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J510FN Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z (2) Build/OPXS27.109-34-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T07 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 5A Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G610F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Redmi Note 2 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FRD-L09 Build/HUAWEIFRD-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L21 Build/HUAWEICUN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VTR-L09 Build/HUAWEIVTR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SLA-L02 Build/HUAWEISLA-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M250 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MI 5 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo K10a40 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320FN Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel Build/OPM4.171019.021.P1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-N7100 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Redmi 3 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SLA-L22 Build/HUAWEISLA-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; SM-G3812B Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G600FY Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SLA-L22 Build/HUAWEISLA-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TAG-L21 Build/HUAWEITAG-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K430 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 5A Prime Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1032 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Mi Note 3 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi S2 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G903F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LLD-L31 Build/HONORLLD-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; RNE-L21 Build/HUAWEIRNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HBXCNCU5801811241S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LENNY3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-A530F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT X18 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi 3S Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9500 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9301I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5321 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-K130 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NX551J Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G355H Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G935F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX3 Build/HUAWEIANE-LX3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPS26.118-19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris X5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J120H Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D6503 Build/23.5.A.1.291) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700H Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BND-L21 Build/HONORBND-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6003 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J701F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3312 Build/43.0.A.7.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 Build/OPM2.171026.006.G1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Nexus 6 Build/N6F27M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9195 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; E2115 Build/24.0.B.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N9500 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Lenovo S850 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-795 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S70 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Lenovo K50-t5 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BND-L21 Build/HONORBND-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel XL Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5122 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.157) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TIT-U02 Build/HUAWEITIT-U02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LEX720 Build/WAXCNFN5902801241S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G600FY Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J710GN Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAN-L11 Build/HUAWEICAN-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D802 Build/KOT49I.D80220b) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi 3X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; HTC U11 Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00UD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1562 Build/NPD26.48-24-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV8000Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H4113 Build/50.1.A.10.51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00AD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Apollo Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J500M Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi MIX 2 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3112 Build/33.3.A.1.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/M4B30Z) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; C103 Build/ZIXOSOP5801803011S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; A7 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K220 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X00GD Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A310F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RNE-L21 Build/HUAWEIRNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; MI 4S Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 820 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M200 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo A7010a48 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HTC Desire 728G dual sim Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1033 Build/LPBS23.13-56-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DC Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A720F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; XT1650 Build/OPLS27.76-51-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; WIM Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo K50-t5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LLD-AL10 Build/HONORLLD-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3311 Build/37.0.A.2.248) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5321 Build/34.4.A.2.85) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ZTE BLADE V0820 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 626 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A605FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9500 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G928F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; KIW-L21 Build/HONORKIW-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Nexus 5 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8441 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930P Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D5803 Build/23.5.A.1.291) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Max2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Prime_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9500 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STV100-3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NXT-L29 Build/HUAWEINXT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RAINBOW 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K430 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J327VPP Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Alcatel_5044R Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LUA-L21 Build/HUAWEILUA-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ONEPLUS A5000 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 4027D Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; X5max_PRO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SAMSUNG-SM-G930A Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-92-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MPIS24.107-55-2-17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H930 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J510FN Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A520F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3211 Build/36.1.A.1.86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3111 Build/33.3.A.1.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-142) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041D Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-87) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5s Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G903F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-142) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Redmi Note 4X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C5 PRO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; W_C800 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-N7505 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PGN611 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo P1a42 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G925V Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Ulefone S7 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G318H Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; General Mobile 4G Build/N0F27E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 4 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE BLADE A0622 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; STF-L09 Build/HUAWEISTF-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One_M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VFD 610 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRA-LX1 Build/HONORPRA-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L09 Build/HUAWEICLT-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5503 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One M9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MEIZU_M5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L29 Build/HUAWEIEML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1572 Build/NPHS25.200-15-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-S7582 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X007D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoG3 Build/MPIS24.65-25.1-19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; CUBOT_NOTE_S Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo X3a40 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; T03 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BLN-L24 Build/HONORBLN-L24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-L29 Build/HUAWEIVKY-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3311 Build/37.0.A.2.108) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500H Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01HD Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 7 plus Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9506 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC U Play Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1663 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3112 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PLK-L01 Build/HONORPLK-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SUNSET2 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D5103 Build/18.1.A.2.25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A700FD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; HTC Desire 620 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Infinix X603 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LYO-L21 Build/HUAWEILYO-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S30 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi Note 2 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MIX2 Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6553 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI MAX Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J111F Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-X150 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; LG-D855 Build/LRX21R.A1447064049) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G532M Build/MMB29T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SGH-M919 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900H Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S1 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-15-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Life One X2 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L09 Build/HUAWEIBLA-L09S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G570F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LGL158VL Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2333 Build/26.1.B.3.109) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z (2) Build/OPXS27.109-34-13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LENNY3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300FU Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3212 Build/48.1.A.2.50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LG-X210 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6653 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_X00QD Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ALP-L29 Build/HUAWEIALP-L29S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-N910A Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Infinix Zero 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 825 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H3213 Build/50.1.A.10.51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320M Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9500 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; CPH1605 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FS8010 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Elephone P8000 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A700FD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Nexus 6 Build/N6F27M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SAMSUNG-SM-G890A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) Build/OPS27.82-87) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D415 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G5500 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_A007 Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel 2 XL Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LM-Q710(FGN) Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LM-G710 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 4 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 9001D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MotoE2 Build/LPCS23.13-56-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; CPH1725 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G357FZ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; HUAWEI MT1-U06 Build/HuaweiMT1-U06) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S60 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI SCL-L21 Build/HuaweiSCL-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.6.49) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo K53a48 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1004 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPSS27.76-12-25-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; H30-U10 Build/HuaweiH30-U10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Lenovo P70-A Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3112 Build/40.0.A.6.189) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-152) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 9001X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y625-U43 Build/HUAWEIY625-U43) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HEXCNFN5902012151S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo A2016b30 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G386W Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 8 SE Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X00BD Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-L09 Build/HUAWEIEVA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1575 Build/NPHS25.200-23-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9082 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J510MN Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J727T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530F Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NX541J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-H522 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) plus Build/OPWS27.113-51-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ALE-L23 Build/HuaweiALE-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.171) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5049Z Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3216 Build/36.1.A.1.86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STV100-1 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoE2(4G-LTE) Build/MPI24.65-39-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5s Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MS50S Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G950U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9200 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T08 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935V Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1032 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Lenovo A6000 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi 5 Plus Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00ED Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; JERRY2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930T Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel 2 XL Build/PPR2.180905.005) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; KISSME-DG580 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9505 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 4W Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-156) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M6T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PLK-L01 Build/HONORPLK-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X013D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A710M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L09 Build/HUAWEIVTR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3000 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J327V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J106H Build/MMB29Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L29 Build/HUAWEICLT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6833 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-162) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1580 Build/NPKS25.200-12-9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J250F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-C9000 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4A Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A520F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A720F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; DRA-L21 Build/HUAWEIDRA-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-G611F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BND-L21 Build/HONORBND-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; 6039A Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S16 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L41 Build/HUAWEIMYA-L41) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Pure 2A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L09 Build/HUAWEICLT-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; A1601 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 530 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3-TE Build/MPDS24.107-56-1-28) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Redmi Note 4 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenovo TB-7504X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X018D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H815 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6603 Build/10.7.A.0.228) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Coolpad R108 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U1 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Action-X3 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ASUS_Z00AD Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A310F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI P7-L10 Build/HuaweiP7-L10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; HTC Desire 12 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE B2017G Build/NMF26V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; R7sf Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N9005 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MIX 2S Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; A502_Aurum Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U20 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BlubooS8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J510FN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531BT Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2303 Build/26.1.A.3.111) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1053 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S60 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G935F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HT17Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; C1-U02 Build/ZJXOSOP5801805252S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H635 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5015D Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910T3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Lenovo A1010a20 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV7000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Halo 5 3G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONE A2003 Build/OPM1.171019.018) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI P7-L10 Build/HuaweiP7-L10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S41 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9295 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017DA Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G355HN Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HTC One X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3437.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G9250 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLL-L22 Build/HUAWEIBLL-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J327W Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRA-LX1 Build/HUAWEIPRA-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A320FL Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; vivo 1611 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; RAINBOW LITE Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I8200 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710MN Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Power_3 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; STF-L09 Build/HUAWEISTF-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S5S5IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Max Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; NOS Roya Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Halo3 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SAMSUNG-SM-G900A Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; m2 note Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris U2 Lite Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX722 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; N402 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Mi-4c Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-L19 Build/HUAWEIEVA-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; PSP5550DUO Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A605FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1020 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6653 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT CHEETAH 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P8_Mini Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; RNE-L22 Build/HUAWEIRNE-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-795 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-G900F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Thor Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ZTE BLADE A602 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Le X620 Build/HEXCNFN5801708221S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A310F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; P10 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D620 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARADDICT 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H631 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1058 Build/LPAS23.12-21.7-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3502.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S5IN4GP Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G920I Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Redmi Note 3 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; STARTRAIL 8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; HUAWEI G6-L11 Build/HuaweiG6-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L11 Build/HUAWEIMYA-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5533 Build/29.2.B.0.174) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1033 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q7S55IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HEXCNFN5902606111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris U2 Lite Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 5065X Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G7105 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930K Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N9005 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z1 Build/M4B30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D6543 Build/23.5.A.1.291) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-N960F Build/OPM5.171019.014) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; U_Pro Build/Elephone_U_Pro_20180227) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L31 Build/HUAWEIVNS-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H815 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FIG-LX1 Build/HUAWEIFIG-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; BBA100-2 Build/6.0.1_0.315.0.079) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VFD 500 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320F Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.73 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A8000 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-E500M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; 6045Y Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; X5max_PRO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00UDB Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NDQS26.69-64-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K5000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G955U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G850F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 510 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; LG-X150 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One_M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-V522 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-N9005 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X2 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; WIAM #71 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9082L Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G920F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S30 Build/LMY47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G357FZ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Moto G (5) Build/OPP28.85-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Bluboo S1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPS26.118-19-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FLA-LX3 Build/HUAWEIFLA-LX3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N9005 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A2 Lite Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.033) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HORIZON LITE PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F8331 Build/41.2.A.7.76) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.1; HTC One S Build/JRO03C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P9000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto E Build/OPP27.76) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Lenovo A6020a40 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BV7000 PRO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z017D Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; C6602 Build/10.5.1.A.0.283) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT R9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Moto G Play Build/MPIS24.241-15.3-26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense C30 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J250M Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-72) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1650 Build/NPL25.86-15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoG3 Build/MPI24.65-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G355M Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Armor_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-L09 Build/HUAWEIVKY-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M503 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VIE-L09 Build/HUAWEIVIE-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT X18 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E5823 Build/32.4.A.0.160) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; TECNO-N9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5321 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; 5099D Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; ASUS_Z00A Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J200BT Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Glow Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8232 Build/41.3.A.2.157) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D722 Build/LRX22G.A1525656093) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto e5 Build/OPP27.91-72) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; TA-1020 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5046Y Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1097 Build/MPES24.49-18-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-K420 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; S61 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5044S Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8342 Build/47.1.A.12.235) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1040 Build/LPBS23.13-35.5-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; hi6250 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9506 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H340n Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9350 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; VF695 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Vivo 5R Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Plus Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPS26.116-45) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPS26.116-26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500FN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX2 Build/HUAWEIWAS-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; X20 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G920A Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HBXCNCU5801811241S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; A33f Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; HUAWEI G700-U10 Build/HuaweiG700-U10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930T Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M3 Max Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L01 Build/HUAWEICUN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5025D Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 710 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; SM-G386F Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00AD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3212 Build/42.0.A.4.167) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; IMO Q Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; MI 5X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5049E Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 3 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J106B Build/MMB29Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; STARACTIVE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N900L Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G9250 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J330FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3121 Build/40.0.A.6.189) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H910 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G610F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo Plus Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-H540 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLN-L21 Build/HONORBLN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01KD Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5033X Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; STARTRAIL7 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARXTREM 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; QUANTUM_II_500_N Build/GOCLEVER_20161117_V04) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Fire 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Smart Ultra 6 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SGH-I527M Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3534.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; g5 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HBXCNCU5801811241S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5s Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; IRON_2 Build/1480648876) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; NEM-L51 Build/HONORNEM-L51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto X Play Build/NPD26.48-24-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HUAWEI G750-U10 Build/HuaweiG750-U10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; HTC U11 life Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HT37Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Max Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RNE-L21 Build/HUAWEIRNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Andromax A16C3H Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T04 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VFD 610 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi Note 3 Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Energy Phone Max 2  Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925W8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CUBOT_MANITO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; San Remo Mini Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S9  Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; RIDGE FAB 4G Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; D6616 Build/23.1.C.0.399) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A320FL Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-X150 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 4034A Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A720F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Q5S55IN4G Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 4LTE Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 6044D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5049Z Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S55IN4G2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N9100 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Z220 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N920C Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G355H Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8142 Build/47.1.A.12.75) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M5 Plus Lte Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MPIS24.107-55-2-17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TIT-U02 Build/HUAWEITIT-U02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DC Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00ED Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; GT-I9505 Build/NMF26V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E5 HD Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L29 Build/HUAWEIEML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; 2014813 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KAMBA Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8216 Build/51.1.A.4.265) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Vivo 5R Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; SAMSUNG-SGH-I537 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; TECNO K7 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900H Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; W_K400 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1092 Build/MPES24.49-18-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900FD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z011D Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9515L Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1421911913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; Galaxy Nexus Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Ektra Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 6043D Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320H Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Gigaset GS160 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Aquaris E5 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJ25.75-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G850F Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G920I Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D2212 Build/18.5.B.0.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H870 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Y6_Piano Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ZTE Blade V6 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-G900F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; InnJoo 3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Alpha Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPS27.76-12-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G532MT Build/MMB29T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; BLA-AL00 Build/HUAWEIBLA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2303 Build/26.1.A.2.167) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FRD-AL00 Build/HUAWEIFRD-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J727P Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; NX505J Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Trekker-M1 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-G900M Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CAN-L01 Build/HUAWEICAN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VIE-L09 Build/HUAWEIVIE-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Aquaris M5 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Archos 55 diamond Selfie Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G600FY Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950U1 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; HIGHWAY PURE Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D337 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LYO-L21 Build/HUAWEILYO-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV7000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; A0001 Build/N2G47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T6pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6833 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L01 Build/HUAWEICUN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3475.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E430 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; K6000 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-L09 Build/HUAWEIEVA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5059D Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-TL00 Build/HUAWEIEVA-TL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P8_Mini Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/MHC19Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8141 Build/47.1.A.12.34) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; AIRIS_TM55S Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A310M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900F Build/MOB31K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; GEOTEL G1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; UMI_MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Altice_S31 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5080X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8231 Build/41.3.A.2.75) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; DLI-L22 Build/HONORDLI-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3221 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJ25.93-14.5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo K52e78 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HT70 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.235) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G800H Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-87) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; S2 Build/MMB29U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; BLU STUDIO ONE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1020 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MS50L Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-N9005 Build/OPM1.171019.021) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J510FN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BV6000 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-AL00 Build/HUAWEIEVA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3466.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-N750 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 9001D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris U2 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/MMB29X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K500n Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D955 Build/KOT49I.D95520c) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; GT-N7100 Build/NMF26V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-L09 Build/HUAWEIEVA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MHA-L29 Build/HUAWEIMHA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D6643 Build/23.5.A.1.291) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto e5 Build/OPP27.91-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; PULP 4G Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H870S Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 710 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L41 Build/HUAWEIMYA-L41) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; NOS FIVE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S4S5IN3G Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; STARSHINE_5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 7 plus Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3449.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; NOS NOVU Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A720F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1053 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A520W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H955 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S6293T Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-157) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y625-U21 Build/HUAWEIY625-U21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DC Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FLOW_5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Le X626 Build/HEXCNFN5902303291S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View XL Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U PULSE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; XT1635-02 Build/OPNS27.76-12-22-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ZTE BLADE L110 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris M5.5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; C2105 Build/15.3.A.1.17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E976 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; XT1068 Build/LXB22.99-16) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; OZZY Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris X5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; vernee_M5 Build/vernee_M5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1069 Build/MPB24.65-34) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X017DA Build/NGI77B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H850 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI SCL-L21 Build/HuaweiSCL-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H320 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910C Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Archos 45c Helium Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L31 Build/HUAWEIVNS-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3112 Build/40.0.A.6.175) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5046Y Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M5 mini Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5010E Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Z Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G950F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; XT918 Build/2_330_2009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nokia 3.1 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.205) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-A300FU Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Smart A65 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-S7582 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300FU Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L29 Build/HUAWEIVTR-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; E2333 Build/26.3.B.1.33) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; YU5010A Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-795 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Power_3 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00ED Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J3110 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A510F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Z Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; R1 PLUS Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8552B Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Y100X Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-93-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A1 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L41 Build/HUAWEIMYA-L41) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5056D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LG-D855 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ALBA 5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3212 Build/48.1.A.2.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-92-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3475.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VIENNA Build/HUAWEIVIE) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S6812B Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C1X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0.0; PRO 6 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K500n Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-N9005 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Vodafone Smart ultra 6 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D5503 Build/14.3.A.0.757) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G928F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE B2017G Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A700H Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Vodafone Smart ultra 6 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.157) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B1-790 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; V210101 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; MI 5X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I8200N Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAN-L11 Build/HUAWEICAN-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 1 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel 2 XL Build/PPP3.180510.008) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G920F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Mi Note 3 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; E5603 Build/30.1.A.1.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; C2105 Build/15.3.A.1.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Aquaris X5 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HIGHWAY Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G900A Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Silver_S45 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; AIRIS_TM6SI Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SAMSUNG-SM-G928A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X014D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A910F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KENNY Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI P7-L12 Build/HuaweiP7-L12) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GETAWAY Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A3010 Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One A9s Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NOS NEVA 80 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; P8000 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L09 Build/HUAWEICLT-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo PB2-690M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G9350 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; STARTRAIL 8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris U Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G530W Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B1-733 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5010D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J500FN Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HIGHWAY PURE Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARXTREM 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; C6903 Build/14.2.A.0.290) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J600FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3524.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Plus E Build/MMB29U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L29 Build/HUAWEIVTR-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E467f Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Q4S6IN4G Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; STF-L09 Build/HUAWEISTF-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-45-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935V Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-L09 Build/HUAWEIVKY-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LGMS428 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3537.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2403 Build/18.6.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X013D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; F5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE Blade A522 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K520 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; HTC One_M8 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Note 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D6503 Build/17.1.2.A.0.314) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Y6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 4A Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D722 Build/LRX22G.A1429697571) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G361H Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI G7-L01 Build/HuaweiG7-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HUAWEI ATH-UL01 Build/HUAWEIATH-UL01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; A50TI Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1033 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5023F Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; thl_Knight_1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D405 Build/KOT49I.A1394026094) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ONEPLUS A5010 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; SM-G350 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; JY-G4 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; E39 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935V Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A730F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VS500 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SmartA35 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel XL Build/OPM4.171019.021.P1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041D Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Bluboo Maya Max Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_Z012DE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H850 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G361H Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SUNNY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8730 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3121 Build/40.0.A.6.135) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1032 Build/LPBS23.13-56-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00AD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; JERRY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; X520 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X522 Build/IEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C1001 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E440 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X2 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-L09 Build/HUAWEIGRA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense L678 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPI26.48-43) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D802 Build/KOT49I.D80220c) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; PH602 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1078 Build/MPBS24.65-34-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi Note Pro Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRO 7 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MI 6 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Tommy2 Plus Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo PB2-670M Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9195I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CUBOT MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPP26.102-16) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Armor_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto e5 Build/OPP27.91-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nokia 6.1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A5000 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; ZUK Z2131 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.9 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X5 Plus Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5086D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SWEGUE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; thor Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3221 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Nexus 4 Build/LMY48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H850 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SHT-AL09 Build/HUAWEISHT-AL09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; bq Aquaris 5.7 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ASUS_Z002 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MI 6 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Wooze_L Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Z971 Build/NMF26V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; DRA-L21 Build/HUAWEIDRA-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1053 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TAG-L21 Build/HUAWEITAG-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Vodafone Smart 4G Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G850F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H930 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) plus Build/OPWS27.113-45-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; WIM Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; XT1225 Build/MPGS24.107-70.2-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F5121 Build/34.3.A.0.194) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3533.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Hisense U989 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.73 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Redmi Note 3 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/M4B30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.8 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8730 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L01 Build/HUAWEICUN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/M4B30Z) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LG-D855 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SUNNY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Zenfone 2 Laser Build/NJH47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Smart Ultra 6 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; E5563 Build/29.1.B.0.101) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Gigaset GS160 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L29 Build/HUAWEICLT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Infinix HOT 4 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 6040D Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RNE-L21 Build/HUAWEIRNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; 5058I Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K200 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-E500M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Redmi Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8441 Build/47.1.A.2.374) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00DDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; D2303 Build/18.0.C.1.17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; TA-1032 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5090Y Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FIG-LA1 Build/HUAWEIFIG-LA1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ZTE BLADE A612 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Phone Build/OPM1.171019.011-CKH-180509) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Power_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C1001 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1 Build/HUAWEIWAS-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris U2 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-C7000 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955N Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-162) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FRD-L09 Build/HUAWEIFRD-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MAGIC Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Max Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; NOS NOVU II Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; NX531J Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; X15 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MIX 2 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C5_PRO Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI SCL-L01 Build/HuaweiSCL-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-G611F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; DRA-L21 Build/HUAWEIDRA-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; S50 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ASUS_T00J Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi MIX 2 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S5280 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; KIW-L21 Build/HONORKIW-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1024 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; KYOCERA-E6560 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Energy 400S Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-T380 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300FU Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; SM-G3812B Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5012G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Max3_Pro_LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920K Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H870I Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5603 Build/30.2.A.1.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; DRA-L21 Build/HUAWEIDRA-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H342 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI P7-L10 Build/HuaweiP7-L10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3434.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; JERRY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoE2(4G-LTE) Build/MPI24.65-39-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U20_Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoG3-TE Build/MPDS24.65-33-1-30) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z01HD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; GS370_Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A3003 Build/OPM2.171019.029) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; E2333 Build/26.3.B.1.33) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X018D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; FEVER Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; neon Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q7S5IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500FN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00UDB Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenny4 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M8 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5603 Build/30.2.A.1.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Moto G Play Build/MPIS24.241-2.47-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ZE553KL Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1040 Build/LPBS23.13-35.5-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; U_Pro Build/Elephone_U_Pro_20180423) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1580 Build/NPKS25.200-12-9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920P Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Power Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J105M Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D405 Build/KOT49I.A1413474368) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5010X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; RAINBOW 4G Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; JY-G4 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 900 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K500 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3111 Build/33.3.A.1.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-E975 Build/KOT49I.E97520a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z012DE Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300M Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3211 Build/36.1.A.1.90) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Mi-4c Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G610F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KENNY Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Andy C5VP Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_A002 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ONE E1003 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G530T Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8314 Build/51.1.A.11.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Moto G Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6533 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Le 1 Pro Build/CHXOSOP5501405221S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5303 Build/19.4.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3534.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Omega-X Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9550 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 Build/OPM2.171026.006.G1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-696 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TECNO LA7 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 4024D Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S5IN4GP Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 8 SE Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G930A Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LAIQ Dubai Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI CRR-L09 Build/HUAWEICRR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BND-L21 Build/HONORBND-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J700M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D405 Build/LRX22G.A1454843669) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I8200N Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; TECNO W5 Lite Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HUAWEI RIO-L01 Build/HuaweiRIO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MEDION X5020 Build/MEDION-X5020) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SAMSUNG-SM-G930A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; KIW-L21 Build/HONORKIW-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-72-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Andy 5EL2 LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H525n Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; K6000 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; UNO Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; S480 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3448.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-2-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; S60 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX2 Build/HUAWEIANE-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9105P Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HT7 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D802 Build/KOT49I.D80220h) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; A1 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MI 5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X2 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEAGOO M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Active Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-N7502 Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; CUBOT_X18_Plus Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ROBBY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; UMI_SUPER Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Archos 50 Cobalt Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-X150 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Archos 55 Diamond 2 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FRD-L09 Build/HUAWEIFRD-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View XL Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-696 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-N7505 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Fire 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8216 Build/51.1.A.11.26) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Vodafone Smart 4 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; E2115 Build/24.0.B.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris_E5_FHD Build/MOB31T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; I7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; D1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9600 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One M9 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H960 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BLN-L24 Build/HONORBLN-L24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3226 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPSS26.118-19-1-6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; GEOTEL G1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire EYE Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris V Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 5A Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VFD 200 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MT7-J1 Build/HuaweiMT7-J1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1635-02 Build/NPN26.118-22-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1032 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CAN Build/HUAWEICAN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G390F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320P Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01GS Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; ODYSSEY_Plus Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HUAWEI RIO-L01 Build/HuaweiRIO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L29 Build/HUAWEIEML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Quantum MUV UP Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; WIM Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-M153 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1572 Build/MPH24.49-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5017X Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z2121 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MHA-L09 Build/HUAWEIMHA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; RIDGE 4G Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 5A Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoG3 Build/MPI24.65-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; RAINBOW 4G Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320F Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6070K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; BLU Advance 5.0 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STH100-2 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX820 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; WIM Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_A007 Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BL5000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1454569215) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530FZ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F8331 Build/39.2.A.0.442) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Energy400LTE Build/Energy_400_LTE_FR_6.0_HW1.0_V7_20170110) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F5121 Build/34.2.A.0.333) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S6S5IN3G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L09 Build/HUAWEIBLA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M bot 60 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo 4 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-L19 Build/HUAWEIEVA-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Neffos Y5s Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-C7108 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ROBBY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; BBB100-2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T1 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5012G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) plus Build/OPW27.113-89) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; U_Pro Build/Elephone_U_Pro_20180227) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320F Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q7S5IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-G900F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S9 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L11 Build/HUAWEIMYA-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3417.3 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01HD Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX722 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenovo K33b36 Build/NRD90N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-33-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 4W Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3313 Build/37.0.A.2.248) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K350 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930V Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MI MAX 2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI MLA-AL10 Build/HUAWEIMLA-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI MAX 3 Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H4113 Build/50.1.A.10.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Alcatel_7049D Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 4G Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14.7-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CAM-L21 Build/HUAWEICAM-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MHA-L29 Build/HUAWEIMHA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T09 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; XT1635-02 Build/MCN24.104-35.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; S11 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAZ-AL10 Build/HUAWEICAZ-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HT70 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; S2 Build/MMB29U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1650 Build/NPLS26.118-20-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D405 Build/KOT49I.A1413474368) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930W8 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LGMS550 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 3 Build/OPM5.171019.015) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; San Remo Mini Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Z520 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MSP95013 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; LEX820 Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S30 Build/LMY47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FRD-L19 Build/HUAWEIFRD-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Build/N6F26U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; X15 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1020 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D2005 Build/20.1.A.2.19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1413907724) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense C1 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G360BT Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A1_PRO Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5046D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E4 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C5 PRO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D2303 Build/18.3.C.0.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H933 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-UL00 Build/HUAWEIGRA-UL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5095K Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6943 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E4.5 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-36-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ECHO MOSS Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J327T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 791 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC 2PXH3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S6S55IN3G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G130BT Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; STUDIO 6 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ASUS_Z00ADB Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; thor Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; XT1572 Build/OPM2.171019.029) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3416.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; vivo Y28 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X5 Plus Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; m3 note Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; 5058I Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NX575J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9060I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; NOS NOVU Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H4113 Build/50.1.A.3.141) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D410 Build/LRX22G.A1425182538) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G388F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Lenovo A936 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; P6000 Pro Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E4.5 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ANDY_55EI Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-H901 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Quantum Fly Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Vodafone Smart ultra 6 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; XT1635-01 Build/ODNS27.76-12-30-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Mi-4c Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; D6603 Build/23.1.A.1.28) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A500F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G920W8 Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_X00QD Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N900V Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5343 Build/27.3.B.0.173) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J510MN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J200M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris U Lite Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ALP-TL00 Build/HUAWEIALP-TL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5010 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense C30 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1020 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D852 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5026D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G800F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M320 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ONEPLUS A5010 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A600FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6603 Build/10.7.A.0.228) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Y6 Max Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H840 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MI 5 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X522 Build/IFXNAOP5802102141S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J120FN Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910W8 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/MOB31K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo TB-7703F Build/S100) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ASUS_T00J Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3535.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6603 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Paris Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-UL10 Build/HUAWEIGRA-UL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 3 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPS27.76-12-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J701F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LEX829 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; 2014813 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.73 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T710 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H870I Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U15 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Energy Phone Pro HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9192 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SUPER Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D2303 Build/18.3.C.0.37) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G390W Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MS60 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P8_Max Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Archos 55 Cobalt Plus Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ZTE Axon 7 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo Plus Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3437.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Platinum_5.0M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_A007 Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; HUAWEI G700-U10 Build/HuaweiG700-U10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 5054X Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q9S55IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H990 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Build/NRD90M.049) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; City Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1064 Build/MPBS24.65-34-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6045K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Build/NRD90M.007) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense C30 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00DD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3466.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MOVE_6_mini Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_Z012D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Mi-4c Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A520L Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5085N Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6683 Build/32.4.A.0.160) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; XT1068 Build/LXB22.99-16.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; H60-L02 Build/HDH60-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J250F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K500n Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K430 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G800F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J106H Build/MMB29Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.73 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; K4000_PRO Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.61-14-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; VF-895N Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi 4X Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Max Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT R9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Gigaset GS160 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Z981 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/M4B30Z) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-69) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-6-12) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Monaco Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D5503 Build/14.3.A.0.757) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H955 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M503 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; UMI TOUCH Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BND-L21 Build/HONORBND-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto X Play Build/NPD26.48-24-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE BLADE A0621 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X00GD Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K350 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Archos 50f Neon Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1032 Build/LPBS23.13-56-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; DIG-L01 Build/HUAWEIDIG-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HT16 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Hisense L695 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MX6 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00UDB Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1033 Build/LPBS23.13-56-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3486.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; E5533 Build/29.1.B.0.87) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X829 Build/FIXNAOP5801607292S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; RIDGE 4G Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G389F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N915W8 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; C1905 Build/15.4.A.1.9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Aquaris_M4.5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel XL Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9082 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Fire 4 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A700YD Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y360-U61 Build/HUAWEIY360-U61) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FEVER Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A700F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3212 Build/48.1.A.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H420 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; A1 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NDRS26.58-33-9-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531H Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM3.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI CRR-L09 Build/HUAWEICRR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC 10 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6633 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5051X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ANDY_5T Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; ZUK Z1 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.128) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Alcatel One Touch 7041X Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Y100X Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAN-L01 Build/HUAWEICAN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01BS Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D6603 Build/23.0.A.2.93) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-28) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3475.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense C30 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SGH-T399N Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; PULP 4G Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1021 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; E2105 Build/24.0.A.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX722) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3522.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D802 Build/KOT49I.D80220d) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3430.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; H3311 Build/49.0.A.2.62) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo3 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1413907724) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D405 Build/LRX22G.A1454843669) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.7.25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8141 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Pro Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; G8142 Build/45.0.A.7.137) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.3; D5303 Build/19.1.1.A.0.165) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930T Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; S60 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LG-M250 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8142 Build/47.1.A.8.49) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3420.1 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H736 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A2 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-G3502L Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H735 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G955U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos C5s Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.171) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; G735-L03 Build/HuaweiG735-L03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Platinum_5.0M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; VF-696 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5085D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; MEO Smart A75 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6003 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; eSmart 7&quot; Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; L-EMENT_403 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nexus 5X Build/OPR4.170623.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.73 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-K420 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5303 Build/19.4.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K450 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; PULP FAB 4G Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Vodafone 985N Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris M5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; HTC One_M8 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-S7580 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D325 Build/KOT49I.A1464618841) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M200 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900P Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 600 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01GS Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; DARKNIGHT Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7040A Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 4013X Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G925V Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1068 Build/MPB24.65-34) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Redmi Note 4X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; COL-L29 Build/HUAWEICOL-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NX573J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L21 Build/HUAWEICUN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos X1 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9060 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9650 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View XL Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Mars pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; HM NOTE 1S Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1635-02 Build/NPN25.137-15-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A9100 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z2132 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01BD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 4W Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; K00G Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-E975 Build/KOT49I.E97520a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G903F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1650 Build/NPLS26.118-20-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6943 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; T01 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI NOTE Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel Build/OPM4.171019.021.P1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; OPPO R11s Plus Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G800F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X820 Build/FGXOSOP5801910121S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI NXT-L29 Build/HUAWEINXT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T719Y Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F8132 Build/41.2.A.7.76) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3452.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense F20 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5321 Build/34.4.A.2.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SELFY Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H3113 Build/50.1.A.10.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Archos Diamond S Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500G Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X2 Pro Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3311 Build/37.0.A.2.248) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00DD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1068 Build/MPB24.65-34-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX720 Build/OPM1.171019.021) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3111 Build/33.3.A.1.115) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLU STUDIO SELFIE 2 Build/S230Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K8000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5303 Build/19.4.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6853 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-T820 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J110M Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1041 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ONE E1003 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos C5a Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 8 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3529.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAN-L01 Build/HUAWEICAN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9350 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 4034F Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VTR-L09 Build/HUAWEIVTR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5045D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1020 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CAN-L11 Build/HUAWEICAN-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5047Y Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S40 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; QUANTUM_470_RUGGED_PRO Build/GC_V01_171016) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.171) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5080X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910C Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LAIQ New York Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D6653 Build/23.5.A.0.575) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Luno Build/HuaweiY330-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HIGHWAY Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300YZ Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3221 Build/48.1.A.2.50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VFD 700 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00ID Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1650 Build/NPLS26.118-20-5-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D955 Build/KOT49I.D95520c) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Infinix X603 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G935F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3430.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; JERRY MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE A2017 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris M5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Pro Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC_D10i Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3221 Build/48.1.A.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531H Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H3113 Build/50.1.A.10.51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Classic_Pro Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; vivo Xplay5A Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FRD-L04 Build/HUAWEIFRD-L04) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320F Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D325 Build/KOT49I.A1407941488) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Elephone P8000 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KIICAA POWER Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BG2-W09 Build/HuaweiBAGGIO2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A1 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; C6603 Build/10.4.B.0.569) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A320FL Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM6.171019.030.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ALE-L23 Build/HuaweiALE-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CUBOT_MANITO Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto X Play Build/NPD26.48-24-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530H Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; TAB92QC-8GB Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_Z012DA Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; vivo 1603 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900S Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PLUS Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00UD Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VKY-L29 Build/HUAWEIVKY-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; MotoG3 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; TA-1024 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto x4 Build/OPWS27.57-40-17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PIC-LX9 Build/HUAWEIPIC-LX9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; QUANTUM_II_500_N Build/GOCLEVER_20161222_V07) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.3; D5322 Build/19.1.1.C.0.56) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X2 Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-6-9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; W_C800 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HTC Desire 820G PLUS dual sim Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 FHD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Y635-L01 Build/HuaweiY635-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-A730F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; HTC 10 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ZTE Axon 7 Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Z7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Max Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L01 Build/HUAWEICUN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; RAINBOW JAM Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; SM-G900F Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J500M Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE BLADE A512 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1032 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H873 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D722 Build/KOT49I.A1414136480) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910L Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HS-U961 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J120FN Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5012G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K10000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; OnePlus 5 Build/PPP3.180510.008) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SUNSET Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3421 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-X240 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; 7071D Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00ADA Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; TAB92QC-8GB Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D722 Build/KOT49I.A1422964302) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPS26.118-19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; E2115 Build/24.0.B.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; m3 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3421 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5059D Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-G870F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-D852 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-J500F Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180620.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H818 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Power Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U PULSE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ONE E1003 Build/MOB31K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Z Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPSS27.76-12-25-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X008DC Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.7.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-T378V Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ZP953 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5503 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NOS NEVA 80 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-AL00 Build/HUAWEIBLA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A720F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D335E Build/KOT49I.A1421913078) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BTV-DL09 Build/HUAWEIBEETHOVEN-DL09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1700 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KIICAA MIX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G9350 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J330F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; D6503 Build/23.1.A.0.690) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ZTE Blade L3 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris M5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H870 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3475.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G955F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TIT-U02 Build/HUAWEITIT-U02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; A8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX650 Build/KPXCNFN5902611041S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5503 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D852 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-64-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01GS Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI P6-U06 Build/HuaweiP6-U06) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D690n Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris V Plus Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX820 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S60 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H870 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8314 Build/51.1.A.4.225) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ZTE Blade L5 Plus Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Coolpad A8 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; KONNECT_510 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K10000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; S8 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ENERGY S600 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; STK_Sync_5e Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; MI 3W Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; WAX Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SAMSUNG-SM-N910A Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; KF6 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-46) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Classic_LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Aquaris E5 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG-SM-G925A Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Alba_8in Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; A50TI Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPP26.102-49) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; S99 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00AD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5046D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9506 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H422 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-45-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S40 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H850 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LEX829 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3522.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; D2005 Build/20.1.A.2.19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; HM 1S Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE BLADE A512 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Le 1 Pro Build/BGXNAOP5501103071S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC Desire 10 pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-12) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3522.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D802 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-A530F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H650 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G570M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MEIZU_M5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLN-L21 Build/HONORBLN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D722 Build/LRX22G.A1440040099) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2303 Build/26.1.A.3.111) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; A0001 Build/NOF27B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-61-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.235) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3212 Build/42.0.A.4.133) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX720 Build/OPM1.171019.021) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; BBF100-1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI MLA-L03 Build/HUAWEIMLA-L03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Note8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SUNNY2 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1033 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L31 Build/HUAWEIVNS-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 791 Build/LMY49F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5017X Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; P8000 6.0 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L22 Build/HUAWEIVNS-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Power Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One_M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Lenovo A936 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1024 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC One A9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-G611M Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Glow Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L03 Build/HUAWEIMYA-L03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense L678 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 6058D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-P875 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC Desire 816 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Nexus 5 Build/KTU84Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.7.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Grand M Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-TL10 Build/HUAWEIWAS-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F8331 Build/41.2.A.7.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 510 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; Nexus 4 Build/JWR66Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z016D Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9506 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NOS NEVA 80 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Monaco Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.149) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T11 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G388F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H4113 Build/50.1.A.10.51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM3.171019.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; TALENT Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-TL10 Build/HUAWEIWAS-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; KIW-L21 Build/HONORKIW-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; TECNO AX8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H320 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi Note 2 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00LDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1033 Build/LPBS23.13-56-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Aquaris_A4.5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ASUS_T00J Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C NOTE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Smart A65 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M503 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H500 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; X4_Soul_Infinity_Plus Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 900 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C8 4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC_M10f Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; W_K300 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531BT Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VR BOT 552 PLUS Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J720M Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-G900F Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MLA-AL10 Build/HUAWEIMLA-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5060D Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LT26i Build/6.2.B.1.96) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A320FL Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VFD 610 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N9208 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; MotoG3 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8331 Build/41.3.A.2.149) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01BD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LEX722 Build/WIXCNFN5802001232S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; XT1069 Build/LXB22.99-16.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VIE-L09 Build/HUAWEIVIE-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X520 Build/IEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; TECNO Camon CX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; STARNAUTE4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J510MN Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPSS27.76-12-25-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9515L Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9195I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X015D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5333 Build/27.3.B.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BL7000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; KING 7 Build/PP60002D_68I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F5121 Build/34.2.A.2.47) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H930 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5015D Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DC Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV7000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Archos 50c Platinum Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H872 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-L19 Build/HUAWEIEVA-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J200F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E4 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-E700M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G610M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A520F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M5 Build/M5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5026D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1635-02 Build/NPNS26.118-22-2-17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D5103 Build/18.1.A.2.25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HM NOTE 1 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-G7105 Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RNE-L22 Build/HUAWEIRNE-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 4X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-M153 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-N750 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G530H Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LYO-L21 Build/HUAWEILYO-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Z2131 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SAMSUNG-SM-G928A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S3 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5533 Build/29.2.B.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; WAX Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo3 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9505 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8231 Build/41.3.A.2.75) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONE A2003 Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; ONE LTE HD (ONE TORO) Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ZTE BLADE L7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S56 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI TIT-U02 Build/HUAWEITIT-U02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6603 Build/10.7.A.0.228) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U15S Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6045Y Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530MU Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Plus Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3493.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8190 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ONE A2003 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M250 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC_D10i Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Elite Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.3; HTC One_M8 Build/KTU84L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; T1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; XT1685 Build/OPM4.171019.021.Y1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ZTE Blade L3 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G965F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5010 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; thor Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S6IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Vivo 5R Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris M5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q8S55IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MHA-L29 Build/HUAWEIMHA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NCQS26.69-64-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; CINK FIVE Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H440n Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01HD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Venus_V3_5580 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC U Ultra Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U10 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-C900F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Advance 4.0 L3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z016D Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-33-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos X1 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H420 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HT7 Pro Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; bq Aquaris 5 HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Positivo Twist M Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8216 Build/51.1.A.4.225) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5010 Build/OPM4.171019.021.Y1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Pocket Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 600 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; vivo 1611 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LENNY Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; X6pro Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; K5 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 Build/OPM4.171019.021.Q1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; XT1069 Build/LXB22.99-16.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G570F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 3 Pro Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930V Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-12) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3212 Build/36.1.A.1.86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K373 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-36-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; S455 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G903M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S6312 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV6000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900T3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi Note 2 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; C6603 Build/10.6.A.0.454) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SAMSUNG-SM-G920A Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ONEPLUS A3003 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 710 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; a70 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 4S Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; TREKKER-X3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-AL00 Build/HUAWEIEVA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 6044D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; YU FLY F8 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1421911913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S10 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HTC One_M8 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3522.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; O1 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J600GT Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H815 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01MD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRA-LX3 Build/HUAWEIPRA-LX3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900M Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P9000 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-M160 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; K10 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5 Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S6S4IN3G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; C6903 Build/14.5.A.0.242) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; Mobile4you S999 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1033 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; G735-L03 Build/HuaweiG735-L03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-C115 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-72-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S41 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; D5 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; BBB100-4 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; CUBOT_P20 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01HD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S6IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5085D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI CUN-L01 Build/HUAWEICUN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 4LTE Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PrimuxDelta6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Vodafone Smart ultra 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3503.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; L-EGANTONE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Bluboo Maya Max Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D722 Build/LRX22G.A1525656093) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BAC-L22 Build/HUAWEIBAC-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S525_MT6735_65U_M Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Lenovo A936 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; SGH-I337M Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; SM-G920F Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI MAX 3 Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LGUS991 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SUNSET2 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00AD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Alcatel_7049D Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Apollo Lite Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 4049D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MEO Smart A80 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CRR-L09 Build/HUAWEICRR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L09 Build/HUAWEIVTR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-X230 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; HTC U11 Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H873 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G930F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/NHG47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5056D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; X4_Soul_Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-142) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; TREKKER-X3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; R2 LTE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Pixel Build/OPM4.171019.021.P1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Lenovo A806 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; FEVER Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L23 Build/HUAWEIMYA-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A1_PRO Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MIX Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Pro Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_Z012D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Pro Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E6853 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GETAWAY Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-A710M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.2.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3121 Build/40.0.A.6.189) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J710MN Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; DARKNIGHT Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5047Y Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6603 Build/10.7.A.0.228) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5052D Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MI 5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Gemini Pro Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; STF-AL00 Build/HUAWEISTF-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris X5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-E700F Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G389F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; BLA-AL00 Build/HUAWEIBLA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MYA-L22 Build/HUAWEIMYA-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Vodafone 985N Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; RAINBOW JAM Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J600FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S9 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6070K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Infinix X511 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; FEVER Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9505 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710MN Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-61-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z2 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; WIAM #24 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H650 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G5700 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-152) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; itel A15 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1078 Build/MPBS24.65-34-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1572 Build/NPHS25.200-15-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenny4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM4.171019.016.A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y625-U51 Build/HUAWEIY625-U51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00LDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Metal Build/F5B_GQ3030AH1_ulefone_20161226) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3420.1 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3522.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Pro Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5122 Build/34.4.A.2.85) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J105M Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1025 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Aquaris E5 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A3003 Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Mi A1 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; TECNO-C9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Z8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A300FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BLN-L24 Build/HONORBLN-L24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1092 Build/MPES24.49-18-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HIGHWAY STAR Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; HUAWEI G6-L11 Build/HuaweiG6-L11C02B120) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; HUAWEI GRA-L09 Build/HUAWEIGRA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris M5 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; G620S-L02 Build/HuaweiG620S-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G928G Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Paris Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E410i Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D855 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-72-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Blade L5 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Allure M1 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G928P Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-L29 Build/HUAWEIVKY-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MX4 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D325 Build/KOT49I.A1464618841) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; thor Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; FREDDY Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KIICAA MIX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Neffos Y50 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 5s Build/OPM2.171019.029.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G950F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T09 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ZTE T84 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; GT-I9195 Build/LMY48Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; NOS FIVE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BKL-L09 Build/HUAWEIBKL-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X820 Build/FEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F7pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One M9 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E5 HD Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LAIQ Dubai Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL LITE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Pro Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; CUBOT_X18_Plus Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; X4_Soul_Infinity_Z Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LGMS210 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-N910C Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE A2017G Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X821 Build/FGXOSOP5801910121S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LAIQ New York Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G901F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo Plus Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Mi A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3529.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L31 Build/HUAWEIVNS-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L09 Build/HUAWEIEML-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.9 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Q7T9INK Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LG-H811 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935T Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MI 5 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HIGHWAY SIGNS Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U1 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CRR-L09 Build/HUAWEICRR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1580 Build/NPKS25.200-12-9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N920P Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S60 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; AC7_DJ17 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; A0001 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A320FL Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6903 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-2-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense T5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_A002 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI RIO-L01 Build/HuaweiRIO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI CAM-L21 Build/HUAWEICAM-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K520 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; RNE-L21 Build/HUAWEIRNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ONE E1001 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925V Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; G7-L01 Build/HuaweiG7-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G930F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S6 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A720F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; H3311 Build/49.0.A.5.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K600 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-36-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Micromax Q450 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-64-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.2.50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I8200N Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1635-02 Build/NPN26.118-22-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-92-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G570F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000 Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; 5099D Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HTC Desire 830 dual sim Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; I7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2403 Build/18.6.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VFD 700 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BQru-5059 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D852 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N920G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel 2 Build/PPR2.180905.005) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Future Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00UD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; H3311 Build/49.0.A.5.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-G611MT Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6039K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X00BD Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; E2306 Build/26.3.A.1.33) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-L09 Build/HUAWEIGRA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9190 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9505 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Build/NJH47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; RIDGE 4G Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D855 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; H3311 Build/49.0.A.5.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE BLADE A512 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3517.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910L Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.7.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos X1 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D320 Build/KOT49I.A1467038589) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D405 Build/LRX22G.A1454843669) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; PH-1 Build/PPR1.180610.091) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00HD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A700YD Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T819 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Hisense U989 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; LG-D855 Build/LRX21R.A1439892504) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MHA-L09 Build/HUAWEIMHA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; U_Pro Build/Elephone_U_Pro_20180511) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris X5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; IM5 Build/KODAK-IM5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Apollo X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; A0001 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3491.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-L09 Build/HUAWEIGRA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Archos 50e Neon Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-N910F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D686 Build/KOT49I.D68620e) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ASUS_T00J Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Pure1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRISM Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; OPPO A83 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Hisense U989 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z017D Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8332 Build/41.3.A.2.171) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3-TE Build/MPDS24.107-56-1-28) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-A7100 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Plus E Build/MMB29U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T813 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G389F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VKY-L09 Build/HUAWEIVKY-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J5108 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-93-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 510 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Energy Phone Pro 3 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-N9002 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K520 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ASUS_T00I Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U22 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B1-780 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-E700M Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G9550 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T1 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris U2 Lite Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; VF-895N Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LUA-L02 Build/HUAWEILUA-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S6IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S6IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 6058X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.7.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955N Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; A0001 Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LG-H810 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3221 Build/42.0.A.4.133) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-S7582L Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; MediaPad X1 7.0 Build/HuaweiMediaPad) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5049W Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X00AD Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P8_Mini Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 710 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-43-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; JERRY2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A500H Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3211 Build/36.1.A.1.90) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Infinix HOT 3 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8190 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; P9000 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; DLI-L22 Build/HONORDLI-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MIX 2S Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8341 Build/47.1.A.12.145) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_A002 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-AL10 Build/HUAWEIEVA-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; E5803 Build/32.4.A.1.54) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-33) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Le X820 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NEM-UL10 Build/HONORNEM-UL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Power_3 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.4.46) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q7S5IN4GP Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5046G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A605F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I8200 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G950W Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H500 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B BOT 50 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y360-U61 Build/HUAWEIY360-U61) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MS70 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J500F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; NOS NOVU II Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; V370 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B1-733 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Pixel XL Build/NOF27C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Sunny2 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; LG-D955 Build/JDQ39B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Lenovo A916 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Redmi Note 4 Build/PPR1.180610.009) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI CRR-L09 Build/HUAWEICRR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.23 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; COR-L29 Build/HUAWEICOR-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; S57 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z2131 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 6042D Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9100 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC One A9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi 5 Plus Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MT7-L09 Build/HuaweiMT7-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MotoG3 Build/MPIS24.65-33.1-2-16) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NX551J Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DA Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1575 Build/NPHS25.200-23-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LENNY3 MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6003 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI KII-L22 Build/HUAWEIKII-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Blade S6 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; C6833 Build/14.5.A.0.242) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H3213 Build/50.1.A.10.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 6.1 Plus Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G920F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S6310N Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L09 Build/HUAWEIEML-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FS8002 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2302 Build/18.6.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H420 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J320M Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; RAINBOW JAM 4G Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One max Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Picasso Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Redmi Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Moto G Play Build/MPIS24.241-15.3-16) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; RAINBOW 4G Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-G611MT Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.3; HTC Desire 510 Build/KTU84L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z016D Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NX597J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Archos 55 diamond Selfie Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI CAN-L11 Build/HUAWEICAN-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLU STUDIO SELFIE 2 Build/S230Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-87) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CLT-L09 Build/HUAWEICLT-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VFD 610 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; P9000 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8T7IN Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9305 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A300FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M5 Note Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N9208 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X520 Build/IEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE BLADE V0800 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U10 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5023F Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; m3 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5026D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-N9005 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-64-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GT-I9301I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D722 Build/LRX22G.A1440040099) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HEXCNFN5902812161S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Nexus 4 Build/NZH54D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 7 plus Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G928T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARXTREM 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; GOA Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris E5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5303 Build/27.3.A.0.173) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; thor Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ZTE B2016 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ASUS_T00F Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530Y Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 4X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G600FY Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HTC Desire 626 Build/LMY47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ALP-AL00 Build/HUAWEIALP-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1024 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Armor Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; RAINBOW Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530FZ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; EML-L09 Build/HUAWEIEML-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530BT Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HARRY Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; RAINBOW UP Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C NOTE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJS25.93-14-6-7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; XT1032 Build/KXB20.9-1.10-1.18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6000 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE A2017G Build/NMF26V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; LG-D855 Build/LRX21R.A1449202626) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3425.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5051X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Pixel XL Build/OPR6.170623.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 5s Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (4) Build/NPJ25.93-14.5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel 2 Build/PPP3.180510.008) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; A8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Mi Note 2 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1020 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-J110H Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Z1 PRO Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; 6037Y Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris E5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5010D Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J727T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D335E Build/KOT49I.A1455181819) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Redmi Note 4X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MI 3W Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X520 Build/IEXCNFN5801708172S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LIVE4_KM0438 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-E450f Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; S57 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MI 2 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI P8max Build/HUAWEIDAV-701L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-K200 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Vivo XL3 Plus Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; bq Elcano 2 Quad Core Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-H850 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Neffos C5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; S61 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ZE553KL Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SELFY Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; KING 7 Build/PP60002D_68I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; RAINBOW LITE Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One_M8 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SUNNY MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; TA-1024 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-X230 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI P8max Build/HUAWEIDAV-701L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G850F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; PULP Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HT16Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; HTC 10 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K6000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1020 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G530FZ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FS8002 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 4 Build/OPM5.171019.017) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LEX820 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Moto Z3 Play Build/OPW28.70-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3489.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Archos 50 Saphir Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-N7505 Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-L29 Build/HUAWEIEVA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1058 Build/LPAS23.12-21.7-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1069 Build/MPB24.65-34-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L23 Build/HUAWEIVNS-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenny4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CPH1819 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ANSELMO_ONE Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L22 Build/HUAWEIVNS-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; STARNAUTE4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H818 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LAIQ Dubai Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-J100H Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A500FU Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U7 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-L09 Build/HUAWEIGRA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X017DA Build/NGI77B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; M12 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NXT-L29 Build/HUAWEINXT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2303 Build/18.6.A.0.175) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Orange-Rise31 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3116 Build/48.1.A.2.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SM-450 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-T550 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S6S55IN3G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q9T7IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Vertis 5510 Aim Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330FN Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ECHO_NOTE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; BBD100-6 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI RIO-L01 Build/HuaweiRIO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G920F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI P7-L10 Build/HuaweiP7-L10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; m8 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3221 Build/48.1.A.2.50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MX5 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; BLU STUDIO C SUPER CAMERA Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510S Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610M Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Mars pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S9 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI MAX 2 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; One Build/NZH54D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J110M Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3508.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo K33a48 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S55IN4G2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-A500H Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; A1-734 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9082 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; B BOT 50 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S40 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; PULP Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; QUANTUM_400_LITE Build/GC_20150204.V01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D955 Build/KOT49I.D95520c) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV8000Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; LG-D802 Build/JDQ39B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3437.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NDQS26.69-64-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; A0001 Build/PPR1.180610.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U PULSE LITE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8T7IN4G Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; COL-L29 Build/HUAWEICOL-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Vodafone Smart ultra 6 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J330F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; X5 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500M Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 10 lifestyle Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G930A Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 6058D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Seoul 5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; NOS NOVU III Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9350 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; WOLDER_34 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Lenovo K33a48 Build/NRD90N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 820 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/MMB29X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FRD-L19 Build/HUAWEIFRD-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X007D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-I8190L Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Z6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9060 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi Note Pro Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI NXT-L29 Build/HUAWEINXT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531BT Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto e5 Build/OPP27.91-72) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 6044D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J701F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; NX505J Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3311 Build/37.0.A.1.105) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01RD Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; RIDGE 4G Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; X Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MEO Smart A80 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Le X820 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One_M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F5121 Build/34.3.A.0.194) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-N9005 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 600 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MS50M Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos X1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X820 Build/FEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; QUANTUM_470_RUGGED_PRO Build/GC_V01_171016) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLN-L21 Build/HONORBLN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H873 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E6 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; bq Aquaris 5 HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; 5058I Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D2303 Build/18.3.1.C.0.21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; K7 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-2-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SGH-M919N Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2303 Build/18.6.A.0.182) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 4X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris U Plus Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9060 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3493.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-E986 Build/KOT49I.E98620a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-I9195 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; MW_5056 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LUA-L01 Build/HUAWEILUA-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; XT1635-02 Build/NPNS26.118-22-2-17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ZP955 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1413907724) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Mi MIX 2 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Andy 3.5EI Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ODYSEA 5HD Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9505 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View Prime Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Phone Build/OPM1.171019.011-RZR-180509) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-64-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MotoE2(4G-LTE) Build/LXI22.50-53) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; EVA-L19 Build/HUAWEIEVA-L19) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F8332 Build/41.3.A.2.149) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX720 Build/OPM1.171019.021) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; GT-I9505 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG-SM-N915A Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.85) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9550 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM2.171019.029.A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; MEO Smart A40 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; STARNAUTE3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1004 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; R7 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; VFD 720 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LOOK Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Z2 Plus) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3535.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N920T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LYO-L01 Build/HUAWEILYO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-D855 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M bot 60 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; FREDDY Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5033X Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; STARNAUTE4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S10 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G530T Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Xperia Z1 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; LG-P875 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STH100-2 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01KD Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; STARSHINE III Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-79) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; vivo X7 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; BLU Life One Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; S30 Build/LMY47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PIC-LX9 Build/HUAWEIPIC-LX9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPSS26.118-19-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; BV9000-F Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPIS26.48-36-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5046Y Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HT3 Pro Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; MI 5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ZTE Blade A522 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.24 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G800F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI SCL-L01 Build/HuaweiSCL-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; HUAWEI GRA-UL10 Build/HUAWEIGRA-UL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A500H Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BLA-L29 Build/HUAWEIBLA-L29S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; PULP Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A5100 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G361F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; NX591J Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-33-10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G928F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-S7710 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H960 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; RIDGE 4G Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; FP2 Build/FP2-gms-18.04.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 4049X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FIG-LX1 Build/HUAWEIFIG-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; NOS FIVE Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01KD Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5321 Build/34.4.A.2.85) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Vodafone Smart ultra 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-C115 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L29 Build/HUAWEIVTR-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BKL-L09 Build/HUAWEIBKL-L09S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 6060X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925I Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 7070X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-A310M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense C30 Lite Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D2403 Build/18.6.A.0.175) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 FHD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-G930F Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WIN4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; QUANTUM_3_500_Lite Build/GOCLEVER_170524.V03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; M-PPxS470 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 9; Pixel Build/PPR2.180905.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; GT-I9505 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; BBB100-2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.023) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-AL00 Build/HUAWEIEVA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HIGHWAY Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; FRD-L14 Build/HUAWEIFRD-L14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-113) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Plus Build/NMA26.42-11-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HIGHWAY SIGNS Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Just5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G935U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; POWER BOT Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T580 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ROBBY Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPR5.170623.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) play Build/OPP27.91-25) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; MX6 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; LEX722 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; LG-D855 Build/LRX21R.A1445306351) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Plus E Build/MMB29U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H420 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ATU-L21 Build/HUAWEIATU-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-26-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G930F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; AC7_DJ02 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Monaco Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Vodafone Smart ultra 6 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D120 Build/KOT49I.A1442298380) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G935T Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1021 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3003 Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; RAINBOW JAM Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J700M Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HTC One_M8 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G928C Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H960 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Z2131 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; HUAWEI Y530-U00 Build/HuaweiY530-U00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NXT-L29 Build/HUAWEINXT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; D2303 Build/18.3.1.C.1.17) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G386W Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; STARTRAIL7 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; HUAWEI G535-L11 Build/HuaweiG535-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MHA-L09 Build/HUAWEIMHA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9305 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H525n Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; R9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nexus 5X Build/OPR6.170623.023) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-K420 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; C6802 Build/14.4.A.0.108) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; ZUK Z1 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Vertis 5510 Aim Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SAMSUNG-SGH-I337 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.07.010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-H960 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; OPPO R11 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; 5022D Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Aquaris_A4.5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LEX720 Build/WAXCNFN5902303282S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G925F Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense F20 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI MT7-TL10 Build/HuaweiMT7-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; U Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Apollo Lite Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Crystal Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; moto g(6) plus Build/OPW27.113-27) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Glow Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Archos Core 55S Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LUA-L03 Build/HUAWEILUA-L03) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; MI 8 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; A1Q Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G903W Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ENERGY_500_LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A605FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; myPhone_AXE_LTE Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; force_2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; C6903 Build/14.3.A.0.757) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-64-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Primux_ioxphone Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-92-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; S31 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Moto Z2 Play Build/OPSS27.76-12-28-3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; XT1033 Build/KXB20.25-1.37) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; RAINBOW Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H4113 Build/50.1.A.10.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U7 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Aquaris X5 Plus Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Andy C5V Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Fire 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; MI 5X Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 5 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HUAWEI Y330-U11 Build/HuaweiY330-U11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5045J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925W8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z016D Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; XT1032 Build/KXB21.14-L1.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; JR6-B606 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; Archos 50b Oxygen Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Plus Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3475.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; E5823 Build/32.2.A.4.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; F5121 Build/34.4.A.2.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6039K Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J710MN Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; JIMMY Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; STARTRAIL7 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01KD Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Q8S5IN4GP Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; SM-N910F Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-N950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC_M10f Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N9100 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Nexus 4 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; BAC-L21 Build/HUAWEIBAC-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Le 2 Pro Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM6.171019.030.B1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610Y Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N910C Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; NOS NOVU III Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NXT-AL10 Build/HUAWEINXT-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Comet Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MIX 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LGMP260 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1572 Build/NPHS25.200-15-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Redmi 3 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-C105 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D337 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-N915FY Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5056X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-A730F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; D1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; C1_Max Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; XT1635-02 Build/OPN27.76-12-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; STARTRAIL 4 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Z1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VFD 700 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; 77X Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI RIO-L01 Build/HuaweiRIO-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G920I Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; LG-H500 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; WAS-TL10 Build/HUAWEIWAS-TL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; S5.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Twist XL Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J500F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi A1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; X4_Soul_Infinity_Z Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z01KD Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; ASUS_Z00AD Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-33) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G531H Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto C Plus Build/NRD90M.04.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARADDICT 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; MOCHE SMART A16 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J330F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; SmartA35 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G903M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 FHD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PRA-LX1 Build/HUAWEIPRA-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3417.3 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VFD 700 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; AX570 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MotoE2(4G-LTE) Build/LXI22.50-53.8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; A0001 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Elite Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SM-G360F Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D331 Build/KOT49I.A1421911913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; S3 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U FEEL LITE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X820 Build/FEXCNFN5902303111S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Quantum Go Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L23 Build/HUAWEIVNS-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ALE-L21 Build/HuaweiALE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 530 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A9000 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; L-EMENT_403 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; XT1580 Build/NPKS25.200-12-9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; VFD 510 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; STARADDICT 4 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G920F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; moto x4 Build/OPWS28.46-21-4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LENNY3 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T07 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G Play Build/NPI26.48-43) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; F3111 Build/33.2.A.4.70) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G530FZ Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.3; HTC One_M8 Build/KTU84L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; BLU NEO 5.5 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM6.171019.030.E1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; M-PPxS470 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; Nokia 3.1 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; 5059D Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G935S Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3535.2 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; PRA-LX1 Build/HUAWEIPRA-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Archos 55 Diamond 2 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; XT1635-02 Build/OPN27.76-12-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3478.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 FHD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S51 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 5X Build/OPM4.171019.016.A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; SM-G900F Build/OPM1.171019.018) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Bluboo S1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X522 Build/IFXNAOP5802012141S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; CUBOT_J3 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-K371 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; 6037K Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Nokia 2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Glow Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G900H Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Cynus_F10 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-J250F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; E5303 Build/27.3.A.0.129) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; RunboQ5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-2-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Heart Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; HUAWEI G6-L11 Build/HuaweiG6-L11C02B120) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9060I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; E2124 Build/24.0.B.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Eluga_A2 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MI 5s Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; P5_Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi Note 3 Build/OPM4.171019.021.R1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris M5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; BV7000 Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5095K Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00HD Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi 3S Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.53 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; VTR-L09 Build/HUAWEIVTR-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00LDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Armor_X Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F5121 Build/34.3.A.0.244) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; TA-1021 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3121 Build/40.0.A.5.14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; XT1032 Build/LPBS23.13-56.1-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG-SM-N915A Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo 4 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; NX575J Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3121 Build/48.1.A.2.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5056D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; RIDGE 4G Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J120FN Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900T3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5010 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3420.1 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730FM Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HIGHWAY STAR Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SLA-L22 Build/HUAWEISLA-L22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; GT-I9515 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G850F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ML-JI-M7_3G_PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LEX626 Build/HBXCNCU5801811241S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris X5 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6055B Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G8441 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.1; HUAWEI G610-U20 Build/HuaweiG610-U20) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; HUAWEI Y560-L01 Build/HUAWEIY560-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-A600FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI Y625-U21 Build/HUAWEIY625-U21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; VFD 900 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Nokia 2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CUBOT_NOTE_S Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-N9500 Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-J120M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; M3 Max Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; UMI_SUPER Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Nexus 5X Build/N4F26T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6045Y Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; D5303 Build/19.0.1.A.0.207) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Halo Plus Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 9008N Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; C1905 Build/15.1.C.2.8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; vivo 1725 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Plus Build/NPSS26.116-26-18) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; F5321 Build/34.3.A.0.238) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; XT1032 Build/KTU84P.M003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; NEM-L51 Build/HONORNEM-L51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; UMI_MAX Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; BLOOM Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; BV9000Pro-F Build/N4F26M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5010 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ZUK Z2131 Build/OPM6.171019.030.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1662 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E4 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TA-1033 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; vernee_M5 Build/vernee_M5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-N9200 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z017D Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; K1 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-N915G Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Thor Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3441.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MPI24.107-55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi 3 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense C1 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; MEO Smart A40 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG-SM-G920A Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo3 Plus Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; D5803 Build/23.4.A.1.232) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; BLU STUDIO X Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G935T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; S11 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; CHC-U01 Build/HuaweiCHC-U01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; W_K300 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Le X622 Build/HEXCNFN5902303291S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 791 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U22 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00UD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ASUS_Z012DA Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPSS26.118-19-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Archos 55b Platinum Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris X Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3456.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; LG-D337 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-J600FN Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-J120A Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X00BD Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HUAWEI G750-U10 Build/HuaweiG750-U10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1021) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ALP-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3529.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI GRA-UL10 Build/HUAWEIGRA-UL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3410.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Lenovo K50a40 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nexus 6P Build/OPM2.171026.006.C1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HUAWEI G6-L11 Build/HuaweiG6-L11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Swift 2 X Build/OPM5.171019.017) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HTC One M9 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 6055P Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Alpha Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1024 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; itel A12 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G610F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/MOB31E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Studio J1 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; vivo X9s Plus L Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense L678 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-AL00 Build/HUAWEIVKY-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X015D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Selecline-854599 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1039 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; LG-H650 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; MotoG3 Build/MPIS24.107-55-2-12) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; TA-1032 Build/O00623) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HIGHWAY STAR Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Orange Dive 71 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z1 Build/M4B30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.33 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HUAWEI LUA-L02 Build/HUAWEILUA-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NOS NEVA 80 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G800H Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z017D Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H810 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; LG-M400 Build/NRD90U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; SM-J500FN Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D802 Build/KOT49I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 7 plus Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ASUS_Z00YD Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z1 Build/M4B30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A700F Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HT16 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LG-D405 Build/KOT49I.A1396450705) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950W Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Mi Note 2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; SAMSUNG-SGH-I337 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; ASUS_Z00VD Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J320F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HTC One ME dual sim Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; U FEEL PRIME Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PLK-L01 Build/HONORPLK-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; JERRY MAX Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; AX570 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Neffos X1 Max Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M5 Build/M5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; XT1072 Build/MPBS24.65-34-5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X017DA Build/NGI77B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 4047F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; BBF100-6 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Vienna Build/F5B_GQ3028AF2_13E_MM_VIENNA_20160830) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; A0001 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; CUBOT X18 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G900H Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J327V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G313F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-A300H Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; SM-G900F Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Infinix X571 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G935A Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-M151 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; SAMSUNG-SM-G900A Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H870 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3493.4 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; SM-G3815 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A6003 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H815 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; GT-I9060C Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Z Play Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX1 Build/HUAWEIANE-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.41 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; E5603 Build/30.1.A.1.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T5_Lite Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G950F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; HT7 Pro Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S8 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris U Plus Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Le X527 Build/IMXOSOP5801607291S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; BLL-L23 Build/HUAWEIBLL-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5026D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; T5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; S60 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; Boston 4G Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Redmi 3X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZUK Z2121 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; MOVE_6_mini Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; X5 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; LAIQ Dubai Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Moto G (4) Build/MPJ24.139-50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG-SM-N910A Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; KOB-W09 Build/HUAWEIKOB-W09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510M Build/Alpha Centauri) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; EVA-L09 Build/HUAWEIEVA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Lenovo K920 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; D5803 Build/23.5.A.1.291) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; ASUS_X00DD Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3532.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Vodafone Smart ultra 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Vivo 5R Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Q7S55IN4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-A510M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; Redmi Note 3 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.40 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; U008 Pro Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Lenovo A6020a41 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Lenovo A806 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Energy Phone Pro HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 5009D Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; TOMMY2 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3112 Build/48.1.A.2.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; MotoE2(4G-LTE) Build/LXI22.50-50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Z2 Build/O11019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Le X620 Build/HEXCNFN5601405171S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X550 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; K00G Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FRD-L04 Build/HUAWEIFRD-L04) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G388F Build/LMY48B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-A310M Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2303 Build/26.1.A.3.111) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto G (5S) Build/NPPS26.102-49-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; T1700Q Build/NHG47K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M503 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HIGHWAY PURE Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; M5 Build/M5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-92-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Halo X Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G925I Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Aquaris U Lite Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S2 LITE Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 5085Y Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ONEPLUS A3000 Build/OPR6.170623.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Mi Note 2 Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3212 Build/42.0.A.4.153) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; MIX Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 6.1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Omega-K Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; VIE-L29 Build/HUAWEIVIE-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; C5-S508 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00UD Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ECHO_SMART_4G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LEX722 Build/WIXCNFN5902607031S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; OV-Vertis 5011 Expi Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Aquaris E5 HD Build/MOB30R) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8216 Build/51.1.A.2.213) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-A300FU Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Moto G Build/MOB31K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; WAS-LX1A Build/HUAWEIWAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ANDY_55EI Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Lenovo A5000 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HARRY Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00LDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; GT-N7000 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; 562 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G925F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; Lenovo P1c72 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPSS26.118-19-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Bluboo  D2 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; ASUS_Z00UD Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; W5 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; G3311 Build/43.0.A.5.79) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Ektra Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Energy Phone Pro HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; View Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; Vodafone 890N Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; X9076 Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Nokia 8 Sirocco Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; HIGHWAY Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; F3116 Build/33.3.A.1.97) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G9500 Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; LG-D955 Build/JDQ39B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STH100-1 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G360M Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; HTC 10 Build/OPR1.170623.027) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_X00ID Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; T07 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; LG-H500 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ZTE BLADE V0710 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.2.2; GT-S7275R Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LG-H915 Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G920F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ASUS_X007D Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; Aquaris M4.5 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-J327P Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; ASUS_Z01MDA Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; SM-G928V Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 6044D Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; 0PJA10 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Hisense C1 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto Z2 Play Build/NPSS26.118-19-11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; Aquaris E5 HD Build/LRX21M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; HTC One mini 2 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; VKY-L29 Build/HUAWEIVKY-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Archos Sense 50 DC Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; E2333 Build/26.1.B.3.109) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; NX506J Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; FEVER Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LON-L29 Build/HUAWEILON-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Infinix-X521 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; ENERGY E520 LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; JERRY2 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Andy 5EL2 LTE Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; PHICOMM ENERGY 3  Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.4; 2014819 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; ASUS_X00BD Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; SM-G350H Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; PULP FAB Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G9250 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.1.2; CINK FIVE Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Hisense F24 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; 6016D Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; TA-1004 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3473.0 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; AGM X2 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; C6903 Build/14.6.A.1.236) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Moto G 2014 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPP25.137-93-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G950F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; LG-US998 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; Aquaris E5 HD Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Colors P45 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J730F Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-J530YM Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; SM-G610Y Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; LGM-G600S Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Redmi 4 Prime Build/OPM2.171026.006.H1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; iris870 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S70 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; S8 Build/NMF26O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Infinix Zero 4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; ASUS_X008DB Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; HUAWEI NMO-L31 Build/HUAWEINMO-L31) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; itel S32 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; ZTE BLADE A512 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; H8116 Build/51.1.A.4.147) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; C1001 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-T815 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; HIGHWAY PURE Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; G3212 Build/48.1.A.2.35) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.2; Redmi Note 3 Build/NJH47F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-M153 Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; 6045Y Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Plus Build/NPNS25.137-93-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; A0001 Build/M4B30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; U16 Max Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; XT1635-02 Build/MCN24.104-35.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SUPER Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.2; PLK-L01 Build/HONORPLK-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Sunny2 Plus Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; LG-H840 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; SM-G920T Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; 2014813 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.14 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1.1; Coolpad E561_EU Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; STARADDICT 6 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; S6S5IN3G Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; AJM Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.4.2; ZP999 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; FRD-AL10 Build/HUAWEIFRD-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; SM-N950F Build/NMF26X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; S8_Pro Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; 5056X Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 4.3; SM-G3502T Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0; Energy 400 LTE Build/Energy_400_LTE_FR_6.0_HW1.0_V9_20170222) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; M653 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; Aquaris X Pro Build/OPM1.171019.019) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NMA26.42-157) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.76 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; T03 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.70 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.1; LG-X190 Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.85 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0.1; ALE-L23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.1.1; REVVLPLUS C3701A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.1.0; RMX1805) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 6.0.1; NX551J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.105 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36
Mozilla/5.0 (Linux; Android 7.0; RNE-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Mobile Safari/537.36
`
	userAgentIos = `Mozilla/5.0 (iPhone; CPU iPhone OS 9_0 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 FingerBrowser/1.4
Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.28.0 main/1.0 baiduboxapp/12.9.0.14 (Baidu; P2 14.3) NABar/1.0 themeUA=Theme/default webCore=0x149f6ca80
Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.28.0 main/1.0 baiduboxapp/12.9.0.14 (Baidu; P2 14.3) NABar/1.0 themeUA=Theme/default webCore=0x127cf9d90
Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.18(0x1700122d) NetType/4G Language/zh_CN
Mozilla/5.0 (iPhone; CPU iPhone OS 14_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/18A393 UCBrowser/13.1.3.1385 Mobile  AliApp(TUnionSDK/0.1.20.3)
Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.27.0 main/1.0 baiduboxapp/12.8.5.17 (Baidu; P2 14.0.1) NABar/1.0 webCore=0x11b16b3c0
Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.28.0 main/1.0 baiduboxapp/12.9.0.14 (Baidu; P2 13.6) NABar/1.0 themeUA=Theme/default webCore=0x1468402b0
Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 13_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_9 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/16H5 UCBrowser/13.1.5.1393 Mobile  AliApp(TUnionSDK/0.1.20.3)
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_9 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_8 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/16G201 UCBrowser/13.1.4.1377 Mobile  AliApp(TUnionSDK/0.1.20.3)
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/16F203 UCBrowser/12.9.5.1306 Mobile  AliApp(TUnionSDK/0.1.20.3)
Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/16F203 UCBrowser/12.9.5.1306 Mobile  AliApp(TUnionSDK/0.1.20.3)
Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.2(0x1800022c) NetType/WIFI Language/zh_CN
Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 AliApp(DingTalk/6.0.5) com.laiwang.DingTalk/14539441 Channel/201200 language/zh-Hans-CN UT4Aplus/0.0.6 WK
Mozilla/5.0 (iPhone; CPU iPhone OS 12_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/16C101 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 11_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 11_0_3 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Mobile/15A432 SP-engine/2.25.0 main/1.0 baiduboxapp/12.3.0.10 (Baidu; P2 11.0.3) NABar/1.0 webCore=0x12a514dc0
Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.0 Mobile/14G60 Safari/602.1
Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1
Mozilla/5.0 (iPhone; CPU iPhone OS 10_1_1 like Mac OS X) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0 Mobile/14B100 Safari/602.1
Mozilla/5.0 (Linux; Android 10; 16s Build/QKQ1.191222.002) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.3359.203 mobile Safari/537.36baiduboxapp/3.2.5.10 SearchCraft/2.8.2 (Baidu; P1 10)
Mozilla/5.0 (iPhone; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.18(0x1700122d) NetType/WIFI Language/zh_CN
Mozilla/5.0 (iPhone; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.2(0x18000223) NetType/4G Language/zh_CN
Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X; zh-cn) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/18D52 Quark/4.6.5.1146 Mobile
Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/18D52 UCBrowser/13.2.5.1434 Mobile  AliApp(TUnionSDK/0.1.20.4)
Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1
Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 QQ/8.5.5.648 V1_IPH_SQ_8.5.5_1_APP_A P`
)

var (
	_userAgentPc      = make([]string, 0)
	_userAgentAndroid = make([]string, 0)
	_userAgentIos     = make([]string, 0)
	_getS             = sync.Mutex{}
)

func init() {
	_userAgentPc = getLines(userAgentPc)
	_userAgentAndroid = getLines(userAgentAndroid)
	_userAgentIos = getLines(userAgentIos)
}

func getLines(s string) []string {
	s = gstr.Replace(s, "\r", "")
	return gstr.SplitAndTrim(s, "\n")
}

func GetRandAgentPc() string {
	_getS.Lock()
	defer _getS.Unlock()
	return _userAgentPc[grand.N(0, len(_userAgentPc)-1)]
}

func GetRandAgentAndroid() string {
	_getS.Lock()
	defer _getS.Unlock()
	return _userAgentAndroid[grand.N(0, len(_userAgentAndroid)-1)]
}

func GetRandAgentIos() string {
	_getS.Lock()
	defer _getS.Unlock()
	return _userAgentIos[grand.N(0, len(_userAgentIos)-1)]
}
