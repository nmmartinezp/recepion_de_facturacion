# **RECEPCION DE FACTURACIÓN (MICROSERVICIO)**

Recibe las facturas electrónicas de las empresas, las valida y genera un registro de facturas con su reporte de estado.

### **Funciones del servicio**

- Recepcion de facturas
- Validación de facturas
- Registro de facturas
- Registro de estado de facturas

### **Comunicación con otros servicios**

- Validación de autenticacón de JWT con servicio "auth"
- Validación de CUFD con servicio "empresas cufd"
- Mensajeria sobre registro de facturas con serivio "notificaciones"

## **Ejecución del proyecto (Docker)**

Asegurate de tener instalado docker en tu entorno de desarrollo, sino puedes descargalo desde la pagina oficial [aqui](https://www.docker.com/products/docker-desktop/).

Instala el proyecto en tu maquina local:

```bash
git clone https://github.com/nmmartinezp/recepion_de_facturacion.git
```

### **Metodo 1**: Ejecución con uso de `Dockerfile` (Construcción de binario)

- Construcción de binaria main de la api

### **Metodo 2**: Ejecución con uso de `Dockerfile.dev` (Desarrollo en caliente)

- Uso de la herramienta air para crear un binario temporal actualizable en caliente

### **Variables de Entorno para el servicio**:

```env
# Environment configuration
ENV=
PORT=

# db postgres
DB_POSTGRES_HOST=
DB_POSTGRES_USER=
DB_POSTGRES_PASSWORD=
DB_POSTGRES_NAME=
DB_POSTGRES_PORT=

# RabbitMQ
RABBITMQ_URL=
RABBITMQ_EXCHANGE=
```

Tambien puede consultar el archivo `.env.example`

## **Base de datos Postgres**

Puede usar cualquier herramienta para crear la siguiente base de datos:

```sql
CREATE TYPE estado_factura AS ENUM ('ACTIVO', 'ANULADO');
-- ======================================
-- Tabla: facturas
-- ======================================
CREATE TABLE IF NOT EXISTS facturas (
    id BIGSERIAL PRIMARY KEY,
    cuf VARCHAR(100) NOT NULL UNIQUE,
    cufd VARCHAR(100) NOT NULL UNIQUE,
    nit_emisor VARCHAR(100) NOT NULL,
    codigo_sucursal VARCHAR(100) NOT NULL,
    codigo_pv VARCHAR(100) NOT NULL,
    razon_social_emisor VARCHAR(100) NOT NULL,
    fecha_emision TIMESTAMP NOT NULL,
    nit_empresa VARCHAR(20) NOT NULL,
    monto_total NUMERIC(10,2) NOT NULL,
    codigo_control VARCHAR(50),
    estado estado_factura NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- ======================================
-- Tabla: detalles
-- ======================================
CREATE TABLE IF NOT EXISTS detalles (
    id BIGSERIAL PRIMARY KEY,
    factura_id BIGINT NOT NULL REFERENCES facturas(id) ON DELETE CASCADE ON UPDATE CASCADE,
    descripcion VARCHAR(200) NOT NULL,
    cantidad INT NOT NULL,
    precio_unitario NUMERIC(10,2) NOT NULL
);

-- ======================================
-- Índices
-- ======================================
CREATE INDEX IF NOT EXISTS idx_facturas_cuf ON facturas (cuf);
CREATE INDEX IF NOT EXISTS idx_detalles_factura_id ON detalles (factura_id);
```

## **Documentacion de API del servicio**

La api se encuentra protegida con JWT que debe ser proporcionado por un servicio auth.

El servicio contiene documentación `Swagger` que se puede consultar con la ruta base:

```http
/swagger/index.html
```

Las rutas documentadas son:

```http
GET /api/v1/facturacion/facturas
GET /api/v1/facturacion/facturas/:cuf
POST /api/v1/facturacion/facturas
```

## **Tencologias Usadas**

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin_Framework-16a085?style=for-the-badge&logo=go&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black" alt="Swagger" />
  <img src="https://img.shields.io/badge/gRPC-5282FF?style=for-the-badge&logo=grpc&logoColor=white" alt="gRPC" />
  <img src="https://img.shields.io/badge/RabbitMQ-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white" alt="RabbitMQ" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>
