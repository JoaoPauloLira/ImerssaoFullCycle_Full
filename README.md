# ImersaoFullCycle_01
Implementação do Projeto Go-kafka


1 - Rodar o docker-compose do Kafka
  * cd .docker/kafka
  * docker compose up -d
  * docker exec -it kafka-kafka-1 bash
  * (Console 1) kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position
  * (Console 2) kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
    - {"clienteId":"1","routeId":"1"}
  
2 - Rodar o docker-compose do Simulador
  * docker compose up
  
3 - Rodar o GO
  * docker compose exec app bash
  * go run main.go
