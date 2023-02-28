echo
echo Toggle size to context width inside the terminal, press: "alt+z"
echo
echo To check answers go to: https://github.com/01-edu/public/blob/master/subjects/lem-in/examples/README.md
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo example00
echo
go run . examples/example00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example01
echo
go run . examples/example01.txt
echo
read -n1 -r -p "Press any key to continue..." key 
echo
echo --------------------------------------
echo
echo example02
echo
go run . examples/example02.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example03
echo
go run . examples/example03.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example04
echo
go run . examples/example04.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example05
echo
go run . examples/example05.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example06
echo
go run . examples/example06.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo example07
echo
go run . examples/example07.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo badexample00
echo
go run . examples/badexample00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo badexample01
echo
go run . examples/badexample01.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo myexample00
echo
go run . examples/myexample00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
cat ascii.txt
echo
echo --------------------------------------
echo Thank you for the audit! Have a nice day!
echo --------------------------------------

exit 0
