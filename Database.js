create or replace function create_database(name text)
returns void
language plpgsql
as $$
begin
  execute format('create database %I', name);
end;
$$;
