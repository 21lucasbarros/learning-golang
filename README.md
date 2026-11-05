# Learning Golang

<div align="center">
  <p align="center">
    <img src="https://skillicons.dev/icons?i=go,docker,postgres,nginx,bash" />
  </p>

  <p align="center">
    <strong>Repositório dedicado ao aprendizado e domínio do ecossistema Go (Golang).</strong><br />
    Da construção de APIs REST eficientes no back-end à conteinerização com Docker, com foco em concorrência, performance, boas práticas e arquitetura escalável.
  </p>
</div>

## Sobre

Este projeto reúne exemplos, exercícios e pequenas aplicações para estudar e praticar Go em diferentes contextos. O objetivo é desenvolver familiaridade com a sintaxe da linguagem, compreender seus mecanismos de concorrência (goroutines e channels), construir serviços back-end e explorar a conteinerização de aplicações utilizando Docker.

## Objetivos

- Explorar os principais recursos do Go (estruturas de dados, ponteiros, interfaces e concorrência).
- Aplicar Go no desenvolvimento de APIs RESTful eficientes e escaláveis.
- Compreender e aplicar goroutines e channels para programação concorrente.
- Implementar conteinerização com Docker, incluindo builds multi-stage.
- Integrar aplicações com bancos de dados relacionais, como PostgreSQL.
- Aplicar boas práticas de código, organização e arquitetura de software.

## Estrutura do repositório

- `api/` - Código-fonte da aplicação back-end e implementação das rotas da API REST em Go.
- `docker/` - Dockerfiles e configurações de conteinerização e infraestrutura.
- `exercises/` - Exercícios isolados para praticar conceitos da linguagem e concorrência.

> Nota: a estrutura será expandida conforme o aprendizado avança.

## Tecnologias

- Go (Golang)
- Docker & Docker Compose
- PostgreSQL (planejado)
- Nginx (planejado)

## Como começar

Pré-requisitos: Go (>= 1.21) e Docker instalados na máquina.

1. Clonar o repositório e acessar a pasta da API:

```bash
git clone https://github.com/21lucasbarros/learning-golang.git
cd learning-golang/api
```

2. Executar a aplicação:

```bash
go run .
```

3. Para executar a aplicação utilizando Docker, consulte os arquivos de configuração disponíveis no diretório `docker/`.

## Boas práticas recomendadas

- Seguir as convenções e o formato oficial da linguagem (`gofmt`).
- Priorizar simplicidade, legibilidade e organização do código.
- Utilizar interfaces para desacoplar componentes quando necessário.
- Tratar erros explicitamente, seguindo as convenções do Go.
- Escrever testes para validar o comportamento das aplicações.

## Contribuição

- Abra uma issue para sugerir melhorias estruturais ou reportar bugs.
- Pull requests são bem-vindos — foque em adições limpas e bem documentadas.

## Licença

MIT — veja o arquivo LICENSE para detalhes.
