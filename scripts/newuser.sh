### add user to passwd via argv

#echo "creating user account..."
#sudo useradd -m $1
#printf $1:$2 | sudo chpasswd
#
#echo "user account successfully created"


# add user iteractively

echo "Enter new account name: "
read name

echo "Enter new account password: "
read -s password

sudo useradd -m $name
sudo chpasswd <<< $name:$password
echo "Account successfully created!"
