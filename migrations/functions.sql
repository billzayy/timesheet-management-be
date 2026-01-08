CREATE OR REPLACE FUNCTION get_children(p_id bigINT)
RETURNS JSONB LANGUAGE SQL AS $$
  SELECT COALESCE(
    jsonb_agg(
      jsonb_build_object(
        'name', c.name,
		    'display_name', c.display_name,
        'children', get_children(c.id)
      )
    ), NULL
  )
  FROM permissions c
  WHERE c.parent_id = p_id;
$$;

WITH RECURSIVE tree AS (
  SELECT
    name,
    display_name,
    parent_id
  FROM permissions
)
SELECT jsonb_pretty(
  jsonb_agg(
    jsonb_build_object(
      'name', p.name,
	    'display_name',p.display_name,
      'children', get_children(p.id)
    )
  )
) AS tree
FROM permissions p
WHERE parent_id IS NULL;
