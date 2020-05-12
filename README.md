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

Hat Log-Nachrichten aboniert und gibt diese auf der Console aus.

## Starten

Zum Ausprobieren steht ein Client zur Verfügung, der mit dem Greeter-Service kommuniziert.

### Starten mit lokalen Sourcen

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

### Starten mit bereits erzeugten Images

Wenn Sie beim `docker-compose`-Kommando keine Datei angeben, wird die Datei `docker-compose.yaml`
verwendet. Darin sind die Docker-Images aus der GitHub-Registry angegeben. Dass heisst, Sie benötigen nur die Datei `docker-compose.yaml` und überhaupt keine Sourcen.

Zum Starten geben Sie daher einfach folgendes ein:

```
docker-compose up
```
