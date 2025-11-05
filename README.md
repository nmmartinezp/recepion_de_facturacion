# **RECEPCION DE FACTURACIÓN (MICROSERVICIO)**

## **Ejecución del proyecto (Docker)**

Asegurate de tener instalado docker en tu entorno de desarrollo, sino puedes descargalo desde la pagina oficial [aqui](https://www.docker.com/products/docker-desktop/).

Instala el proyecto en tu maquina local:

```bash
git clone https://github.com/nmmartinezp/recepion_de_facturacion.git
```

### **Metodo 1**: Ejecución con uso de `Dockerfile` (Construcción de binario)

1.  Ejecuta el comando en la raiz del proyecto para contruir la imagen docker, puedes modificar el nombre de la imagen si es necesario

    ```bash
    docker build -t microservicio-rf:1.0 .
    ```

2.  Ejecuta el contenedor para la imagen que se contruyo
    ```bash
    docker run microservicio-rf:1.0
    ```

### **Metodo 2**: Ejecución con uso de `Dockerfile.dev` (Desarrollo en caliente)

Para el desarrollo en caliente se recomienda un ambiente linux dado que windows tiene problemas con la sincronizacion de volumenes al usar desarrollo en caliente, en windows puedes usar una maquina virtual o WSL(recomendado) para desarrollo en caliente.

1.  Ejecuta el comando en la raiz del proyecto para contruir la imagen docker, contiene el uso de la herramienta para desarrollo caliente `air`, puedes modificar el nombre de la imagen si es necesario

    ```bash
    docker build -f Dockerfile.dev -t microserviciodev-rf:1.0 .
    ```

2.  Ejecuta el contenedor para la imagen que se contruyo con un volumen sincronizado con el directorio del proyecto y destino de la app
    ```bash
    docker run -v .:/app microserviciodev-rf:1.0
    ```

## **Tencologias Usadas**

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin_Framework-16a085?style=for-the-badge&logo=go&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black" alt="Swagger" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>
