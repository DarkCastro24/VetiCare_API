# VetiCare_API 

**VetiCare_API** es una API RESTful desarrollada en lenguaje Go con arquitectura en capas. Este proyecto fue creado como parte de la asignatura de **Programación N-Capas** en la Universidad Centroamericana José Simeón Cañas (UCA).

## Descripción

La API gestiona las funcionalidades básicas de un sistema de veterinaria, implementando buenas prácticas como separación de responsabilidades, validaciones, servicios reutilizables y middleware.

## Tecnologías

- **Lenguaje**: Go 
- **Estilo de arquitectura**: RESTful API en capas
- **Gestión de dependencias**: `go mod`
- **Despliegue**: Puede desplegarse localmente o en servicios cloud
- **JWT**: Para mejorar la seguridad en las peticiones.
- **Bycrypt**: Utilizado para encriptacion de
- **Deploy**: Sistema desplegado en la nube mediante Vercel y Railway.
  
## Arquitectura por Capas

El proyecto está organizado en las siguientes carpetas:

- `controllers/`: Maneja las peticiones HTTP y respuestas.
- `data/`: Contiene la configuración de base de datos (conexión, seed, migraciones).
- `entities/`: Estructuras de datos o modelos utilizados en el sistema.
- `middlewares/`: Funciones para autenticar, loggear o interceptar peticiones.
- `repositories/`: Encapsulan el acceso a datos (consultas SQL).
- `services/`: Contienen la lógica de negocio.
- `validators/`: Validaciones estructurales para entradas del usuario.
- `utils/`: Funciones auxiliares reutilizables.
- `main.go`: Punto de entrada de la aplicación.

## Variables de entorno

Se incluye el archivo `.env.example` como referencia para definir tus variables de configuración necesarias (puerto, DB, etc.).

## Autor

**Diego Eduardo Castro Quintanilla** 
Carnet: 00117322  
UCA 2025


