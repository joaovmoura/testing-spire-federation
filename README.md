# Spire federation test
Testando o Spire federation para fornecer SVID's para a comunicação entre um server e um client rodando em diferentes clusters kubernets.

Algumas informações úteis podem ser encontradas aqui:

- [Sobre o spire](https://maize-fish-44c.notion.site/Spiffe-Spire-33b0c52abacb4694b1e1437814845ab4?pvs=4)

- [Sobre o teste](https://maize-fish-44c.notion.site/Federation-test-fd4ae3b9dc99484eb0ba2f57a2419f8d?pvs=4)

# Setup

## Requerimentos
 - Golang(usei a versão 1.18.1)
 - Docker(usei a versão 24.0.4)

1- Clonar o repositório

 `git clone https://github.com/joaovmoura/testing-spire-federation.git`

2- Entrar no diretório

 `cd testing-spire-federation`

3- Rodar o script "setup.sh"

OBS.: A ideia inicial da configuração demandava a criação de um registry local para subir as imgens do server e do client.
No entanto tive problemas para configurar localmente o registry para ser acessado na criação dos PODS. Portanto, por hora não seria bom usar os
scripts de setup.
