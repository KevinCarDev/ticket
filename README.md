# Ticket App

## 🚀 Levantar el proyecto con Docker

1. **Clona el repositorio** y entra a la carpeta del proyecto.

2. **Construye y levanta los servicios** (app y base de datos) con Docker Compose:

   ```sh
   docker-compose up --build
   ```

   Esto levantará:
   - Un contenedor con PostgreSQL (`db`) en el puerto `5432`.
   - Un contenedor con la aplicación Go (`app`) en el puerto `8080`.

3. **Detener los servicios**:

   ```sh
   docker-compose down
   ```

---

## 📚 Endpoints y ejemplos de prueba

### Obtener todos los tickets

```sh
curl http://localhost:8080/tickets
```

### Obtener un ticket por ID

```sh
curl http://localhost:8080/tickets/1
```

### Crear un ticket

```sh
curl -X POST http://localhost:8080/tickets \
  -H "Content-Type: application/json" \
  -d '{"Cliente":"Juan","Origen":"A","Destino":"B","Price":10}'
```

### Eliminar un ticket

```sh
curl -X DELETE http://localhost:8080/tickets/1
```

> También puedes usar Postman importando estos endpoints manualmente.

---

## 🗄️ Base de datos y migraciones

- **Base de datos:** PostgreSQL, configurada en el servicio `db` de Docker Compose.
- **Credenciales por defecto:**
  - Usuario: `user`
  - Contraseña: `pass`
  - Base de datos: `ticket`
  - Host: `db` (interno en Docker), `localhost:5432` (externo)
- **Migraciones:**  
  Al iniciar la aplicación, se ejecuta automáticamente la migración de la tabla `Ticket` usando GORM (`db.AutoMigrate(models.Ticket{})`).  
  No necesitas ejecutar migraciones manualmente.

---

## 📝 Notas

- Si cambias la estructura del modelo `Ticket`, reinicia los contenedores para aplicar los cambios en la base de datos.
- Los datos de la base de datos se almacenan en un volumen Docker llamado `db_data` para persistencia entre reinicios.
- Si tienes problemas de conexión, asegúrate de que los servicios estén levantados y que las variables de entorno coincidan.

---