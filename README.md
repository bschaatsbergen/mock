<p align="center"><img src="https://twemoji.maxcdn.com/2/svg/1f4dd.svg" height="64" alt="Project Logo"></p>
<h3 align="center">mock</h3>
<p align="center">📡  mock is a simple, cross-platform, cli app to simulate HTTP-based APIs. </p>
<p align="center">
    <a href="https://github.com/bschaatsbergen/mock/releases"><img src="https://img.shields.io/github/downloads/bschaatsbergen/mock/total.svg" alt="GitHub Downloads"></a>
    <a href="https://github.com/bschaatsbergen/mock/releases/latest"><img src="https://img.shields.io/github/release/bschaatsbergen/mock.svg" alt="Latest Release"></a>
    <a href="https://github.com/bschaatsbergen/mock/actions/workflows/go.yaml"><img src="https://img.shields.io/github/workflow/status/bschaatsbergen/mock/Go" alt="Build Status"></a>
    <a href="https://github.com/bschaatsbergen/mock/issues"><img src="https://img.shields.io/badge/contributions-welcome-ff69b4.svg" alt="Contributions Welcome"></a>
</p>

## About mock

Mock allows you to spin up a local http server based of a `.mock.yaml` file. Preferably mock is used per project and a `.mock.yaml` file is created at the root of a repository.

## Installation

### Binaries 

If you prefer grabbing mock its binaries, download the latest from the the **[GitHub releases](https://github.com/bschaatsbergen/mock/releases)** page.

### Brew

```sh
❯ brew tap bschaatsbergen/mock
❯ brew install mock
```

### Chocolatey (Not supported yet)
```cmd
❯ choco install mock
```

## Usage

Running mock is very easy, by simply running the below command you tell mock to translate the `.mock.yaml` in the current working directory to an http server that by default runs under port 7070. Use the `-p` flag to change the default port.

```sh
❯ mock serve

📝 using config file: .mock.yaml
📡 starting mock on :7070, exit to stop
```

To bring you up to speed, we allow you to run the below command in order to generate an exampe `.mock.yaml` in the current working directory.

```sh
❯ mock init

🎉 generated '.mock.yaml' in: /home/bruno/Projects/MyNewApi
```

The generated `.mock.yaml` configuration looks as following. It's made up of a list of `endpoints`, each containing a resource path, http method, response and http status code. 

```yaml
# Example .mock.yaml config
Endpoints:
  - Resource: /city/1
    Method: GET
    Response: '{ "Id": 1, "Name": "Albuquerque", "Population": 559.374, "State": "New Mexico" }'
    StatusCode: 200

  - Resource: /city
    Method: POST
    Response: '{ "Name": "Albuquerque", "Population": 559.374, "State": "New Mexico" }'
    statusCode: 200

  - Resource: /city/1
    Method: PUT
    Response: '{ "Id": 1, "Name": "Albuquerque", "Population": 601.255, "State": "New Mexico" }'
    StatusCode: 200

  - Resource: /city/1
    Method: DELETE
    StatusCode: 204

```

After running `mock serve` you can verify whether the routes are available

```sh
❯ curl -i -X GET 'http://localhost:7070/city/1'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 19 Oct 2021 13:25:47 GMT
Content-Length: 80

{ "Id": 1, "Name": "Albuquerque", "Population": 559.374, "State": "New Mexico" }
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
