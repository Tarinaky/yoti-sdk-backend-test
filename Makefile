restart: stop start

build:
	make -C ysbt_api
	make -C ysbt_db


start: build
	docker stack deploy -c ysbt-compose.yml ysbt

stop:
	-docker service rm ysbt_db
	-docker service rm ysbt_api
	-docker stack down ysbt

