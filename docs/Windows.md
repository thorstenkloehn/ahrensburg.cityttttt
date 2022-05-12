
## Windows Gis Software

* QGIS
* PostgreSQL


## Windows

## Einstellung

```


    setx TEMP e:\AppData\Local\Temp
    setx TMP e:\AppData\Local\Temp
     setx GOPATH e:\gopath
   
```

## Grundlagenn Programme


### Anti Virus Programmieren

* [Bitdefender](https://login.bitdefender.com/central/login.html?lang=de_DE&redirect_url=https:%2F%2Fcentral.bitdefender.com%2Factivity%3FbrowserLang%3Dde_DE)

### Browser

* [Microsoft Edge](https://www.microsoft.com/en-us/edge)
* [Google Chrome](https://www.google.de/chrome)
### Versionsverwaltungssoftware


* [Git Installieren](https://git-scm.com/)
### Programmiersprachen
### Allgemeine Module


* [Python](https://www.python.org/downloads/)

## IDE
* [IntelliJ IDEA](https://www.jetbrains.com/idea/) Erste Installieren und Starten
* [Clion](https://www.jetbrains.com/de-de/clion/)

## Editor


## Allgemeine Software
## Windows CD
### PID.txt

```

[PID]
Value=XXXXX-XXXXX-XXXXX-XXXXX-XXXXX

```

### ei.cfg

````

[EditionID]
Professional
[Channel]
Retail

````

## Webassembly
### Emscripten
#### Einführung

* Windows 10
* CLion 2019.1 oder höher Version.
* Emscripten
* Chrome

#### Vorbereitung

#### Emscripten SDK

##### Installieren Sie das Emscripten SDK

```

git clone https://github.com/emscripten-core/emsdk.git
cd emsdk
emsdk install latest
emsdk activate latest

```

##### Clion Einstellung.

###### PATH Variables

* Öffnet File -->> Settings -->> PATH Variables
* Erstellen Sie eine Variable EMSCRIPTEN.

##### Konfigurieren Sie die Build-Einstellungen für Emscripten

* Öffnet Sie File -->> Settings -->> Build, Execution, Deployment -->Cmake
* CMake options eingeben :
```
-DCMAKE_TOOLCHAIN_FILE=${EMSCRIPTEN}/upstream/emscripten/cmake/Modules/Platform/Emscripten.cmake


