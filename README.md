## Deploy Docker Container
```bash
# Build container
docker build -t golang-animeapi:multistage -f Dockerfile.multistage .

# Run container
docker run -dp 81:8080 golang-animeapi:multistage
```

## Variables & Syntax
 - `!` means the variable and the leading slash can be left out 
 - `{}` contains possible value

|Variable Name|Possible Value|Datatype or Regex|Default Value|
|---        |---                |:-----:|:-----:|
|**version**|{ v1 }             |-      |-      |
|**index**  |{ 0, 1, 2, 3, ... }|uint   |0      |

## Animes
### List all Animes
#### Syntax
```
GET /api/{version}/anime/list/{!index}

http://localhost:8080/api/v1/anime/list
http://localhost:8080/api/v1/anime/list/0
```
#### Response
```js
// GOOD RESPONSE
{
    "data": [...],
    "error": null,
    "len": 99,
    "success": 1
}
```
- **data**      | lists 100 elements
- **len**       | the amount of data the server has
- **error**     | error description if sth failes
- **success**   | if is success

```js
// BAD RESPONSE
{
    "data": null,
    "error": "Bad Request",
    "len": null,
    "success": 0
}
```
- **error** | error description if sth failes
- **success** | if is success 1 else 0

## Users
### Register
#### Syntax
```
POST /api/{version}/user/register

http://localhost:8080/api/v1/user/register
```
#### Response
```js
// GOOD RESPONSE
{
    "data": null,
    "error": null,
    "len": null,
    "success": 1
}
```
- **success**   | if is success 1 else 0

```js
// BAD RESPONSE
{
    "data": null,
    "error": "Bad Request",
    "len": null,
    "success": 0
}
```
- **error** | error description if sth failes
- **success** | if is success 1 else 0
