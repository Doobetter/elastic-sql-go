> 
> todo: 
> where 条件解析+生成 query 需要两次 自顶向下解析 
> 
优化方向：

### 1.解析过程优化
自顶向下 visitLogicalExpr只负责解析顶层 顶层以下交给另一个函数，统一生成一个PathTree
分情况
1. 单个compare expr
2. 复杂logic-expr
3. 括号 包裹的 compare expr
4. 括号 包裹的复杂 logic-expr

### 2.生成query优化
在解析语法的时候生成、组装query