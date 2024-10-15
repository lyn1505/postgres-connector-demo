# postgres-connector-demo

## To run the demo
1. Clone the repository
2. Run the databases. This will run `docker-compose up` under the hood, and set up 2 masters and 2 slaves:
```bash
make up master=1 slave=1 # can set different numbers to scale
```

3. Run the tests in [`db_test.go`](https://github.com/lyn1505/postgres-connector-demo/blob/main/db_test.go). The databases will log SQL queries it executes, so to see the logs, run any of these:
```bash
docker-compose logs -f pg-demo-postgresql-master-users-1
docker-compose logs -f pg-demo-postgresql-master-games-1
docker-compose logs -f pg-demo-postgresql-slave-users-1
docker-compose logs -f pg-demo-postgresql-slave-games-1
```
Read the comments in the test file to understand what each test does.

4. When you're done:
```bash
make clean # remove containers and volumes
```
