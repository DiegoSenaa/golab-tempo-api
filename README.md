# Desafio Sistema de temperatura por CEP

### Execução Local

1. Clone o repositório e acesse a subpasta do desafio:

   ```bash
   git clone https://github.com/mbombonato/pos-go-expert.git
   cd pos-go-expert/challenge-weather-by-cep
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Execute o servidor:

   ```bash
   go run main.go
   ```

4. Execute a curl em outro terminal ou abra a url no navegador:

   ```bash
   curl --request GET --url 'http://localhost:8080/<CEP>'
   ou
   http://localhost:8080/<CEP>
   ```

### Execução Local via Docker-Compose
1. Suba o Docker Compose
   ```bash
   docker compose up
   ```

2. Execute a curl em outro terminal ou abra a url no navegador
   ```bash
   curl --request GET --url 'http://localhost:8080/<CEP>'
   ou
   http://localhost:8080/<CEP>
   ```

### Execução no Cloud-Run
1. Abra a URL abaixo:

   ```bash
   # invalid zip code
   https://tempo-api-p6bxy5owja-rj.a.run.app/666
   # zip code not found
   https://tempo-api-p6bxy5owja-rj.a.run.app/02253022

   # success
   https://tempo-api-p6bxy5owja-rj.a.run.app/02034002
   ```
