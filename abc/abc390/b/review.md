<!-- <script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/3.2.2/es5/tex-mml-chtml.min.js">
</script>
<script type="text/x-mathjax-config">
 MathJax.Hub.Config({
 tex2jax: {
 inlineMath: [['$', '$'] ],
 displayMath: [ ['$$','$$'], ["\\[","\\]"] ]
 }
 });
</script> -->
# ABC390 B問題
## URL
https://atcoder.jp/contests/abc390/tasks/abc390_b
## 解説
### 考え方
公比の比較で考えると、float型の誤差によって正答できない。以下のように式を変形する。
```math
\frac{a_{i+1}}{a_i}=\frac{ai}{a_{i-1}}
```
```math
a_i^2=a_{i+1}a_{i-1}
```
こうすると`int64`で解ける。

### 使用したアルゴリズムやデータ構造
- 式変形
- 配列
## メモ
floatを使う時は誤差のことを考える。
