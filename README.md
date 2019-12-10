# Greeter-Service und Client als getrennte Apps

1. Starten Sie `etcd`.

   ```
   make start-registry
   ```

2. Starten Sie den Greeter-Service.

   ```
   make start-counter
   ```

3. Starten Sie den Counter-Service.

   ```
   make start-greeter
   ```

4. Starten Sie den Client

   ```
   make start-client
   ```

Wenn Sie die Registry stoppen, gibt es einen Timeout vom Client, aber er scheint den alt bekannten Service erfolgreich
zu connecten.
Wenn Sie die Server stoppen und neu starten, connected der Client an den neu gestarteten Service (auch auf anderem Port).
