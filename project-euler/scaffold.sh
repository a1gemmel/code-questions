
PROBLEM=$1

DIR="problem-${PROBLEM}"

mkdir "${DIR}"

echo "
package main

func main() {

}
" > "${DIR}/main.go"
