DO
$$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_publication
      WHERE pubname = 'students_publication') THEN

      CREATE PUBLICATION students_publication FOR ALL TABLES;
   ELSE
      RAISE NOTICE 'Publicação "nome_da_publicacao" já existe. Ignorando.';
   END IF;
END
$$;
