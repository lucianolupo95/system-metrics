# system-metrics

Proyecto educativo y práctico para aprender monitoreo de sistemas, mensajería y arquitectura desacoplada, inspirado en escenarios reales de IoT.

El proyecto se desarrolla en **versiones incrementales**, priorizando comprensión, simplicidad y estabilidad.

---

## Objetivo general

Medir métricas reales de la **PC local** (CPU, RAM, uptime) y transmitirlas mediante sistemas de mensajería, comenzando de forma simple y evolucionando paso a paso.

---

## Versiones del proyecto

### Versión 1 — Base (sin UI)
- Un **agent local** que:
  - Lee métricas del sistema.
  - Publica datos usando **MQTT**.
- Un **backend** que:
  - Se suscribe a MQTT.
  - Recibe y procesa las métricas.
- No hay frontend.
- Enfoque: entender el flujo de datos y la mensajería.

### Versión 2 — Dashboard
- Se agrega un **frontend en React + MUI**.
- El backend expone datos para el frontend.
- Visualización simple y en tiempo casi real.
- Enfoque: consumo de datos y visualización.

### Versión 3 — Mensajería avanzada
- Se incorpora **RabbitMQ** además de MQTT.
- Cada sistema cumple un rol claro.
- Componentes desacoplados.
- Las versiones anteriores siguen funcionando.
- Enfoque: arquitectura y comparación de tecnologías.

---

## Estructura del repositorio

