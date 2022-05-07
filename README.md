# ahrensburg.city

## Installieren
```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install git-core
apt-get install  gnupg2
sudo apt-get install tomcat9
```
* [OSM](https://switch2osm.org)
* [Geoserver](https://geoserver.org/)

## Apache Einstellung Ändern

```

nano /etc/apache2/ports.conf

```

## Geoserver Einstellung Ändern
```
nano /var/lib/tomcat9/webapps/geoserver/WEB-INF/web.xml

Zeile hinzufügen 


<context-param>
  <param-name>GEOSERVER_CSRF_WHITELIST</param-name>
  <param-value>example.org</param-value>
</context-param>
```

## Apache neu Starten

```
systemctl restart apache2

```

## Tomcat9 neu Starten

```
systemctl restart tomcat9

```

