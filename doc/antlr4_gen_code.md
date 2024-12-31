
项目根目录下
java -cp ./src/parser/antlr-4.13.2-complete.jar org.antlr.v4.Tool \
    -Dlanguage=Go -visitor -package parser  \
   ./src/parser/ElasticSQL.g4