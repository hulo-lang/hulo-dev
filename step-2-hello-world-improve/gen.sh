#!/bin/bash

cd grammar && java -jar ../../antlr.jar -Dlanguage=Go -o ../generated -package generated *.g4