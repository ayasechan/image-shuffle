

# 混淆图片

将图片 8x8 分块后
随机打乱

不能被整除的图片边缘会被舍去！

如果要保留边缘 需要先对图像做 resize
确保宽高都能被 8 整除


# 效果 



原图(1060x1500)

![origin](origin.jpg)

混淆之后(1056x1496)

![encrypt](encrypt.jpg)

反混淆(1056x1496)

![decrypt](decrypt.jpg)

# usage

see [here](block/block_test.go)


# ref

https://github.com/xfgryujk/weibo-img-crypto

https://ybzjdsd.gitee.io/tphx/