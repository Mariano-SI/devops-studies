# Guia de Governança e Fluxo DevOps - Health Checker

Este documento detalha as regras de automação, proteção de branches e o fluxo de trabalho configurado neste repositório.

## Configurações do Repositório (GitHub)

Para garantir a integridade do código, o repositório segue as seguintes diretrizes:

1. **Branch Default**: Definida como `develop`. Todo o desenvolvimento deve partir desta branch.
2. **Proteção de Branches (Rulesets)**:
    * **Branches Protegidas**: `master` e `develop`.
    * **Bloqueio de Push Direto**: Proibido enviar código diretamente para estas branches via terminal. Todo código deve ser submetido via **Pull Request**.
    * **Status Checks**: O merge só é liberado se os testes de Go e o Build do Docker forem concluídos com sucesso.
3. **Permissões de Actions**: Configurado para `Read and Write permissions` para permitir a publicação de imagens e criação automática de PRs de sincronia.

---

## Regras do CI/CD (GitHub Actions)

O pipeline de Integração e Entrega Contínua (`ci.yml`) executa três pilares de validação:

### 1. Validação de Nomenclatura (Gitflow)

| Destino (Base) | Origem Permitida (Head) | Objetivo |
| :--- | :--- | :--- |
| **master** | `release/*`, `hotfix/*` | Produção. Apenas versões estáveis ou correções críticas. |
| **develop** | `feature/*`, `bugfix/*`, `release/*`, `hotfix/*` | Integração. Base para o desenvolvimento diário. |

### 2. Verificação de Código (Go)
* `go mod tidy`: Validação de integridade das dependências.
* `go test -v ./...`: Execução de testes unitários. Falhas neste estágio bloqueiam o merge.

### 3. Containerização e Entrega (Docker)
O fluxo de trabalho do Docker é dividido em dois níveis de permissão:
* **Validação (Qualquer Branch/PR)**: O CI executa um *Smoke Test* para garantir que o `Dockerfile` compila corretamente. **Nenhuma imagem é salva ou publicada neste estágio.**
* **Entrega (Exclusivo Master)**: Apenas quando ocorre um push ou merge na branch **master**, o CI gera a imagem final e realiza o push para o **GitHub Container Registry (GHCR)**.

---

## Registro de Imagens (GHCR)

As imagens oficiais são geradas e armazenadas no GitHub Packages **apenas após o merge em produção (master)**. Para baixar a versão estável mais recente:

```bash
docker pull ghcr.io/${{ github.repository }}:latest