# Eine simple Beispielanwendung mit [go-micro](https://github.com/micro/go-micro)

## Verwendete Technologien/Plattformen

-   Synchrone Kommunikation über [gRPC](https://grpc.io/)
-   Registry: [etcd](https://etcd.io/)
-   Broker für asynchrone Kommunikation: [NATS](https://nats.io/)
-   Key/Value-Store: [Redis](https://redis.io/)
-   Logging: [zerolog](https://github.com/rs/zerolog)

## Services

### Greeter

Empfängt einen `HelloRequest` per gRPC, fragt einen Zähler-Wert beim Counter-Service an
und sendet eine Antwort an den Absender.

### Counter

Speichert einen Zähler pro Namen und antwortet auf einen `IncRequest` mit einem
`SumResponse`, wobei der Zähler inkrementiert wird und die Info als asynchrone
Nachricht veröffentlicht wird.

### Logger

Hat Log-Nachrichten abonniert und gibt diese auf der Console aus. In den Key/Value-Store
fügt er eine Zufallszahl mit dem Schlüssel `sleep` ein.

## Starten

Zum Ausprobieren steht ein Client zur Verfügung, der mit dem Greeter-Service kommuniziert.
Zwischen zwei Anfragen wartet der Client immer eine gewisse Zeit. Die Zeit versucht
er aus einem Key/Value-Store auszulesen.

### Starten mit wenig Docker oder ohne Docker

**[wenig Docker]**

Starten Sie die drei benötigten Server mit folgendem Kommando:

```
docker-compose -f docker-compose-deps.yaml up
```

oder

**[kein Docker]**

Installieren Sie etcd, NATS und Redis und starten Sie in je einem Terminal
die 3 benötigten Server:

```
# jede Zeile in einem anderen Terminal
etcd
nats-server
redis-server
```

**[und jetzt für beide Ansätze]**

die drei Services und den Client:

```
# jede Zeile in einem anderen Terminal
go run greeter-service/main.go
go run counter-service/main.go
go run logger-service/main.go
go run client/main.go
```

### Starten mit Docker aus lokalen Sourcen

**[Achung: go-micro findet in der aktuellen Version 2.6.0 den Redis nicht, wenn
er nicht auf localhost läuft. Daher funktioniert der Zugriff auf den Store nicht,
wenn auch die Services und der Client als Docker-Conatainer laufen. Mit den
Entwicklern habe ich diesbezüglich schon Verbindung aufgenommen.]**

Die Datei `docker-compose-local.yaml` definiert die Services über die `Dockerfiles`.
Erzeugen Sie die Docker-Images lokal mit `...build` und starten Sie alles mit `...up`.
Wenn die Dockerfiles beim ersten Start nicht vorhanden sind, werden sie automatisch
erzeugt bzw. heruntergeladen.

```
docker-compose -f docker-compose-local.yaml up
```

Wenn Sie neuere Vesionen der Images verwenden wollen, nutzen Sie `pull` bzw. `build`.
Wenn Sie ein neues Image, das nur für den Build-Step verwendet wird, herunterladen
wollen, verwenden Sie `build --pull`.

```
docker-compose -f docker-compose-local.yaml pull
docker-compose -f docker-compose-local.yaml build --pull
# oder
docker-compose -f docker-compose-local.yaml build
```

Wenn Sie an einem einzelnen Service etwas ändern, können Sie auch nur für diesen
ein neues Docker-Image erzeugen, z.B.

```
docker-compose -f docker-compose-local.yaml build client
```

Denken Sie also daran, wenn Docker ein Image hat, sucht es weder in der Registry
nach einem aktuelleren noch wird, bei geänderten Sourcen, ein neueres erzeugt.
Sie müssen dass immer explizit mit `pull` oder `build` anstossen.

### Starten mit Docker aus bereits erzeugten Images

**[Achung: go-micro findet in der aktuellen Version 2.6.0 den Redis nicht, wenn
er nicht auf localhost läuft. Daher funktioniert der Zugriff auf den Store nicht,
wenn auch die Services und der Client als Docker-Conatainer laufen. Mit den
Entwicklern habe ich diesbezüglich schon Verbindung aufgenommen.]**

Wenn Sie beim `docker-compose`-Kommando keine Datei angeben, wird die Datei `docker-compose.yaml`
verwendet. Darin sind die Docker-Images aus der GitHub-Registry angegeben. Dass heisst, Sie benötigen nur die Datei `docker-compose.yaml` und überhaupt keine Sourcen.

Zum Starten geben Sie daher einfach folgendes ein:

```
docker-compose up
```
