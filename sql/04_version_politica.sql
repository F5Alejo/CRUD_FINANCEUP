-- Agrega a una base ya creada la version de la politica de datos que acepto cada usuario.
-- No borra nada: solo agrega la columna y marca con la version 1.0 a quienes ya aceptaron.
-- Ejecutar una vez: psql -d financeup -f 04_version_politica.sql

ALTER TABLE "auth"."usuario" ADD COLUMN IF NOT EXISTS "version_politica" VARCHAR(10);

UPDATE "auth"."usuario"
   SET "version_politica" = '1.0'
 WHERE "acepta_terminos" = true AND "version_politica" IS NULL;
