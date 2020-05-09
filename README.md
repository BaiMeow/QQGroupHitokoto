## 一言onQQ
这是一个酷q插件，在Q群实现了一言

#### 特性

- [x] 支持httpAPI源
- [x] 获取随机一言
- [x] 获取某一分类的随机一言，并分群
- [x] 支持本地json文件源
- [x] 支持sqlite源
- [ ] 支持群员添加一言条目（针对sqlite源以及json源
  
#### 配置文件示例

```json
{
	"Sources": [
		{
			"Name": "动画",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=a"
		},
		{
			"Name": "漫画",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=b"
		},
		{
			"Name": "游戏",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=c"
		},
		{
			"Name": "文学",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=d"
		},
		{
			"Name": "原创",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=e"
		},
		{
			"Name": "来自网络",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=f"
		},
		{
			"Name": "其他",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=g"
		},
		{
			"Name": "影视",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=h"
		},
		{
			"Name": "诗词",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=i"
		},
		{
			"Name": "网易云",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=j"
		},
		{
			"Name": "哲学",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=k"
		},
		{
			"Name": "抖机灵",
			"SourceType": "HTTP",
			"Source": "https://v1.hitokoto.cn/?c=l"
		},
		{
			"Name": "Msc",
			"SourceType": "SQLITE",
			"Source": "test.db"
		},
		{
			"Name": "Msc2",
			"SourceType": "JSON",
			"Source": "msc.json"
		}
	]
}
```

#### 一言源说明

##### json

```json
[
	{
		"hitokoto":"我们服就不会蹦",
		"from":"柏喵"
	}
]
```
##### sqlite
```
sqlite> .schema hitokoto
CREATE TABLE IF NOT EXISTS "hitokoto" (
  "hitokoto" text,
  "from" text
);
```