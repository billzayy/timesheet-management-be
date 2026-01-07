/* User Daily View */
CREATE OR REPLACE VIEW user_daily_summary_v AS
SELECT
    u.id AS user_id,

    COALESCE(u.sur_name, '') || ' ' || COALESCE(u.last_name, '') AS full_name,
    u.email,
    u.dob,
    u.gender,
    u.phone,
    u.current_address,
    u.address,
    u.avatar_path,
    u.bank_account,
    u.identify_number,
    u.identify_issue_date,
    u.identify_place,
    u.emergency_contact,
    u.emergency_contact_phone,
    u.tax_code,
    u.mezon_id,

    u.level_id,
    u.branch_id,
    u.position_id,
    u.user_type_id,

    b.name  AS branch_name,
    l.name  AS level_name,
    p.name  AS position_name,
    ut.name AS user_type_name,

    /* Morning shift */
    MIN(wt.start_time)
        FILTER (WHERE wt.shift_name = 'morning') AS morning_start_at,

    MAX(wt.end_time)
        FILTER (WHERE wt.shift_name = 'morning') AS morning_end_at,

    ROUND(
        (
            SUM(
                EXTRACT(EPOCH FROM
                    GREATEST(wt.end_time - wt.start_time, INTERVAL '0')
                ) / 3600
            )
            FILTER (WHERE wt.shift_name = 'morning')
        )::numeric,
        1
    ) AS morning_working_time,

    /* Afternoon shift */
    MIN(wt.start_time)
        FILTER (WHERE wt.shift_name = 'afternoon') AS afternoon_start_at,

    MAX(wt.end_time)
        FILTER (WHERE wt.shift_name = 'afternoon') AS afternoon_end_at,

    ROUND(
        (
            SUM(
                EXTRACT(EPOCH FROM
                    GREATEST(wt.end_time - wt.start_time, INTERVAL '0')
                ) / 3600
            )
            FILTER (WHERE wt.shift_name = 'afternoon')
        )::numeric,
        1
    ) AS afternoon_working_time

FROM users u
LEFT JOIN branches   b  ON b.id  = u.branch_id
LEFT JOIN levels     l  ON l.id  = u.level_id
LEFT JOIN positions  p  ON p.id  = u.position_id
LEFT JOIN user_type  ut ON ut.id = u.user_type_id
LEFT JOIN working_times wt
       ON wt.entity_id   = u.id
      AND wt.entity_type = 'user'

GROUP BY
    u.id,
    b.name,
    l.name,
    p.name,
    ut.name;

/* Branch Working Time View */
CREATE OR REPLACE VIEW branch_working_time_summary_v AS
SELECT
	b.id,
  b.name,
	b.display_name,
	b.code,
	b.color,

    /* Morning shift */
    MIN(wt.start_time)
        FILTER (WHERE wt.shift_name = 'morning') AS morning_start_at,

    MAX(wt.end_time)
        FILTER (WHERE wt.shift_name = 'morning') AS morning_end_at,

    ROUND(
        (
            SUM(
                EXTRACT(EPOCH FROM
                    GREATEST(wt.end_time - wt.start_time, INTERVAL '0')
                ) / 3600
            )
            FILTER (WHERE wt.shift_name = 'morning')
        )::numeric,
        1
    ) AS morning_working_time,

    /* Afternoon shift */
    MIN(wt.start_time)
        FILTER (WHERE wt.shift_name = 'afternoon') AS afternoon_start_at,

    MAX(wt.end_time)
        FILTER (WHERE wt.shift_name = 'afternoon') AS afternoon_end_at,

    ROUND(
        (
            SUM(
                EXTRACT(EPOCH FROM
                    GREATEST(wt.end_time - wt.start_time, INTERVAL '0')
                ) / 3600
            )
            FILTER (WHERE wt.shift_name = 'afternoon')
        )::numeric,
        1
    ) AS afternoon_working_time

FROM branches b
LEFT JOIN working_times wt ON wt.entity_id = b.uuid AND wt.entity_type = 'branch'
GROUP BY b.id;
