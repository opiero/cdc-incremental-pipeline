DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'replicator') THEN

      RAISE NOTICE 'Role "replicator" already exists. Skipping.';
   ELSE
      CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'password';
   END IF;
END
$do$;
