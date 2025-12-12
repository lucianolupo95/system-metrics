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
