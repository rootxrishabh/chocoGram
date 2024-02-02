sudo apt-get update

sudo apt-get install --assume-yes gnupg apt-utils gnupg2 curl gcc

sudo apt-get update

sudo apt-get install --assume-yes mysql-server-8.0

sudo usermod -d /var/lib/mysql mysql

sudo mkdir -p /var/run/mysqld

sudo chown -R mysql: /var/run/mysqld

sudo mkdir -p /data/db
sudo chown -R $USER:$USER /data/db

echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

# Start MySql
sudo mysqld_safe --skip-grant-tables &
sleep 2
sudo mysql -e "FLUSH PRIVILEGES"
sudo mysql -e "ALTER USER 'root'@'localhost' IDENTIFIED BY 'admin'"
sudo service mysql start
sleep 2
sudo service mysql restart
sleep 2
sudo mysql -u root --password=admin -e "create database socialNetwork"