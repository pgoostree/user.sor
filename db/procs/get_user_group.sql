-- Stored procedure to get user_group records
 CREATE OR REPLACE FUNCTION users.get_user_group(
	group_name_in text
 )
 RETURNS SETOF text as $$
BEGIN
    RETURN QUERY
    SELECT
        user_id
    FROM users.user_group
    WHERE group_name = group_name_in;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;