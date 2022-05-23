## Mein Projekt Installieren
````
sudo -i 
adduser ahrensburg
usermod -aG sudo ahrensburg
exit
ssh ahrensburg@ahrensburg.city
sudo -u postgres -i
createuser ahrensburg # answer yes for superuser (although this isn't strictly necessary)
createdb -E UTF8 -O ahrensburg ahrensburg
psql
\c ahrensburg
CREATE EXTENSION postgis;
ALTER TABLE geometry_columns OWNER TO ahrensburg;
ALTER TABLE spatial_ref_sys OWNER TO ahrensburg;
git clone https://github.com/thorstenkloehn/ahrensburg.city.git



