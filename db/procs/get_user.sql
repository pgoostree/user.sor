-- Stored procedure to get a user record
 CREATE OR REPLACE FUNCTION users.get_user(
	user_id_in text
 )
 RETURNS SETOF users.user as $$
BEGIN
    RETURN QUERY
    SELECT
        user_id,
        first_name,
        last_name
    FROM users.user
    WHERE user_id = user_id_in;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;