全字庫 playground
===

解決中文罕見字問題

說明
---

全字庫的正宋體字體檔案總共分為三個檔案，分別如下：

| 檔案名稱               | 字型名稱      | 中文字型名稱       | Unicode Code Point Range | 中文規格用語[^1]  |
| ---------------------- | ------------- | ------------------ | ------------------------ | ----------------- |
| TW-Sung-98_1.ttf       | TW-Sung       | 全字庫正宋體       | U+0000  - U+FFFF         | BMP字面  / 字面0  |
| TW-Sung-Ext-B-98_1.ttf | TW-Sung-Ext-B | 全字庫正宋體 Ext-B | U+20000 - U+2FFFF        | 第2字面  / 字面2  |
| TW-Sung-Plus-98_1.ttf  | TW-Sung-Plus  | 全字庫正宋體 Plus  | U+F0000 - U+FFFFF        | 第15字面 / 字面15 |

[^1]: [維基百科 - Unicode字元平面對映](https://zh.wikipedia.org/wiki/Unicode字符平面映射)

因此在使用時，須判斷當下的中文字是在哪個 Unicode Code Point Range 來使用對應的字體檔

範例
---

### Golang

[main.go](./go/main.go)

![screenshots](https://i.imgur.com/R1EBhMd.png)

### CSS

因為 CSS 的 font-family 可以設置多個字體、且有 fallback 特性，因此可以這樣：

```css
html, body {
  font-family: 'TW-Sung', 'TW-Sung-Ext-B', 'TW-Sung-Plus';
}
```

相關資源
---

- [全字庫官網](https://www.cns11643.gov.tw/)
- [政府資料開放平台：全字庫字型下載](https://data.gov.tw/dataset/5961)
