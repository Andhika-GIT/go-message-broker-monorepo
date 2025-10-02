Got it 🚀 — here’s a concise but complete **README.md** draft for your project, covering the key points we discussed:

---

# Go Message Broker Monorepo (Dev Environment)

## Requirements

* Docker & Docker Compose
* RabbitMQ running on your host machine (Linux Mint in this setup)
* Go 1.25+ (for local dev if needed)
* Node.js (for frontend local dev if needed)

---

## 1. RabbitMQ Setup

By default, RabbitMQ only allows the `guest/guest` user from `localhost`.
Since our workers run inside Docker, we need a dedicated user.

Run these commands on your host (Linux Mint terminal):

```bash
# Enter RabbitMQ management (assuming it's running locally)
rabbitmqctl add_user devuser devpass
rabbitmqctl set_user_tags devuser administrator
rabbitmqctl set_permissions -p / devuser ".*" ".*" ".*"
```

Update your `.env` files:

```env
RABBITMQ_CONNECTION_URL=amqp://devuser:devpass@host.docker.internal:5672/
```

---

## 2. Docker Compose Services

We use three main services:

```yaml
version: "3.9"
services:
  be:                # Backend (Go + Air for hot reload)
    ports: ["3005:3005"]
    extra_hosts: ["host.docker.internal:host-gateway"]

  upload-worker:     # Worker service (Go + Air)
    ports: ["3009:3009"]
    extra_hosts: ["host.docker.internal:host-gateway"]

  fe:                # Frontend (Next.js)
    ports: ["4000:4000"]
    volumes:
      - ./fe:/app
      - /app/node_modules
```

### Important:

* `extra_hosts: host.docker.internal:host-gateway`
  → lets containers connect to host services (e.g., RabbitMQ running on Linux Mint).
* `volumes: ./xxx:/app`
  → mounts your source code into the container for live development.
* **Do not ignore `.air.toml`** in `.dockerignore`, otherwise `air` won’t run properly.

---

## 3. Environment Variables

### Backend & Worker

```env
RABBITMQ_CONNECTION_URL=amqp://devuser:devpass@host.docker.internal:5672/
```

### Frontend

* `/fe/.env`

```env
NEXT_PUBLIC_SERVER_BASE_URL_FOR_CLIENT=http://localhost:3005
NEXT_PUBLIC_SERVER_BASE_URL=http://be:3005
```

* **Rule of thumb:**

  * `NEXT_PUBLIC_*` → used in client-side code (browser).
  * Non-`NEXT_PUBLIC` → used only in server-side (Next.js container).

---

## 4. Common Issues

* **`air.toml not found`** → make sure `.air.toml` is not excluded in `.dockerignore`.
* **RabbitMQ 403 (guest user blocked)** → create your own RabbitMQ user (see step 1).
* **`localhost` inside containers** → always use:

  * `be:3005` to reach backend from other containers.
  * `host.docker.internal` to reach services on the host machine.

---

## 5. Run the Environment

```bash
docker-compose up --build
```

Frontend → [http://localhost:4000](http://localhost:4000)
Backend → [http://localhost:3005](http://localhost:3005)
RabbitMQ UI → [http://localhost:15672](http://localhost:15672) (login with `devuser/devpass`)

---

Would you like me to also include **example .env files** for each service (fe, be, upload-worker) in the README so new devs can just copy-paste and run immediately?
