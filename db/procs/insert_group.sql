-- Stored procedure to insert a group record
 CREATE OR REPLACE FUNCTION users.insert_group(
	group_name_in text
 )
 RETURNS SETOF users.group AS $$
 BEGIN
     RETURN QUERY
     INSERT INTO users.group (
       group_name
     )
     VALUES (
       group_name_in
     )
        RETURNING group_name;
 END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;