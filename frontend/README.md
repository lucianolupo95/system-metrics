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
---

## Versión 2 — Estado final ✅ (Dashboard)

La Versión 2 agrega una interfaz web para visualizar métricas del sistema.

### Qué incluye

- Frontend en **React + Vite**.
- UI con **MUI (Material UI)**.
- Dashboard simple con tarjetas:
  - Uptime (formateado).
  - Memoria RAM usada.
  - Uso de CPU con color según carga.
- Actualización periódica mediante **polling HTTP**.
- Backend expone métricas vía `/metrics` con CORS habilitado.

### Qué NO incluye

- WebSockets o tiempo real push.
- Persistencia histórica.
- Autenticación.
- Docker o cloud.

### Flujo completo

