#### Creado por Violeth Valmont Azahar

# API REST GO OATI

En un contexto de aprendizaje virtual, un tutorial de un tema en
específico tiene unos detalles en puntual; por ejemplo, la fecha en qué se creó y
quién lo creó. Se desea registrar los tutoriales con información como
descripción, título, estado de publicación (visible/oculto) y los detalles de este:
día de creación, usuario que lo creó. Nótese que un tutorial tiene únicamente un
detalle de tutorial. Se debe permitir listar los tutoriales, el detalle asociado al
tutorial, así como permitirse la modificación y eliminación de un tutorial y su
detalle.



## Deployment



```bash
  cd database
  docker build . 
  docker run 
```

```bash
  go get github.com/gin-gonic/gin
```

```bash
  go run main.go
```
## Tech Stack

 Go, Gin, Dockerfile



## Demo

https://github.com/LethSphere/api-rest-go-oati/blob/main/assets/postTutorial.png

https://github.com/LethSphere/api-rest-go-oati/blob/main/assets/postDetalles.png

https://github.com/LethSphere/api-rest-go-oati/blob/main/assets/getTutorial.png

https://github.com/LethSphere/api-rest-go-oati/blob/main/assets/getDetalles.png

https://github.com/LethSphere/api-rest-go-oati/blob/main/assets/deleteTutorial.png

## Features

- Agregar, modificar o eliminar tutoriales
- Agregar, modificar o eliminar detalles asociados a un tutorial
- Listar tutoriales y sus detalles asociados

## API Reference


#### InsertTutorial

```http
  POST /tutorials
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `titulo` | `string` | **Required** |
| `descripcion` | `string` | **Required** |
| `estado` | `string` | **Required** |

#### GetTutorialById

```http
  GET /tutorials/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.  |


#### UpdateTutorial

```http
  PUT /tutorials/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.  |
| `titulo` | `string` | **Required** |
| `descripcion` | `string` | **Required** |
| `estado` | `string` | **Required** |

#### DeleteTutorial

```http
  DELETE /tutorials/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.  |

#### ListTutorial

```http
  GET /detalles
```

#### InsertDetalle

```http
  POST /detalles
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `tutorialId`      | `string` | **Required**.  |
| `Autor` | `string` | **Required** |

#### GetDetallelById

```http
  GET /detalles/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.  |


#### UpdateDetalle

```http
  PUT /detalles/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tutorialId`      | `string` | **Required**.  |
| `Autor` | `string` | **Required** |

#### DeleteDetalle

```http
  DELETE /detalles/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**.  |

#### ListDetalle

```http
  GET /detalles
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |




## Database
- Crea todas las tablas necesarias según el modelo ER.
- Definir las claves primarias y claves foráneas.

#### 1. tutorials

- **id_Tutorials (PK)**
- titulo 
- descripcion 
- estado 


#### 2. detalles
- **id_detalles (PK)**
- autor 
- created_at 
- tutorials_id (FK)
