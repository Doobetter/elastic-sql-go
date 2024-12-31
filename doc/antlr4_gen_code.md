

java -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" \
   org.antlr.v4.Tool -visitor -Dlanguage=Go \
   -o /Users/liucheng/workspace_go/elastic-sql-go/src/parser \
   /Users/liucheng/workspace_go/elastic-sql-go/src/parser/ElasticSQL.g4

项目根目录下
java -cp ./src/parser/antlr-4.13.2-complete.jar org.antlr.v4.Tool \
    -Dlanguage=Go -visitor -package parser  \
   ./src/parser/ElasticSQL.g4

    -o ./src/parsing \