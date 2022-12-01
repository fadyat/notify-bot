create or replace function calc_next_cron_time(t timestamp, f frequency_type) returns timestamp as
$$
begin
    case f
        when 'daily' then return t + interval '1 day';
        when 'weekly' then return t + interval '1 week';
        when 'monthly' then return t + interval '1 month';
        else return t;
        end case;
end;
$$ language plpgsql;

