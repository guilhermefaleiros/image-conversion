# Image Conversion

Este é um projeto construído como entregável do projeto prático da disciplina de Linguagens e Paradigmas de Programação da Universidade Federal de Goiás, de modo a explorar as características da linguagem Golang para resolução do problema de aplicação de efeitos e conversão de formatos de imagens através da construção de uma API RESTful em que é possível se cadastrar e utilizar os serviços oferecidos.

## Funcionalidades

### Cadastro de Clientes

- **Descrição**: É possível se cadastrar como cliente da API, fornecendo um nome de usuário, de forma que as credenciais para obtenção de um token de acesso serão fornecidas como resposta do serviço.
- **URL**: `/client`
- **Método**: `POST`
- **Estrutura**:
    - **Corpo da requisição**:
      ```json
      {
          "name": "<NOME DO CLIENTE>"
      }
      ```
    - **Corpo da resposta**:
      ```json
      {
          "access_key": "<CHAVE DE ACESSO GERADA>",
          "id": "<ID DO CLIENTE GERADO>",
          "name": "<NOME DO CLIENTE>"
      }
      ```

### Obtenção de Token de Acesso

- **Descrição**: Uma vez com um cliente gerado, é necessário obter tokens de acesso para que seja possível utilizar os serviços oferecidos pela API. Para isso, é necessário fornecer o ID do cliente e a chave de acesso gerada no cadastro.
- **URL**: `/generate-token`
- **Método**: `POST`
- **Estrutura**:
    - **Corpo da requisição**:
      ```json
      {
          "client_id": "<ID DO CLIENTE>",
          "access_key": "<CHAVE DE ACESSO>"
      }
      ```
    - **Corpo da resposta**:
      ```json
      {
          "access_token": "<TOKEN DE ACESSO GERADO>"
      }
      ```

### Upload de Imagem

- **Descrição**: Antes de realizar o processamento de efeitos nas imagens, é necessário fazer upload previamente das imagens que serão utilizadas. Para isso, é necessário fornecer o ID do cliente e a chave de acesso gerada no cadastro.
- **URL**: `/image`
- **Método**: `POST`
- **Estrutura**:
    - **Corpo da requisição**:
        - Headers:
            ```plaintext
            Content-Type: multipart/form-data
            Authorization: <TOKEN DE ACESSO>
            ```
        - Multipart Form:
            ```plaintext
            image: <ARQUIVO DE IMAGEM>
            ```
    - **Corpo da resposta**:
      ```json
      {
          "id": "<ID DA IMAGEM GERADA>"
      }
      ```

### Aplicação de efeitos em imagens

- **Descrição**: É possível aplicar efeitos em imagens previamente cadastradas na API, dessa forma é necessário passar o ID da imagem e uma lista de efeitos que devem ser aplicados. É importante ressaltar que este serviço permite processar um lote de imagens fornecidas.
- **URL**: `/image/process`
- **Método**: `POST`
    - **Headers**:
    ```plaintext
    Content-Type: application/json
    Authorization: <TOKEN DE ACESSO>
   ```
    - **Corpo da requisição**:
      ```json
      [
         {
             "image_id": "<ID DA IMAGEM>",
             "effects": ["invert_colors", "png"]
         }
      ]
      ```
  No campo effects devem ser passados todos os efeitos que devem ser aplicados à imagem, os seguintes efeitos estão disponíveis:
    - `invert_colors`: Inverte as cores da imagem
    - `grayscale`: Aplica escala de cinza na imagem
    - `png`: Converte a imagem para o formato PNG
    - `jpg`: Converte a imagem para o formato JPG
    - `sepia`: Aplica o efeito de sépia na imagem

    - **Corpo da resposta**:
    ```json
    {
        "processed_images": [
            {
                 "image_id": "<ID DA IMAGEM>",
                 "status": "success",
                 "effects": [
                     "invert_colors",
                     "png"
                 ],
                 "processed_image_id": "<ID DA NOVA IMAGEM PROCESSADA>"
            }
        ]
    }
    ```

### Download de imagens processadas

- **Descrição**: Com este serviço é possível fazer o download de imagens previamente processadas, para isso é necessário fornecer o ID da imagem processada e as credenciais de acesso do cliente.
- **URL**: `/processed-images/<ID DA IMAGEM PROCESSADA>/download`
- **Método**: `GET`
- **Estrutura**:
  - **Headers**:
    ```plaintext
    Content-Type: application/json
    Authorization: <TOKEN DE ACESSO>
    ```

### Tecnologias utilizadas

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [GoORM](https://gorm.io/)
- [Docker](https://www.docker.com/)
