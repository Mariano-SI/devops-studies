# 🏗️ Guia de Governança e Fluxo DevOps - Health Checker

Este documento detalha as regras de automação, proteção de branches e o fluxo de trabalho profissional configurado neste repositório.

## 📌 Configurações do Repositório (GitHub)

Para garantir a integridade do código, o repositório foi configurado com as seguintes diretrizes:

1.  **Branch Default:** Definida como `develop`. Todo o desenvolvimento parte daqui.
2.  **Proteção de Branches (Rulesets):**
    * **Branches Protegidas:** `master` e `develop`.
    * **Bloqueio de Push Direto:** É impossível enviar código diretamente para estas branches via terminal. Todo código deve passar por um **Pull Request**.
    * **Status Checks:** O merge só é liberado se todos os testes do GitHub Actions (CI) passarem com sucesso.
3.  **Permissões de Actions:**
    * Configurado em `Settings > Actions > General`.
    * **Workflow permissions:** `Read and Write permissions`.
    * **Allow GitHub Actions to create and approve pull requests:** Ativado (necessário para o robô de Backport).

---

## 🚦 Regras do CI (GitHub Actions)

O pipeline de Integração Contínua (`ci.yml`) executa duas validações críticas:

### 1. Validação de Nomenclatura (Gitflow)
O robô analisa os nomes das branches de origem e destino para garantir a organização do projeto:

| Destino (Base) | Origem Permitida (Head) | Objetivo |
| :--- | :--- | :--- |
| **`master`** | `release/*`, `hotfix/*` | Estabilidade total. A master só recebe versões testadas ou correções críticas. |
| **`develop`** | `feature/*`, `bugfix/*`, `release/*`, `hotfix/*` | Integração. Aceita novas funcionalidades e sincronizações de volta. |

### 2. Verificação de Código (Go)
Para garantir que o software está funcional, o CI executa:
* `go mod tidy`: Valida as dependências do projeto.
* `go test -v ./...`: Roda os testes unitários. Se houver falha, o merge é bloqueado.

---

## 🔄 Automação de Sincronia (Backport)

Configuramos o workflow `Automatic Backport (Full Sync)` para resolver o problema de "desincronização" entre produção e desenvolvimento.

**Como funciona:**
1.  Sempre que ocorre um merge na branch **`master`** (seja via Release ou Hotfix).
2.  O robô detecta a mudança e cria automaticamente um Pull Request da `master` para a `develop`.
3.  Isso garante que qualquer correção feita diretamente na `master` (Hotfix) seja levada de volta para a `develop` sem que o desenvolvedor precise lembrar de fazer isso manualmente.

---

## 🛠️ Como trabalhar no projeto (Guia Prático)

### Criar uma nova funcionalidade
```bash
git checkout develop
git pull origin develop
git checkout -b feature/nome-da-funcionalidade
# desenvolve o código...
git push origin feature/nome-da-funcionalidade
# No GitHub, abre PR para 'develop'