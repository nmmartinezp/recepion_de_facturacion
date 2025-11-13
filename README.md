# **RECEPCION DE FACTURACIÓN (MICROSERVICIO)**

Recibe las facturas electrónicas de las empresas, las valida y genera un registro de facturas válidas.

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
# variables para el entorno del servicio
ENV= # PROD o DEV
PORT=

# variables para la base de datos postgres
DB_POSTGRES_HOST=
DB_POSTGRES_USER=
DB_POSTGRES_PASSWORD=
DB_POSTGRES_NAME=
DB_POSTGRES_PORT=

# Json web token clave
JWT_SECRET=
```

Tambien puede consultar el archivo `.env.example`

## **Base de datos Postgres**

Puede usar cualquier herramienta para crear la siguiente base de datos:

```sql
-- Conectarse a la base de datos recién creada
\connect db_postgres;

-- ======================================
-- Tabla: facturas
-- ======================================
CREATE TABLE IF NOT EXISTS facturas (
    id BIGSERIAL PRIMARY KEY,
    cuf VARCHAR(50) NOT NULL UNIQUE,
    nit_emisor VARCHAR(20) NOT NULL,
    razon_social_emisor VARCHAR(100) NOT NULL,
    fecha_emision TIMESTAMP NOT NULL,
    nit_receptor VARCHAR(20) NOT NULL,
    monto_total NUMERIC(10,2) NOT NULL,
    monto_descuento NUMERIC(10,2) DEFAULT 0,
    codigo_metodo_pago INT NOT NULL,
    codigo_control VARCHAR(50),
    firma_digital TEXT,
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

Puede agregar algunos datos de muestra:

```sql
-- ======================================
-- Datos de ejemplo
-- ======================================

-- 🔹 Factura 1: Venta de laptop
INSERT INTO facturas (cuf, nit_emisor, razon_social_emisor, fecha_emision, nit_receptor, monto_total, monto_descuento, codigo_metodo_pago, codigo_control, firma_digital)
VALUES (
    '3E4F56A7B9CDE1234567890123456789',
    '1234567011',
    'EMPRESA S.A.',
    NOW(),
    '9876543012',
    1500.50,
    0.00,
    1,
    'A1B2C3D4E5F6',
    'MIIC...QAB'
);

INSERT INTO detalles (factura_id, descripcion, cantidad, precio_unitario)
VALUES
((SELECT id FROM facturas WHERE cuf = '3E4F56A7B9CDE1234567890123456789'), 'Laptop Lenovo', 1, 1500.50);

-- 🔹 Factura 2: Venta de artículos de oficina
INSERT INTO facturas (cuf, nit_emisor, razon_social_emisor, fecha_emision, nit_receptor, monto_total, monto_descuento, codigo_metodo_pago, codigo_control, firma_digital)
VALUES (
    '7A9B12C34D56E78F901234567890ABCD',
    '1234567011',
    'EMPRESA S.A.',
    NOW(),
    '9988776655',
    345.75,
    10.00,
    2,
    'Z9Y8X7W6V5U4',
    'MIIC...XYZ'
);

INSERT INTO detalles (factura_id, descripcion, cantidad, precio_unitario)
VALUES
((SELECT id FROM facturas WHERE cuf = '7A9B12C34D56E78F901234567890ABCD'), 'Paquete de hojas A4 (500)', 2, 25.50),
((SELECT id FROM facturas WHERE cuf = '7A9B12C34D56E78F901234567890ABCD'), 'Impresora HP DeskJet 2700', 1, 295.00);

-- 🔹 Factura 3: Venta de mobiliario
INSERT INTO facturas (cuf, nit_emisor, razon_social_emisor, fecha_emision, nit_receptor, monto_total, monto_descuento, codigo_metodo_pago, codigo_control, firma_digital)
VALUES (
    'ABC123DEF456GHI789JKL012MNO345PQ',
    '1234567011',
    'EMPRESA S.A.',
    NOW(),
    '1112223334',
    1020.00,
    0.00,
    3,
    'CTRL987654321',
    'MIIC...LMN'
);

INSERT INTO detalles (factura_id, descripcion, cantidad, precio_unitario)
VALUES
((SELECT id FROM facturas WHERE cuf = 'ABC123DEF456GHI789JKL012MNO345PQ'), 'Escritorio de madera', 1, 450.00),
((SELECT id FROM facturas WHERE cuf = 'ABC123DEF456GHI789JKL012MNO345PQ'), 'Silla ergonómica', 2, 285.00);
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
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>
