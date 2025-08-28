# API DeliveKeep.

## Estrutura da api:
* Routes: definição de rotas (nome, método http e função)
* Controllers: funções das rotas (recebe requisição e chama outros pacotes para enviar resposta)
* Repositories: funções de interação com banco de dados
* Models: classes para validar dados
* Middlewares: funções a serem executadas entre a requisição e chamar funções das rotas de fato
* Auth: funções que envolvem autorização/jwt
* Config: pacote de inicialização de variáveis de ambiente (.env)
* Database: abertura da conexão com banco de dados
* Response: formatação de respostas a serem devolvidas
* Secutiry: funções de segurança/hash
* Utils: funções de utilidades diversas que não se encaixam em nenhum do pacotes

## Siga os passos abaixo para clonar e executar o projeto localmente:
* 1. Clone o repositório
```
git clone https://github.com/RafaellaMasutti/API_DeliveKeep
cd API_DeliveKeep
```
* 2. Crie um banco de dados postgresql da maneira que preferir e rode o script dbinit.sql nele

* 3. Crie um .env na raiz do projeto contendo
```
DB_USER=usuario_do_banco
DB_PASSWORD=senha_do_banco
DB_NAME=nome_do_banco
DB_PORT=porta_do_banco
DB_HOST=servidor_do_banco
API_PORT=porta_da_api
SECRET_KEY=chave_secreta
```
* 4. Instale as dependências
```
cd api
go mod download
go mod verify
```
* 5. Compile o código
```
go build -o nome_executavel .
```
* 6. Rode
```
./nome_executavel
nome_executavel.exe # Windows
```