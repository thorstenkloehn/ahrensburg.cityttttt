## Ubuntu 20.4

[Ubuntu](index) - [Windows](Windows) - [Lernen](Lernen) 

### Installieren 

```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install git-core
apt-get install  gnupg2
sudo apt install python3 python3-dev curl python-is-python3 -y
sudo apt install python3-pip -y
```
## Quellangabe 
Dieser Text basiert auf dem Artikel [Ubuntu 20.4](https://switch2osm.org/serving-tiles/manually-building-a-tile-server-20-04-lts/)  aus der OpenStreetMap and contributor und steht unter der Lizenz [CC BY-SA](http://creativecommons.org/licenses/by-sa/2.0/)
und Maschinelle Übersetzung
#### OSM Installieren

Auf dieser Seite wird beschrieben, wie Sie alle erforderliche Software installieren, einrichten und konfigurieren, um Ihren eigenen Kachelserver zu betreiben. Diese Schritt-für-Schritt-Anleitung wurde für Ubuntu Linux 20.04 LTS (Focal Fossa) geschrieben und im Mai 2020 getestet.

##### Software Installation

Der OSM-Tileserver-Stack ist eine Sammlung von Programmen und Bibliotheken, die zusammenarbeiten, um einen Tileserver zu erstellen. Wie so oft bei OpenStreetMap gibt es viele Wege, um dieses Ziel zu erreichen, und für fast alle Komponenten gibt es Alternativen, die verschiedene spezifische Vor- und Nachteile haben. Dieses Tutorial beschreibt die Standardversion, die derjenigen ähnelt, die auf den Hauptkachelservern von OpenStreetMap.org verwendet wird.

Es besteht aus 5 Hauptkomponenten: mod_tile, renderd, mapnik, osm2pgsql und einer postgresql/postgis-Datenbank. Mod_tile ist ein Apache-Modul, das zwischengespeicherte Kacheln bereitstellt und entscheidet, welche Kacheln neu gerendert werden müssen - entweder weil sie noch nicht zwischengespeichert wurden oder weil sie veraltet sind. Renderd bietet ein Prioritätswarteschlangensystem für verschiedene Arten von Anfragen, um die Last von Renderanfragen zu verwalten und zu glätten. Mapnik ist die Softwarebibliothek, die das eigentliche Rendern durchführt und von renderd verwendet wird.

Beachten Sie, dass diese Anweisungen für einen neu installierten Ubuntu 20.04-Server geschrieben und getestet wurden. Wenn Sie bereits andere Versionen einiger Software installiert haben (vielleicht haben Sie ein Upgrade von einer früheren Ubuntu-Version durchgeführt oder einige PPAs zum Laden eingerichtet), müssen Sie möglicherweise einige Anpassungen vornehmen.

Diese Anleitung geht davon aus, dass Sie alles von einem Nicht-Root-Benutzer über „sudo“ ausführen. Der unten standardmäßig verwendete Nicht-Root-Benutzername ist „thorsten“ – Sie können diesen lokal erstellen, wenn Sie möchten, oder Skripte bearbeiten, um auf einen anderen Benutzernamen zu verweisen, wenn Sie möchten. Wenn Sie den „thorsten“-Benutzer erstellen, müssen Sie ihn der Gruppe von Benutzern hinzufügen, die sudo zum Rooten verwenden können. Von Ihrem normalen Nicht-Root-Benutzerkonto:

```
sudo -i 
adduser thorsten 
usermod -aG sudo thorsten
exit
```

Um diese Komponenten zu bauen, müssen zunächst verschiedene Abhängigkeiten installiert werden:

```
ssh thorsten@deine_Hompage
sudo apt install libboost-all-dev git tar unzip wget bzip2 build-essential autoconf libtool libxml2-dev libgeos-dev libgeos++-dev libpq-dev libbz2-dev libproj-dev munin-node munin protobuf-c-compiler libfreetype6-dev libtiff5-dev libicu-dev libgdal-dev libcairo2-dev libcairomm-1.0-dev apache2 apache2-dev libagg-dev liblua5.2-dev ttf-unifont lua5.1 liblua5.1-0-dev

```

Sag Ja zur Installation. Das wird eine Weile dauern, also geh und trink eine Tasse Tee. Diese Liste enthält verschiedene Dienstprogramme und Bibliotheken, den Apache-Webserver und „carto“, das verwendet wird, um Carto-CSS-Stylesheets in etwas zu konvertieren, das „mapnik“ der Karten-Renderer verstehen kann. Wenn dies abgeschlossen ist, installieren Sie den zweiten Satz von Voraussetzungen:

##### Installation von postgresql / postgis

Unter Ubuntu gibt es vorgefertigte Versionen von postgis und postgresql, sodass diese einfach über den Ubuntu-Paketmanager installiert werden können.

```

sudo apt install postgresql postgresql-contrib postgis postgresql-12-postgis-3 postgresql-12-postgis-3-scripts


```

Hier ist „postgresql“ die Datenbank, in der wir Kartendaten speichern werden, und „postgis“ fügt ihr zusätzliche grafische Unterstützung hinzu. Sagen Sie erneut Ja zur Installation.

Jetzt müssen Sie eine Postgis-Datenbank erstellen. Die Standardeinstellungen verschiedener Programme gehen davon aus, dass die Datenbank gis heißt, und wir werden in diesem Tutorial dieselbe Konvention verwenden, obwohl dies nicht notwendig ist. Ersetzen Sie thorsten durch Ihren Benutzernamen, wo dieser unten verwendet wird. Dies sollte der Benutzername sein, der Karten mit Mapnik rendert.

```
sudo -u postgres -i
createuser thorsten # answer yes for superuser (although this isn't strictly necessary)
createdb -E UTF8 -O thorsten gis

```

Während Sie immer noch als „postgres“-Benutzer arbeiten, richten Sie PostGIS in der PostgreSQL-Datenbank ein (ersetzen Sie wieder Ihren Benutzernamen für thorsten unten):

```
psql
```

(dadurch gelangen Sie zu einer „postgres=#“-Eingabeaufforderung)
```
\c gis
```
(Es wird geantwortet: „Sie sind jetzt mit der Datenbank „gis“ als Benutzer „postgres“ verbunden“.)
```

CREATE EXTENSION postgis;

```
(es wird CREATE EXTENSION antworten)

```
CREATE EXTENSION hstore;

```

(es wird CREATE EXTENSION antworten)

```
ALTER TABLE geometry_columns OWNER TO thorsten;


```

(es wird ALTER TABLE antworten)

```
ALTER TABLE spatial_ref_sys OWNER TO thorsten;

```
(es wird ALTER TABLE antworten)

```

\q
```

(Es wird psql beenden und zu einer normalen Linux-Eingabeaufforderung zurückkehren.)

```
exit

```
(um wieder der Benutzer zu sein, der wir waren, bevor wir oben „sudo -u postgres -i“ gemacht haben)

Wenn Sie noch keinen erstellt haben, erstellen Sie auch für diesen Benutzer einen Unix-Benutzer und wählen Sie ein Passwort, wenn Sie dazu aufgefordert werden:

```
sudo useradd -m thorsten
sudo passwd thorsten

```

Ersetzen Sie oben erneut „thorsten“ durch den von Ihnen gewählten Nicht-Root-Benutzernamen.

##### Installing osm2pgsql
Als nächstes installieren wir Mapnik. Wir verwenden die Standardversion in Ubuntu 20.04:
```

sudo apt install osm2pgsql

```
##### Mapnik

```
sudo apt install autoconf apache2-dev libtool libxml2-dev libbz2-dev libgeos-dev libgeos++-dev libproj-dev gdal-bin libmapnik-dev mapnik-utils python3-mapnik python3-psycopg2 python3-yaml

```

Wir überprüfen, ob Mapnik korrekt installiert wurde:

``` 
python3
>>> import mapnik
>>>
```

Wenn Python mit dem zweiten Chevron-Prompt »> und ohne Fehler antwortet, wurde die Mapnik-Bibliothek von Python gefunden. Herzliche Glückwünsche! Sie können Python mit diesem Befehl verlassen:

```
>>> quit()

```
##### Installieren Sie mod_tile und rendern Sie

Als nächstes installieren wir mod_tile und rendern. „mod_tile“ ist ein Apache-Modul, das Anfragen für Kacheln verarbeitet; „renderd“ ist ein Daemon, der Kacheln tatsächlich rendert, wenn „mod_tile“ sie anfordert. Wir verwenden den „switch2osm“-Zweig von https://github.com/SomeoneElseOSM/mod_tile, der selbst von https://github.com/openstreetmap/mod_tile gegabelt, aber so modifiziert wurde, dass er Ubuntu 20.04 unterstützt, und mit ein paar andere Änderungen, um auf einem Standard-Ubuntu-Server statt auf einem der Rendering-Server von OSM zu funktionieren.

###### Kompilieren Sie den mod_tile-Quellcode:

```
mkdir ~/src
cd ~/src
git clone -b switch2osm https://github.com/SomeoneElseOSM/mod_tile.git
cd mod_tile
./autogen.sh

```

(Das sollte mit „autoreconf: Leaving directory '.'“ enden.)
```
./configure

```

(das sollte mit „config.status: libtool-Befehle ausführen“ enden)

```
make
```

Beachten Sie, dass einige „besorgniserregende“ Nachrichten hier auf dem Bildschirm nach oben scrollen. Es sollte jedoch mit „make[1]: Leaving directory '/home/ahrensburg/src/mod_tile'“ enden.

```
sudo make install
```
(das sollte mit „make[1]: Leaving directory '/home/thorsten/src/mod_tile' enden“)

```
sudo make install-mod_tile

```

(das sollte mit „chmod 644 /usr/lib/apache2/modules/mod_tile.so“ enden)

```
sudo ldconfig
```

###### Stylesheet-Konfiguration
Nachdem die gesamte erforderliche Software installiert ist, müssen Sie ein Stylesheet herunterladen und konfigurieren.

Der Stil, den wir hier verwenden, ist derjenige, der von der „Standard“-Karte auf der Website openstreetmap.org verwendet wird. Es wurde ausgewählt, weil es gut dokumentiert ist und überall auf der Welt funktionieren sollte (einschließlich an Orten mit nicht-lateinischen Ortsnamen). Es gibt jedoch ein paar Nachteile - es ist ein Kompromiss, der darauf ausgelegt ist, global zu funktionieren, und es ist ziemlich kompliziert zu verstehen und zu ändern, falls Sie dies tun müssen.

Die Heimat von „OpenStreetMap Carto“ im Web ist https://github.com/gravitystorm/openstreetmap-carto/ und es gibt eine eigene Installationsanleitung unter https://github.com/gravitystorm/openstreetmap-carto/blob/master /INSTALL.md, obwohl wir hier alles behandeln, was getan werden muss.

Hier gehen wir davon aus, dass wir die Stylesheet-Details in einem Verzeichnis unterhalb von „src“ unterhalb des Home-Verzeichnisses des „thorsten“-Benutzers (oder eines anderen von Ihnen verwendeten) speichern.

```
cd ~/src
git clone https://github.com/gravitystorm/openstreetmap-carto
cd openstreetmap-carto

```
Als nächstes installieren wir eine passende Version des „carto“-Compilers.

```
sudo apt install npm
sudo npm install -g carto
carto -v
```

Das sollte mit einer Zahl antworten, die mindestens so hoch ist wie:

```

1.2.0

```
Dann konvertieren wir das Carto-Projekt in etwas, das Mapnik verstehen kann:

````
carto project.mml > mapnik.xml

````
Sie haben jetzt ein XML-Stylesheet von Mapnik unter /home/thorsten/src/openstreetmap-carto/mapnik.xml .

##### Zunächst laden wir ahrensburg Karte hinunterladen

```

mkdir ~/data
cd ~/data
git clone https://github.com/thorstenkloehn/ahrensburg_karte.git .


```

Der folgende Befehl fügt die zuvor heruntergeladenen OpenStreetMap-Daten in die Datenbank ein. Dieser Schritt ist sehr I/O-intensiv auf der Festplatte; Das Importieren des gesamten Planeten kann je nach Hardware viele Stunden, Tage oder Wochen dauern. Bei kleineren Extrakten ist die Importzeit entsprechend viel kürzer, und Sie müssen möglicherweise mit verschiedenen -C-Werten experimentieren, um sie in den verfügbaren Speicher Ihres Computers zu integrieren

```

osm2pgsql -d gis --create --slim  -G --hstore --tag-transform-script ~/src/openstreetmap-carto/openstreetmap-carto.lua -C 2500 --number-processes 1 -S ~/src/openstreetmap-carto/openstreetmap-carto.style ~/data/ahrensburg.osm.pbf

```

###### Indizes erstellen
Seit Version v5.3.0 müssen nun einige zusätzliche Indizes manuell angewendet werden .

```

cd ~/src/openstreetmap-carto/
psql -d gis -f indexes.sql

```
Es sollte 14 Mal mit „CREATE INDEX“ antworten.

###### Shapefile-Download
Obwohl die meisten Daten, die zum Erstellen der Karte verwendet werden, direkt aus der OpenStreetMap-Datendatei stammen, die Sie oben heruntergeladen haben, werden noch einige Shapefiles für Dinge wie Ländergrenzen mit niedrigem Zoom benötigt. Um diese herunterzuladen und zu indizieren:
```
pip3 install requests
cd ~/src/openstreetmap-carto/
scripts/get-external-data.py
```

Dieser Vorgang beinhaltet einen beträchtlichen Download und kann einige Zeit dauern - es wird nicht viel auf dem Bildschirm angezeigt, wenn er ausgeführt wird. Es wird tatsächlich ein „data“-Verzeichnis unterhalb von „openstreetmap-carto“ auffüllen.

###### Schriftarten

Die Namen für Orte auf der ganzen Welt werden nicht immer mit lateinischen Buchstaben (das bekannte westliche Alphabet az) geschrieben. Gehen Sie wie folgt vor, um die erforderlichen Schriftarten zu installieren:

```
sudo apt install fonts-noto-cjk fonts-noto-hinted fonts-noto-unhinted ttf-unifont
```

Die eigenen Installationsanweisungen von OpenSteetMap Carto schlagen auch vor, „Noto Emoji Regular“ von der Quelle zu installieren. Das wird offenbar für die Emojis in einem amerikanischen Ladennamen benötigt. Alle anderen wahrscheinlich benötigten internationalen Schriftarten (einschließlich der häufig nicht unterstützten) sind in der gerade installierten Liste enthalten.

## Einrichten Ihres Webservers

### Rendern konfigurieren
```
sudo nano /usr/local/etc/renderd.conf
```

### Apache konfigurieren

```
sudo mkdir /var/lib/mod_tile
sudo chown thorsten /var/lib/mod_tile

sudo mkdir /var/run/renderd
sudo chown thorsten /var/run/renderd

```
Wir müssen Apache nun über „mod_tile“ informieren, also mit nano (oder einem anderen Editor):
```
sudo nano /etc/apache2/conf-available/mod_tile.conf
```

Fügen Sie dieser Datei die folgende Zeile hinzu:

```
LoadModule tile_module /usr/lib/apache2/modules/mod_tile.so

```

und speichern Sie es und führen Sie dann aus:

```
sudo a2enconf mod_tile

```
Das bedeutet, dass Sie „service apache2 reload“ ausführen müssen, um die neue Konfiguration zu aktivieren; das machen wir noch nicht.

Wir müssen Apache jetzt über „renderd“ informieren. Mit nano (oder einem anderen Editor):
```
sudo nano /etc/apache2/sites-available/000-default.conf
```
Fügen Sie zwischen den Zeilen „ServerAdmin“ und „DocumentRoot“ Folgendes hinzu:

```
LoadTileConfigFile /usr/local/etc/renderd.conf
ModTileRenderdSocketName /var/run/renderd/renderd.sock
# Timeout before giving up for a tile to be rendered
ModTileRequestTimeout 0
# Timeout before giving up for a tile to be rendered that is otherwise missing
ModTileMissingRequestTimeout 30
```

Und Apache zweimal neu laden:
```
sudo service apache2 reload
sudo service apache2 reload

```

(Ich vermute, dass es zweimal gemacht werden muss, weil Apache „verwirrt“ wird, wenn es beim Ausführen neu konfiguriert wird.)

Wenn Sie mit einem Webbrowser auf: http://yourserveripaddress/index.html zeigen, sollten Sie Ubuntu / Apaches „It works!“ erhalten. Seite.

(Wenn Sie nicht wissen, welche IP-Adresse ihm zugewiesen wurde, können Sie wahrscheinlich „ifconfig“ verwenden, um es herauszufinden - wenn die Netzwerkkonfiguration nicht zu kompliziert ist, ist es wahrscheinlich die „inet-Adresse“, die nicht „127.0. 0,1"). Wenn Sie einen Server bei einem Hosting-Provider verwenden, ist es wahrscheinlich, dass sich die interne Adresse Ihres Servers von der Ihnen zugewiesenen externen Adresse unterscheidet, aber diese externe IP-Adresse wurde Ihnen bereits gesendet und wird es wahrscheinlich derjenige sein, auf dem Sie gerade auf den Server zugreifen.

Beachten Sie, dass dies nur die Seite „http“ (Port 80) ist – Sie müssen etwas mehr Apache konfigurieren, wenn Sie https aktivieren möchten, aber das würde den Rahmen dieser Anleitung sprengen. Wenn Sie jedoch „Let's Encrypt“ verwenden, um Zertifikate auszustellen, kann der Einrichtungsprozess auch die Apache HTTPS-Site konfigurieren.

### Render zum ersten Mal ausführen
Als Nächstes führen wir render aus, um zu versuchen, einige Kacheln zu rendern. Zunächst führen wir es im Vordergrund aus, damit wir alle auftretenden Fehler sehen können:
```
renderd -f -c /usr/local/etc/renderd.conf
```

Möglicherweise sehen Sie hier einige Warnungen - machen Sie sich vorerst keine Sorgen darüber. Sie sollten keine Fehler erhalten. Wenn Sie dies tun, speichern Sie die vollständige Ausgabe in einem Pastebin und stellen Sie eine Frage zu dem Problem irgendwo wie help.openstreetmap.org (Link zum Pastebin - geben Sie nicht den gesamten Text in die Frage ein).

Zeigen Sie mit einem Webbrowser auf: http://yourserveripaddress/hot/0/0/0.png

Sie sollten eine Weltkarte in Ihrem Browser und weitere Debug-Funktionen in der Befehlszeile sehen, einschließlich „DEBUG: START TILE“ und „DEBUG: DONE TILE“. Ignorieren Sie die Meldung „DEBUG: Failed to read cmd on fd“ – es handelt sich nicht um einen Fehler. Wenn Sie keine Kachel erhalten und wieder andere Fehler erhalten, speichern Sie die vollständige Ausgabe in einem Pastebin und stellen Sie eine Frage zu dem Problem irgendwo wie help.openstreetmap.org.

Wenn das alles funktioniert, drücken Sie Strg-c, um den Vordergrund-Rendering-Prozess zu stoppen.

```
nano ~/src/mod_tile/debian/renderd.init
sudo cp ~/src/mod_tile/debian/renderd.init /etc/init.d/renderd
sudo chmod u+x /etc/init.d/renderd
sudo cp ~/src/mod_tile/debian/renderd.service /lib/systemd/system/
```

Die Datei „renderd.service“ ist eine „systemd“-Dienstdatei. Die hier verwendete Version ruft nur alte Init-Befehle auf. Um zu testen, ob der Startbefehl funktioniert:

```
sudo /etc/init.d/renderd start
```

(Das sollte mit „[ ok ] Starting render (via systemctl): renderd.service“ antworten.)

Damit es jedes Mal automatisch startet:

```
sudo systemctl enable renderd

```

## Quellangabe
Dieser Text basiert auf dem Artikel [Ubuntu 20.4](https://switch2osm.org/serving-tiles/manually-building-a-tile-server-20-04-lts/)  aus der OpenStreetMap and contributor und steht unter der Lizenz [CC BY-SA](http://creativecommons.org/licenses/by-sa/2.0/)
und Maschinelle Übersetzung

## Tomcat 9 installieren

sudo apt-get install tomcat9

## Certbot Installieren
### Quellangabe
Dieser Text basiert auf dem Artikel [certbot instructions](https://certbot.eff.org/instructions?ws=other&os=ubuntufocal) aus der [Cerbot](https://certbot.eff.org/) und steht unter der Lizenz[[Creative Commons](https://www.eff.org/copyright)
und Maschinelle Übersetzung

### SSH in den Server
SSH in den Server, auf dem Ihre HTTP-Website als Benutzer mit sudo-Berechtigungen ausgeführt wird.

### snapd installieren
Sie müssen snapd installieren und sicherstellen, dass Sie alle Anweisungen befolgen, um die klassische Snap-Unterstützung zu aktivieren.

```
sudo apt install snapd
```

#### Stellen Sie sicher, dass Ihre Version von snapd auf dem neuesten Stand ist

Führen Sie die folgenden Anweisungen in der Befehlszeile auf dem Computer aus, um sicherzustellen, dass Sie über die neueste Version von snapd.
```
sudo snap install core; sudo snap refresh core
```

#### Certbot installieren
Führen Sie diesen Befehl in der Befehlszeile auf dem Computer aus, um Certbot zu installieren.

```
sudo snap install --classic certbot
```

##### Bereiten Sie den Certbot-Befehl vor
Führen Sie die folgende Anweisung in der Befehlszeile auf dem Computer aus, um sicherzustellen, dass der certbotBefehl ausgeführt werden kann.

```

sudo ln -s /snap/bin/certbot /usr/bin/certbot

```

#### Cerbort ausführen

```

sudo certbot certonly --standalone

```

### Quellangabe
Dieser Text basiert auf dem Artikel [certbot instructions](https://www.eff.org/copyright)
und Maschinelle Übersetzung

## nginx

```

sudo apt-get install nginx -y

```

## ahrensburg.city hinunderladen

```

git clone https://github.com/thorstenkloehn/ahrensburg.city.git /Server
sudo apt-get install unzip
unzip /Server/externe_daten/geoserver/geoserver-2.20.4-war.zip -d /var/lib/tomcat9/webapps

```

## Nginx config Datei

```

sudo cp -u /Server/Server_Einstellung/start.conf /etc/nginx/conf.d/start.conf
rm /etc/nginx/sites-enabled/default
sudo systemctl restart nginx

```

## Geoserver
```
nano /var/lib/tomcat9/webapps/geoserver/WEB-INF/web.xml

Zeile hinzufügen


<context-param>
  <param-name>GEOSERVER_CSRF_WHITELIST</param-name>
  <param-value>example.org</param-value>
</context-param>

```

## Tomcat neu starten

```

systemctl restart tomcat9

```

### Go und Rust Installieren

```
mkdir ~/download
cd ~/download
wget https://go.dev/dl/go1.18.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
sudo curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sudo sh
cd /Server
make build
cd config
cp demo.config.yaml config.yaml
nano config.yaml
cd ..

 sudo cp -u Server_Einstellung/ahrensburg.service /etc/systemd/system/ahrensburg.service
 sudo  systemctl enable ahrensburg.service
sudo  systemctl start ahrensburg.service

```
## C und C++ Installieren

```

sudo apt-get install cmake gcc clang gdb build-essential

```


