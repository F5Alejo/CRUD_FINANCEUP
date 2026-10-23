-- Tablas requeridas por los codigos de verificacion del login y el acceso social.
-- Ejecutar una vez contra la base financeup.

CREATE TABLE IF NOT EXISTS "auth"."codigo_verificacion" (
    "id_codigo"           SERIAL PRIMARY KEY,
    "id_usuario"          INT NOT NULL REFERENCES "auth"."usuario"("id_usuario") ON DELETE CASCADE,
    "proposito"           VARCHAR(40) NOT NULL,
    "canal"               VARCHAR(20) NOT NULL,
    "codigo_hash"         VARCHAR(255) NOT NULL,
    "fecha_expiracion"    TIMESTAMP NOT NULL,
    "fecha_uso"           TIMESTAMP,
    "intentos"            INT NOT NULL DEFAULT 0,
    "activo"              BOOLEAN NOT NULL DEFAULT true,
    "fecha_creacion"      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "fecha_modificacion"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS ix_codigo_verificacion_usuario_proposito
    ON "auth"."codigo_verificacion" ("id_usuario", "proposito", "activo");

CREATE TABLE IF NOT EXISTS "auth"."cuenta_social" (
    "id_cuenta_social"  SERIAL PRIMARY KEY,
    "id_usuario"        INT NOT NULL REFERENCES "auth"."usuario"("id_usuario") ON DELETE CASCADE,
    "proveedor"         VARCHAR(40) NOT NULL,
    "id_externo"        VARCHAR(255) NOT NULL,
    "email_proveedor"   VARCHAR(150) NOT NULL,
    "activo"            BOOLEAN NOT NULL DEFAULT true,
    UNIQUE ("proveedor", "id_externo")
);

