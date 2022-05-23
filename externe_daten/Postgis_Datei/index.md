## exportieren Windows
```
pg_dump -U postgres ahrensburg > ahrensburg.sql

```

## Importiert Ubuntu

```

psql -U ahrensburg ahrensburg < ahrensburg.sql

```