go run . examples/example00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
go run . examples/example01.txt
echo
read -n1 -r -p "Press any key to continue..." key 
echo
echo --------------------------------------
echo
go run . examples/example02.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
go run . examples/example03.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
go run . examples/example04.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
go run . examples/example05.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
go run . examples/example06.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
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
read -n1 -r -p "Press any key to continue..." key
echo
go run . examples/badexample01.txt
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