-- Stored procedure to delete a group record
 CREATE OR REPLACE FUNCTION users.delete_group(
	group_name_in text
 )
RETURNS INT as $$
DECLARE out_rowsaffected INT;
BEGIN
    DELETE FROM users.group
    WHERE group_name = group_name_in;
    GET DIAGNOSTICS out_rowsaffected = ROW_COUNT;
    RETURN out_rowsaffected;
END;
$$ LANGUAGE 'plpgsql' SECURITY DEFINER;