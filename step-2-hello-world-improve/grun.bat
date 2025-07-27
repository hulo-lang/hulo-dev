@echo off

cd grammar
echo Compiling Java files...
javac -cp "../../antlr.jar" .antlr/*.java
java -Xmx500M -cp "../../antlr.jar;.antlr;%CLASSPATH%" org.antlr.v4.gui.TestRig Hello program -gui