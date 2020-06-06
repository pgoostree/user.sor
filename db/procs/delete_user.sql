-- Stored procedure to delete a user record
 CREATE OR REPLACE FUNCTION users.delete_user(
	user_id_in text
 )
RETURNS INT as $$
DECLARE out_rowsaffected INT;
BEGIN
    DELETE FROM users.user
    WHERE user_id = user_id_in;
    GET DIAGNOSTICS out_rowsaffected = ROW_COUNT;
    RETURN out_rowsaffected;
END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;