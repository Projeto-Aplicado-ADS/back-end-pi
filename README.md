# Back-end API - Como utilizar


- Dar git clone do projeto
- docker-compose up -d (Docker deve estar aberto)
- No mysql é so conectar

### Como Rodar o Boneless
Baixar packages:
```
  go get -u -d github.com/golang-migrate/migrate/cmd/migrate
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
  go install github.com/renanbastos93/boneless/cmd/boneless@latest
```
### Subir as migrations e rodar o projeto
```
  boneless migrate (nome da migrate) up 
  boneless build 
  boneless run
```
