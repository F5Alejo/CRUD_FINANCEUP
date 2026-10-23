-- Historial de respuestas de las PQR para crud_soporte.
-- Ejecutar una vez contra la base financeup.

CREATE TABLE IF NOT EXISTS "soporte"."respuesta_pqr" (
    "id_respuesta"     SERIAL PRIMARY KEY,
    "id_pqr"           INT NOT NULL REFERENCES "soporte"."pqr"("id_pqr") ON DELETE CASCADE,
    "id_usuario"       INT REFERENCES "auth"."usuario"("id_usuario") ON DELETE SET NULL,
    "autor"            VARCHAR(30) NOT NULL,
    "mensaje"          TEXT NOT NULL,
    "fecha_respuesta"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "activo"           BOOLEAN NOT NULL DEFAULT true
);

CREATE INDEX IF NOT EXISTS ix_respuesta_pqr_pqr_activo
    ON "soporte"."respuesta_pqr" ("id_pqr", "activo", "id_respuesta");

