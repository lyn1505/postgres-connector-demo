up:
	docker-compose up --detach --scale postgresql-master-users=$(master) --scale postgresql-master-games=$(master) --scale postgresql-slave-users=$(slave) --scale postgresql-slave-games=$(slave)

down:
	docker-compose down

clean:
	docker-compose down --volumes
