[table partitioning](https://dba.stackexchange.com/questions/106014/how-to-partition-existing-table-in-postgres)

https://blog.heroku.com/handling-very-large-tables-in-postgres-using-partitioning

```shell
pg_dump postgres://postgres:postgres@localhost:15432/partitioning -F t > partitioning.tar
```

```shell
pg_restore -d 'postgres://postgres:postgres@localhost:15432/partitioning' --no-privileges --no-owner -c partitioning.tar
```
 
