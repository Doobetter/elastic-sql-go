> antlr4-go 对visitor模式支持有限
> 因为golang的多态限制
> 

# 1. 问题

无法Visit(some-ctx)这样，当前实现是空的

# 2.实现方法