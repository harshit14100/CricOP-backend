ALTER TABLE deliveries
    DROP CONSTRAINT deliveries_inning_id_fkey,
    DROP CONSTRAINT deliveries_striker_id_fkey,
    DROP CONSTRAINT deliveries_non_striker_id_fkey,
    DROP CONSTRAINT deliveries_bowler_id_fkey,
    DROP CONSTRAINT deliveries_fielder_id_fkey,
    DROP CONSTRAINT deliveries_player_out_id_fkey;