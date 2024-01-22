printf "The amount of code written in go: "
find . -name "*.go" | xargs cat | wc -l

printf "The amount of code written in cpp: "
find . -name "*.cpp" | xargs cat | wc -l

printf "The amount of code all of the above: "
find . -name "*.go" -o -name "*.cpp" | xargs cat | wc -l

echo ""
read -n 1 -p "请按任意键继续..."