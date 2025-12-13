---

## Versión 1 — Estado final ✅

La Versión 1 del proyecto está completa y funcional.

### Qué hace el sistema

- Un **agent local**:
  - Lee métricas reales del sistema (uptime, RAM usada, CPU).
  - Publica métricas cada 5 segundos.
  - Envía los datos como mensajes JSON vía MQTT.

- Un **backend**:
  - Se suscribe al tópico MQTT.
  - Consume y parsea los mensajes.
  - Muestra las métricas recibidas.

### Qué NO incluye esta versión

- Frontend o dashboard.
- Persistencia en base de datos.
- Autenticación.
- Docker, cloud o servicios externos.

### Objetivo cumplido

Entender el flujo completo:
**Sistema → Agent → MQTT → Backend**

Con métricas reales y arquitectura desacoplada.
---

## Versión 3 — Estado final ✅ (Mensajería avanzada)

La Versión 3 introduce una arquitectura de mensajería desacoplada,
incorporando RabbitMQ sin romper versiones anteriores.

### Qué incluye

- **MQTT** como sistema de ingreso de métricas desde agentes.
- **RabbitMQ** como sistema de distribución interna.
- El backend actúa como **bridge** entre MQTT y RabbitMQ.
- Soporte para **múltiples consumidores internos**.
- Consumidor independiente de ejemplo (logger).

### Flujo completo

Agent
└─ MQTT
└─ Backend (ingest + bridge)
├─ HTTP → Frontend
└─ RabbitMQ → Consumers

markdown
Copiar código

### Qué NO incluye

- Persistencia histórica.
- Orquestación o autoescalado.
- Seguridad avanzada.
- Cloud o contenedores.

### Objetivo cumplido

Demostrar una arquitectura realista y evolutiva:
- Entrada liviana de datos (MQTT).
- Distribución interna desacoplada (RabbitMQ).
- Frontend independiente del sistema de mensajería.